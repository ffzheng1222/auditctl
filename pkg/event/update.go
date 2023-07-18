package event

import (
	"fmt"

	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "auditctl event update",
	Long:  "update k8s apiserver audit event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run auditctl event update ...")
	},
}
