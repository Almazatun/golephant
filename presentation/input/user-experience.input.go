package input

type UserExperienceInput struct {
	UserExperienceID string `json:"user_experience_id"`
	StartDate        string `json:"start_date" validate:"required"`
	EndDate          string `json:"end_date" validate:"required"`
	CompanyName      string `json:"company_name" validate:"required"`
	Position         string `json:"position" validate:"required"`
	City             string `json:"city" validate:"required"`
}
