package menus

import (
	"adb/devices"
	"regexp"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InfoView(_ fyne.Window) fyne.CanvasObject {
	if devices.Device == "" {
		return emptyDevice()
	}
	brandLabel := widget.NewLabelWithStyle("品牌: ", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	nameLabel := widget.NewLabelWithStyle("名称: ", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	modelLabel := widget.NewLabelWithStyle("型号: ", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	releaseLabel := widget.NewLabelWithStyle("Android版本: ", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	sizeLabel := widget.NewLabelWithStyle("分辨率: ", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	ipLabel := widget.NewLabelWithStyle("IP: ", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	go func() {
		brandCmd := newCmdDevice("shell", "getprop", "ro.product.brand")
		brand := brandCmd.SyncReadString()
		nameCmd := newCmdDevice("shell", "getprop", "ro.product.name")
		name := nameCmd.SyncReadString()
		modelCmd := newCmdDevice("shell", "getprop", "ro.product.model")
		model := modelCmd.SyncReadString()
		releaseCmd := newCmdDevice("shell", "getprop", "ro.build.version.release")
		release := releaseCmd.SyncReadString()
		sizeCmd := newCmdDevice("shell", "wm", "size")
		size := sizeCmd.SyncReadString()
		ipCmd := newCmdDevice("shell", "ip", "addr", "show ", "wlan0")
		ip := ipCmd.SyncReadString()
		ipReg := regexp.MustCompile(`inet(.*?)brd`)

		brandLabel.SetText(brandLabel.Text + strings.ReplaceAll(brand, "\n", ""))
		nameLabel.SetText(nameLabel.Text + strings.ReplaceAll(name, "\n", ""))
		modelLabel.SetText(modelLabel.Text + strings.ReplaceAll(model, "\n", ""))
		releaseLabel.SetText(releaseLabel.Text + strings.ReplaceAll(release, "\n", ""))
		sizeLabel.SetText(sizeLabel.Text + strings.ReplaceAll(strings.ReplaceAll(size, "Physical size: ", ""), "\n", ""))
		ipLabel.SetText(ipLabel.Text + strings.ReplaceAll(strings.ReplaceAll(ipReg.FindString(ip), "inet ", ""), "brd", ""))
	}()
	return container.NewCenter(container.NewVBox(brandLabel, nameLabel, modelLabel, releaseLabel, sizeLabel, ipLabel))
}
