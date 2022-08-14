package cmd

import (
	"fmt"
	"os"

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
		file, err := os.OpenFile("/Users/jerzy/.bookmarks", os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()

		bms, err = model.Load(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		var f model.FirstExactMatchFinder
		bm, err := bms.Get(args[0], f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "%s", bm.Path)
	},
}
