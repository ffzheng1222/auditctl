package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	defLogFile string
	rootCmd    = &cobra.Command{
		Use:   "auditctl",
		Short: "Tool for audit log operation",
		Long:  "Tool used to auditctl, look and analysis k8s apiserver audit log.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Use \"./auditctl <event>\"")
			err := fmt.Errorf("\nSee './%s -h' for help and examples", cmd.CommandPath())
			fmt.Printf("%s \n", err)
		},
	}
)

func init() {
	cobra.OnInitialize(getAuditLogFile)
	rootCmd.PersistentFlags().StringVarP(&defLogFile, "file", "f",
		"", "audit log file (default is $(PWD)/audit.log).")
}

func getAuditLogFile() {
	var configFilePath string
	if defLogFile != "" {
		// Use audit log file from the flag.
		viper.SetConfigFile(defLogFile)
	} else {
		// Check the current workspace
		pwdPath, err := os.Getwd()
		cobra.CheckErr(err)
		_, err = os.Stat(pwdPath + "/audit.log")
		if err == nil {
			configFilePath = pwdPath
		} else {
			fmt.Printf("Not found audit.log in %s", pwdPath)
			os.Exit(1)
		}

		// Search config in home directory with name "audit.log" (without extension).
		viper.AddConfigPath(configFilePath)
		viper.SetConfigType("log")
		viper.SetConfigName("audit")
	}

	viper.AutomaticEnv() // read in environment variables that match
}

func NewCommand() *cobra.Command {
	eventcmd := NewEventCmd()
	rootCmd.AddCommand(eventcmd)

	return rootCmd
}

func Execute() {
	rootcmd := NewCommand()
	if err := rootcmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
