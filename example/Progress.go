package example

import (
	"github.com/GUMI-golang/gumi"
	"time"
	"github.com/GUMI-golang/gumi/gcore"
)

func Progress() gumi.GUMI {
	var progress = gumi.MTProgress0(gumi.Material.Pallette.White)
	return gumi.LinkingFrom(
		gumi.NBackground0(gumi.Material.Pallette.BackgroundDrawer()),
		gumi.NEventer3(nil, nil, time.Second, func(self *gumi.NEventer, t time.Time) {
			if progress.Get() <= 0{
				progress.Set(1)
			}else if progress.Get() >= 1 {
				progress.Set(0)
			}
		}),
		gumi.LCenter0(
			gumi.LinkingFrom(
				gumi.NSize0(gcore.Size{
					Horizontal: gcore.MinLength(400),
					Vertical:gcore.AUTOLENGTH,
				}),
				progress,
			),
		),
	)
}
