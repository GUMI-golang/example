package common

import (
	"github.com/GUMI-golang/gumi"
	"fmt"
)

func ParseSize(val string) (width, height int) {
	width, height = gumi.DefinedResolutions.Get(val)
	if width == 0 || height == 0{
		fmt.Scanf("(%d, %d)", &width, &height)
	}
	return width, height
}
