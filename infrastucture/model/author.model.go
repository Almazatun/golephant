package model

type Author struct {
	Id    int    `json:"id" gorm:"primary_key";"AUTO_INCREMENT"`
	Name  string `json:"name" gorm:"type:varchar(100)"`
	Email string `json:"email" gorm:"type:varchar(100)"`
	Books []Book `json:"books" gorm:"foreig_key:BookID`
}
