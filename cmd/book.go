package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kedadiannao220/bible_reader/model"
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

	books, _ := model.FindBookList()
	fmt.Println(books)
}
