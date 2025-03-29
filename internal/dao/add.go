package dao

import (
	"github.com/Chandan94f/mb-utility/utility/logger"
	"github.com/Chandan94f/mb-view/internal/models"
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
