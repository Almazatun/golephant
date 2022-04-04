package input

type UpdateAboutMeResumeInput struct {
	AboutMe string `json:"about_me" validate:"required,max=300,min=1"`
}
