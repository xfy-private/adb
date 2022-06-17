package menus

import (
	"adb/cmd"
	"adb/devices"
	"fmt"
	"os/exec"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func checkDevice() bool {
	if devices.Device == "" {
		fyne.CurrentApp().SendNotification(&fyne.Notification{
			Title:   "警告",
			Content: "没有选择操作设备",
		})
		return false
	}
	return true
}

func newBtn(label string, cb func()) *widget.Button {
	btn := widget.NewButton(label, nil)
	btn.OnTapped = func() {
		btn.Disable()
		if devices.Device == "" {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "警告",
				Content: "没有选择操作设备",
			})
		} else {
			cb()
		}
		btn.Enable()
	}
	return btn
}

func runCmdDevice(arg ...string) error {
	args := append([]string{"-s", devices.Device}, arg...)
	return cmd.New(args...).SyncRun()
}

func newCmdDevice(arg ...string) *cmd.Command {
	args := append([]string{"-s", devices.Device}, arg...)
	return cmd.New(args...)
}

func openFile(path string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", path).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", path).Start()
	case "darwin":
		err = exec.Command("open", path).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}
