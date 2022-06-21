package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"adb/cmd"
	"adb/data"
	"adb/devices"
	"adb/logger"
	"adb/tool"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
)

const selectdPreferencesMenuId = "selectdPreferencesMenuId"
const selectdPreferencesMenuParentId = "selectdPreferencesMenuParentId"
const charsetPreferencesEncoding = "charsetPreferencesEncoding"
const adbPreferencesPath = "adbPreferencesPath"

var charsetMap = map[string]int{
	"UTF-8":   0,
	"GB18030": 1,
}

func init() {
	cmd.AdbPath = filepath.Join("adb-tool", "adb.exe")
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "simkai.ttf") || strings.Contains(path, "simhei.ttf") {
			os.Setenv("FYNE_FONT", path)
			return
		}
	}
	os.Setenv("FYNE_FONT", filepath.Join("simkai.ttf"))
}

func main() {
	os.Setenv("FYNE_THEME", "dark")
	a := app.NewWithID("www.adb.com")
	logo, _ := fyne.LoadResourceFromPath("icon.png")
	a.SetIcon(logo)
	logLifecycle(a)

	mainWin := a.NewWindow("ADB可视化工具")

	mainWin.SetMainMenu(createMenu(mainWin))
	mainWin.SetMaster()
	content := container.NewMax()
	devicesSelect := devices.Select()
	buttonRefresh := widget.NewButton("刷新", nil)
	buttonRefresh.OnTapped = func() {
		devices.Refresh(mainWin, buttonRefresh)
	}

	buttonUnlock := widget.NewButton("解锁", nil)
	buttonUnlock.OnTapped = func() {
		devices.Unlock(mainWin, buttonUnlock)
	}

	topContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewLabel("设备:"), devicesSelect, buttonRefresh, buttonUnlock, layout.NewSpacer())
	rightContainer := container.NewBorder(container.NewVBox(topContainer, widget.NewSeparator()), nil, nil, nil, content)
	split := container.NewHSplit(newNav(mainWin, content), rightContainer)
	split.SetOffset(0.2)
	mainWin.SetContent(split)

	mainWin.SetCloseIntercept(func() {
		dialog.NewConfirm("提示", "确定退出主程序吗？", func(b bool) {
			if b {
				a.Quit()
			}
		}, mainWin).Show()
	})
	mainWin.SetOnClosed(func() {
		os.Unsetenv("FYNE_FONT")
		os.Unsetenv("FYNE_THEME")
	})
	mainWin.CenterOnScreen()
	mainWin.Resize(fyne.NewSize(850, 600))
	mainWin.ShowAndRun()
}

func newNav(w fyne.Window, content *fyne.Container) fyne.CanvasObject {
	a := fyne.CurrentApp()
	tree := &widget.Tree{
		ChildUIDs: func(id string) []string {
			return data.MenuIndex[id]
		},
		IsBranch: func(id string) bool {
			children, ok := data.MenuIndex[id]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return &widget.Label{}
		},
		UpdateNode: func(id string, branch bool, obj fyne.CanvasObject) {
			m, ok := data.Menus[id]
			if !ok {
				fyne.LogError("缺少该页面: "+id, nil)
				return
			}
			obj.(*widget.Label).SetText(m.Id)
		},
		OnSelected: func(id string) {
			if m, ok := data.Menus[id]; ok {
				a.Preferences().SetString(selectdPreferencesMenuId, id)
				a.Preferences().SetString(selectdPreferencesMenuParentId, m.ParentId)
				content.Objects = []fyne.CanvasObject{m.View(w)}
				content.Refresh()
			}
		},
	}
	tree.OnBranchOpened = func(id widget.TreeNodeID) {
		for _, menuId := range data.ParentMenus {
			if menuId != id {
				tree.CloseBranch(menuId)
			}
		}
	}
	selectdId := a.Preferences().StringWithFallback(selectdPreferencesMenuId, "info")
	selectdParentId := a.Preferences().StringWithFallback(selectdPreferencesMenuParentId, "")
	tree.Select(selectdId)
	tree.OpenBranch(selectdParentId)

	return container.NewBorder(nil, nil, nil, nil, tree)
}

func createMenu(w fyne.Window) *fyne.MainMenu {
	a := fyne.CurrentApp()

	adbItem := fyne.NewMenuItem("ADB路径", func() {
		tool.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
			if tool.CheckErr(w, err) {
				if uc != nil {
					defer uc.Close()
					cmd.AdbPath = uc.URI().Path()
					a.Preferences().SetString(adbPreferencesPath, cmd.AdbPath)
				}
			}
		}, w)
	})
	ipItem := fyne.NewMenuItem("设备Ip", func() {
		tool.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
			if tool.CheckErr(w, err) {
				if uc != nil {
					devices.IpOpened(uc)
				}
			}
		}, w)
	})
	logItem := fyne.NewMenuItem("日志路径", func() {
		ex, err := os.Executable()
		if tool.CheckErr(w, err) {
			exPath := filepath.Dir(ex)
			folderOpen := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
				if tool.CheckErr(w, err) {
					if lu != nil {
						logger.LogFilePath = lu.Path()
					}
				}
			}, w)
			folderOpen.Resize(fyne.NewSize(640, 460))
			luri, _ := storage.ListerForURI(storage.NewFileURI(exPath + `\`))
			folderOpen.SetLocation(luri)

			folderOpen.Show()
			folderOpen.Refresh()
		}
	})

	encoding := fyne.NewMenuItem("编码", nil)
	encoding.ChildMenu = fyne.NewMenu("", fyne.NewMenuItem("UTF-8", func() {
		encoding.ChildMenu.Items[0].Checked = true
		encoding.ChildMenu.Items[1].Checked = false
		tool.Charset = "UTF-8"
		a.Preferences().SetString(charsetPreferencesEncoding, "UTF-8")
	}), fyne.NewMenuItem("GB18030", func() {
		encoding.ChildMenu.Items[0].Checked = false
		encoding.ChildMenu.Items[1].Checked = true
		tool.Charset = "GB18030"
		a.Preferences().SetString(charsetPreferencesEncoding, "GB18030")
	}))

	charset := a.Preferences().StringWithFallback(charsetPreferencesEncoding, "UTF-8")
	tool.Charset = charset
	encoding.ChildMenu.Items[charsetMap[charset]].Checked = true

	cmd.AdbPath = a.Preferences().String(adbPreferencesPath)
	os.Setenv("ADB", cmd.AdbPath)

	versionItem := fyne.NewMenuItem("版本", func() {
		adbCmd := cmd.New("version")
		version := adbCmd.SyncReadString()
		dialog.NewInformation("版本信息", version, w).Show()
	})

	file := fyne.NewMenu("文件", adbItem, ipItem, logItem)
	settings := fyne.NewMenu("设置", encoding)
	help := fyne.NewMenu("帮助", versionItem)

	return fyne.NewMainMenu(file, settings, help)
}

func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}
