package example

import (
	"fmt"
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gumi/gcore"
)

func Radios() gumi.GUMI {

	return gumi.LinkingFrom(
		gumi.NBackground0(gumi.Material.Pallette.BackgroundDrawer()),
		gumi.LCenter0(
			gumi.LHorizontal1(
				gumi.MTRadio1(gumi.Material.Pallette.White, gumi.Material.Pallette.White, func(self *gumi.MTRadio, active bool) {
					fmt.Println(self.GetToMaterialColor(), active)
				}),
				gumi.ASpacer1(gcore.MinLength(10)),
				gumi.MTRadio1(gumi.Material.Pallette.White, gumi.Material.Pallette.Red, func(self *gumi.MTRadio, active bool) {
					fmt.Println(self.GetToMaterialColor(), active)
				}),
				gumi.ASpacer1(gcore.MinLength(10)),
				gumi.MTRadio1(gumi.Material.Pallette.White, gumi.Material.Pallette.Green, func(self *gumi.MTRadio, active bool) {
					fmt.Println(self.GetToMaterialColor(), active)
				}),
				gumi.ASpacer1(gcore.MinLength(10)),
				gumi.MTRadio1(gumi.Material.Pallette.White, gumi.Material.Pallette.Blue, func(self *gumi.MTRadio, active bool) {
					fmt.Println(self.GetToMaterialColor(), active)
				}),
				gumi.ASpacer1(gcore.MinLength(10)),
				gumi.MTRadio1(gumi.Material.Pallette.White, gumi.Material.Pallette.Yellow, func(self *gumi.MTRadio, active bool) {
					fmt.Println(self.GetToMaterialColor(), active)
				}),
			),
		),
	)
}
