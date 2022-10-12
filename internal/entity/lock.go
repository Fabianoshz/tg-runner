package entity

import (
	"github.com/google/uuid"
)

type Lock struct {
	Id uuid.UUID
}

func AcquireLock() Lock {
	l := Lock{Id: uuid.New()}

	// TODO create lock logic

	return l
}

func (l Lock) Release() {
	// TODO release the lock
}
