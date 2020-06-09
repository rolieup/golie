/*
Copyright © 2020 Rolie and Golie Contributors. See LICENSE for license.
*/

package main

import (
	"fmt"
	"os"
	"path"

	"github.com/rolieup/golie/cmd/golie/new"
	golie "github.com/rolieup/golie/golie/client"
	"github.com/rolieup/golie/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const loglevel = "loglevel"

type globalOpts struct {
	Debug    bool
	CfgFile  string
	Loglevel string
	Version  string
}

var (
	globalFlags globalOpts
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   path.Base(os.Args[0]),
	Short: "A client implementation of ROLIE",
	Long: `A client implementation of the Resource-Oriented Lightweight
Information Exchange (ROLIE) specification.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return setup(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		golie.Client()
		return nil
	},
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
	})

	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(new.Cmd)
	rootCmd.TraverseChildren = true
	rootCmd.Version = fmt.Sprintf("%s, build: %s, date: %s", version.Version, version.Commit, version.Date)
	rootCmd.PersistentFlags().BoolVar(&globalFlags.Debug, "debug", false, "Run in debug mode")
	rootCmd.PersistentFlags().StringVar(&globalFlags.Loglevel, loglevel, "error", `Set log verbosity. Options are "debug", "info", "warn" or "error".`)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

}

func setup(cmd *cobra.Command) error {
	loglevelOpt, err := cmd.Flags().GetString(loglevel)
	if err != nil {
		return err
	}

	loglevel, err := log.ParseLevel(loglevelOpt)
	if err != nil {
		return fmt.Errorf("unable to parse log level: %v", err)
	}

	log.SetLevel(loglevel)
	if globalFlags.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Running client in DEBUG mode")
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
