package model

type (
	BibleBook struct {
		ID     int    `gorm:"column:id"`
		Litter string `gorm:"column:litter"`
		Name   string `gorm:"column:name"`
		Count  int    `gorm:"column:count"`
	}

	BibleText struct {
		ID        int    `gorm:"column:id"`
		BookIndex int    `gorm:"column:orig_book_index"`
		Chapter   int    `gorm:"column:orig_chapter"`
		Verse     int    `gorm:"column:orig_verse"`
		Text      string `gorm:"column:text"`
	}

	BibleInfo struct {
		TextID  string `gorm:"column:id" json:"id"`
		Name    string `gorm:"column:name" json:"name"`
		Litter  string `gorm:"column:litter" json:"litter"`
		Chapter int    `gorm:"column:chapter" json:"chapter"`
		Verse   int    `gorm:"column:verse" json:"verse"`
		Text    string `gorm:"column:text" json:"text"`
	}
)

func (BibleBook) TableName() string {
	return "bible_book"
}

func (BibleText) TableName() string {
	return "bible_text"
}
