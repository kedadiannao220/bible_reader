package model

var allBooks []BibleBook

func ConvertBookID(str string) int {
	book, err := FindBookByName(str)
	if err != nil {
		logger.Fatal("fail to query book by", str, err.Error())
	}
	return book.ID
}
