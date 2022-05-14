package usecase

import (
	"errors"

	repository "github.com/Almazatun/golephant/internal/infrastucture"
	"github.com/Almazatun/golephant/internal/infrastucture/entity"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
)

type positionUseCase struct {
	positionRepo repository.PositionRepo
}

type PositionUseCase interface {
	UpdateStatus(
		companyId, positionId string,
	) (positionDB *entity.Position, err error)
	Delete(
		companyId, positionId string,
	) (str string, err error)
}

func NewPositionUseCase(positionRepositor repository.PositionRepo) PositionUseCase {
	return &positionUseCase{
		positionRepo: positionRepositor,
	}
}

func (uc *positionUseCase) UpdateStatus(
	companyId, positionId string,
) (positionDB *entity.Position, err error) {
	position, err := uc.positionRepo.GetById(positionId)

	if err != nil {
		return nil, err
	}

	if position.PositionID.String() != positionId &&
		position.CompanyID.String() != companyId {
		err = errors.New(error_message.BAD_REGUEST)
		return nil, err
	}

	if position.Status == "OPEN" {
		position.Status = "CLOSE"
	} else {
		position.Status = "OPEN"
	}

	res, err := uc.positionRepo.Save(*position)

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

	if position.CompanyID.String() != companyId && position.CompanyID.String() != companyId {
		newErr := errors.New(error_message.BAD_REGUEST)
		return "", newErr
	}

	res, err := uc.positionRepo.DeleteById(positionId)

	if err != nil {
		return "", err
	}

	return res, nil
}
