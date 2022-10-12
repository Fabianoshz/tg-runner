package usecase

import (
	"github.com/google/uuid"
)

type Lock struct {
	id uuid.UUID
}

func AcquireLock() Lock {
	l := Lock{id: uuid.New()}

	// TODO create lock logic

	return l
}

func (l Lock) Release() {
	// TODO release the lock
}
