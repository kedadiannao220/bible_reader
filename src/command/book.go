package command

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"


	"model"
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

	books, err := model.FindBookList()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"序号", "书卷", "简称", "章数"})
	for _, book := range books {
		val := []string{fmt.Sprintf("%d", book.ID), book.Name, book.Litter, fmt.Sprintf("%d", book.Count)}
		table.Append(val)
	}
	table.Render()
}
