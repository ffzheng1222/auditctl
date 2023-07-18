package event

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "auditctl event list",
	Long:  "list k8s apiserver audit event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run auditctl event list ...")
	},
}
