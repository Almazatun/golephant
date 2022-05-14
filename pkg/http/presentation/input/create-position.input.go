package input

type CreatePositionInput struct {
	Description      string   `json:"description" validate:"required"`
	Requirements     []string `json:"requirements" validate:"required"`
	Responsibilities []string `json:"responsibilities" validate:"required"`
	PositionType     string   `json:"position_type" validate:"required"`
	Salary           int      `json:"salary"`
}
