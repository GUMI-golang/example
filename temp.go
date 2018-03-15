package main

import (
	"github.com/GUMI-golang/gumi/gcore"
	"fmt"
)

func main() {

	for i := 0; i <= 100; i ++{
		fmt.Println(gcore.Animation.Functions.Cubic.EasingOut(float64(i) / 100))
	}

}
