package usecase

import (
	"errors"
	"time"

	repository "github.com/Almazatun/golephant/internal/infrastucture"
	"github.com/Almazatun/golephant/internal/infrastucture/entity"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	"github.com/Almazatun/golephant/pkg/http/presentation/input"
	"gopkg.in/go-playground/validator.v9"
)

type resumeUseCase struct {
	resumeRepo           repository.ResumeRepo
	userRepo             repository.UserRepo
	resumeEducationRepo  repository.ResumeEducationRepo
	resumeExperienceRepo repository.ResumeExperienceRepo
	// use cases
	resumeEducationUseCase  ResumeEducationUseCase
	resumeExperienceUseCase ResumeExperienceUseCase
}

type ResumeUseCase interface {
	Create(
		userId string,
		createResumeInput input.CreateResumeInput,
	) (createResume *entity.Resume, err error)
	UpdateBasicInfo(
		userId, resumeId string,
		updateBasicInfoResumeInput input.UpdateBasicInfoResume,
	) (updateResume *entity.Resume, err error)
	UpdateAboutMe(
		userId, resumeId string,
		updateAboutMeResumeInput input.UpdateAboutMeResumeInput,
	) (updateResume *entity.Resume, err error)
	UpdateCitizenship(
		userId, resumeId string,
		updateCitizenshipResumInput input.UpdateCitizenshipResumeInput,
	) (updateResume *entity.Resume, err error)
	UpdateTags(
		userId, resumeId string,
		updateTagsResume input.UdateTagsResumeInput,
	) (updateResume *entity.Resume, err error)
	UpdateEducation(
		userId, resumeId string,
		updateResumeEducationResumeInput input.UpdateResumeEducationInput,
	) (updateResume *entity.Resume, err error)
	UpdateExperience(
		userId, resumeId string,
		updateResumeExperienceResumeInput input.UpdateResumeExperienceInput,
	) (updateResume *entity.Resume, err error)
	UpdateDesiredPosition(
		userId, resumeId string,
		updateDesiredPositionResumeInput input.UpdateDesiredPositionResumeInput,
	) (updateResume *entity.Resume, err error)
	Delete(resumeId string) (str string, err error)
	DeleteExperience(resumeId, experienceId string) (str string, err error)
	DeleteEducation(resumeId, educationId string) (str string, err error)
	fillResumeToCreate(
		createResumeInput input.CreateResumeInput,
	) (resume entity.Resume)
	getCreateAndUpdateResumeExperience(
		updateResumeExperienceInput input.UpdateResumeExperienceInput,
		resumeExperienceChannel chan input.ResumeInput[input.ResumeExperienceInput],
	)
	getCreateAndUpdateResumeEducation(
		updateResumeEducationInput input.UpdateResumeEducationInput,
		resumeEducationChannel chan input.ResumeInput[input.ResumeEducationInput],
	)
	setGender(gender string) string
}

func NewResumeUseCase(
	resumeRepo repository.ResumeRepo,
	userRepo repository.UserRepo,
	resumeEducationRepo repository.ResumeEducationRepo,
	resumeExperienceRepo repository.ResumeExperienceRepo,
	// use cases
	resumeEducationUseCase ResumeEducationUseCase,
	resumeExperienceUseCase ResumeExperienceUseCase,
) ResumeUseCase {
	return &resumeUseCase{
		resumeRepo:           resumeRepo,
		userRepo:             userRepo,
		resumeEducationRepo:  resumeEducationRepo,
		resumeExperienceRepo: resumeExperienceRepo,
		// use cases
		resumeEducationUseCase:  resumeEducationUseCase,
		resumeExperienceUseCase: resumeExperienceUseCase,
	}
}

const layoutISO = "2006-01-02"

