package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tiulpin/qodana/pkg"
)

func NewInitCommand() *cobra.Command {
	options := &pkg.LinterOptions{}
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Create qodana.yaml",
		Long:  "Prepare Qodana configuration file",
		PreRun: func(cmd *cobra.Command, args []string) {
			pkg.EnsureDockerRunning()
		},
		Run: func(cmd *cobra.Command, args []string) {
			pkg.PrintProcess(
				func() { pkg.ConfigureProject(options) },
				"Configuring project",
				"project configuration. Check qodana.yaml.")
			pkg.Primary.Println("🚀  Run `qodana scan` to analyze the project")
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&options.ProjectPath, "project-dir", "i", ".", "Root directory of the inspected project")
	return cmd
}
