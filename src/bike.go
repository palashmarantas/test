package src

import (
	"fmt"
	"github.com/test/src/core"
)
type Bike struct {
	name string
}

func New() core.Vechicle {
	bike := &bike{name: "bike"}
	return bike
}

func (bike *Bike) Start()
{
   fmt.Println("bike started")
}

func (bike *Bike) Stop()
{
   fmt.Println("bike stopped")
}
