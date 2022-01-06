package model

type Book struct {
	Id     int    `json:"id gorm:"primar_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Person Person
}
