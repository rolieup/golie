package clone

import (
	"errors"

	"github.com/rolieup/golie/pkg/rolie"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "clone URI",
	Short: "Clone remote rolie resource with all the referenced files",
	Long:  `Reads the rolie resource from the given URI and traverses any referenced documents. Everything is fetched locally with the same directory structure.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Please provide URI to remote ROLIE resource as command-line argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := cmd.Flags().GetString("dir")
		if err != nil {
			return err
		}
		return rolie.Clone(args[0], dir)
	},
}

func init() {
	Cmd.Flags().String("dir", "./", "Directory to clone the feed into.")
}
