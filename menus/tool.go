package menus

import (
	"adb/cmd"
	"adb/devices"
	"adb/gifted"
	"adb/tool"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/disintegration/gift"
	"github.com/disintegration/imageorient"
)

var (
	imageCanvas      *canvas.Image
	imageFile        string
	originalImage    image.Image
	editedImage      *image.RGBA
	imgGifteds       *gifted.Gifted
	imgContainer     *fyne.Container
	imgRedoBtn       *widget.Button
	imgUndoBtn       *widget.Button
	imgHorizontalBtn *widget.Button
	imgVerticaBtn    *widget.Button
	imgSaveBtn       *widget.Button
	imgDeleteBtn     *widget.Button

	videoFile string
)

func ToolView(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("1. 屏幕截图", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("2. 录制屏幕", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	))
}

func refreshImage(w fyne.Window) {
	file, err := os.Open(imageFile)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}

	defer file.Close()
	originalImage, _, err = imageorient.Decode(file)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}

	imageCanvas.Image = originalImage
	imageCanvas.Refresh()

	imgContainer.Objects = []fyne.CanvasObject{imageCanvas}
	imgContainer.Refresh()

	imgRedoBtn.Enable()
	imgUndoBtn.Enable()
	imgHorizontalBtn.Enable()
	imgVerticaBtn.Enable()
	imgSaveBtn.Enable()
	imgDeleteBtn.Enable()
}

func ScreencapView(w fyne.Window) fyne.CanvasObject {
	imgContainer = container.NewMax(container.NewCenter(widget.NewLabel("暂无截图展示")))
	imageCanvas = &canvas.Image{}
	imageCanvas.FillMode = canvas.ImageFillContain
	imgGifteds = &gifted.Gifted{}
	imgGifteds.GIFT = gift.New()

	oldBtn := newBtn("老设备", func() {
		now := time.Now()
		sdPath := fmt.Sprintf("/sdcard/%s.png", now.Format("20060102150405"))
		imageFile = fmt.Sprintf("image/%s.png", now.Format("20060102150405"))
		if runCmdDevice("shell", "screencap", "-p", sdPath) == nil {
			runCmdDevice("pull", sdPath, imageFile)
			runCmdDevice("shell", "rm", "-f", sdPath)
			refreshImage(w)
		}
	})

	newBtn := newBtn("新设备", func() {
		now := time.Now()
		imageFile = fmt.Sprintf("image/%s.png", now.Format("20060102150405"))
		runCmdDevice("exec-out", "screencap", "-p", ">", imageFile)
		refreshImage(w)
	})

	topContainer := container.NewHBox(layout.NewSpacer(), oldBtn, newBtn, layout.NewSpacer())
	mainContainer := container.NewBorder(container.NewVBox(topContainer), loadStatusBar(w), nil, nil, imgContainer)
	return container.NewMax(mainContainer)
}

func loadStatusBar(w fyne.Window) *fyne.Container {
	imgRedoBtn = widget.NewButtonWithIcon("", theme.ContentRedoIcon(), func() {
		addParameter(gift.Rotate(-90, color.Transparent, gift.LinearInterpolation))
	})
	imgUndoBtn = widget.NewButtonWithIcon("", theme.ContentUndoIcon(), func() {
		addParameter(gift.Rotate90())
	})
	imgHorizontalBtn = widget.NewButtonWithIcon("", theme.MoreHorizontalIcon(), func() {
		addParameter(gift.FlipHorizontal())
	})
	imgVerticaBtn = widget.NewButtonWithIcon("", theme.MoreVerticalIcon(), func() {
		addParameter(gift.FlipVertical())
	})
	imgSaveBtn = widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		saveImageDialog(w)
	})

	imgDeleteBtn = widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		dialog.NewConfirm("提示", "确定删除该张图片吗？", func(b bool) {
			if b {
				err := os.Remove(imageFile)
				if err != nil {
					dialog.NewError(err, w).Show()
				} else {
					imgDisable()
					imgContainer.Objects = []fyne.CanvasObject{container.NewCenter(widget.NewLabel("暂无截图展示"))}
					imgContainer.Refresh()
				}
			}
		}, w).Show()
	})

	openBtn := widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {
		openFile("image")
	})
	imgDisable()
	statusBar := container.NewVBox(
		widget.NewSeparator(),
		container.NewHBox(
			layout.NewSpacer(),
			imgRedoBtn,
			imgUndoBtn,
			imgHorizontalBtn,
			imgVerticaBtn,
			imgSaveBtn,
			imgDeleteBtn,
			openBtn,
		),
	)
	return statusBar
}

