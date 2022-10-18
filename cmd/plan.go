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
		Short: "Run an execution plan based on a changelist",
		Run: func(cmd *cobra.Command, args []string) {
			changelistFile, _ := cmd.Flags().GetString("changelist")

			if changelistFile != "" {
				p.executionPlanner.PlanExecution(changelistFile)
			}
		},
	}

	cmd.PersistentFlags().String("changelist", "", "The changelist file you wish to generate a plan")

	return cmd
}
