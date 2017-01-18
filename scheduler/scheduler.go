package scheduler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"github.com/memlis/boat/health"
	"github.com/memlis/boat/mesosproto/mesos"
	sched "github.com/memlis/boat/mesosproto/sched"
	"github.com/memlis/boat/scheduler/client"
	"github.com/memlis/boat/store"
	"github.com/memlis/boat/types"
)

// Scheduler represents a Mesos scheduler
type Scheduler struct {
	master       string
	framework    *mesos.FrameworkInfo
	store        store.Store
	client       *client.Client
	doneCh       chan struct{}
	handlers     map[sched.Event_Type]EventHandler
	apiserver    *api.Server
	pendingTasks []*task.Task
        pendingJobs  []*job.Job
        strategy     Strategy
}

// NewScheduler returns a pointer to new Scheduler
func NewScheduler(master string, fw *mesos.FrameworkInfo, store store.Store) *Scheduler {
	scheduler := &Scheduler{
		master:    master,
		client:    client.New(master, "/api/v1/scheduler"),
		framework: fw,
		store:     store,
		doneCh:    make(chan struct{}),
	}

	scheduler.initHandlers()

	return scheduler
}

func (s *Scheduler) initHandlers() {
	s.handlers = map[sched.Event_Type]EventHandler{
		sched.Event_SUBSCRIBED: handler.NewSubscribedHandler(s),
		sched.Event_OFFERS:     handler.NewOfferHandler(s),
		sched.Event_RESCIND:    handler.NewRescindHandler(s),
		sched.Event_UPDATE:     handler.NewUpdateHandler(s),
		sched.Event_MESSAGE:    handler.NewMessageHandler(s),
		sched.Event_FAILURE:    handler.NewFailureHandler(s),
		sched.Event_ERROR:      handler.NewErrorHandler(s),
		sched.Event_HEARTBEAT:  handler.NewHeartBeatHandler(s),
		sched.Event_UNKNOWN:    handler.NewUnknownHandler(s),
	}
}

// start starts the scheduler and subscribes to event stream
// returns a channel to wait for completion.
func (s *Scheduler) Start() <-chan struct{} {
	if err := s.subscribe(); err != nil {
		logrus.Error(err)
		close(s.doneChan)
	}
	return s.doneCh
}

func (s *Scheduler) stop() {
	for _, event := range s.events {
		close(event)
	}
}

func (s *Scheduler) send(call *sched.Call) (*http.Response, error) {
	payload, err := proto.Marshal(call)
	if err != nil {
		return nil, err
	}
	return s.client.Send(payload)
}

// Subscribe subscribes the scheduler to the Mesos cluster.
// It keeps the http connection opens with the Master to stream
// subsequent events.
func (s *Scheduler) subscribe() error {
	logrus.Infof("Subscribe with mesos master %s", s.master)
	call := &sched.Call{
		Type: sched.Call_SUBSCRIBE.Enum(),
		Subscribe: &sched.Call_Subscribe{
			FrameworkInfo: s.framework,
		},
	}

	if s.framework.Id != nil {
		call.FrameworkId = &mesos.FrameworkID{
			Value: proto.String(s.framework.Id.GetValue()),
		}
	}

	resp, err := s.send(call)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Subscribe with unexpected response status: %d", resp.StatusCode)
	}

	logrus.Info(s.client.StreamID)
	go s.handleEvents(resp)

	return nil
}

func (s *Scheduler) handleEvents(resp *http.Response) {
	defer func() {
		resp.Body.Close()
		close(s.doneChan)
		for _, event := range s.events {
			close(event)
		}
	}()
	dec := json.NewDecoder(resp.Body)
	for {
		event := new(sched.Event)
		if err := dec.Decode(event); err != nil {
			if err == io.EOF {
				return
			}
			continue
		}

		s.handlers[event.GetType()].Handle(event)
	}
}
