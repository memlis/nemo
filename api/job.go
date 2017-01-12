package api

type Job struct {
	ID           string
	Name         string
	Type         string
	Priority     int
	User         string
	Cluster      string
	Replicas     int
	Status       string
	Created      time.Time
	Updated      time.Time
	UpdatedTimes int
	LiveTime     string
	Health       string

	Update  *UpdateStrategy
	Restart *RestartStrategy
	Kill    *KillStrategy
}

func NewJob(spec *JobSpec) *Job {
	job := &Job{
		Name:     spec.Name,
		Type:     spec.Type,
		Priority: spec.Priority,
		User:     spec.User,
		Replicas: spec.Replicas,
	}

	job.ID = NexId()

	return job
}

func (job *Job) Status() string {
	for _, tasks := ranage app.Tasks {
		if task.Status == mesos.TASK_STATUS_RUNNING {
			return JOB_STATUS_RUNNING
		}
	}

	return JOB_STATUS_FAILED
}

func (job *Job) Stop() {
	for _, task := range job.Tasks {
		task.Stop()
	}
}

func (job *Job) Start() {
	for _, task := range job.Tasks {
		task.Start()
	}
}

//func (job *Job) Status() string {
//	return ""
//}
//
//func (job *Job) Healthy() string {
//	return ""
//}
//
//
//func (job *Job) RunningInstance() int {
//	return 0
//}
//
//func (job *Job) TotalCpus() float64 {
//	return 0.0
//}
//
//func (job *Job) TotalGpus() float64 {
//	return 0.0
//}
//
//func (job *Job) TotalMem float64 {
//	return 0.0
//}
//
//func (job *Job) TotalDisk float64 {
//	return 0.0
//}
//
//func (job *Job) Update() error {
//	return nil
//}
//
//func (job *Job) Delete() error {
//	return nil
//}
