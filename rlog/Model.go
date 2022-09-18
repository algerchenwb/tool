package rlog

import (
	"os"
	"sync"
)

// 默认的日志格式 [日期] [时间] [时区] [标签] [代码文件路径:行号] [:] [日志内容]
// 比如：2022/09/09 20:22:09 +0800 INFO C:/go/src/go1.19.1/main.go:09 : 日志内容

type LogLeve uint8
type PrintModel uint8

const (
	ModelNoPath      PrintModel = 1 //不打印代码文件路径及行号
	ModelNoTime      PrintModel = 2 //不打印日期、时间、时区
	ModelNoPathTime  PrintModel = 3 //不打印代码文件路径及行号、不打印日期、时间、时区
	ModelMustConsole PrintModel = 4

	LogLevelDebug    LogLeve = 0 //日志等级debug，该等级打印所有级别的日志
	LogLevelInfo     LogLeve = 1 //日志等级info，该等级仅不打印debug日志
	LogLeveSuccess   LogLeve = 2 //日志等级success，该等级仅不打印debug日志
	LogLevelWarn     LogLeve = 3 //日志等级warn，该等级仅不打印debug,info日志
	LogLevelError    LogLeve = 4 //日志等级Error，该等级仅不打印debug,info,warn日志
	LogLevelPanic    LogLeve = 5 //日志等级Panic，该等级仅不打印debug,info,warn,error日志
	LogLevelNotPrint LogLeve = 6 //日志等级最高，该等级不打印任何日志
)

type defaultConfig struct {
	LogLevelConsole   LogLeve //打印到控制台的日志等级
	LogLevelFile      LogLeve //打印到文件的日志等级
	CropPathLength    uint    //去除日志代码文件路径前面多少个字符
	PrintCodeFilePath bool    //是否打印代码文件全路径及行号
	AutoRemoveFile    uint    //自动删除多少天以前的日志，0则不删除任何日志
}

type logConfig struct {
	LogDir      string //存放日志文件的文件夹
	IsOneFile   bool   //是否保存成1个文件
	FileMaxSize int64  //日志文件最大大小，如果为0，则无限大日志。不会进行日志切割。

	setOK bool //是否已经设置过了

	lock        sync.Mutex
	logFile     *os.File
	logSize     int64
	debugFile   *os.File
	debugSize   int64
	infoFile    *os.File
	infoSize    int64
	successFile *os.File
	successSize int64
	warnFile    *os.File
	warnSize    int64
	errorFile   *os.File
	errorSize   int64
	panicFile   *os.File
	panicSize   int64
}
