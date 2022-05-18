package usecase

import (
	"errors"

	repository "github.com/Almazatun/golephant/internal/infrastucture"
	"github.com/Almazatun/golephant/internal/infrastucture/entity"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	"github.com/Almazatun/golephant/pkg/http/presentation/input"
)

type positionUseCase struct {
	positionRepo repository.PositionRepo
}

type PositionUseCase interface {
	UpdateStatus(
		companyId, positionId string,
	) (positionDB *entity.Position, err error)
	UpdateResponsibilities(
		companyId, positionId string,
		updatePositionResponsobilitesInput input.UpdatePositionResponsobilitesInput,
	) (positionDB *entity.Position, err error)
	UpdateRequirements(
		companyId, positionId string,
		updatePositionRequirementsInput input.UpdatePositionRequirementsInput,
	) (positionDB *entity.Position, err error)
	Update(
		companyId, positionId string,
		updatePositionInput input.UpdatePositionInput,
	) (positionDB *entity.Position, err error)
	Delete(
		companyId, positionId string,
	) (str string, err error)
	isCompanyPosition(
		companyId, positionId string,
		position entity.Position,
	) bool
}

func NewPositionUseCase(positionRepositor repository.PositionRepo) PositionUseCase {
	return &positionUseCase{
		positionRepo: positionRepositor,
	}
}

func (uc *positionUseCase) Update(
	companyId, positionId string,
	updatePositionInput input.UpdatePositionInput,
) (positionDB *entity.Position, err error) {
	position, err := uc.positionRepo.GetById(positionId)

	if err != nil {
		return nil, err
	}

	if !uc.isCompanyPosition(companyId, positionId, *position) {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	if updatePositionInput.Description != "" {
		position.Description = updatePositionInput.Description
	}

	if updatePositionInput.Salary != *position.Salary {
		position.Salary = &updatePositionInput.Salary
	}

	if updatePositionInput.PositionType != "" {
		position.PositionType = updatePositionInput.PositionType
	}

	res, err := uc.positionRepo.Save(*position)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (uc *positionUseCase) UpdateResponsibilities(
	companyId, positionId string,
	updatePositionResponsobilitesInput input.UpdatePositionResponsobilitesInput,
) (positionDB *entity.Position, err error) {
	position, err := uc.positionRepo.GetById(positionId)

	if err != nil {
		return nil, err
	}

	if !uc.isCompanyPosition(companyId, positionId, *position) {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	position.Responsibilities = updatePositionResponsobilitesInput.Responsobilities

	res, err := uc.positionRepo.Save(*position)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *positionUseCase) UpdateRequirements(
	companyId, positionId string,
	updatePositionRequirementsInput input.UpdatePositionRequirementsInput,
) (positionDB *entity.Position, err error) {
	position, err := uc.positionRepo.GetById(positionId)

	if err != nil {
		return nil, err
	}

	if !uc.isCompanyPosition(companyId, positionId, *position) {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	position.Requirements = updatePositionRequirementsInput.Requirements

	res, err := uc.positionRepo.Save(*position)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *positionUseCase) UpdateStatus(
	companyId, positionId string,
) (positionDB *entity.Position, err error) {
	position, err := uc.positionRepo.GetById(positionId)

	if err != nil {
		return nil, err
	}

	if !uc.isCompanyPosition(companyId, positionId, *position) {
		newErr := errors.New(error_message.BAD_REGUEST)
		return nil, newErr
	}

	res, err := uc.positionRepo.Save(uc.updateStatus(*position))

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *positionUseCase) Delete(
	companyId, positionId string,
) (str string, err error) {
	position, err := uc.positionRepo.GetById(positionId)

	if err != nil {
		return "", err
	}

	if !uc.isCompanyPosition(companyId, positionId, *position) {
		newErr := errors.New(error_message.BAD_REGUEST)
		return "", newErr
	}

	res, err := uc.positionRepo.DeleteById(positionId)

	if err != nil {
		return "", err
	}

	return res, nil
}

func (uc *positionUseCase) isCompanyPosition(
	companyId, positionId string,
	position entity.Position,
) bool {
	return position.PositionID.String() == positionId &&
		position.CompanyID.String() == companyId
}

func (uc *positionUseCase) updateStatus(
	position entity.Position,
) entity.Position {
	if position.Status == "OPEN" {
		position.Status = "CLOSE"
	} else {
		position.Status = "OPEN"
	}

	return position
}
