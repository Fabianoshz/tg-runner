package app

import (
	"github.com/fabianoshz/tg-runner/internal/repository"
	"github.com/fabianoshz/tg-runner/internal/usecase"
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
