package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kedadiannao220/bible_reader/model"
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

	if qword == "" {
		if book != "" && chapter != 0 && verse != 0 {
			texts, _ := model.FindTextByBookAndChapterAndVerse(model.ConvertBookID(book), chapter, verse)
			fmt.Println(texts)
		}

		if book != "" && chapter != 0 && verse == 0 {
			texts, _ := model.FindTextByBookAndChapter(model.ConvertBookID(book), chapter)
			fmt.Println(texts)
		}
	} else {
		if book != "" && chapter != 0 {
			texts, _ := model.FindFuzzyBookChapterText(model.ConvertBookID(book), chapter, fmt.Sprintf("%%%s%%", qword))
			fmt.Println(texts)
		}

		if book != "" && chapter == 0 {
			texts, _ := model.FindFuzzyBookText(model.ConvertBookID(book), fmt.Sprintf("%%%s%%", qword))
			fmt.Println(texts)
		}

		if book == "" {
			texts, _ := model.FindFuzzyText(fmt.Sprintf("%%%s%%", qword))
			fmt.Println(texts)
		}
	}

}
