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
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update name path",
	Short: "Updates a bookmark in the bookmark list",
	Args:  cobra.RangeArgs(2, 2),
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := config.StdConfigPath()
		bms, err := config.Read(path)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}

		b := model.Bookmark{
			Name: args[0],
			Path: args[1],
		}

		if err := bms.Update(b); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		if err := config.Write(path, bms); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	},
}
