package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kedadiannao220/bible_reader/model"
	"strconv"
	"strings"
)

func GetOutCommand() *cobra.Command {
	outCommand := &cobra.Command{
		Use:   "out ",
		Short: "bible out ctl",
		Run:   OutAction,
	}

	outCommand.Flags().StringVarP(&book, "book", "b", "", OutBookUsage)
	outCommand.Flags().IntVarP(&chapter, "chapter", "c", 0, OutBookChapterUsage)
	outCommand.Flags().StringVarP(&verse, "verse", "v", "", OutBookChapterVerseUsage)
	outCommand.Flags().StringVarP(&qword, "query", "q", "", OutBookUsage)

	return outCommand
}

func OutAction(cmd *cobra.Command, args []string) {

	var outs []model.BibleInfo
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

			outs, err = model.FindTextByBookAndChapterAndVerse(model.ConvertBookID(book), chapter, ver)
		}

		if book != "" && chapter != 0 && verse == "" {
			outs, err = model.FindTextByBookAndChapter(model.ConvertBookID(book), chapter)
		}
	} else {
		if book != "" && chapter != 0 {
			outs, err = model.FindFuzzyBookChapterText(model.ConvertBookID(book), chapter, fmt.Sprintf("%%%s%%", qword))
		}

		if book != "" && chapter == 0 {
			outs, err = model.FindFuzzyBookText(model.ConvertBookID(book), fmt.Sprintf("%%%s%%", qword))
		}

		if book == "" {
			outs, err = model.FindFuzzyText(fmt.Sprintf("%%%s%%", qword))
		}
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, book := range outs {
		fmt.Println(fmt.Sprintf("%s %d:%d  %s", book.Name, book.Chapter, book.Verse, book.Text))
	}
}
