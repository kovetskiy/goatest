package main

import (
	"errors"

	"github.com/goadesign/goa"
	"github.com/kovetskiy/goatest/app"
)

// SomeController implements the some resource.
type SomeController struct {
	*goa.Controller
}

// NewSomeController creates a some controller.
func NewSomeController(service *goa.Service) *SomeController {
	return &SomeController{Controller: service.NewController("SomeController")}
}

// Get runs the get action.
func (c *SomeController) Get(ctx *app.GetSomeContext) error {
	return errors.New("something went wrong")
}
