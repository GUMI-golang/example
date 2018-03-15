package example

import (
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gumi/gcore"
	"fmt"
)

func Editing() gumi.GUMI {
	ed := gumi.MTEdit0()
	ed.OnChange(func(self *gumi.MTEdit, text string) {
		fmt.Println(text)
	})

	return gumi.LinkingFrom(
		gumi.NBackground0(gumi.Material.Pallette.BackgroundDrawer()),
		gumi.LCenter0(
			gumi.LinkingFrom(
				gumi.NSize0(gcore.Size{
					Horizontal:gcore.MinLength(300),
					Vertical:gcore.AUTOLENGTH,
				}),
				ed,
			),
		),
	)
}