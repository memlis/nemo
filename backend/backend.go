package backend

import (
	"github.com/memlis/boat/scheduler"
	. "github.com/memlis/boat/store"
)

type Backend struct {
	sched *scheduler.Scheduler
	store Store
}

func NewBackend(sched *scheduler.Scheduler, store Store) *Backend {
	return &Backend{
		sched: sched,
		store: store,
	}
}

func (b *Backend) ClusterId() string {
	return b.sched.ClusterId
}
