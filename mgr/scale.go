package mgr

type ScaleManager struct {
	stopCh chan struct{}
}

func NewScaleManager() *ScaleManager {
}

func (sm *ScaleManager) Start() {
	for {
		select {
		case <-sm.stopCh:
			return
		}
	}
}

func (sm *ScaleManager) Stop() {
	close(sm.stopCh)
}
