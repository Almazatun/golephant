package input

type RegisterUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=20,min=7"`
	Age      string `json:"age"`
	Mobile   string `json:"mobile"`
	Status   string `json:"status"`
	City     string `json:"city"`
}
