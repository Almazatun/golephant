package input

type UserEducationInput struct {
	UserEducationID string `json:"user_education_id"`
	StartDate       string `json:"start_date" validate:"required"`
	EndDate         string `json:"end_date" validate:"required"`
	DegreePlacement string `json:"degree_placement" validate:"required"`
	City            string `json:"city" validate:"required"`
}