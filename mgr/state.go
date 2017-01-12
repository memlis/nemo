package mgr

//
// StateManager period sync state of task between mesos and local storage
//

type StateManager struct {
	stopCh chan struct{}
}

func NewStateManager() *StateManager {

}

func (sm *StateManager) Start() {
	for {
		select {
		case <-sm.stopCh:
			return
		}
	}
}

func (sm *StateManager) Stop() {
	close(sm.stopCh)
}
