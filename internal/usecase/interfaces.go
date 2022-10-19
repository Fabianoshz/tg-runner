package usecase

import "github.com/fabianoshz/tg-runner/internal/entity"

type PlanExecutionInterface interface {
	PlanExecution(string) bool
}

type CalculateDependenciesInterface interface {
	CalculateDependencies([]entity.Resource, string) [][]entity.Resource
}

type LoadResourcesInterface interface {
	LoadResources(string) []entity.Resource
}