func (uc *resumeUseCase) Create(userId string, createResumeInput input.CreateResumeInput) (createResume *entity.Resume, err error) {
	// Validate create resume input
	v := validator.New()
	e := v.Struct(createResumeInput)

	if e != nil {
		return nil, e
	}

	user, err := uc.userRepo.GetById(userId)

	if err != nil {
		return nil, err
	}

	resume := uc.fillResumeToCreate(createResumeInput)

	userDB := *user

	resume.User = userDB
	resume.UserID = userDB.UserID

	// Experience in resume
	var experience []entity.ResumeExperience
	if createResumeInput.IsHaveExperience &&
		len(createResumeInput.Experience) > 0 {

		createExperience, err := uc.
			resumeExperienceUseCase.
			Create(createResumeInput.Experience)

		if err != nil {
			return nil, err
		}
		experience = createExperience

	}

	resume.Experience = experience

	res, err := uc.resumeRepo.Create(resume)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeUseCase) UpdateBasicInfo(
	userId, resumeId string,
	updateBasicInfoResumeInput input.UpdateBasicInfoResume,
) (updateResume *entity.Resume, err error) {

	resumeDB, err := uc.resumeRepo.GetById(resumeId)

	if err != nil {
		return nil, err
	}

	if resumeDB.UserID.String() != userId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	if updateBasicInfoResumeInput.FirstName != "" {
		resumeDB.FirstName = updateBasicInfoResumeInput.FirstName
	}

	if updateBasicInfoResumeInput.LastName != "" {
		resumeDB.LastName = updateBasicInfoResumeInput.LastName
	}

	if updateBasicInfoResumeInput.Gender != "" {
		resumeDB.Gender = updateBasicInfoResumeInput.Gender
	}

	if updateBasicInfoResumeInput.DateOfBirght != "" {

		date_of_birght, err := time.Parse(layoutISO, string(updateBasicInfoResumeInput.DateOfBirght))

		if err != nil {
			return nil, err

		}

		resumeDB.DateOfBirght = date_of_birght

	}

	now := time.Now()
	resumeDB.CreationTime = now
	resumeDB.UpdateTime = now

	res, err := uc.resumeRepo.Save(*resumeDB)

	if err != nil {
		return nil, err

	}

	return res, nil
}

func (uc *resumeUseCase) UpdateAboutMe(
	userId, resumeId string,
	updateAboutMeResumeInput input.UpdateAboutMeResumeInput,
) (updateAboutMeResume *entity.Resume, err error) {
	resumeDB, err := uc.resumeRepo.GetById(resumeId)

	if err != nil {
		return nil, err
	}

	if resumeDB.UserID.String() != userId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	if updateAboutMeResumeInput.AboutMe != "" {
		resumeDB.About = updateAboutMeResumeInput.AboutMe
	}

	now := time.Now()
	resumeDB.CreationTime = now
	resumeDB.UpdateTime = now

	res, err := uc.resumeRepo.Save(*resumeDB)

	if err != nil {
		return nil, err

	}

	return res, nil

}

func (uc *resumeUseCase) UpdateDesiredPosition(
	userId, resumeId string,
	updateDesiredPositionResumeInput input.UpdateDesiredPositionResumeInput,
) (updateAboutMeResume *entity.Resume, err error) {
	resumeDB, err := uc.resumeRepo.GetById(resumeId)

	if err != nil {
		return nil, err
	}

	if resumeDB.UserID.String() != userId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	if updateDesiredPositionResumeInput.DesiredPosition != "" {
		resumeDB.DesiredPosition = updateDesiredPositionResumeInput.DesiredPosition
	}

	if updateDesiredPositionResumeInput.Specialization != "" {
		resumeDB.Specialization = updateDesiredPositionResumeInput.Specialization
	}

	if updateDesiredPositionResumeInput.WorkMode != "" {
		resumeDB.WorkMode = updateDesiredPositionResumeInput.WorkMode
	}

	now := time.Now()
	resumeDB.CreationTime = now
	resumeDB.UpdateTime = now

	res, err := uc.resumeRepo.Save(*resumeDB)

	if err != nil {
		return nil, err

	}

	return res, nil

}

