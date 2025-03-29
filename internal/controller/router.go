package controller

import (
	// "github.com/mb-view/internal/controller"
	"github.com/mb-view/internal/literals"
	"github.com/mb-view/internal/models"

	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) SetupRoutes(router *mux.Router) *models.ServiceRouter {

	return &models.ServiceRouter{
		Routes: []models.CommonRouter{
			{
				Endpoint: literals.AddStudent,
				Method:   http.MethodPost,
				Handler:  c.addStudent,
			},
			// {
			// 	Endpoint: literals.GetStudent,
			// 	Method:   http.MethodPost,
			// 	Handler:  c.getStudent,
			// },
		},
	}

	// router.HandleFunc("/students", AddStudentHandler).Methods("POST")
	// router.HandleFunc("/students/{id}", GetStudentHandler).Methods("GET")
}
