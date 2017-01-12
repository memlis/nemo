package health

import (
	"github.com/memlis/boat/types"
)

type Store interface {
	ListChecks() ([]*types.Check, error)
}
