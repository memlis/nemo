package mesos

type EventHandler interface {
	Handle(sched.Event)
}
