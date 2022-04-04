package input

type UpdateCitizenshipResumeInput struct {
	City string `json:"city" validate:"required"`
}
