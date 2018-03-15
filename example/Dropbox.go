package example

import (
	"fmt"
	"github.com/GUMI-golang/gumi"
)

func Dropbox() gumi.GUMI {
	return gumi.LinkingFrom(
		gumi.NBackground0(gumi.Material.Pallette.BackgroundDrawer()),
		gumi.LCenter0(gumi.MTDropbox3(func(self *gumi.MTDropbox, selected string) {
			fmt.Println("Select : ", selected)
		}, "Hello", "World", "Good", "Bye")),
	)
}
