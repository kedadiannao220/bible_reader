package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func GetTextCommand() *cobra.Command {
	textCommand := &cobra.Command{
		Use:   "text ",
		Short: "bible text ctl",
		Run:   TextAction,
	}

	textCommand.Flags().StringVarP(&book, "book", "b", "", TextBookUsage)
	textCommand.Flags().IntVarP(&chapter, "chapter", "c", 0, TextBookUsage)
	textCommand.Flags().IntVarP(&verse, "verse", "v", 0, TextBookUsage)
	textCommand.Flags().StringVarP(&qword, "like", "l", "", TextBookUsage)

	return textCommand
}

func TextAction(cmd *cobra.Command, args []string) {

	if qword != "" {
		fmt.Println("bible_reader text -l 创")
	}
}
