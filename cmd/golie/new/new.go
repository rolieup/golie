package new

import (
	"errors"
	"fmt"
	"os"

	"github.com/rolieup/golie/pkg/rolie"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "new SCAP-CONTENT-PATH",
	Short: "Generate new Rolie feed based on pre-existing SCAP content.",
	Long:  `Rolie files will be placed to the same directory enabling easy serving by http server software of your choice.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Please provide path to your pre-existing SCAP content as command-line argument")
		}
		path := args[0]
		stat, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("Could not stat '%s': %v", path, err)
		}
		if !stat.IsDir() {
			return fmt.Errorf("Provided path '%s' does not point to a directory", path)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		rootUri, err := cmd.Flags().GetString("root-uri")
		if err != nil {
			return err
		}
		builder := rolie.Builder{
			RootURI:       rootUri,
			DirectoryPath: args[0],
		}
		return builder.New()
	},
}

func init() {
	Cmd.Flags().String("root-uri", "", "URI to the feed itself. Example 'https://acme.org/my_rolie_content/")
}
