package application

import (
	"github.com/memlis/boat/types"
)

type Backend interface {
	ClusterId() string
	// RegisterApplication register application in consul.
	SaveApplication(*types.Application) error

	// RegisterApplicationVersion register application version in consul.
	SaveVersion(string, *types.Version) error

	// LaunchApplication launch applications
	LaunchApplication(*types.Version) error

	// DeleteApplication will delete all data associated with application.
	DeleteApplication(string) error

	// DeleteApplicationTasks delete all tasks belong to appcaiton but keep that application exists.
	DeleteApplicationTasks(string) error

	ListApplications() ([]*types.Application, error)

	FetchApplication(string) (*types.Application, error)

	ListApplicationTasks(string) ([]*types.Task, error)

	DeleteApplicationTask(string, string) error

	ListApplicationVersions(string) ([]string, error)

	FetchApplicationVersion(string, string) (*types.Version, error)

	UpdateApplication(string, int, *types.Version) error

	ScaleApplication(string, int) error

	RollbackApplication(string) error
}
