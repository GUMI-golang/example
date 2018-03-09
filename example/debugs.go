package example

import (
	"github.com/GUMI-golang/gumi"
	"regexp"
)
var re = regexp.MustCompile(`debug\((?P<name>.*)\)`)
var DebugFPS = gumi.NDrawing0(gumi.Drawing.FPS())

var DebugFPSResolutions = gumi.NDrawing0(
	gumi.Drawing.FPS(),
	gumi.Drawing.Ruler.Screen(),
)

var DebugFPSRuler = gumi.NDrawing0(
	gumi.Drawing.FPS(),
	gumi.Drawing.Ruler.Graduation.Horizontal(100),
	gumi.Drawing.Ruler.Graduation.Vertical(100),
)

func debugAppend(name string) gumi.GUMI {
	res := re.FindStringSubmatch(name)
	if res == nil || len(res) <2{
		return DebugFPS
	}

	switch res[1]{
	default:
		fallthrough
	case "FPS":
		return DebugFPS
	case "FPSResolutions":
		return DebugFPSResolutions
	case "FPSRuler":
		return DebugFPSRuler
	}
}