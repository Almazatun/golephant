package input

type CreateResumeInput struct {
	Title          string   `json:"title" validate:"required,max=100,min=1"`
	Specialization string   `json:"specialization" validate:"required,min=7"`
	WorkMode       string   `json:"work_mode" validate:"required,min=7" gorm:"type:text;default:0; not NULL"`
	About          string   `json:"about"`
	Tags           []string `json:"tags" validate:"required"`
	UserEducation  []userEducationInput
	UserExperience []userExperienceInput
}

// Local types
type userEducationInput struct {
	StartDate       string `json:"start_date" validate:"required"`
	EndDate         string `json:"end_date" validate:"required"`
	DegreePlacement string `json:"degree_placement" validate:"required"`
	City            string `json:"city" validate:"required"`
}

type userExperienceInput struct {
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
	Position    string `json:"position" validate:"required"`
	City        string `json:"city" validate:"required"`
}
