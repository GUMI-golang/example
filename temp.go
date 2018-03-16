package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/GUMI-golang/gumi/gcore"
)

func main() {

	r, radian := gcore.ToPolar(1,-1)
	fmt.Println(r, gg.Degrees(radian), radian)
}