func (uc *resumeUseCase) UpdateCitizenship(
	userId, resumeId string,
	updateCitizenshipResumInput input.UpdateCitizenshipResumeInput,
) (updateCitizenshipResume *entity.Resume, err error) {
	resumeDB, err := uc.resumeRepo.GetById(resumeId)

	if err != nil {
		return nil, err
	}

	if resumeDB.UserID.String() != userId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	if updateCitizenshipResumInput.City != "" {
		resumeDB.Citizenship = updateCitizenshipResumInput.City
	}

	if updateCitizenshipResumInput.SubwayStation != "" {
		resumeDB.SubwayStation = updateCitizenshipResumInput.SubwayStation
	}

	now := time.Now()
	resumeDB.CreationTime = now
	resumeDB.UpdateTime = now

	res, err := uc.resumeRepo.Save(*resumeDB)

	if err != nil {
		return nil, err

	}

	return res, nil
}

func (uc *resumeUseCase) UpdateTags(
	userId, resumeId string,
	updateTagsResume input.UdateTagsResumeInput,
) (updateTagsResum *entity.Resume, err error) {
	resumeDB, err := uc.resumeRepo.GetById(resumeId)

	if err != nil {
		return nil, err
	}

	if resumeDB.UserID.String() != userId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	if len(updateTagsResume.Tags) > 0 {
		resumeDB.Tags = updateTagsResume.Tags
	}

	now := time.Now()
	resumeDB.CreationTime = now
	resumeDB.UpdateTime = now

	res, err := uc.resumeRepo.Save(*resumeDB)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeUseCase) UpdateEducation(
	userId, resumeId string,
	updateEducationInput input.UpdateResumeEducationInput,
) (updateResume *entity.Resume, err error) {
	resumeDB, err := uc.resumeRepo.GetById(resumeId)

	if err != nil {
		return nil, err
	}

	if resumeDB.UserID.String() != userId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	if len(updateEducationInput.Education) > 0 {
		var createResumeEducationInput []input.ResumeEducationInput
		var updateResumeEducationInput []input.ResumeEducationInput

		resumeEducationСhannel := make(chan input.ResumeInput[input.ResumeEducationInput], 2)

		go uc.getCreateAndUpdateResumeEducation(
			updateEducationInput,
			resumeEducationСhannel,
		)

		for userEducation := range resumeEducationСhannel {
			if userEducation.IsUpdate {
				updateResumeEducationInput = append(
					updateResumeEducationInput,
					userEducation.Data,
				)
			}
			createResumeEducationInput = append(
				createResumeEducationInput,
				userEducation.Data,
			)
		}

		// createUserEducationDB, err := uc.createUserEducation(createUserEducationInput)
		createResumeEducationDB, err := uc.
			resumeEducationUseCase.
			Create(createResumeEducationInput)

		if err != nil {
			return nil, err
		}

		updateResumeEducationDB, err := uc.
			resumeEducationUseCase.
			Update(*resumeDB, updateResumeEducationInput)

		if err != nil {
			return nil, err
		}

		resumeDB.Education = nil

		resumeDB.Education = append(resumeDB.Education, createResumeEducationDB...)
		resumeDB.Education = append(resumeDB.Education, updateResumeEducationDB...)
	}

	res, err := uc.resumeRepo.Save(*resumeDB)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeUseCase) UpdateExperience(
	userId, resumeId string,
	updateExperienceInput input.UpdateResumeExperienceInput,
) (updateUserEducationResum *entity.Resume, err error) {
	resumeDB, err := uc.resumeRepo.GetById(resumeId)

	if err != nil {
		return nil, err
	}

	if resumeDB.UserID.String() != userId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	if len(updateExperienceInput.Experience) > 0 {
		var createResumeExperienceInput []input.ResumeExperienceInput
		var updateResumeExperienceInput []input.ResumeExperienceInput

		experienceChannel := make(chan input.ResumeInput[input.ResumeExperienceInput], 2)

		go uc.getCreateAndUpdateResumeExperience(
			updateExperienceInput,
			experienceChannel,
		)

		for experience := range experienceChannel {
			if experience.IsUpdate {
				updateResumeExperienceInput = append(
					updateResumeExperienceInput,
					experience.Data,
				)
			}
			createResumeExperienceInput = append(
				createResumeExperienceInput,
				experience.Data,
			)
		}

		createResumeEducationDB, err := uc.
			resumeExperienceUseCase.
			Create(createResumeExperienceInput)

		if err != nil {
			return nil, err
		}

		updateResumeExperienceDB, err := uc.
			resumeExperienceUseCase.
			Update(resumeDB, updateResumeExperienceInput)

		if err != nil {
			return nil, err
		}

		resumeDB.Experience = nil

		resumeDB.Experience = append(resumeDB.Experience, createResumeEducationDB...)
		resumeDB.Experience = append(resumeDB.Experience, updateResumeExperienceDB...)
	}

	res, err := uc.resumeRepo.Save(*resumeDB)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeUseCase) Delete(resumeId string) (str string, err error) {
	res, err := uc.resumeRepo.DeleteById(resumeId)

	if err != nil {
		return "", err
	}

	return res, nil
}

