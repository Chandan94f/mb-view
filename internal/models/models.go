package models

import (
	"net/http"
	"sync"
)

type ServiceRouter struct {
	Routes       []CommonRouter
	wg           sync.WaitGroup
	shutdownHTTP func() error
	// HealthChecks HealthCheck // to be implemented
}

type CommonRouter struct {
	Method   string
	Endpoint string
	Handler  http.HandlerFunc
	// Middlewares []mux.MiddlewareFuncstudent // to be implemented
}

// type Response struct {
// 	authentication.Authentication
// 	errorutils.ErrorEnricherconst
// }

type student struct {
	Name   string `json:"name"`
	RollNo string `json:"entryNo"`
	Age    string `json:"age"`
}

type AddStudentRequest struct {
	Name   string `json:"name" validate:"required"`
	RollNo string `json:"rollNo" validate:"required"`
	Age    string `json:"age"`
}

type AddStudentResponse struct {
	Name    string `json:"name"`
	RollNo  string `json:"rollNo"`
	Age     string `json:"age"`
	EntryNo string `json:"entryNo" validate:"required"`
}

type GetStudentRequest struct {
	EntryNo string `json:"entryNo" validate:"required"`
}

type GetStudentResponse struct {
	EntryNo string `json:"entryNo"`
	Name    string `json:"name"`
	RollNo  string `json:"rollNo"`
	Age     string `json:"age"`
}
