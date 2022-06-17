package menus

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func emptyDevice() fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("请选择要操作的设备", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
