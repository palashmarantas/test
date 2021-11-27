package main

import (
	"fmt"

	"github.com/test/src"
)

func main() {
	fmt.Println("hello")

	car := src.New()
	bike := src.New()

	car.Start()
	car.Stop()

	bike.Start()
	bike.Stop()

}
