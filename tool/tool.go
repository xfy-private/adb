package tool

import (
	"errors"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unsafe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"golang.org/x/text/encoding/simplifiedchinese"
)

var Charset = "GB18030"

func CheckErr(w fyne.Window, err error) bool {
	if err != nil {
		dialog.ShowError(err, w)
		return false
	}
	return true
}

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func CheckStoreAddr(addr string) error {
	if addr == "" {
		return errors.New("addr is empty")
	}
	address := net.ParseIP(addr)
	if address == nil {
		addrs := strings.Split(addr, ":")
		if len(addrs) != 2 {
			return errors.New("addr is error")
		}
		ip := addrs[0]
		port := addrs[1]
		address = net.ParseIP(ip)
		if address == nil {
			return errors.New("addr is error")
		}
		if ok, _ := regexp.Match(`/^([1-9](\d{0,3}))$|^([1-5]\d{4})$|^(6[0-4]\d{3})$|^(65[0-4]\d{2})$|^(655[0-2]\d)$|^(6553[0-5])$/`, Str2Bytes(port)); ok {
			return errors.New("addr is error")
		}
	}
	return nil
}

func ShowFileOpen(cb func(uc fyne.URIReadCloser, err error), w fyne.Window) {
	ex, err := os.Executable()
	if err == nil {
		exPath := filepath.Dir(ex)
		fileOpen := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			cb(uc, err)
		}, w)
		fileOpen.Resize(fyne.NewSize(640, 460))
		luri, _ := storage.ListerForURI(storage.NewFileURI(exPath + `\`))
		fileOpen.SetLocation(luri)

		fileOpen.Show()
		fileOpen.Refresh()
		return
	}
	cb(nil, err)
}

func ConvertByte2String(byte []byte) string {
	var str string
	switch Charset {
	case "GB18030":
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case "UTF-8":
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
