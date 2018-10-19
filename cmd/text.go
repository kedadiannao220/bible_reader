package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"kedadiannao220/bible_reader/model"
	"os"
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
	textCommand.Flags().StringVarP(&qword, "query", "q", "", TextBookUsage)

	return textCommand
}

func TextAction(cmd *cobra.Command, args []string) {

	var texts []model.BibleInfo
	var err error

	if qword == "" {
		if book != "" && chapter != 0 && verse != 0 {
			texts, err = model.FindTextByBookAndChapterAndVerse(model.ConvertBookID(book), chapter, verse)
		}

		if book != "" && chapter != 0 && verse == 0 {
			texts, err = model.FindTextByBookAndChapter(model.ConvertBookID(book), chapter)
		}
	} else {
		if book != "" && chapter != 0 {
			texts, err = model.FindFuzzyBookChapterText(model.ConvertBookID(book), chapter, fmt.Sprintf("%%%s%%", qword))
		}

		if book != "" && chapter == 0 {
			texts, err = model.FindFuzzyBookText(model.ConvertBookID(book), fmt.Sprintf("%%%s%%", qword))
		}

		if book == "" {
			texts, err = model.FindFuzzyText(fmt.Sprintf("%%%s%%", qword))
		}
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"书卷", "简称", "章", "节", "内容"})
	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor})
	for _, book := range texts {
		val := []string{book.Name, book.Litter, fmt.Sprintf("%d", book.Chapter), fmt.Sprintf("%d", book.Verse), book.Text}
		table.Append(val)
	}
	table.Render()
}
