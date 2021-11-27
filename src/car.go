package src

import (
	"fmt"

	"github.com/test/core"
)

type Car struct {
	Name string
}

func NewCar(name string) core.Vehicle {
	return &Car{Name: name}
}

func (car *Car) Start() int {
	fmt.Println("car started")
	return 0
}

func (car *Car) Stop() int {
	fmt.Println("car stopped")
	return 0
}
