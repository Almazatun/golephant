package repository

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Almazatun/golephant/pkg/http/presentation/_type"
)

type positionTypeRepository struct{}

type PositionTypeRepo interface {
	List() (positionTypes []_type.PositonType, err error)
}

func NewPositionTypeRepo() PositionTypeRepo {
	return &positionTypeRepository{}
}

// File-path no the same current folder structure, because the app running inside docker
var filePathPositionTypeList = "../app/pkg/common/constant/position-type-list.json"

func (r *positionTypeRepository) List() (positionTypes []_type.PositonType, err error) {
	fileBytes, err := ioutil.ReadFile(filePathPositionTypeList)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileBytes, &positionTypes)

	if err != nil {
		return nil, err
	}

	return positionTypes, nil
}
