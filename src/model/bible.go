package model

import (
	"strconv"
)

func ConvertBookID(str string) int {
	var book BibleBook

	ID, err := strconv.Atoi(str)
	if err != nil {
		book, err = FindBookByName(str)
	} else {
		book, err = FindBookByID(ID)
	}

	if err != nil {
		logger.Fatal("fail to query book by ", str, err.Error())
	}

	return book.ID
}
