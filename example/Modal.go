package example

import (
	"github.com/GUMI-golang/gumi"
	"image/jpeg"
	"bytes"
	"github.com/GUMI-golang/example/asset"
	"github.com/GUMI-golang/gumi/media"
)

func Modal() gumi.GUMI {
	img, err := jpeg.Decode(bytes.NewBuffer(asset.MustAsset("fireworks-2585843_1920.jpg")))
	if err != nil {
		panic(err)
	}
	var modal *gumi.ALModal
	modal = gumi.ALModal1(
		gumi.LinkingFrom(

			gumi.LCenter0(gumi.MTButton0("Hello, Modal!", func(self *gumi.MTButton) {
				modal.SetShow(false)
			})),
		),
	)
	return gumi.LinkingFrom(
		gumi.NBackground0(gumi.Material.Pallette.BackgroundDrawer()),
		modal,
		gumi.NClicker0(func(self *gumi.NClicker) {
			modal.SetShow(true)
		}),
		gumi.Tool.MarginMinRegular(10,
			gumi.AImage0(media.NewFillup(img, media.FillupNearest)),
		),
	)
}