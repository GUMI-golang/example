package example

import "github.com/GUMI-golang/gumi"

func HelloWorld() gumi.GUMI {
	return  gumi.LinkingFrom(
		//gumi.NDrawing0(gumi.Drawing.FPS()),
		gumi.NBackground0(gumi.Material.Pallette.BackgroundDrawer()),
		gumi.LCenter0(
			gumi.AText0("Hello, World!"),
		),
	)
}