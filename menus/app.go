package menus

import (
	"adb/devices"
	"adb/tool"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func AppView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("1. 应用列表", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("2. 安装应用", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}
func AppListView(w fyne.Window) fyne.CanvasObject {
	args := []string{"shell", "pm", "list", "packages", "", "", "", ""}

	packNameLabel := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	packVersionLabel := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	infoContainer := container.NewCenter(container.NewVBox(
		packNameLabel,
		packVersionLabel,
	))

	infoContainer.Hide()

	unBtn := widget.NewButton("卸载", func() {
		if packNameLabel.Text != "" {
			packName := strings.ReplaceAll(packNameLabel.Text, "包名: ", "")
			keepCache := false
			dialog.ShowCustomConfirm("卸载", "确定", "取消",
				container.NewVBox(widget.NewLabel(fmt.Sprintf("确定卸载包: %s 吗?", packName)),
					widget.NewCheck("清除数据和缓存", func(b bool) {
						keepCache = b
					})), func(b bool) {
					if b {
						cmds := tool.IfThree(keepCache, []string{"uninstall", "-k", packName}, []string{"uninstall", packName})
						if runCmdDevice(cmds...) == nil {
							fyne.CurrentApp().SendNotification(&fyne.Notification{
								Title:   "提示",
								Content: fmt.Sprintf("%s 卸载成功!", packName),
							})
						}
					}
				}, w)
		}
	})
	clearBtn := widget.NewButton("清除", func() {
		if packNameLabel.Text != "" {
			packName := strings.ReplaceAll(packNameLabel.Text, "包名: ", "")
			dialog.ShowConfirm("提示", fmt.Sprintf("确定清除: %s 安数据与缓存吗?", packName), func(b bool) {
				if runCmdDevice("shell", "pm", "clear", packName) == nil {
					fyne.CurrentApp().SendNotification(&fyne.Notification{
						Title:   "提示",
						Content: fmt.Sprintf("%s 数据缓存清除成功!", packName),
					})
				}
			}, w)
		}
	})

	barBtnContainers := container.NewHBox(
		layout.NewSpacer(), unBtn, clearBtn, layout.NewSpacer(),
	)

	rightContainers := container.NewBorder(barBtnContainers, nil, nil, nil, infoContainer)

	packList := binding.BindStringList(nil)

	dataList := widget.NewListWithData(packList,
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(item binding.DataItem, obj fyne.CanvasObject) {
			nameData := item.(binding.String)
			text := obj.(*widget.Label)
			name, _ := nameData.Get()
			text.SetText(name)
		})

	searchFunc := func() {
		infoContainer.Hide()
		packNameLabel.SetText("")

		packCmd := newCmdDevice(args...)
		packStr := packCmd.SyncReadString()
		dataList.UnselectAll()
		packList.Set(nil)
		for _, pack := range strings.Split(packStr, "\n") {
			if pack != "" {
				packList.Append(strings.ReplaceAll(pack, "package:", ""))
			}
		}
		if packList.Length() != 0 {
			dataList.Select(0)
		}
	}

	appSideSelect := widget.NewSelect([]string{"系统应用", "第三方"}, func(s string) {
		args[4] = tool.IfThree(s == "系统应用", "-s", "-3")
		if devices.Device != "" {
			searchFunc()
		}
	})
	appSideSelect.PlaceHolder = "选择应用方"

	appStateSelect := widget.NewSelect([]string{"禁用的应用", "启用的应用"}, func(s string) {
		args[5] = tool.IfThree(s == "禁用的应用", "-d", "-e")
		if devices.Device != "" {
			searchFunc()
		}
	})
	appStateSelect.PlaceHolder = "选择应用状态"

	unCheck := widget.NewCheck("已卸载", func(on bool) {
		args[6] = tool.IfThree(on, "-u", "")
	})

	searchText := widget.NewEntry()
	searchText.SetPlaceHolder("输入要过滤的包")
	searchText.OnChanged = func(s string) {
		args[7] = tool.IfThree(s != "", s, "")
	}

	textContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(200, 36)), searchText)

	searchBtn := newBtn("搜索", func() {
		searchFunc()
	})

	topContainer := container.NewHBox(layout.NewSpacer(), appSideSelect, appStateSelect, unCheck, textContainer, searchBtn, layout.NewSpacer())

	dataList.OnSelected = func(id widget.ListItemID) {
		dataItem, err := packList.GetItem(id)
		if err == nil {
			nameData := dataItem.(binding.String)
			packName, err := nameData.Get()
			if err == nil {
				packNameLabel.SetText(fmt.Sprintf("包名: %s", packName))
				packCmd := newCmdDevice("shell", "dumpsys", "package", packName)
				info := packCmd.SyncReadString()
				versions := tool.MatchOne(info, `versionName=((?s).+?)splits`)
				if len(versions) == 2 {
					packVersionLabel.SetText(fmt.Sprintf("版本号: %s", strings.ReplaceAll(versions[1], "\n", "")))
				} else {
					packVersionLabel.SetText("版本号: 未知")
				}
				infoContainer.Show()
			}
		}
	}

	split := container.NewHSplit(dataList, rightContainers)
	split.Offset = 0.5

	return container.NewMax(container.NewBorder(container.NewVBox(topContainer, widget.NewSeparator()), nil, nil, nil, split))
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
