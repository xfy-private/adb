package devices

import (
	"adb/cmd"
	"adb/tool"
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	ipList      = []string{}
	selects     *widget.Select
	connectWin  fyne.Window
	isClose     bool = false
	adbCmd      *cmd.Command
	connectDone bool = false
	Device      string
	deviceMap   map[string]string = make(map[string]string)
)

func connect() {
	a := fyne.CurrentApp()

	connectWin = a.NewWindow("连接设备")
	ticker := time.NewTicker(time.Second * 1)
	startTime := time.Now()

	connectWin.SetCloseIntercept(func() {
		dialog.NewConfirm("提示", "有正在运行的命令，确定强制退出程序吗？", func(b bool) {
			if b {
				if connectDone {
					connectWin.Close()
					ticker.Stop()
					return
				}
				err := adbCmd.Quit()
				if err != nil {
					dialog.NewError(err, connectWin)
				} else {
					isClose = true
					ticker.Stop()
					connectWin.Close()
				}
			}
		}, connectWin).Show()
	})

	go func() {
		for range ticker.C {
			timeCost := time.Since(startTime).Seconds()
			connectWin.SetTitle(fmt.Sprintf("连接设备 %.0f秒", timeCost))
		}
	}()

	connectWin.SetFixedSize(true)
	connectWin.CenterOnScreen()
	connectWin.Resize(fyne.NewSize(500, 320))
	connectWin.Show()

	logEntry := widget.NewMultiLineEntry()
	logEntry.Wrapping = fyne.TextWrapWord
	logEntry.Disable()
	connectWin.SetContent(logEntry)
	logEntry.SetText("开始连接设备......\n")
	for _, ip := range ipList {
		if !isClose {
			adbCmd = cmd.New("connect", ip)
			log := adbCmd.SyncReadString()
			logEntry.SetText(log + "\n" + logEntry.Text)
		} else {
			return
		}
	}
	connectDone = true
	ticker.Stop()
	logEntry.SetText("设备连接结束......" + "\n" + logEntry.Text)
}

func IpOpened(uc fyne.URIReadCloser) {
	defer uc.Close()
	ipList = []string{}
	br := bufio.NewReader(uc)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		ip := tool.Bytes2Str(line)
		if tool.CheckStoreAddr(ip) == nil {
			ipList = append(ipList, ip)
		}
	}
	if len(ipList) != 0 {
		connect()
	}
}

func Select() *widget.Select {
	selects = widget.NewSelect([]string{}, func(device string) {
		Device = deviceMap[device]
		selects.Disable()
		selects.Refresh()
	})
	selects.PlaceHolder = "请选择在线设备操作"
	selects.Disable()
	return selects
}

func refresh(btn *widget.Button) {
	selects.ClearSelected()
	deviceCmd := cmd.New("devices")
	devices := deviceCmd.SyncReadString()

	devicesList := []string{}
	for index, device := range strings.Split(devices, "\n") {
		if index != 0 && device != "" && strings.Contains(device, "device") {
			device = strings.ReplaceAll(device, "	device", "")
			nameCmd := cmd.New("-s", device, "shell", "getprop", "ro.product.name")
			name := nameCmd.SyncReadString()
			brandCmd := cmd.New("-s", device, "shell", "getprop", "ro.product.brand")
			brand := brandCmd.SyncReadString()

			key := fmt.Sprintf("%s%s", strings.ReplaceAll(brand, "\n", ""), strings.ReplaceAll(name, "\n", ""))
			deviceMap[key] = device
			devicesList = append(devicesList, key)
		}
	}
	if len(devicesList) > 0 {
		selects.Options = devicesList
		selects.Enable()
		selects.Refresh()
	}
	if btn != nil {
		btn.Enable()
	}
}

func Refresh(w fyne.Window, btn *widget.Button) {
	btn.Disable()
	dialog.NewConfirm("提示", "已选择设备正在执行命令，确定刷新设备列表吗？", func(b bool) {
		if b {
			refresh(btn)
		} else {
			btn.Enable()
		}
	}, w).Show()
}

func Unlock(w fyne.Window, btn *widget.Button) {
	btn.Disable()
	if Device == "" {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "警告",
			Content: "没有选择操作设备",
		})
	} else {
		err := cmd.New("-s", Device, "shell", "input", "keyevent", "224").SyncRun()
		if err == nil {
			time.Sleep(100 * time.Millisecond)
			cmd.New("-s", Device, "shell", "input", "swipe", "300", "1000", "300", " 500").SyncRun()
		}
	}
	btn.Enable()
}

func Resat() {
	refresh(nil)
}
