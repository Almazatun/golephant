package input

type UpdateCitizenshipResumeInput struct {
	City          string `json:"city" validate:"required"`
	SubwayStation string `json:"subway_station" validate:"required"`
}
