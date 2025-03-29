package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mb-view/internal/literals"
	"github.com/mb-view/internal/models"
)

func (c *Controller) addStudent(w http.ResponseWriter, r *http.Request) {
	var (
		request  models.AddStudentRequest
		response models.AddStudentResponse
	)

	// Check if the request body is empty
	if r.Body == nil {
		c.resp.HTTPFail(w, http.StatusBadRequest, fmt.Sprintf(literals.RequestValidationFailed, "Add Student"), "request body is empty")
		return
	}

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.resp.HTTPFail(w, http.StatusBadRequest, fmt.Sprintf(literals.RequestValidationFailed, "Add Student"), err.Error())
		return
	}

	// Call the service layer to add student
	response, err := c.service.AddStudent(request)
	if err != nil {
		c.resp.HTTPFail(w, http.StatusInternalServerError, "failed to add student", err.Error())
		return
	}

	// Respond with success
	c.resp.HTTPSuccess(r, w, http.StatusOK, response)
}
