package input

type CreateResumeInput struct {
	FirstName        string `json:"first_name" validate:"required"`
	LastName         string `json:"last_name" validate:"required"`
	DateOfBirght     string `json:"date_of_birght" validate:"required"`
	Gender           string `json:"gender" validate:"required"`
	IsHaveExperience bool   `json:"is_have_experience"`
	UserExperience  []UserExperienceInput
}
