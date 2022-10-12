package main

import (
	"github.com/fabianoshz/tg-runner/cmd"
	"github.com/fabianoshz/tg-runner/internal/app"
)

func main() {
	app, err := app.Start()
	if err != nil {
		panic(err)
	}

	rootCoomand := cmd.NewCLIRootCommand(app)
	rootCoomand.Execute()

	// persistenceRepository := repository.NewPersistenceClient()

	// executionPlanner := usecase.NewExecutionPlannerService(persistenceRepository)

	// executionPlanner.PlanExecution("internal/usecase/testdata/changelist-plan.yaml")
}
