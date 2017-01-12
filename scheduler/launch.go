package scheduler

type LaunchManager struct {
	tasks []*Task
}

func NewLaunchManager() *LaunchManager {
	return &LaunchManager{}
}