func imgDisable() {
	imgRedoBtn.Disable()
	imgUndoBtn.Disable()
	imgHorizontalBtn.Disable()
	imgVerticaBtn.Disable()
	imgSaveBtn.Disable()
	imgDeleteBtn.Disable()
}

func saveImageDialog(w fyne.Window) {
	if originalImage == nil {
		dialog.ShowError(errors.New("no image opened"), w)
		return
	}
	if editedImage == nil {
		apply()
	}

	ex, err := os.Executable()
	if tool.CheckErr(w, err) {
		fileSaveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, _ error) {
			err := saveImage(writer)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
		}, w)
		exPath := filepath.Dir(ex)
		fileSaveDialog.Resize(fyne.NewSize(640, 460))
		luri, _ := storage.ListerForURI(storage.NewFileURI(exPath + `\`))
		fileSaveDialog.SetLocation(luri)
		fileSaveDialog.SetFileName(filepath.Base(imageFile))
		fileSaveDialog.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpeg", ".jpg", ".gif"}))
		fileSaveDialog.Show()
	}
}

func saveImage(writer fyne.URIWriteCloser) error {
	if writer == nil {
		return nil
	}
	switch writer.URI().Extension() {
	case ".jpeg":
		jpeg.Encode(writer, editedImage, nil)
	case ".jpg":
		jpeg.Encode(writer, editedImage, nil)
	case ".png":
		png.Encode(writer, editedImage)
	case ".gif":
		gif.Encode(writer, editedImage, nil)
	default:
		os.Remove(writer.URI().String()[7:])
		return errors.New("unsupported file extension\n supported extensions: .jpg, .png, .gif")
	}
	return nil
}

func addParameter(filter gift.Filter) {
	if originalImage == nil {
		return
	}
	imgGifteds.Add(filter)
	go apply()
}

func apply() {
	editedImage = image.NewRGBA(imgGifteds.Bounds(originalImage.Bounds()))
	imgGifteds.Draw(editedImage, originalImage)

	imageCanvas.Image = editedImage
	imageCanvas.Refresh()
}

func ScreenrecordView(w fyne.Window) fyne.CanvasObject {
	var (
		videoCmd  *cmd.Command
		recordBtn *widget.Button
		stopBtn   *widget.Button
		sdPath    string
		isScrcpy  bool
	)

	labelProgress := widget.NewLabel("")
	labelProgress.Hide()
	progress := widget.NewProgressBarInfinite()
	progress.Hide()

	recordBtn = widget.NewButton("录制", nil)
	recordBtn.OnTapped = func() {
		if checkDevice() {
			recordBtn.Disable()
			stopBtn.Enable()
			progress.Show()
			ticker := time.NewTicker(time.Second * 1)
			startTime := time.Now()

			go func() {
				labelProgress.SetText("当前已录制: 0s")
				labelProgress.Show()
				for range ticker.C {
					timeCost := time.Since(startTime).Seconds()
					labelProgress.SetText(fmt.Sprintf("当前已录制: %.0f秒", timeCost))
				}
			}()

			go func() {
				defer func() {
					labelProgress.Hide()
					progress.Hide()
					ticker.Stop()
					stopBtn.Disable()
					recordBtn.Enable()
				}()
				now := time.Now()
				sdPath = fmt.Sprintf("/sdcard/%s.mp4", now.Format("20060102150405"))
				videoFile = fmt.Sprintf("video/%s.mp4", now.Format("20060102150405"))

				videoCmd = newCmdDevice("shell", "screenrecord", sdPath)
				err := videoCmd.SyncRun()
				if err != nil && err.Error() == "exit status 127" {
					videoCmd = cmd.NewScrcpy("-s", devices.Device, "--always-on-top", "--record", videoFile)
					if videoCmd != nil {
						isScrcpy = true
						videoCmd.SyncRun()
					}
				}
			}()
		}
	}

	stopBtn = widget.NewButton("停止", nil)
	stopBtn.OnTapped = func() {
		if checkDevice() {
			err := videoCmd.Quit()
			if err == nil && !isScrcpy {
				runCmdDevice("pull", sdPath, videoFile)
				runCmdDevice("shell", "rm", "-f", sdPath)
			}
			if err != nil {
				dialog.ShowError(err, w)
			}
			stopBtn.Disable()
		}
		labelProgress.Hide()
		progress.Hide()
	}
	stopBtn.Disable()

	openBtn := widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {
		openFile("video")
	})

	topContainer := container.NewHBox(layout.NewSpacer(), recordBtn, stopBtn, layout.NewSpacer(), openBtn)
	mainContainer := container.NewBorder(topContainer, container.NewVBox(progress), nil, nil, container.NewCenter(container.NewVBox(labelProgress)))
	return container.NewMax(mainContainer)
}
