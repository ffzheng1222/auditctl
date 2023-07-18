package event

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "auditctl event create",
	Long:  "create k8s apiserver audit event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run auditctl event create ...")
	},
}
