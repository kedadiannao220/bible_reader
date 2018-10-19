package main

import (
	"kedadiannao220/bible_reader/cmd"
	"kedadiannao220/bible_reader/model"
)

func main() {
	model.Init()
	cmd.Execute()
}
