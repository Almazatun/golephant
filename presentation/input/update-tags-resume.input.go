package input

type UdateTagsResumeInput struct {
	Tags []string `json:"tags" validate:"required"`
}
