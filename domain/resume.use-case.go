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
	UpdateBasicInfoResume(
		userId string,
		resumeId string,
		updateBasicInfoResumeInput input.UpdateBasicInfoResume,
	) (updateResume *entity.Resume, err error)
	UpdateAboutMeResume(
		userId string,
		resumeId string,
		updateAboutMeResumeInput input.UpdateAboutMeResumeInput,
	) (updateAboutMeResume *entity.Resume, err error)
	UpdateCitizenshipResume(
		userId string,
		resumeId string,
		updateCitizenshipResumInput input.UpdateCitizenshipResumeInput,
	) (updateCitizenshipResume *entity.Resume, err error)
	UpdateTagsResumeInput(
		userId string,
		resumeId string,
		updateTagsResume input.UdateTagsResumeInput,
	) (updateTagsResum *entity.Resume, err error)
	UpdateUserEducationResume(
		userId string,
		resumeId string,
		updateUserEducationsResumeInput input.UpdateUserEducationsResumeInput,
	) (updateUserEducationsResum *entity.Resume, err error)
	UpdateUserExperiencesResume(
		userId string,
		resumeId string,
		updateUserExperiencesResumeInput input.UpdateUserExperiencesResumeInput,
	) (updateUserEducationsResum *entity.Resume, err error)
	UpdateDesiredPositionResume(
		userId string,
		resumeId string,
		updateDesiredPositionResumeInput input.UpdateDesiredPositionResumeInput,
	) (updateAboutMeResume *entity.Resume, err error)
	DeleteResume(resumeId string) (str string, err error)
	DeleteUserExperienceInResume(resumeId string, userExperienceId string) (str string, err error)
	DeleteUserEducationInResume(resumeId string, userEducationId string) (str string, err error)
}

type UserEducationChannel struct {
	UserEducationInput    input.UserEducationInput
	isUserEducationUpdate bool
}

type UserExperienceChannel struct {
	UserExperienceInput    input.UserExperienceInput
	isUserExperienceUpdate bool
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

	user, err := uc.userRepo.GetById(userId)

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

	resume.UserExperiences = userExperiences

	res, err := uc.resumeRepo.Create(resume)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeUseCase) UpdateBasicInfoResume(
	userId string,
	resumeId string,
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

	res, err := uc.resumeRepo.Update(*resumeDB)

	if err != nil {
		return nil, err

	}

	return res, nil
}

func (uc *resumeUseCase) UpdateAboutMeResume(
	userId string,
	resumeId string,
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

	res, err := uc.resumeRepo.Update(*resumeDB)

	if err != nil {
		return nil, err

	}

	return res, nil

}

func (uc *resumeUseCase) UpdateDesiredPositionResume(
	userId string,
	resumeId string,
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

	res, err := uc.resumeRepo.Update(*resumeDB)

	if err != nil {
		return nil, err

	}

	return res, nil

}

func (uc *resumeUseCase) UpdateCitizenshipResume(
	userId string,
	resumeId string,
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

	res, err := uc.resumeRepo.Update(*resumeDB)

	if err != nil {
		return nil, err

	}

	return res, nil
}

