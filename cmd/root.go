package cmd

import (
	"fmt"
	"os"

	"github.com/mrjerz/bookmarks/model"
	"github.com/spf13/cobra"
)

var bms model.Bookmarks

var rootCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "Bookmark manages bookmarks in the cli",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
