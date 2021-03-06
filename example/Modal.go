package example

import (
	"github.com/GUMI-golang/example/asset"
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gumi/media"
)

func Modal() gumi.GUMI {
	var modal *gumi.ALModal
	modal = gumi.ALModal1(
		gumi.LinkingFrom(
			gumi.NBackground0(gumi.Material.Pallette.SilluetDrawer()),
			gumi.LCenter0(gumi.MTButton0("Hello, Modal!", func(self *gumi.MTButton) {
				modal.SetShow(false)
			})),
		),
	)
	return gumi.LinkingFrom(
		gumi.NBackground0(gumi.Material.Pallette.BackgroundDrawer()),
		modal,
		gumi.NEventer1(func(self *gumi.NEventer) {
			modal.SetShow(true)
		}),
		gumi.Tool.MarginMinRegular(10,
			gumi.AImage0(media.NewFillup(
				media.MustImageDecode(asset.MustAsset("fireworks-2585843_1920.jpg")),
				media.FillupNearest,
			)),
		),
	)
}
