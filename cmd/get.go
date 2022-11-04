package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mrjerz/bookmarks/config"
	"github.com/mrjerz/bookmarks/model"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "retrieve path for a given bookmark",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := config.StdConfigPath()
		bms, err := config.Read(path)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}

		var f model.FirstExactMatchFinder
		bm, err := bms.Get(args[0], f)
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}

		fmt.Fprintf(os.Stdout, "%s", bm.Path)
	},
}
