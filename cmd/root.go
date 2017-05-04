package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "sigstat",
	Short: "Sigstat let you monitor long running process by sending process status over a REST api.",
	Long: `Sigstat is a cli command wrapper that take the command passed as argument,
	execute it and send informations such as the status (running, stopped), the exit code, the stderr, the stdout, etc.
	to a server using http calls. The server catch those datas and store them in order to be queried later.
	This is usefull for monitoring the actual execution of long running processes.

	By getting informations on a command at a given timestamp it becomes possible to easily plug monitoring and alerting tools o n top of the REST API.

	Since the command informations are send throw http, it's easy to have multiple wrapper over multiple server that send command information
	to a centralized monitoring server.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sigstat.yaml)")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".sigstat") // name of config file (without extension)
	viper.AddConfigPath("$HOME")    // adding home directory as first search path
	viper.AutomaticEnv()            // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
