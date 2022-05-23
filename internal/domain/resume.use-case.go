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
	) (updateResume *entity.Resume, err error)
	UpdateCitizenship(
		userId, resumeId string,
		updateCitizenshipResumInput input.UpdateCitizenshipResumeInput,
	) (updateResume *entity.Resume, err error)
	UpdateTags(
		userId, resumeId string,
		updateTagsResume input.UdateTagsResumeInput,
	) (updateResume *entity.Resume, err error)
	UpdateUserEducation(
		userId, resumeId string,
		updateUserEducationResumeInput input.UpdateUserEducationResumeInput,
	) (updateResume *entity.Resume, err error)
	UpdateUserExperience(
		userId, resumeId string,
		updateUserExperienceResumeInput input.UpdateUserExperienceResumeInput,
	) (updateResume *entity.Resume, err error)
	UpdateDesiredPosition(
		userId, resumeId string,
		updateDesiredPositionResumeInput input.UpdateDesiredPositionResumeInput,
	) (updateResume *entity.Resume, err error)
	Delete(resumeId string) (str string, err error)
	DeleteUserExperience(resumeId, userExperienceId string) (str string, err error)
	DeleteUserEducation(resumeId, userEducationId string) (str string, err error)
	fillResumeToCreate(
		createResumeInput input.CreateResumeInput,
	) (resume entity.Resume)
	createUserEducation(
		createUserEducationResumeInput []input.UserEducationInput,
	) (res []entity.UserEducation, err error)
	updateUserEducation(
		resumeDB entity.Resume,
		updateUserEducation []input.UserEducationInput,
	) (res []entity.UserEducation, err error)
	updateUserExperience(
		resumeDB *entity.Resume,
		updateUserExperience []input.UserExperienceInput,
	) (res []entity.UserExperience, err error)
	createUserExperienceToUpdate(
		createUserExperienceInput []input.UserExperienceInput,
	) (res []entity.UserExperience, err error)
	createUserExperience(
		createResumeInput input.CreateResumeInput,
	) (res []entity.UserExperience, err error)
	getCreateAndUpdateUserExperience(
		updateUserExperienceResumeInput input.UpdateUserExperienceResumeInput,
		userExperienceChannel chan input.ResumeInput[input.UserExperienceInput],
	)
	getCreateAndUpdateUserEducation(
		updateUserEducationResumeInput input.UpdateUserEducationResumeInput,
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
	userExperience, err := uc.createUserExperience(createResumeInput)

	if err != nil {
		return nil, err
	}

	resume.UserExperience = userExperience

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
	updateUserEducationResumeInput input.UpdateUserEducationResumeInput,
) (updateUserEducationResum *entity.Resume, err error) {
	resumeDB, err := uc.resumeRepo.GetById(resumeId)

	if err != nil {
		return nil, err
	}

	if resumeDB.UserID.String() != userId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	if len(updateUserEducationResumeInput.UserEducation) > 0 {
		var createUserEducationInput []input.UserEducationInput
		var updateUserEducationInput []input.UserEducationInput

		userEducationСhannel := make(chan input.ResumeInput[input.UserEducationInput], 2)

		go uc.getCreateAndUpdateUserEducation(
			updateUserEducationResumeInput,
			userEducationСhannel,
		)

		for userEducation := range userEducationСhannel {
			if userEducation.IsUpdate {
				updateUserEducationInput = append(
					updateUserEducationInput,
					userEducation.Data,
				)
			}
			createUserEducationInput = append(
				createUserEducationInput,
				userEducation.Data,
			)
		}

		createUserEducationDB, err := uc.createUserEducation(createUserEducationInput)

		if err != nil {
			return nil, err
		}

		updateUserEducationDB, err := uc.updateUserEducation(*resumeDB, updateUserEducationInput)

		if err != nil {
			return nil, err
		}

		resumeDB.UserEducation = nil

		resumeDB.UserEducation = append(resumeDB.UserEducation, createUserEducationDB...)
		resumeDB.UserEducation = append(resumeDB.UserEducation, updateUserEducationDB...)
	}

	res, err := uc.resumeRepo.Save(*resumeDB)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *resumeUseCase) UpdateUserExperience(
	userId, resumeId string,
	updateUserExperienceResumeInput input.UpdateUserExperienceResumeInput,
) (updateUserEducationResum *entity.Resume, err error) {
	resumeDB, err := uc.resumeRepo.GetById(resumeId)

	if err != nil {
		return nil, err
	}

	if resumeDB.UserID.String() != userId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	if len(updateUserExperienceResumeInput.UserExperience) > 0 {
		var createUserExperienceInput []input.UserExperienceInput
		var updateUserExperienceInput []input.UserExperienceInput

		userExperienceChannel := make(chan input.ResumeInput[input.UserExperienceInput], 2)

		go uc.getCreateAndUpdateUserExperience(
			updateUserExperienceResumeInput,
			userExperienceChannel,
		)

		for userExperience := range userExperienceChannel {
			if userExperience.IsUpdate {
				updateUserExperienceInput = append(
					updateUserExperienceInput,
					userExperience.Data,
				)
			}
			createUserExperienceInput = append(
				createUserExperienceInput,
				userExperience.Data,
			)
		}

		createUserEducationDB, err := uc.createUserExperienceToUpdate(createUserExperienceInput)

		if err != nil {
			return nil, err
		}

		updateUserExperienceDB, err := uc.updateUserExperience(resumeDB, updateUserExperienceInput)

		if err != nil {
			return nil, err
		}

		resumeDB.UserExperience = nil

		resumeDB.UserExperience = append(resumeDB.UserExperience, createUserEducationDB...)
		resumeDB.UserExperience = append(resumeDB.UserExperience, updateUserExperienceDB...)
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

func (uc *resumeUseCase) createUserEducation(
	createUserEducationResumeInput []input.UserEducationInput,
) (res []entity.UserEducation, err error) {

	if len(createUserEducationResumeInput) > 0 {
		for _, userEducation := range createUserEducationResumeInput {
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

func (uc *resumeUseCase) updateUserEducation(
	resumeDB entity.Resume,
	updateUserEducation []input.UserEducationInput,
) (res []entity.UserEducation, err error) {
	for _, userEducation := range updateUserEducation {
		for _, userEducationDB := range resumeDB.UserEducation {
			if userEducation.UserEducationID == userEducationDB.UserEducationID.String() {
				userEducationDB.City = userEducation.City
				userEducationDB.DegreePlacement = userEducation.DegreePlacement

				startDate, e := time.Parse(layoutISO, userEducation.StartDate)

				if e != nil {
					err = e
					break
				}

				userEducationDB.StartDate = startDate

				endDate, e := time.Parse(layoutISO, userEducation.EndDate)

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

func (uc *resumeUseCase) updateUserExperience(
	resumeDB *entity.Resume,
	updateUserExperience []input.UserExperienceInput,
) (res []entity.UserExperience, err error) {
	for _, userExperience := range updateUserExperience {
		for _, userExperienceDB := range resumeDB.UserExperience {
			if userExperience.UserExperienceID == userExperienceDB.UserExperienceID.String() {
				userExperienceDB.City = userExperience.City
				userExperienceDB.CompanyName = userExperience.CompanyName
				userExperienceDB.Position = userExperience.Position

				startDate, e := time.Parse(layoutISO, userExperience.StartDate)

				if e != nil {
					err = e
					break
				}

				userExperienceDB.StartDate = startDate

				endDate, e := time.Parse(layoutISO, userExperience.EndDate)

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

func (uc *resumeUseCase) createUserExperienceToUpdate(
	createUserExperienceInput []input.UserExperienceInput,
) (res []entity.UserExperience, err error) {
	if len(createUserExperienceInput) > 0 {
		for _, userExperience := range createUserExperienceInput {
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

func (uc *resumeUseCase) createUserExperience(
	createResumeInput input.CreateResumeInput,
) (res []entity.UserExperience, err error) {

	if len(createResumeInput.UserExperience) > 0 && createResumeInput.IsHaveExperience {
		for _, userExperience := range createResumeInput.UserExperience {
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

func (uc *resumeUseCase) getCreateAndUpdateUserExperience(
	updateUserExperienceResumeInput input.UpdateUserExperienceResumeInput,
	userExperienceChannel chan input.ResumeInput[input.UserExperienceInput],
) {
	for _, userExperience := range updateUserExperienceResumeInput.UserExperience {
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

func (uc *resumeUseCase) getCreateAndUpdateUserEducation(
	updateUserEducationResumeInput input.UpdateUserEducationResumeInput,
	userEducationChannel chan input.ResumeInput[input.UserEducationInput],
) {
	for _, userEducation := range updateUserEducationResumeInput.UserEducation {
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
