package main

import (
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/example/common"
	"flag"
	"fmt"
)



func main() {
	flag.Set("size", "VGA")
	flag.Set("example", "HelloImage")

	gcore.Must(common.CheckFlags())
	var width, height = common.GetSize()
	var example = common.GetExample()
	//
	scr := gumi.NewScreen(width, height)
	scr.Root(example)
	scr.Init()
	scr.Update(gumi.Information{Dt:0})
	scr.Draw()
	scr.Update(gumi.Information{Dt:0})
	scr.Draw()
	gumi.Capture("out", scr.Frame())
}
