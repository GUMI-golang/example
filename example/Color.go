package example

import (
	"github.com/GUMI-golang/gumi"
)

func Color() gumi.GUMI {
	return gumi.LinkingFrom(
		gumi.NBackground0(gumi.Material.Pallette.BackgroundDrawer()),
		gumi.LCenter0(gumi.MTColor0()),
	)
}
