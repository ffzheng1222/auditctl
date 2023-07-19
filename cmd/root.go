package cmd

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "file", "f",
		"", "audit log file (default is $(PWD)/audit.log).")
}

func initConfig() {
	var configPath string
	if cfgFile != "" {
		// Use audit log file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Check the current workspace
		pwdPath, err := os.Getwd()
		cobra.CheckErr(err)
		_, err = os.Stat(pwdPath + "/audit.yaml")
		if err == nil {
			configPath = pwdPath
		} else {
			fmt.Printf("Not found audit.yaml in %s", pwdPath)
			os.Exit(1)
		}

		// Search config in home directory with name "audit.yaml" (without extension).
		viper.AddConfigPath(configPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName("audit")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		err = fmt.Errorf("error to open audit.yaml file with error: %s", err)
		fmt.Printf("%s \n", err)
	}

	validateDefaultConfig()
}

func validateDefaultConfig() {
	var allAuditConfigKeys = []string{"logFile"}
	for _, key := range allAuditConfigKeys {
		value := viper.GetString(key)
		if !reflect.ValueOf(value).IsValid() {
			cobra.CheckErr(fmt.Errorf("parameter %s is not configed in audit.yaml", key))
		}
		//fmt.Printf("Audit config: %s \n", value)
	}
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
