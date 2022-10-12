package usecase

import "github.com/fabianoshz/iflantis/internal/entity"

type ExecutionPlanner interface {
	PlanExecution(entity.Changelist) bool
}

type ExecutionApplier interface {
	ApplyExecution(entity.Changelist) bool
}

type DependencyCalculator interface {
	CalculateDependencies([]entity.Resource) []entity.Resource
}
