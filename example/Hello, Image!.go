package example

import (
	"github.com/GUMI-golang/gumi"
	"image/jpeg"
	"bytes"
	"github.com/GUMI-golang/example/asset"
	"github.com/GUMI-golang/gumi/media"
)


func HelloImage() gumi.GUMI {
	img, err := jpeg.Decode(bytes.NewBuffer(asset.MustAsset("helloImage.jpg")))
	if err != nil {
		panic(err)
	}
	return gumi.LinkingFrom(
		gumi.NBackground0(gumi.Material.Pallette.BackgroundDrawer()),
		gumi.Tool.MarginMinRegular(10,
			gumi.AImage0(media.NewFillup(img, media.FillupNearest)),
		),
	)
}