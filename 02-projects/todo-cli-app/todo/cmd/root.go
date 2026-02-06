/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var dataFile string

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo app with Cobra and Viper",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo-cli-app.yaml)")

	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", "todo.json", "data file to store todos")

	viper.BindPFlag("datafile", rootCmd.PersistentFlags().Lookup("datafile"))
}

func initConfig() {

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(".")
		viper.AddConfigPath(home + "/.todo")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("todo")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}
