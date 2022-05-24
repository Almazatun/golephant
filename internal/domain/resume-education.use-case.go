package usecase

import (
	"time"

	"github.com/Almazatun/golephant/internal/infrastucture/entity"
	"github.com/Almazatun/golephant/pkg/http/presentation/input"
)

type resumeEducationUseCase struct{}

type ResumeEducationUseCase interface {
	Create(
		createResumeEducationInput []input.ResumeEducationInput,
	) (res []entity.ResumeEducation, err error)
	Update(
		resumeDB entity.Resume,
		updateEducation []input.ResumeEducationInput,
	) (res []entity.ResumeEducation, err error)
}

func NewResumeEducationUseCase() ResumeEducationUseCase {
	return &resumeEducationUseCase{}
}

func (uc *resumeEducationUseCase) Create(
	createResumeEducationInput []input.ResumeEducationInput,
) (res []entity.ResumeEducation, err error) {
	if len(createResumeEducationInput) > 0 {
		for _, education := range createResumeEducationInput {
			var createResumeEducation entity.ResumeEducation

			createResumeEducation.City = education.City
			createResumeEducation.DegreePlacement = education.DegreePlacement

			startDate, e := time.Parse(layoutISO, string(education.StartDate))

			if e != nil {
				err = e
				break
			}

			createResumeEducation.StartDate = startDate

			endDate, e := time.Parse(layoutISO, string(education.EndDate))

			if e != nil {
				err = e
				break
			}

			createResumeEducation.EndDate = endDate

			res = append(res, createResumeEducation)
		}
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeEducationUseCase) Update(
	resumeDB entity.Resume,
	updateEducation []input.ResumeEducationInput,
) (res []entity.ResumeEducation, err error) {
	for _, education := range updateEducation {
		for _, educationDB := range resumeDB.Education {
			if education.ResumeEducationID == educationDB.ResumeEducationID.String() {
				educationDB.City = education.City
				educationDB.DegreePlacement = education.DegreePlacement

				startDate, e := time.Parse(layoutISO, education.StartDate)

				if e != nil {
					err = e
					break
				}

				educationDB.StartDate = startDate

				endDate, e := time.Parse(layoutISO, education.EndDate)

				if e != nil {
					err = e
					break
				}

				educationDB.EndDate = endDate

				res = append(res, educationDB)
			}
		}

		if err != nil {
			break
		}

	}

	if err != nil {
		return nil, err
	}

	return res, nil
}
