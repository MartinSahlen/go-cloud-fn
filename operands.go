package main

import (
	"log"
	"strconv"

	"github.com/MartinSahlen/go-cloud-fn/app"
	"github.com/goadesign/goa"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("OperandsController")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	// OperandsController_Add: start_implement
	log.Println("helo")
	// Put your logic here
	return ctx.OK([]byte(strconv.Itoa(ctx.Left + ctx.Right)))
	// OperandsController_Add: end_implement
}
