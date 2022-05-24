package usecase

import (
	"time"

	"github.com/Almazatun/golephant/internal/infrastucture/entity"
	"github.com/Almazatun/golephant/pkg/http/presentation/input"
)

type resumeExperienceUseCase struct{}

type ResumeExperienceUseCase interface {
	Create(
		createResumeExperienceInput []input.ResumeExperienceInput,
	) (res []entity.ResumeExperience, err error)
	Update(
		resumeDB *entity.Resume,
		updateResumeExperience []input.ResumeExperienceInput,
	) (res []entity.ResumeExperience, err error)
}

func NewResumeExperienceUseCase() ResumeExperienceUseCase {
	return &resumeExperienceUseCase{}
}

func (uc *resumeExperienceUseCase) Create(
	createResumeExperienceInput []input.ResumeExperienceInput,
) (res []entity.ResumeExperience, err error) {
	if len(createResumeExperienceInput) > 0 {
		for _, experience := range createResumeExperienceInput {
			var createExperience entity.ResumeExperience

			createExperience.City = experience.City
			createExperience.CompanyName = experience.CompanyName
			createExperience.Position = experience.Position

			startDate, e := time.Parse(layoutISO, experience.StartDate)

			if e != nil {
				err = e
				break
			}

			createExperience.StartDate = startDate

			endDate, e := time.Parse(layoutISO, experience.EndDate)

			if e != nil {
				err = e
				break
			}

			createExperience.EndDate = endDate

			res = append(res, createExperience)
		}
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeExperienceUseCase) Update(
	resumeDB *entity.Resume,
	updateResumeExperience []input.ResumeExperienceInput,
) (res []entity.ResumeExperience, err error) {
	for _, experience := range updateResumeExperience {
		for _, experienceDB := range resumeDB.Experience {
			if experience.ResumeExperienceID == experienceDB.ResumeExperienceID.String() {
				experienceDB.City = experience.City
				experienceDB.CompanyName = experience.CompanyName
				experienceDB.Position = experience.Position

				startDate, e := time.Parse(layoutISO, experience.StartDate)

				if e != nil {
					err = e
					break
				}

				experienceDB.StartDate = startDate

				endDate, e := time.Parse(layoutISO, experience.EndDate)

				if e != nil {
					err = e
					break
				}

				experienceDB.EndDate = endDate

				res = append(res, experienceDB)
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