func (uc *resumeUseCase) DeleteExperience(resumeId, experienceId string) (str string, err error) {

	resumeExperienceDB, err := uc.resumeExperienceRepo.GetById(experienceId)

	if err != nil {
		return "", err
	}

	if resumeExperienceDB.ResumeID.String() != resumeId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return "", newErr
	}

	res, err := uc.resumeExperienceRepo.DeleteById(experienceId)

	if err != nil {
		return "", err
	}

	return res, nil
}

func (uc *resumeUseCase) DeleteEducation(resumeId, educationId string) (str string, err error) {
	resumeEducationDB, err := uc.resumeEducationRepo.GetById(educationId)

	if err != nil {
		return "", err
	}

	if resumeEducationDB.ResumeID.String() != resumeId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return "", newErr
	}

	res, err := uc.resumeEducationRepo.DeleteById(educationId)

	if err != nil {
		return "", err
	}

	return res, nil
}

func (uc *resumeUseCase) fillResumeToCreate(createResumeInput input.CreateResumeInput) (resume entity.Resume) {

	if createResumeInput.FirstName != "" {
		resume.FirstName = createResumeInput.FirstName
	}

	if createResumeInput.LastName != "" {
		resume.LastName = createResumeInput.LastName
	}

	resume.Gender = uc.setGender(createResumeInput.Gender)

	now := time.Now()
	resume.CreationTime = now
	resume.UpdateTime = now

	return resume
}

func (uc *resumeUseCase) getCreateAndUpdateResumeExperience(
	updateResumeExperienceInput input.UpdateResumeExperienceInput,
	resumeExperienceChannel chan input.ResumeInput[input.ResumeExperienceInput],
) {
	for _, experienceData := range updateResumeExperienceInput.Experience {
		if experienceData.ResumeExperienceID == "" {
			// channel out put experience data to create
			createExperience := input.ResumeInput[input.ResumeExperienceInput]{
				Data:     experienceData,
				IsUpdate: false,
			}
			resumeExperienceChannel <- createExperience
		} else {
			// channel out put experience data to update
			updateExperience := input.ResumeInput[input.ResumeExperienceInput]{
				Data:     experienceData,
				IsUpdate: true,
			}
			resumeExperienceChannel <- updateExperience
		}
	}
	// Close channel
	close(resumeExperienceChannel)
}

func (uc *resumeUseCase) getCreateAndUpdateResumeEducation(
	updateResumeEducationInput input.UpdateResumeEducationInput,
	resumeEducationChannel chan input.ResumeInput[input.ResumeEducationInput],
) {
	for _, educationData := range updateResumeEducationInput.Education {
		if educationData.ResumeEducationID == "" {
			// channel out put education data to create
			createEducation := input.ResumeInput[input.ResumeEducationInput]{
				Data:     educationData,
				IsUpdate: false,
			}
			resumeEducationChannel <- createEducation
		} else {
			// channel out put education data to udpate
			updateEducation := input.ResumeInput[input.ResumeEducationInput]{
				Data:     educationData,
				IsUpdate: true,
			}
			resumeEducationChannel <- updateEducation
		}
	}

	// Close channel
	close(resumeEducationChannel)
}

func (uc *resumeUseCase) setGender(gender string) string {
	switch {
	case gender == "MALE":
		return gender
	case gender == "FEMALE":
		return gender
	default:
		return ""
	}
}
