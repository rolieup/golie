package info

import (
	"errors"
	"fmt"
	"os"

	"github.com/rolieup/golie/pkg/rolie"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "info ROLIE-RESOURCE",
	Short: "Print summary information about given ROLIE resource",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Please provide path to your pre-existing ROLIE file as command-line argument")
		}
		path := args[0]
		stat, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("Could not stat '%s': %v", path, err)
		}
		if stat.IsDir() {
			return fmt.Errorf("Provided path '%s' does not point to a file but to a directory", path)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return rolie.Info(args[0])
	},
}
