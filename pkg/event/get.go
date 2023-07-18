package event

import (
	"fmt"

	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "auditctl event get",
	Long:  "get k8s apiserver audit event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run auditctl event get ...")
	},
}
