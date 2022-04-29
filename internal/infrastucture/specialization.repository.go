package repository

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Almazatun/golephant/pkg/http/presentation/_type"
)

type specializationRepository struct{}

type SpecializationRepo interface {
	List() (positionTypes []_type.Specialization, err error)
}

func NewSpecializationRepo() SpecializationRepo {
	return &specializationRepository{}
}

// File-path no the same current folder structure, because the app running inside docker
var filePathSpecializations = "../app/pkg/common/constant/specialization-list.json"

func (r *specializationRepository) List() (specializations []_type.Specialization, err error) {
	fileBytes, err := ioutil.ReadFile(filePathSpecializations)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileBytes, &specializations)

	if err != nil {
		return nil, err
	}

	return specializations, nil
}
