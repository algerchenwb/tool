package rlog

import (
	"sync"

	"github.com/mattn/go-colorable"
)

var rCfg defaultConfig                      //基本配置
var rLogFile logConfig                      //日志文件地址配置
var curPool int64                           //当前还剩多少日志未处理
var output = colorable.NewColorableStdout() //默认日志输出设备

func init() {
	rCfg.LogLevelConsole = LogLevelDebug
	rCfg.LogLevelFile = LogLevelDebug
	rCfg.CropPathLength = 0
	rCfg.PrintCodeFilePath = true
	rCfg.AutoRemoveFile = 0
	rLogFile.FileMaxSize = 1 * 1024 * 1024 // 1MB
	rLogFile.lock = sync.Mutex{}
}
