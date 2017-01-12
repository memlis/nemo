package health

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	fifo "github.com/foize/go.fifo"
	"github.com/memlis/boat/types"
)

type HealthCheckManager struct {
	store     Store
	checkers  map[string]Checker
	msgQueue  chan types.ReschedulerMsg
	taskQueue *fifo.Queue
	quit      chan struct{}
}

func NewHealthCheckManager(store Store, queue chan types.ReschedulerMsg) *HealthCheckManager {
	return &HealthCheckManager{
		store:     store,
		msgQueue:  queue,
		taskQueue: fifo.NewQueue(),
		checkers:  make(map[string]Checker),
		quit:      make(chan struct{}),
	}
}

func (m *HealthCheckManager) Init() {
	logrus.Info("Initial health checkers...")
	checks, err := m.store.ListChecks()
	if err != nil {
		logrus.Errorf("Initial health checker failed: %s", err)
		return
	}

	if len(checks) == 0 {
		logrus.Info("No checks defined.Skip")
		return
	}

	for _, check := range checks {
		m.Add(check)
	}
}

func (m *HealthCheckManager) Start() {
	for {
		select {
		case <-m.quit:
			return
		default:
			if checker := m.Next(); checker != nil {
				go func() {
					checker.(Checker).Start()
				}()
			}
		}
	}
}

func (m *HealthCheckManager) Stop() {
	close(m.quit)

	for _, checker := range m.checkers {
		checker.Stop()
	}
}

func (m *HealthCheckManager) Add(check *types.Check) {
	switch check.Protocol {
	case "http", "HTTP":
		logrus.Infof("Register health check for task %s protocol %s url %s",
			check.TaskID,
			"http",
			fmt.Sprintf("http://%s:%d", check.Address, check.Port),
		)
		checker := &HTTPChecker{
			ID:          check.TaskID,
			Url:         fmt.Sprintf("http://%s:%d", check.Address, check.Port),
			Interval:    check.Interval,
			Timeout:     check.Timeout,
			MaxFailures: check.MaxFailures,
			FailedHandler: func(appId, taskId string) error {
				return m.HealthCheckFailedHandler(appId, taskId)
			},
			AppID:  check.AppID,
			TaskID: check.TaskID,
			quit:   make(chan struct{}),
		}
		m.taskQueue.Add(checker)
		m.checkers[check.TaskID] = checker
	case "tcp", "TCP":
		logrus.Infof("Add health check for task %s protocol %s address %s port %d",
			check.TaskID,
			"tcp",
			check.Address,
			check.Port,
		)
		checker := &TCPChecker{
			ID:          check.TaskID,
			Addr:        fmt.Sprintf("%s:%d", check.Address, check.Port),
			Interval:    check.Interval,
			Timeout:     check.Timeout,
			MaxFailures: check.MaxFailures,
			FailedHandler: func(appId, taskId string) error {
				return m.HealthCheckFailedHandler(appId, taskId)
			},
			AppID:  check.AppID,
			TaskID: check.TaskID,
			quit:   make(chan struct{}),
		}
		m.taskQueue.Add(checker)
		m.checkers[check.TaskID] = checker
	}

}

func (m *HealthCheckManager) Next() (item interface{}) {
	return m.taskQueue.Next()
}

func (m *HealthCheckManager) StopCheck(id string) {
	if checker, exist := m.checkers[id]; exist {
		logrus.Infof("Remove health check for task %s", id)
		checker.Stop()
		delete(m.checkers, id)
	}
}

func (m *HealthCheckManager) HasCheck(id string) bool {
	_, exists := m.checkers[id]

	return exists
}
