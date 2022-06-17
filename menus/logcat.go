package menus

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LogcatView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("LogcatView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func ConfigLogcatView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("ConfigLogcatView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func ClearLogcat(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("ClearLogcat", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func ShowLogcatView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("ShowLogcatView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func FileLogcatView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("FileLogcatView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func BufferLogcatView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("BufferLogcatView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func DmesgLogcatView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("DmesgLogcatView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
