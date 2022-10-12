package entity

import "github.com/google/uuid"

type Changelist struct {
	Id        uuid.UUID
	Resources []Resource
}
