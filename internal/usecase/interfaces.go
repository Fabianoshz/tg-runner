package usecase

import "github.com/fabianoshz/tg-runner/internal/entity"

type ExecutionPlanner interface {
	PlanExecution(string) bool
}

type ExecutionApplier interface {
	ApplyExecution(entity.Changelist) bool
}

type DependencyCalculator interface {
	CalculateDependencies([]entity.Resource) []entity.Resource
}
