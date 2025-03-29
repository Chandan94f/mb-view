package service

import (
	"github.com/Chandan94f/mb-view/internal/dao"
	"github.com/Chandan94f/mb-view/internal/models"
)

type Service interface {
	AddStudent(request models.AddStudentRequest) (models.AddStudentResponse, error)
	// GetStudent(request models.GetStudentRequest) (models.GetStudentResponse, error)
}

type serviceImp struct {
	dao dao.DaoService
}

type ServiceParams struct {
	Dao dao.DaoService
}

func NewService(serviceParams ServiceParams) Service {
	return &serviceImp{
		dao: serviceParams.Dao,
	}
}
