package main

type Framework struct {
	FrameworkInfo info
	pendingTasks  map[taskId]TaskInfo
	tasks         map[taskId]Task
	offers        []*mesos.Offer
	scheduler     Scheduler
	apiserver     api.Server
}

func NewFramework() *Framework {
	return &Framework{}
}

func (fw *Framework) Run() {
}

func (fw *Framework) Stop() {
}
