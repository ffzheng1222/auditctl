package event

import (
	"fmt"

	"github.com/spf13/cobra"
)

var PatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "auditctl event patch",
	Long:  "patch k8s apiserver audit event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run auditctl event patch ...")
	},
}
