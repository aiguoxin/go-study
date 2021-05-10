package main

import (
	shaper "agxgin/interfacetest"
	"fmt"
)

func main() {
	//agx.F()

	sq1 := new(shaper.Square)
	sq1.Side = 5

	var areaIntf shaper.Shaper
	areaIntf = sq1
	fmt.Printf("%f\n", areaIntf.Area())

	areaIntf.Ok()
	fmt.Printf("%f\n",sq1.Side)
}





