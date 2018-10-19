package main

import (
	"testing"
	"fmt"
)

func TestFindBookList(t *testing.T) {
	Init()
	output(FindBookList())
}

func TestFindTextByBookAndChapter(t *testing.T) {
	Init()
	output(FindTextByBookAndChapter(ConvertBookID("创"), 1))
}

func TestFindTextByBookAndChapterAndVerse(t *testing.T) {
	Init()
	output(FindTextByBookAndChapterAndVerse(ConvertBookID("创世记"), 1, 1))
}

func TestFindFuzzyText(t *testing.T) {
	Init()
	output(FindFuzzyText("%神爱世人%"))
}

func TestFindFuzzyBookText(t *testing.T) {
	Init()
	output(FindFuzzyBookText(1, "%耶和华%"))
}

func TestFindFuzzyBookChapterText(t *testing.T) {
	Init()
	output(FindFuzzyBookChapterText(1, 1, "%神%"))
}

func output(body interface{}, err error) {
	if err != nil {
		fmt.Printf("errors is %s \n", err.Error())
		return
	}
	fmt.Printf("body is %s", body)
}
