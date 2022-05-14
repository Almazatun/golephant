package repository

import (
	"errors"

	entity "github.com/Almazatun/golephant/internal/infrastucture/entity"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	"gorm.io/gorm"
)

type positionRepository struct {
	db *gorm.DB
}

type PositionRepo interface {
	Save(position entity.Position) (positionDB *entity.Position, err error)
	GetById(positionId string) (positionDB *entity.Position, err error)
	DeleteById(positionId string) (str string, err error)
}

func NewPositionRepo(db *gorm.DB) PositionRepo {
	return &positionRepository{
		db: db,
	}
}

func (r *positionRepository) Save(
	position entity.Position,
) (positionDB *entity.Position, err error) {
	result := r.db.Save(&position)

	e := result.Error

	if e != nil {
		return nil, e
	}

	return &position, nil
}

func (r *positionRepository) GetById(
	positionId string,
) (positionDB *entity.Position, err error) {
	var position entity.Position

	result := r.db.First(&position, "position_id = ?", positionId)

	dbErr := result.Error

	if dbErr != nil {
		err := errors.New(error_message.POSITION_NOT_FOUND)

		return nil, err
	}

	return &position, nil
}

func (r *positionRepository) DeleteById(positionId string) (str string, err error) {
	res := "Position successfully deleted"

	result := r.db.Delete(&entity.Position{}, "position_id = ?", positionId)

	er := result.Error

	if er != nil {
		return "", er
	}

	return res, nil
}
