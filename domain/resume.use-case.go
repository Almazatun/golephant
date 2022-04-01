package usecase

import (
	"errors"
	"time"

	error_message "github.com/Almazatun/golephant/common/error-message"
	repository "github.com/Almazatun/golephant/infrastucture"
	"github.com/Almazatun/golephant/infrastucture/entity"
	"github.com/Almazatun/golephant/presentation/input"
	"gopkg.in/go-playground/validator.v9"
)

type resumeUseCase struct {
	resumeRepo         repository.ResumeRepo
	userRepo           repository.UserRepo
	userEducationRepo  repository.UserEducationRepo
	userExperienceRepo repository.UserExperienceRepo
}

type ResumeUseCase interface {
	CreateResume(userId string, createResumeInput input.CreateResumeInput) (createResume *entity.Resume, err error)
	DeleteResume(resumeId string) (str string, err error)
	DeleteUserExperienceInResume(resumeId string, userExperienceId string) (str string, err error)
	DeleteUserEducationInResume(resumeId string, userEducationId string) (str string, err error)
}

func NewResumeUseCase(
	resumeRepo repository.ResumeRepo,
	userRepo repository.UserRepo,
	userEducationRepo repository.UserEducationRepo,
	userExperienceRepo repository.UserExperienceRepo,
) ResumeUseCase {
	return &resumeUseCase{
		resumeRepo:         resumeRepo,
		userRepo:           userRepo,
		userEducationRepo:  userEducationRepo,
		userExperienceRepo: userExperienceRepo,
	}
}

const layoutISO = "2006-01-02"

func (uc *resumeUseCase) CreateResume(userId string, createResumeInput input.CreateResumeInput) (createResume *entity.Resume, err error) {
	// Validate create resume input
	v := validator.New()
	e := v.Struct(createResumeInput)

	if e != nil {
		return nil, e
	}

	user, err := uc.userRepo.FindById(userId)

	if err != nil {
		return nil, err
	}

	resume := createResumeColums(createResumeInput)

	userDB := *user

	resume.User = userDB
	resume.UserID = userDB.UserID

	// append user experiences in resume
	userExperiences, err := createUserExperiences(createResumeInput)

	if err != nil {
		return nil, err
	}

	resume.UserExperience = userExperiences

	// append user educations in resume
	userEducations, err := createUserEducations(createResumeInput)

	if err != nil {
		return nil, err
	}

	resume.UserEducation = userEducations

	res, err := uc.resumeRepo.Create(resume)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeUseCase) DeleteResume(resumeId string) (str string, err error) {
	res, err := uc.resumeRepo.DeleteById(resumeId)

	if err != nil {
		return "", err
	}

	return res, nil
}

func (uc *resumeUseCase) DeleteUserExperienceInResume(resumeId string, userExperienceId string) (str string, err error) {

	userExperienceDB, err := uc.userExperienceRepo.FindById(userExperienceId)

	if err != nil {
		return "", err
	}

	if userExperienceDB.ResumeID.String() != resumeId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return "", newErr
	}

	res, err := uc.userExperienceRepo.DeleteById(userExperienceId)

	if err != nil {
		return "", err
	}

	return res, nil
}

func (uc *resumeUseCase) DeleteUserEducationInResume(resumeId string, userEducationId string) (str string, err error) {
	userEducationDB, err := uc.userEducationRepo.FindById(userEducationId)

	if err != nil {
		return "", err
	}

	if userEducationDB.ResumeID.String() != resumeId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return "", newErr
	}

	res, err := uc.userEducationRepo.DeleteById(userEducationId)

	if err != nil {
		return "", err
	}

	return res, nil
}

func createResumeColums(createResumeInput input.CreateResumeInput) (resume entity.Resume) {

	if createResumeInput.About != "" {
		resume.About = createResumeInput.About
	}

	if len(createResumeInput.Tags) > 0 {
		resume.Tags = createResumeInput.Tags
	}

	resume.Title = createResumeInput.Title
	resume.Specialization = createResumeInput.Specialization
	resume.WorkMode = createResumeInput.WorkMode

	now := time.Now()
	resume.CreationTime = now
	resume.UpdateTime = now

	return resume
}

func createUserEducations(createResumeInput input.CreateResumeInput) (userEducations []entity.UserEducation, e error) {

	if len(createResumeInput.UserEducation) > 0 {
		for _, userEducation := range createResumeInput.UserEducation {
			var createUserEducation entity.UserEducation

			now := time.Now()
			createUserEducation.CreationTime = now
			createUserEducation.UpdateTime = now

			createUserEducation.City = userEducation.City
			createUserEducation.DegreePlacement = userEducation.DegreePlacement

			startDate, err := time.Parse(layoutISO, string(userEducation.StartDate))

			if err != nil {
				e = err
				break
			}

			createUserEducation.StartDate = startDate

			endDate, err := time.Parse(layoutISO, string(userEducation.EndDate))

			if err != nil {
				e = err
				break
			}

			createUserEducation.EndDate = endDate

			userEducations = append(userEducations, createUserEducation)
		}
	}

	if e != nil {
		return nil, e
	}

	return userEducations, nil
}

func createUserExperiences(createResumeInput input.CreateResumeInput) (userExperiences []entity.UserExperience, e error) {

	if len(createResumeInput.UserExperience) > 0 {
		for _, userExperience := range createResumeInput.UserExperience {
			var createUserExperience entity.UserExperience

			now := time.Now()
			createUserExperience.CreationTime = now
			createUserExperience.UpdateTime = now

			createUserExperience.City = userExperience.City
			createUserExperience.CompanyName = userExperience.CompanyName
			createUserExperience.Position = userExperience.Position

			startDate, err := time.Parse(layoutISO, userExperience.StartDate)

			if err != nil {
				e = err
				break
			}

			createUserExperience.StartDate = startDate

			endDate, err := time.Parse(layoutISO, userExperience.EndDate)

			if err != nil {
				e = err
				break
			}

			createUserExperience.EndDate = endDate

			userExperiences = append(userExperiences, createUserExperience)
		}
	}

	if e != nil {
		return nil, e
	}

	return userExperiences, nil
}
