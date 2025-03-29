package dao

import "github.com/mb-view/internal/models"

type DaoService interface {
	AddStudent(request models.AddStudentRequest) (models.AddStudentResponse, error)
	GetStudent(request models.GetStudentRequest) (models.GetStudentResponse, error)
}
