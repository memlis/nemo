package mgr

type UpdateManager struct {
	store  Store
	sched  Scheduler
	stopCh chan struct{}
}

func NewUpdateManager(store Store) *UpdateManager {
	return &UpdateManager{
		store: store,
	}
}

func (um *UpdateManager) Start() {
	ticker := time.NewTicker(time.Duration(1 * time.Second))
	for {
		select {
		case <-ticker.C:
		case <-um.stopCh:
			return
		}
	}
}

func (um *UpdateManager) Stop() {
	close(um.stopCh)
}
