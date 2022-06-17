package menus

import (
	"adb/devices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func AppView(_ fyne.Window) fyne.CanvasObject {
	if devices.Device == "" {
		return emptyDevice()
	}
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("AppView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func AppListView(_ fyne.Window) fyne.CanvasObject {
	if devices.Device == "" {
		return emptyDevice()
	}
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("AppListView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func InstallAppView(_ fyne.Window) fyne.CanvasObject {
	if devices.Device == "" {
		return emptyDevice()
	}
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("InstallAppView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func UninstallAppView(_ fyne.Window) fyne.CanvasObject {
	if devices.Device == "" {
		return emptyDevice()
	}
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("UninstallAppView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func ClearAppView(_ fyne.Window) fyne.CanvasObject {
	if devices.Device == "" {
		return emptyDevice()
	}
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("ClearAppView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func ServicesAppView(_ fyne.Window) fyne.CanvasObject {
	if devices.Device == "" {
		return emptyDevice()
	}
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("ServicesAppView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func InfoAppView(_ fyne.Window) fyne.CanvasObject {
	if devices.Device == "" {
		return emptyDevice()
	}
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("InfoAppView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func PathAppView(_ fyne.Window) fyne.CanvasObject {
	if devices.Device == "" {
		return emptyDevice()
	}
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("PathAppView", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
