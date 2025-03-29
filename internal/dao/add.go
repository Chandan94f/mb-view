package dao

import (
	"github.com/mb-view/internal/models"
	"github.com/mb-view/internal/utility/logger"
)

type dao struct {
	logger logger.Logger
}

// AddStudent implements DaoService.
func (d *dao) AddStudent(request models.AddStudentRequest) (models.AddStudentResponse, error) {
	panic("unimplemented")
}

// GetStudent implements DaoService.
func (d *dao) GetStudent(request models.GetStudentRequest) (models.GetStudentResponse, error) {
	panic("unimplemented")
}

type DoaParams struct {
	Logger logger.Logger
}

func NewDaoService(doaParams DoaParams) DaoService {
	return &dao{
		logger: doaParams.Logger,
	}
}