func (uc *resumeUseCase) UpdateTagsResumeInput(
	userId string,
	resumeId string,
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

	res, err := uc.resumeRepo.Update(*resumeDB)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeUseCase) UpdateUserEducationResume(
	userId string,
	resumeId string,
	updateUserEducationsResumeInput input.UpdateUserEducationsResumeInput,
) (updateUserEducationsResum *entity.Resume, err error) {
	resumeDB, err := uc.resumeRepo.GetById(resumeId)

	if err != nil {
		return nil, err
	}

	if resumeDB.UserID.String() != userId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	if len(updateUserEducationsResumeInput.UserEducations) > 0 {
		var createUserEducationsInput []input.UserEducationInput
		var updateUserEducationsInput []input.UserEducationInput

		userEducationСhannel := make(chan UserEducationChannel, 2)

		go getCreateAndUpdateUserEducations(
			updateUserEducationsResumeInput,
			userEducationСhannel,
		)

		for userEducationChannelData := range userEducationСhannel {
			if userEducationChannelData.isUserEducationUpdate {
				updateUserEducationsInput = append(
					updateUserEducationsInput,
					userEducationChannelData.UserEducationInput,
				)
			}
			createUserEducationsInput = append(
				createUserEducationsInput,
				userEducationChannelData.UserEducationInput,
			)
		}

		createUserEducationsDB, err := createUserEducations(createUserEducationsInput)

		if err != nil {
			return nil, err
		}

		updateUserEducationsDB, err := updateUserEducations(*resumeDB, updateUserEducationsInput)

		if err != nil {
			return nil, err
		}

		resumeDB.UserEducations = nil

		resumeDB.UserEducations = append(resumeDB.UserEducations, createUserEducationsDB...)
		resumeDB.UserEducations = append(resumeDB.UserEducations, updateUserEducationsDB...)
	}

	res, err := uc.resumeRepo.Update(*resumeDB)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeUseCase) UpdateUserExperiencesResume(
	userId string,
	resumeId string,
	updateUserExperiencesResumeInput input.UpdateUserExperiencesResumeInput,
) (updateUserEducationsResum *entity.Resume, err error) {
	resumeDB, err := uc.resumeRepo.GetById(resumeId)

	if err != nil {
		return nil, err
	}

	if resumeDB.UserID.String() != userId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	if len(updateUserExperiencesResumeInput.UserExperiences) > 0 {
		var createUserExperiencesInput []input.UserExperienceInput
		var updateUserExperiencesInput []input.UserExperienceInput

		userExperienceChannel := make(chan UserExperienceChannel, 2)

		go getCreateAndUpdateUserExperiences(
			updateUserExperiencesResumeInput,
			userExperienceChannel,
		)

		for userExperienceChannelData := range userExperienceChannel {
			if userExperienceChannelData.isUserExperienceUpdate {
				updateUserExperiencesInput = append(
					updateUserExperiencesInput,
					userExperienceChannelData.UserExperienceInput,
				)
			}
			createUserExperiencesInput = append(
				createUserExperiencesInput,
				userExperienceChannelData.UserExperienceInput,
			)
		}

		createUserEducationsDB, err := createUserExperiencesToUpdate(createUserExperiencesInput)

		if err != nil {
			return nil, err
		}

		updateUserExperiencesDB, err := updateUserExperiences(resumeDB, updateUserExperiencesInput)

		if err != nil {
			return nil, err
		}

		resumeDB.UserExperiences = nil

		resumeDB.UserExperiences = append(resumeDB.UserExperiences, createUserEducationsDB...)
		resumeDB.UserExperiences = append(resumeDB.UserExperiences, updateUserExperiencesDB...)
	}

	res, err := uc.resumeRepo.Update(*resumeDB)

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

	userExperienceDB, err := uc.userExperienceRepo.GetById(userExperienceId)

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
	userEducationDB, err := uc.userEducationRepo.GetById(userEducationId)

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

	if createResumeInput.FirstName != "" {
		resume.FirstName = createResumeInput.FirstName
	}

	if createResumeInput.LastName != "" {
		resume.LastName = createResumeInput.LastName
	}

	resume.Gender = setGender(createResumeInput.Gender)

	now := time.Now()
	resume.CreationTime = now
	resume.UpdateTime = now

	return resume
}

func createUserEducations(
	createUserEducationsResumeInput []input.UserEducationInput,
) (userEducations []entity.UserEducation, err error) {

	if len(createUserEducationsResumeInput) > 0 {
		for _, userEducation := range createUserEducationsResumeInput {
			var createUserEducation entity.UserEducation

			createUserEducation.City = userEducation.City
			createUserEducation.DegreePlacement = userEducation.DegreePlacement

			startDate, e := time.Parse(layoutISO, string(userEducation.StartDate))

			if e != nil {
				err = e
				break
			}

			createUserEducation.StartDate = startDate

			endDate, e := time.Parse(layoutISO, string(userEducation.EndDate))

			if e != nil {
				err = e
				break
			}

			createUserEducation.EndDate = endDate

			userEducations = append(userEducations, createUserEducation)
		}
	}

	if err != nil {
		return nil, err
	}

	return userEducations, nil
}

