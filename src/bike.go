package src

import (
	"fmt"

	"github.com/test/core"
)

type Bike struct {
	Name string
}

func NewBike(name string) core.Vehicle {
	return &Bike{Name: name}
}

func (bike *Bike) Start() int {
	fmt.Println("bike started")
	return 0
}

func (bike *Bike) Stop() int {
	fmt.Println("bike stopped")
	return 0
}
