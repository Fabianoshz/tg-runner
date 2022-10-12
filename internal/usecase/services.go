package usecase

import (
	"github.com/fabianoshz/tg-runner/internal/repository"
)

type PlanExecutionService struct {
	persistenceRepository repository.Persistence
}

type LoadResourcesService struct {
}

type CalculateDependenciesService struct {
}

func NewPlanExecutionService(persistenceRepository repository.Persistence) PlanExecutionInterface {
	return &PlanExecutionService{
		persistenceRepository: persistenceRepository,
	}
}

func NewLoadResourcesService() LoadResourcesInterface {
	return &LoadResourcesService{}
}

func NewCalculateDependenciesService() CalculateDependenciesInterface {
	return &CalculateDependenciesService{}
}
