package example

import (
	"fmt"
	"github.com/GUMI-golang/gumi"
)

func HelloButton() gumi.GUMI {
	return gumi.LinkingFrom(
		gumi.NBackground0(gumi.Material.Pallette.BackgroundDrawer()),
		gumi.LCenter0(
			gumi.MTButton0("Hello, Button!", func(self *gumi.MTButton) {
				fmt.Println("Hello, Button!")
			}),
		),
	)
}
