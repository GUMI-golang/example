package main

import (
	"fmt"
	"github.com/GUMI-golang/example/common"
	"github.com/GUMI-golang/glumi"
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/veandco/go-sdl2/sdl"
	"flag"
)


func main() {
	flag.Set("size", "HD")
<<<<<<< Updated upstream
	flag.Set("example", "debug(FPS).Editing")
=======
	flag.Set("example", "debug.HelloImage")
>>>>>>> Stashed changes
	//
	var err error
	gcore.Must(common.CheckFlags())
	var width, height = common.GetSize()
	var example = common.GetExample()
	// common Init
	common.GoRuntimeInit()
	common.SDL2Init()
	// Make SDL OpenGL Window
	wnd := SDL2Window(width, height)
	ctx, err := sdl.GLCreateContext(wnd)
<<<<<<< Updated upstream
	gcore.Must(err)
=======
	gcore.Assert(err)
>>>>>>> Stashed changes
	sdl.GLSetSwapInterval(0)
	defer sdl.GLDeleteContext(ctx)
	sdl.StartTextInput()
	//
	common.GLInit()
	common.GUMIInit()
	common.GLUMIInit()
	// Print OpenGL
	fmt.Println("OpenGL version : ", gl.GoStr(gl.GetString(gl.VERSION)))

	// Create GUMI Screen
	scr := gumi.NewScreen(width, height)
	scr.Root(
		//gumi.LinkingFrom(
		//	gumi.NDrawing0(gumi.Drawing.FPS()),
		//	example,
		//),
		example,
	)
	// Create GLUMI Object allocate
	lumi := glumi.NewGLUMI()
	lumi.SetScreen(scr)
	// glumi Initalize
<<<<<<< Updated upstream
	gcore.Must(lumi.Init(0))
=======
	gcore.Assert(lumi.Init(300))
>>>>>>> Stashed changes
	// Main Loop
	lumi.Loop(
		// Event Process, GL Clearing
		func(lumi *glumi.GLUMIFullScreen) error {
			if err != nil {
				return err
			}
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				err = lumi.Event.Handle(event)
				if err != nil {
					return err
				}
			}
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
			return nil
		},
		// GL Frame Swap
		func(lumi *glumi.GLUMIFullScreen) error {
			sdl.GLSwapWindow(wnd)
			return nil
		},

	)

}

func SDL2Window(w, h int) *sdl.Window {
	var disp sdl.DisplayMode
	gcore.Must(sdl.GetDesktopDisplayMode(0, &disp))
	var windW, windH int32
	if int32(w) > disp.W {
		windW = disp.W
	} else {
		windW = int32(w)
	}
	if int32(h) > disp.H {
		windH = disp.H
	} else {
		windH = int32(h)
	}
	wnd, err := sdl.CreateWindow("GUMI", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, windW, windH,
		sdl.WINDOW_OPENGL,
	)
	gcore.Must(err)
	return wnd
}
