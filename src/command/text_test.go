package command

import (
	"os"
	"testing"

	"github.com/olekukonko/tablewriter"
)

func TestRenderChineseWord(t *testing.T) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"编号", "简称", "章：节", "内容"})
	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor})
	// 自动分隔不支持中文
	table.Append([]string{"1", "2", "3:1", AddSpace("中华人民共和国70周年大庆！中华人民共和国70周年大庆！中华人民共和国70周年大庆！中华人民共和国70周年大庆！中华人民共和国70周年大庆！中华人民共和国70周年大庆！中华人民共和国70周年大庆！中华人民共和国70周年大庆！中华人民共和国70周年大庆！中华人民共和国70周年大庆！中华人民共和国70周年大庆！中华人民共和国70周年大庆！")})
	//table.Append([]string{"1", "2", "3:1", "hello world!hello world!hello world!hello world!hello world!hello world!hello world!hello world!hello world!hello world!hello world!hello world!hello world!hello world!hello world!hello world!hello world!hello world!hello world!"})
	table.Render()
}
