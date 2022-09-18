package rlog

import (
	"fmt"
	"path"
	"sync/atomic"
	"time"
)

// WaitFinish 等待完成，一般用于程序终止前，打印完所有日志再结束程序。
func WaitFinish() {
	for {
		if 0 >= atomic.LoadInt64(&curPool) {
			time.Sleep(time.Millisecond * 10)
			return
		}
		time.Sleep(time.Millisecond * 10)
	}
}

// SetPrintCodeFilePath 设置是否打印代码文件路径，默认打印
func SetPrintCodeFilePath(printCode bool) {
	rCfg.PrintCodeFilePath = printCode
}

// SetCropPathLength 设置去除日志代码文件路径前面多少个字符
func SetCropPathLength(length int) {
	rCfg.CropPathLength = uint(length)
}

// SetLogLevel 设置基本配置信息
// Console：打印到控制台的日志等级
// File：打印到文件的日志等级
func SetLogLevel(Console LogLeve, File LogLeve) {
	rCfg.LogLevelConsole = Console
	rCfg.LogLevelFile = File
	go writeFile()
}

var setRemoveOK bool //是否已经设置过自动删除旧日志文件

// SetAutoRemoveLogFile 设置自动删除多少天以前的日志
func SetAutoRemoveLogFile(day uint) {
	if setRemoveOK {
		fmt.Println("您已设置过 rlog.SetAutoRemoveLogFile 了，不能进行第二次设置。")
		return
	}
	rCfg.AutoRemoveFile = day
	setRemoveOK = true
	go autoRemoveFile()
}

// SetLogFile 设置日志文件路径
func SetLogFile(logDir string, fileMaxSize int64, isOneFile bool) {
	if rLogFile.setOK {
		fmt.Println("您已设置过 rlog.SetLogFile 了，不能进行第二次设置。")
		return
	}
	rLogFile.setOK = true
	rLogFile.lock.Lock()
	defer rLogFile.lock.Unlock()
	rLogFile.FileMaxSize = fileMaxSize
	rLogFile.IsOneFile = isOneFile
	rLogFile.LogDir = logDir
	if rLogFile.LogDir != "" {
		if rLogFile.IsOneFile {
			rLogFile.logFile, rLogFile.logSize = autoCreateFile(rLogFile.logFile, path.Join(rLogFile.LogDir, "app.log"))
		} else {
			rLogFile.debugFile, rLogFile.debugSize = autoCreateFile(rLogFile.debugFile, path.Join(rLogFile.LogDir, "debug", "debug.log"))
			rLogFile.infoFile, rLogFile.infoSize = autoCreateFile(rLogFile.infoFile, path.Join(rLogFile.LogDir, "info", "info.log"))
			rLogFile.successFile, rLogFile.successSize = autoCreateFile(rLogFile.successFile, path.Join(rLogFile.LogDir, "success", "success.log"))
			rLogFile.warnFile, rLogFile.warnSize = autoCreateFile(rLogFile.warnFile, path.Join(rLogFile.LogDir, "warn", "warn.log"))
			rLogFile.errorFile, rLogFile.errorSize = autoCreateFile(rLogFile.errorFile, path.Join(rLogFile.LogDir, "error", "error.log"))
			rLogFile.panicFile, rLogFile.panicSize = autoCreateFile(rLogFile.panicFile, path.Join(rLogFile.LogDir, "panic", "panic.log"))
		}
	}
}
