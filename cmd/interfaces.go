package cmd

import "github.com/spf13/cobra"

type RootCommand interface {
	Execute()
}

type command interface {
	newCommand() *cobra.Command
}
