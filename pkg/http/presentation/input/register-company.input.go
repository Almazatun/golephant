package input

type RegisterCompanyInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=20,min=7"`
}
