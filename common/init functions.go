package common

import (
	"runtime"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/GUMI-golang/example/asset"
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/glumi"
	"github.com/iamGreedy/freetype/truetype"
)

func GoRuntimeInit() {
	runtime.LockOSThread()
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func GLInit() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
}
func SDL2Init() {
	gcore.Must(sdl.Init(sdl.INIT_EVERYTHING))

	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 4)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_FORWARD_COMPATIBLE_FLAG, 1)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)
}
func GUMIInit() {
	f, err := truetype.Parse(asset.MustAsset("NanumSquareRoundR.ttf"))
	if err != nil {
		panic(err)
	}
	gumi.ModifyDefaultStyle(f, 12)
}
func GLUMIInit() {
	err := glumi.Init()
	if err != nil {
		panic(err)
	}
}