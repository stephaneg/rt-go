package main

import (
	"fmt"

	"github.com/stephstephg/rt-go/core"
)

func main() {

	c := core.NewCanvas(10, 2)
	var i, j uint

	for i = 0; i < 10; i++ {
		for j = 0; j < 2; j++ {
			c.Write(i, j, core.NewColor(1.0, 0.8, 0.6))
		}
	}

	s := c.ToPPM()
	fmt.Print(s)

}
