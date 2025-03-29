package service

import (
	"github.com/Chandan94f/mb-view/internal/models"
)

// // AddStudent implements Service.
func (s *serviceImp) AddStudent(request models.AddStudentRequest) (models.AddStudentResponse, error) {
	response := models.AddStudentResponse{
		Name:   request.Name,
		RollNo: request.RollNo,
		Age:    request.Age,
	}

	resp, err := s.dao.AddStudent(request)
	if err != nil {
		return response, err
	}
	response.EntryNo = resp.EntryNo
	return response, nil
}
