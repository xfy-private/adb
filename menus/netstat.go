package menus

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NetstatView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("NetstatView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func InfoNetstatView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("InfoNetstatView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func PingNetstatView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("PingNetstatView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func NetcfgNetstatView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("NetcfgNetstatView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func IpNetstatView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("IpNetstatView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
