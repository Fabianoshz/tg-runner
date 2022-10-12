package cmd

import (
	"github.com/fabianoshz/tg-runner/internal/usecase"
	"github.com/spf13/cobra"
)

type executionPlannerCmd struct {
	executionPlanner usecase.PlanExecutionInterface
}

func newExecutionPlannerCmd(executionPlanner usecase.PlanExecutionInterface) command {
	return &executionPlannerCmd{
		executionPlanner: executionPlanner,
	}
}

func (p executionPlannerCmd) newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan",
		Short: "run plan",
		Run: func(cmd *cobra.Command, args []string) {
			p.executionPlanner.PlanExecution("internal/usecase/testdata/changelist-plan.yaml")
		},
	}

	return cmd
}
