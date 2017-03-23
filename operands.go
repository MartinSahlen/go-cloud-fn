package main

import (
	"strconv"

	"github.com/MartinSahlen/go-cloud-fn/app"
	"github.com/goadesign/goa"
)

// OperandsController implements the Operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a Operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("OperandsController")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	// OperandsController_Add: start_implement

	// Put your logic here
	ctx.OK([]byte(strconv.Itoa(ctx.Left + ctx.Right)))
	// OperandsController_Add: end_implement
	return nil
}
