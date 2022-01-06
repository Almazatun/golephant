package model

type Book struct {
	Id     int    `json:"id" gorm:"primar_key";"AUTO_INCREMENT"`
	Title  string `json:"title" gorm:"type:varchar(100)"`
	Author Author `json:"author" gorm: "foreing_key: AuthorID`
}
