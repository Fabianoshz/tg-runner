package cmd

import (
	"fmt"
	"os"

	"github.com/fabianoshz/tg-runner/internal/app"
	"github.com/spf13/cobra"
)

type CLIRootCommand struct {
	app *app.App
}

func NewCLIRootCommand(a *app.App) RootCommand {
	return &CLIRootCommand{
		app: a,
	}
}

func (r CLIRootCommand) Execute() {
	rootCmd := &cobra.Command{
		Use:   "tg-runner",
		Short: "tg-runner",
	}

	executionPlannerCmd := newExecutionPlannerCmd(r.app.ExecutionPlanner)

	rootCmd.AddCommand(executionPlannerCmd.newCommand())

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
