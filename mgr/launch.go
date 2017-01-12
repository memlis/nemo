package mgr

type LaunchManager struct {
	tasks  []*types.Task
	stopCh chan struct{}
}

func NewLaunchManager() *LaunchManager {
}

func (lm *LaunchManager) Start() {
	for {
		select {
		case <-lm.stopCh:
			return
		}
	}
}

func (lm *LaunchManager) Stop() {
	close(stopCh)
}
