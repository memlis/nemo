package api

type Task struct {
	ID    string
	Name  string
	JobID string

	*TaskSpec

	Status            string
	StatusDescription string

	Reborn int

	Info
	Resource
}

func NewTask(spec *TaskSpec) *Task {
	task := &Task{
		Image:          spec.Image,
		Cpu:            spec.Cpu,
		Mem:            spec.Memory,
		Disk:           spec.Disk,
		Network:        spec.Network,
		ForcePullImage: spec.ForcePullImage,
		Privileged:     spec.Privileged,
		Priority:       spec.Priority,
		Environment:    spec.Environment,
		Command:        spec.Command,
		Labels:         spec.Labels,
		Ports:          spec.Ports,
		Volumes:        spec.Volumes,
		Uris:           spec.Uris,
		Constraints:    spec.Constraints,
	}
}

func (t *Task) IsRunning() {
}

func (t *Task) UpdateStatus(status string) {
}

func (t *Task) Persisdent() {
}