func updateUserEducations(
	resumeDB entity.Resume,
	updateUserEducations []input.UserEducationInput,
) (res []entity.UserEducation, err error) {
	for _, updateUserEducation := range updateUserEducations {
		for _, userEducationDB := range resumeDB.UserEducations {
			if updateUserEducation.UserEducationID == userEducationDB.UserEducationID.String() {
				userEducationDB.City = updateUserEducation.City
				userEducationDB.DegreePlacement = updateUserEducation.DegreePlacement

				startDate, e := time.Parse(layoutISO, updateUserEducation.StartDate)

				if e != nil {
					err = e
					break
				}

				userEducationDB.StartDate = startDate

				endDate, e := time.Parse(layoutISO, updateUserEducation.EndDate)

				if e != nil {
					err = e
					break
				}

				userEducationDB.EndDate = endDate

				res = append(res, userEducationDB)
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

func updateUserExperiences(
	resumeDB *entity.Resume,
	updateUserExperiences []input.UserExperienceInput,
) (res []entity.UserExperience, err error) {
	for _, updateUserExperience := range updateUserExperiences {
		for _, userExperienceDB := range resumeDB.UserExperiences {
			if updateUserExperience.UserExperienceID == userExperienceDB.UserExperienceID.String() {
				userExperienceDB.City = updateUserExperience.City
				userExperienceDB.CompanyName = updateUserExperience.CompanyName
				userExperienceDB.Position = updateUserExperience.Position

				startDate, e := time.Parse(layoutISO, updateUserExperience.StartDate)

				if e != nil {
					err = e
					break
				}

				userExperienceDB.StartDate = startDate

				endDate, e := time.Parse(layoutISO, updateUserExperience.EndDate)

				if e != nil {
					err = e
					break
				}

				userExperienceDB.EndDate = endDate

				res = append(res, userExperienceDB)
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

func createUserExperiencesToUpdate(createUserExperiencesInput []input.UserExperienceInput) (userExperiences []entity.UserExperience, err error) {
	if len(createUserExperiencesInput) > 0 {
		for _, userExperience := range createUserExperiencesInput {
			var createUserExperience entity.UserExperience

			createUserExperience.City = userExperience.City
			createUserExperience.CompanyName = userExperience.CompanyName
			createUserExperience.Position = userExperience.Position

			startDate, e := time.Parse(layoutISO, string(userExperience.StartDate))

			if e != nil {
				err = e
				break
			}

			createUserExperience.StartDate = startDate

			endDate, e := time.Parse(layoutISO, string(userExperience.EndDate))

			if e != nil {
				err = e
				break
			}

			createUserExperience.EndDate = endDate

			userExperiences = append(userExperiences, createUserExperience)
		}
	}

	if err != nil {
		return nil, err
	}

	return userExperiences, nil
}

func createUserExperiences(createResumeInput input.CreateResumeInput) (userExperiences []entity.UserExperience, err error) {

	if len(createResumeInput.UserExperiences) > 0 && createResumeInput.IsHaveExperience {
		for _, userExperience := range createResumeInput.UserExperiences {
			var createUserExperience entity.UserExperience

			createUserExperience.City = userExperience.City
			createUserExperience.CompanyName = userExperience.CompanyName
			createUserExperience.Position = userExperience.Position

			startDate, e := time.Parse(layoutISO, userExperience.StartDate)

			if e != nil {
				err = e
				break
			}

			createUserExperience.StartDate = startDate

			endDate, e := time.Parse(layoutISO, userExperience.EndDate)

			if e != nil {
				err = e
				break
			}

			createUserExperience.EndDate = endDate

			userExperiences = append(userExperiences, createUserExperience)
		}
	}

	if err != nil {
		return nil, err
	}

	return userExperiences, nil
}

func getCreateAndUpdateUserExperiences(
	updateUserExperiencesResumeInput input.UpdateUserExperiencesResumeInput,
	userExperienceChannel chan UserExperienceChannel,
) {
	for _, userExperience := range updateUserExperiencesResumeInput.UserExperiences {
		if userExperience.UserExperienceID == "" {
			// channel out put createUserExperience
			userExperienceChannelOutData := UserExperienceChannel{
				UserExperienceInput:    userExperience,
				isUserExperienceUpdate: false,
			}
			userExperienceChannel <- userExperienceChannelOutData
		} else {
			// channel out put updateUserExperience
			userExperienceChannelOutData := UserExperienceChannel{
				UserExperienceInput:    userExperience,
				isUserExperienceUpdate: true,
			}
			userExperienceChannel <- userExperienceChannelOutData
		}
	}
	// Close channel
	close(userExperienceChannel)
}

func getCreateAndUpdateUserEducations(
	updateUserEducationsResumeInput input.UpdateUserEducationsResumeInput,
	userEducationChannel chan UserEducationChannel,
) {
	for _, userEducation := range updateUserEducationsResumeInput.UserEducations {
		if userEducation.UserEducationID == "" {
			// channel out put createUserEducation
			userEducationChannelOutData := UserEducationChannel{
				UserEducationInput:    userEducation,
				isUserEducationUpdate: false,
			}
			userEducationChannel <- userEducationChannelOutData
		} else {
			// channel out put updateUserEducation
			userEducationChannelOutData := UserEducationChannel{
				UserEducationInput:    userEducation,
				isUserEducationUpdate: true,
			}
			userEducationChannel <- userEducationChannelOutData
		}
	}

	// Close channel
	close(userEducationChannel)
}

func setGender(gender string) string {
	if gender != "" {
		if gender == "MALE" {
			return gender
		} else if gender == "FEMALE" {
			return gender
		} else {
			return ""
		}
	}

	return ""
}
