package scheduler

type Scheduler interface {
	Start()
	Stop()
	UpdateStatus()
	LaunchTask()
	KillTask()
}
