package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func GetBookCommand() *cobra.Command {
	bookCommand := &cobra.Command{
		Use:   "book",
		Short: "bible book ctl",
		Run:   BookAction,
	}

	bookCommand.Flags().BoolVarP(&all, "list", "l", false, BookListUsage)
	return bookCommand
}

func BookAction(cmd *cobra.Command, args []string) {
	if !all {
		return
	}

	fmt.Println("bible_reader book -l")
}
