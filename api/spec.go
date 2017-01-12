package api

type JobSpec struct {
	Name     string
	Type     string
	Priority int
	User     string
	Replicas int

	*TaskSpec

	Restart *RestartStrategy
	Update  *UpdateStrategy
	Kill    *KillStrategy
}

type TaskSpec struct {
	Image          string
	Cpu            float64
	Gpu            float64
	Memory         float64
	Disk           float64
	Network        string
	ForcePullImage string
	Privileged     bool
	Priority       int
	Environment    map[string]string
	Command        string
	Labels         map[string]string
	Ports          []*Port
	Volumes        []*Volume
	Uris           []*URI
	Constraints    []*Constraints
}

func NewTaskSpec(spec *JobSpec) *TaskSpec {
	return &TaskSpec{
		Image:          spec.Image,
		Cpu:            spec.Cpu,
		Gpu:            spec.Gpu,
		Memory:         spec.Memory,
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
