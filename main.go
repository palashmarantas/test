package main

import (
	"fmt"

	"github.com/test/src"
)

func main() {
	fmt.Println("hello")

	c := src.NewCar("car")
	b := src.NewBike("bike")

	c.Start()
	c.Stop()

	b.Start()
	b.Stop()

}
