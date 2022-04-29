package input

type UpdateDesiredPositionResumeInput struct {
	DesiredPosition string `json:"desired_position" validate:"required,max=100,min=1"`
	Specialization  string `json:"specialization" validate:"required,min=7"`
	WorkMode        string `json:"work_mode" validate:"required,min=7"`
	Status          string `json:"status" validate:"required"`
}
