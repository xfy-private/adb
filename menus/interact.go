package menus

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InteractView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("InteractView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func StartView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("StartView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func StartserviceView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("StartserviceView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func StopserviceView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("StopserviceView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func BroadcastView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("BroadcastView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func ForceStopView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("ForceStopView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func SendTrimMemoryView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("SendTrimMemoryView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
