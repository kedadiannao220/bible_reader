package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var book, qword string
var chapter, verse int
var all bool

const (
	BookListUsage = "bible_reader book -l"
	TextBookUsage = "bible_reader text -b 创"
)

var rootCmd = &cobra.Command{
	Use:   "bible_read",
	Short: "read bible at cmd ",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}


func Execute() {
	rootCmd.AddCommand(GetBookCommand())
	rootCmd.AddCommand(GetTextCommand())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
