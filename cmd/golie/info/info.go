package info

import (
	"errors"

	"github.com/rolieup/golie/pkg/rolie"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "info ROLIE-RESOURCE",
	Short: "Print summary information about given ROLIE resource",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Please provide path or URL to your pre-existing ROLIE file as command-line argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return rolie.Info(args[0])
	},
}
