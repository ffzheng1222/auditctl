package event

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "auditctl event list",
	Long:  "list k8s apiserver audit event",
	Run: func(cmd *cobra.Command, args []string) {
		defAuditConfigFile := viper.ConfigFileUsed()
		fmt.Printf("Load k8s audit log file of %s \n\n", defAuditConfigFile)

		listK8sAuditLog()
	},
}

func listK8sAuditLog() {
	auditRecordarr := NewAuditRecordArr()

	for _, auditRecord := range auditRecordarr {
		fmt.Printf("反序列化后 auditLogMap=%v\n", auditRecord)
	}
}
