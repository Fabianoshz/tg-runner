package repository

import "github.com/google/uuid"

type Persistence interface {
	SavePlanfile(string, string, uuid.UUID, string)
	GetPlanfiles(uuid.UUID)
}
