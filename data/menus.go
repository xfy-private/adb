package data

import (
	"adb/menus"

	"fyne.io/fyne/v2"
)

type MenuStu struct {
	Id       string
	ParentId string
	View     func(w fyne.Window) fyne.CanvasObject
}

var (
	ParentMenus = []string{"basic", "devices", "tool", "logcat", "app", "file", "interact", "netstat"}
	Menus       = map[string]MenuStu{
		"info": {"设备信息", "root", menus.InfoView},

		"input": {"模拟按键/输入", "root", menus.InputView},

		"tool":         {"实用功能", "root", menus.ToolView},
		"screencap":    {"屏幕截图", "tool", menus.ScreencapView},
		"screenrecord": {"录制屏幕", "tool", menus.ScreenrecordView},

		"logcat":        {"日志管理", "root", menus.LogcatView},
		"config-logcat": {"日志配置", "logcat", menus.ConfigLogcatView},
		"clear-logcat":  {"清空日志", "logcat", menus.ClearLogcat},
		"show-logcat":   {"显示日志", "logcat", menus.ShowLogcatView},
		"file-logcat":   {"输出到文件", "logcat", menus.FileLogcatView},
		"buffer-logcat": {"缓冲区查看", "logcat", menus.BufferLogcatView},
		"dmesg-logcat":  {"内核日志", "logcat", menus.DmesgLogcatView},

		"app":           {"应用管理", "root", menus.AppView},
		"list-app":      {"应用列表", "app", menus.AppListView},
		"install-app":   {"安装应用", "app", menus.InstallAppView},
		"uninstall-app": {"卸载应用", "app", menus.UninstallAppView},
		"clear-app":     {"清除数据", "app", menus.ClearAppView},
		"services-app":  {"查看Services", "app", menus.ServicesAppView},
		"info-app":      {"应用详细信息", "app", menus.InfoAppView},
		"path-app":      {"应用安装路径", "app", menus.PathAppView},

		"file":       {"文件管理", "root", menus.FileView},
		"pull-file":  {"下载文件", "file", menus.PullFileView},
		"push-file":  {"上传文件", "file", menus.PushFileView},
		"ls-file":    {"列出目录内容", "file", menus.LsFileView},
		"cd-file":    {"切换目录", "file", menus.CdFileView},
		"rm-file":    {"删除文件", "file", menus.RmFileView},
		"mkdir-file": {"创建目录", "file", menus.MkdirFileView},
		"touch-file": {"创建空文件", "file", menus.TouchFileView},
		"cp-file":    {"复制文件", "file", menus.CpFileView},
		"mv-file":    {"移动文件", "file", menus.MvFileView},

		"interact":         {"应用交互", "root", menus.InteractView},
		"start":            {"启动设备应用", "interact", menus.StartView},
		"startservice":     {"调起Service", "interact", menus.StartserviceView},
		"stopservice":      {"停止Service", "interact", menus.StopserviceView},
		"broadcast":        {"发送广播", "interact", menus.BroadcastView},
		"force-stop":       {"强制停止应用", "interact", menus.ForceStopView},
		"send-trim-memory": {"收紧内存", "interact", menus.SendTrimMemoryView},

		"netstat":        {"网络管理", "root", menus.NetstatView},
		"info-netstat":   {"查看网络信息", "netstat", menus.InfoNetstatView},
		"ping-netstat":   {"测试网络", "netstat", menus.PingNetstatView},
		"netcfg-netstat": {"网络配置", "netstat", menus.NetcfgNetstatView},
		"ip-netstat":     {"网络操作", "netstat", menus.IpNetstatView},
	}
	MenuIndex = map[string][]string{
		"":         {"info", "input", "tool", "logcat", "app", "file", "interact", "netstat"},
		"app":      {"list-app", "install-app", "uninstall-app", "clear-app", "services-app", "info-app", "path-app"},
		"interact": {"start", "startservice", "stopservice", "broadcast", "force-stop", "send-trim-memory"},
		"file":     {"pull-file", "push-file", "ls-file", "cd-file", "rm-file", "mkdir-file", "touch-file", "cp-file", "mv-file"},
		"netstat":  {"info-netstat", "ping-netstat", "netcfg-netstat", "ip-netstat"},
		"logcat":   {"config-logcat", "clear-logcat", "show-logcat", "file-logcat", "buffer-logcat", "dmesg-logcat"},
		"tool":     {"screencap", "screenrecord"},
	}
)
