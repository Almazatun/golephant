package input

type UpdatePositionInput struct {
	Description  string `json:"description"`
	Salary       int    `json:"salary"`
	PositionType string `json:"position_type"`
}
