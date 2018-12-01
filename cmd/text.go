package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"kedadiannao220/bible_reader/model"
	"os"
	"strconv"
	"strings"
)

func GetTextCommand() *cobra.Command {
	textCommand := &cobra.Command{
		Use:   "text ",
		Short: "bible text ctl",
		Run:   TextAction,
	}

	textCommand.Flags().StringVarP(&book, "book", "b", "", TextBookUsage)
	textCommand.Flags().IntVarP(&chapter, "chapter", "c", 0, TextBookChapterUsage)
	textCommand.Flags().StringVarP(&verse, "verse", "v", "", TextBookChapterVerseUsage)
	textCommand.Flags().StringVarP(&qword, "query", "q", "", TextBookUsage)

	return textCommand
}

func TextAction(cmd *cobra.Command, args []string) {

	var texts []model.BibleInfo
	var err error

	if qword == "" {
		if book != "" && chapter != 0 && verse != "" {
			verse = strings.Replace(verse, "，", ",", -1)
			verse = strings.Replace(verse, "_", "-", -1)
			verse = strings.Replace(verse, "=", "-", -1)
			verse = strings.Replace(verse, "——", "-", -1)

			var ver []int
			if strings.Contains(verse, "-") {
				verstr := strings.Split(verse, "-")
				start, err := strconv.Atoi(verstr[0])
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				end, err := strconv.Atoi(verstr[1])
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				for i := start; i <= end; i++ {
					ver = append(ver, i)
				}
			}
			if strings.Contains(verse, ",") {
				for _, str := range strings.Split(verse, ",") {
					start, err := strconv.Atoi(str)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
					ver = append(ver, start)
				}
			}

			texts, err = model.FindTextByBookAndChapterAndVerse(model.ConvertBookID(book), chapter, ver)
		}

		if book != "" && chapter != 0 && verse == "" {
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
	table.SetHeader([]string{"编号", "简称", "章：节", "内容"})
	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor})
	for _, book := range texts {
		val := []string{book.TextID, book.Litter, fmt.Sprintf("%d:%d", book.Chapter, book.Verse), book.Text}
		table.Append(val)
	}
	table.Render()
}
