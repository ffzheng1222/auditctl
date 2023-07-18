package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:     "auditctl",
		Short:   "Tool for ssm operation, like upgrade",
		Long:    "Tool used to auditctl, look and analysis k8s apiserver audit log.",
		Example: "auditctl --event=[get,list,create,update,patch,delete]",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("run auditctl...")
		},
	}

	rootCmd.PersistentFlags().StringP("verb", "", "list", "the verb of k8s operate")

	return rootCmd
}

func Execute() {
	rootcmd := NewCommand()
	if err := rootcmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
