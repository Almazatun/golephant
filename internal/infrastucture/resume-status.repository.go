package repository

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Almazatun/golephant/pkg/http/presentation/_type"
)

type resumeStatusRepository struct{}

type ResumeStatusRepo interface {
	List() (resumeStatuses []_type.ResumeStatus, err error)
}

func NewResumeStatusRepo() ResumeStatusRepo {
	return &resumeStatusRepository{}
}

// File-path no the same current folder structure, because the app running inside docker
var filePathResumeStatuses = "../app/pkg/common/constant/resume-status-list.json"

func (r *resumeStatusRepository) List() (resumeStatuses []_type.ResumeStatus, err error) {
	fileBytes, err := ioutil.ReadFile(filePathResumeStatuses)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileBytes, &resumeStatuses)

	if err != nil {
		return nil, err
	}

	return resumeStatuses, nil
}
