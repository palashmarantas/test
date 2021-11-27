package src

import (
	"fmt"
	"github.com/test/src/core"
)
type Car struct {
	name string
}

func New() core.Vechicle{
	car := &car{name: "car"}
	return car
}

func (car *Car) Start()
{
   fmt.Println("car started")
}

func (car *Car) Stop()
{
   fmt.Println("car stopped")
}
