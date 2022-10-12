package app

import (
	"github.com/fabianoshz/iflantis/internal/repository"
	"github.com/fabianoshz/iflantis/internal/usecase"
)

type App struct {
	ExecutionPlanner usecase.ExecutionPlanner
}

func Start() (*App, error) {
	persistenceRepository := repository.NewPersistenceClient()

	executionPlanner := usecase.NewExecutionPlannerService(persistenceRepository)

	return &App{
		ExecutionPlanner: executionPlanner,
	}, nil
}
