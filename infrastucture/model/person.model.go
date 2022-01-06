package model

type Person struct {
	Id    int    `json:"id gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Books []Book
}
