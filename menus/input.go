package menus

import (
	"adb/devices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func InputView(w fyne.Window) fyne.CanvasObject {
	text := widget.NewEntry()
	text.SetPlaceHolder("请输入要发送的内容")
	textContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(320, 36)), text)

	sendBtn := newBtn("发送", func() {
		runCmdDevice("shell", "input", "text", text.Text)
	})

	clearBtn := newBtn("删除", func() {
		runCmdDevice("shell", "input", "keyevent", "--longpress", "KEYCODE_DEL")
	})

	topContainer := container.NewHBox(layout.NewSpacer(), widget.NewLabel("输入文本:"), textContainer, sendBtn, clearBtn, layout.NewSpacer())
	btn1 := newBtn("电源键", func() {
		runCmdDevice("shell", "input", "keyevent", "26")
	})
	btn2 := newBtn("返回键", func() {
		runCmdDevice("shell", "input", "keyevent", "4")
	})
	btn3 := newBtn("HOME键", func() {
		runCmdDevice("shell", "input", "keyevent", "3")
	})
	btn30 := newBtn("上滑解锁", func() {
		runCmdDevice("shell", "input", "swipe", "300", "1000", "300", " 500")
	})
	btn6 := newBtn("重启手机", func() {
		dialog.NewConfirm("提示", "这将会使设备断开连接，确定要重启手机吗？", func(b bool) {
			if b {
				runCmdDevice("reboot")
				devices.Resat()
			}
		}, w).Show()
	})
	btn8 := newBtn("增加音量", func() {
		runCmdDevice("shell", "input", "keyevent", "24")
	})
	btn9 := newBtn("降低音量", func() {
		runCmdDevice("shell", "input", "keyevent", "25")
	})
	btn13 := newBtn("播放/暂停", func() {
		runCmdDevice("shell", "input", "keyevent", "85")
	})
	btn15 := newBtn("播放下一首", func() {
		runCmdDevice("shell", "input", "keyevent", "87")
	})
	btn16 := newBtn("播放上一首", func() {
		runCmdDevice("shell", "input", "keyevent", "88")
	})
	btn17 := newBtn("行首或列表顶部", func() {
		runCmdDevice("shell", "input", "keyevent", "122")
	})
	btn18 := newBtn("行末或列表底部", func() {
		runCmdDevice("shell", "input", "keyevent", "123")
	})
	btn19 := newBtn("静音", func() {
		runCmdDevice("shell", "input", "keyevent", "164")
	})
	btn20 := newBtn("系统设置", func() {
		runCmdDevice("shell", "am", "start", "com.android.settings/com.android.settings.Settings")
	})
	btn21 := newBtn("切换应用", func() {
		runCmdDevice("shell", "input", "keyevent", "187")
	})
	btn26 := newBtn("降低亮度", func() {
		runCmdDevice("shell", "input", "keyevent", "220")
	})
	btn27 := newBtn("提高亮度", func() {
		runCmdDevice("shell", "input", "keyevent", "221")
	})
	btn28 := newBtn("点亮屏幕", func() {
		runCmdDevice("shell", "input", "keyevent", "224")
	})
	btn29 := newBtn("熄灭屏幕", func() {
		runCmdDevice("shell", "input", "keyevent", "223")
	})

	mainBtn := container.New(layout.NewGridLayout(4), btn6, btn1, btn2, btn3, btn28, btn29, btn30, btn26, btn27, btn20,
		btn8, btn9, btn13, btn15, btn16, btn17, btn18, btn19, btn21)
	mainContainer := container.NewBorder(container.NewVBox(topContainer), nil, nil, nil, mainBtn)
	return container.NewMax(mainContainer)
}
