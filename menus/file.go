package menus

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func FileView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("FileView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func PullFileView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("PullFileView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func PushFileView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("PushFileView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func LsFileView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("LsFileView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func CdFileView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("CdFileView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func RmFileView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("RmFileView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func MkdirFileView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("MkdirFileView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func TouchFileView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("TouchFileView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func CpFileView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("CpFileView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func MvFileView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("MvFileView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
