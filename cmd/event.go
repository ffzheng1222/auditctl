package cmd

import (
	"auditctl/pkg/event"
	"fmt"

	"github.com/spf13/cobra"
)

// NewEventCmd ...
func NewEventCmd() *cobra.Command {
	var eventCmd = &cobra.Command{
		Use:     "event",
		Short:   "Tools of this auditctl",
		Long:    "Audit event of [get, list, create, update, patch, delete]",
		Example: "auditctl event list",
		Run: func(cmd *cobra.Command, args []string) {
			fullCmdName := cmd.Parent().CommandPath()
			usageString := fmt.Sprintf("\nUse \"./auditctl event <verb>\" for a detailed description (e.g. %[1]s list).", fullCmdName)
			err := UsageErrorf(cmd, usageString)
			fmt.Printf("%s\n", err)
		},
	}

	eventCmd.AddCommand(event.GetCmd)
	eventCmd.AddCommand(event.ListCmd)
	eventCmd.AddCommand(event.CreateCmd)
	eventCmd.AddCommand(event.UpdateCmd)
	eventCmd.AddCommand(event.PatchCmd)
	eventCmd.AddCommand(event.DeleteCmd)

	return eventCmd
}

func UsageErrorf(cmd *cobra.Command, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s\nSee './%s -h' for help and examples", msg, cmd.CommandPath())
}
