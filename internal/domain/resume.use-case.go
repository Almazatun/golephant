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
	resumeRepo         repository.ResumeRepo
	userRepo           repository.UserRepo
	userEducationRepo  repository.UserEducationRepo
	userExperienceRepo repository.UserExperienceRepo
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
	) (updateAboutMeResume *entity.Resume, err error)
	UpdateCitizenship(
		userId, resumeId string,
		updateCitizenshipResumInput input.UpdateCitizenshipResumeInput,
	) (updateCitizenshipResume *entity.Resume, err error)
	UpdateTags(
		userId, resumeId string,
		updateTagsResume input.UdateTagsResumeInput,
	) (updateTagsResum *entity.Resume, err error)
	UpdateUserEducation(
		userId, resumeId string,
		updateUserEducationsResumeInput input.UpdateUserEducationsResumeInput,
	) (updateUserEducationsResum *entity.Resume, err error)
	UpdateUserExperiences(
		userId, resumeId string,
		updateUserExperiencesResumeInput input.UpdateUserExperiencesResumeInput,
	) (updateUserEducationsResum *entity.Resume, err error)
	UpdateDesiredPosition(
		userId, resumeId string,
		updateDesiredPositionResumeInput input.UpdateDesiredPositionResumeInput,
	) (updateAboutMeResume *entity.Resume, err error)
	Delete(resumeId string) (str string, err error)
	DeleteUserExperience(resumeId, userExperienceId string) (str string, err error)
	DeleteUserEducation(resumeId, userEducationId string) (str string, err error)
	fillResumeToCreate(
		createResumeInput input.CreateResumeInput,
	) (resume entity.Resume)
	createUserEducations(
		createUserEducationsResumeInput []input.UserEducationInput,
	) (res []entity.UserEducation, err error)
	updateUserEducations(
		resumeDB entity.Resume,
		updateUserEducations []input.UserEducationInput,
	) (res []entity.UserEducation, err error)
	updateUserExperiences(
		resumeDB *entity.Resume,
		updateUserExperiences []input.UserExperienceInput,
	) (res []entity.UserExperience, err error)
	createUserExperiencesToUpdate(
		createUserExperiencesInput []input.UserExperienceInput,
	) (res []entity.UserExperience, err error)
	createUserExperiences(
		createResumeInput input.CreateResumeInput,
	) (res []entity.UserExperience, err error)
	getCreateAndUpdateUserExperiences(
		updateUserExperiencesResumeInput input.UpdateUserExperiencesResumeInput,
		userExperienceChannel chan input.ResumeInput[input.UserExperienceInput],
	)
	getCreateAndUpdateUserEducations(
		updateUserEducationsResumeInput input.UpdateUserEducationsResumeInput,
		userEducationChannel chan input.ResumeInput[input.UserEducationInput],
	)
	setGender(gender string) string
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

	// append user experiences in resume
	userExperiences, err := uc.createUserExperiences(createResumeInput)

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

