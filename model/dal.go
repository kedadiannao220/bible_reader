package model

func FindBookList() (books []BibleBook, err error) {
	err = globalDB.Table(BibleBook{}.TableName()).Find(&books).Error
	return
}

func FindBookByName(name string) (book BibleBook, err error) {
	err = globalDB.Table(BibleBook{}.TableName()).Where("litter = ? or name = ?", name, name).Find(&book).Error
	return
}
func FindBookByID(id int) (book BibleBook, err error) {
	err = globalDB.Table(BibleBook{}.TableName()).Where("id  = ?", id).Find(&book).Error
	return
}

func FindTextByBookAndChapter(name, chapter int) (bibleInfo []BibleInfo, err error) {
	err = globalDB.Table("bible_text text").Select("book.name, book.litter,text.orig_chapter as chapter,text.orig_verse as verse,text.text,text.id").
		Joins("left join bible_book  book on book.id = text.orig_book_index ").
		Where("text.orig_chapter =  ? and book.id = ?", chapter, name).
		Find(&bibleInfo).Error
	return
}
func FindTextByBookAndChapterAndVerse(name, chapter int, verse []int) (bibleInfo []BibleInfo, err error) {
	err = globalDB.Table("bible_text text").Select("book.name, book.litter,text.orig_chapter as chapter,text.orig_verse as verse,text.text,text.id").
		Joins("left join bible_book  book on book.id = text.orig_book_index ").
		Where("text.orig_chapter =  ? and text.orig_verse in (?) and book.id = ?", chapter, verse, name).
		Find(&bibleInfo).Error
	return
}

// FindFuzzyText  fuzzy query text
func FindFuzzyText(str string) (bibleInfo []BibleInfo, err error) {
	err = globalDB.Table("bible_text text").Select("book.name, book.litter,text.orig_chapter as chapter,text.orig_verse as verse,text.text,text.id").
		Joins("left join  bible_book  book on book.id = text.orig_book_index ").
		Where("text.text like ?", str).
		Find(&bibleInfo).Error
	return
}
func FindFuzzyBookText(book int, str string) (bibleInfo []BibleInfo, err error) {
	err = globalDB.Table("bible_text text").Select("book.name, book.litter,text.orig_chapter as chapter,text.orig_verse as verse,text.text,text.id").
		Joins("left join  bible_book  book on book.id = text.orig_book_index ").
		Where("book.id = ? and  text.text like ? ", book, str).
		Find(&bibleInfo).Error
	return
}
func FindFuzzyBookChapterText(book, chapter int, str string) (bibleInfo []BibleInfo, err error) {
	err = globalDB.Table("bible_text text").Select("book.name, book.litter,text.orig_chapter as chapter,text.orig_verse as verse,text.text,text.id").
		Joins("left join  bible_book  book on book.id = text.orig_book_index ").
		Where(" book.id = ?  and text.orig_chapter = ? and  text.text like ?", book, chapter, str).
		Find(&bibleInfo).Error
	return
}
