package usecase

import (
	"github.com/fabianoshz/tg-runner/internal/repository"
)

type ExecutionPlannerService struct {
	persistenceRepository repository.Persistence
}

func NewExecutionPlannerService(persistenceRepository repository.Persistence) ExecutionPlanner {
	return &ExecutionPlannerService{
		persistenceRepository: persistenceRepository,
	}
}