func (uc *resumeUseCase) UpdateUserEducation(
	userId, resumeId string,
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

		userEducationСhannel := make(chan input.ResumeInput[input.UserEducationInput], 2)

		go uc.getCreateAndUpdateUserEducations(
			updateUserEducationsResumeInput,
			userEducationСhannel,
		)

		for userEducation := range userEducationСhannel {
			if userEducation.IsUpdate {
				updateUserEducationsInput = append(
					updateUserEducationsInput,
					userEducation.Data,
				)
			}
			createUserEducationsInput = append(
				createUserEducationsInput,
				userEducation.Data,
			)
		}

		createUserEducationsDB, err := uc.createUserEducations(createUserEducationsInput)

		if err != nil {
			return nil, err
		}

		updateUserEducationsDB, err := uc.updateUserEducations(*resumeDB, updateUserEducationsInput)

		if err != nil {
			return nil, err
		}

		resumeDB.UserEducations = nil

		resumeDB.UserEducations = append(resumeDB.UserEducations, createUserEducationsDB...)
		resumeDB.UserEducations = append(resumeDB.UserEducations, updateUserEducationsDB...)
	}

	res, err := uc.resumeRepo.Save(*resumeDB)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeUseCase) UpdateUserExperiences(
	userId, resumeId string,
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

		userExperienceChannel := make(chan input.ResumeInput[input.UserExperienceInput], 2)

		go uc.getCreateAndUpdateUserExperiences(
			updateUserExperiencesResumeInput,
			userExperienceChannel,
		)

		for userExperience := range userExperienceChannel {
			if userExperience.IsUpdate {
				updateUserExperiencesInput = append(
					updateUserExperiencesInput,
					userExperience.Data,
				)
			}
			createUserExperiencesInput = append(
				createUserExperiencesInput,
				userExperience.Data,
			)
		}

		createUserEducationsDB, err := uc.createUserExperiencesToUpdate(createUserExperiencesInput)

		if err != nil {
			return nil, err
		}

		updateUserExperiencesDB, err := uc.updateUserExperiences(resumeDB, updateUserExperiencesInput)

		if err != nil {
			return nil, err
		}

		resumeDB.UserExperiences = nil

		resumeDB.UserExperiences = append(resumeDB.UserExperiences, createUserEducationsDB...)
		resumeDB.UserExperiences = append(resumeDB.UserExperiences, updateUserExperiencesDB...)
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

func (uc *resumeUseCase) DeleteUserExperience(resumeId, userExperienceId string) (str string, err error) {

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

func (uc *resumeUseCase) DeleteUserEducation(resumeId, userEducationId string) (str string, err error) {
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

func (uc *resumeUseCase) createUserEducations(
	createUserEducationsResumeInput []input.UserEducationInput,
) (res []entity.UserEducation, err error) {

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

			res = append(res, createUserEducation)
		}
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeUseCase) updateUserEducations(
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

func (uc *resumeUseCase) updateUserExperiences(
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

func (uc *resumeUseCase) createUserExperiencesToUpdate(
	createUserExperiencesInput []input.UserExperienceInput,
) (res []entity.UserExperience, err error) {
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

			res = append(res, createUserExperience)
		}
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeUseCase) createUserExperiences(
	createResumeInput input.CreateResumeInput,
) (res []entity.UserExperience, err error) {

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

			res = append(res, createUserExperience)
		}
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeUseCase) getCreateAndUpdateUserExperiences(
	updateUserExperiencesResumeInput input.UpdateUserExperiencesResumeInput,
	userExperienceChannel chan input.ResumeInput[input.UserExperienceInput],
) {
	for _, userExperience := range updateUserExperiencesResumeInput.UserExperiences {
		if userExperience.UserExperienceID == "" {
			// channel out put createUserExperience
			userExperienceCreate := input.ResumeInput[input.UserExperienceInput]{
				Data:     userExperience,
				IsUpdate: false,
			}
			userExperienceChannel <- userExperienceCreate
		} else {
			// channel out put updateUserExperience
			userExperienceUpdate := input.ResumeInput[input.UserExperienceInput]{
				Data:     userExperience,
				IsUpdate: true,
			}
			userExperienceChannel <- userExperienceUpdate
		}
	}
	// Close channel
	close(userExperienceChannel)
}

func (uc *resumeUseCase) getCreateAndUpdateUserEducations(
	updateUserEducationsResumeInput input.UpdateUserEducationsResumeInput,
	userEducationChannel chan input.ResumeInput[input.UserEducationInput],
) {
	for _, userEducation := range updateUserEducationsResumeInput.UserEducations {
		if userEducation.UserEducationID == "" {
			// channel out put createUserEducation
			userEducationCreate := input.ResumeInput[input.UserEducationInput]{
				Data:     userEducation,
				IsUpdate: false,
			}
			userEducationChannel <- userEducationCreate
		} else {
			// channel out put updateUserEducation
			userEducationChannelOutData := input.ResumeInput[input.UserEducationInput]{
				Data:     userEducation,
				IsUpdate: true,
			}
			userEducationChannel <- userEducationChannelOutData
		}
	}

	// Close channel
	close(userEducationChannel)
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
