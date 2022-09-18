package rlog

import (
	"path"
	"sync/atomic"
)

type logContent struct {
	LogLevel LogLeve //日志等级
	Content  string  //内容
}

var chContent = make(chan logContent, 4096) //日志内容管道

// writeFile 将日志内容写入文件
func writeFile() {
	for d := range chContent {
		atomic.AddInt64(&curPool, 1)
		rLogFile.lock.Lock()
		if rLogFile.IsOneFile {
			if rLogFile.logFile != nil {
				size, _ := rLogFile.logFile.WriteString(d.Content)
				rLogFile.logSize += int64(size)
				if rLogFile.logSize > rLogFile.FileMaxSize {
					rLogFile.logFile, rLogFile.logSize = autoCreateFile(rLogFile.logFile, path.Join(rLogFile.LogDir, "app.log"))
				}
			}
		} else {
			switch d.LogLevel {
			case LogLevelDebug:
				if rLogFile.debugFile != nil {
					size, _ := rLogFile.debugFile.WriteString(d.Content)
					rLogFile.debugSize += int64(size)
					if rLogFile.debugSize > rLogFile.FileMaxSize {
						rLogFile.debugFile, rLogFile.debugSize = autoCreateFile(rLogFile.debugFile, path.Join(rLogFile.LogDir, "debug", "debug.log"))
					}
				}
			case LogLevelInfo:
				if rLogFile.infoFile != nil {
					size, _ := rLogFile.infoFile.WriteString(d.Content)
					rLogFile.infoSize += int64(size)
					if rLogFile.infoSize > rLogFile.FileMaxSize {
						rLogFile.infoFile, rLogFile.infoSize = autoCreateFile(rLogFile.infoFile, path.Join(rLogFile.LogDir, "info", "info.log"))
					}
				}
			case LogLeveSuccess:
				if rLogFile.successFile != nil {
					size, _ := rLogFile.successFile.WriteString(d.Content)
					rLogFile.successSize += int64(size)
					if rLogFile.successSize > rLogFile.FileMaxSize {
						rLogFile.successFile, rLogFile.successSize = autoCreateFile(rLogFile.successFile, path.Join(rLogFile.LogDir, "success", "success.log"))
					}
				}
			case LogLevelWarn:
				if rLogFile.warnFile != nil {
					size, _ := rLogFile.warnFile.WriteString(d.Content)
					rLogFile.warnSize += int64(size)
					if rLogFile.warnSize > rLogFile.FileMaxSize {
						rLogFile.warnFile, rLogFile.warnSize = autoCreateFile(rLogFile.warnFile, path.Join(rLogFile.LogDir, "warn", "warn.log"))
					}
				}
			case LogLevelError:
				if rLogFile.errorFile != nil {
					size, _ := rLogFile.errorFile.WriteString(d.Content)
					rLogFile.errorSize += int64(size)
					if rLogFile.errorSize > rLogFile.FileMaxSize {
						rLogFile.errorFile, rLogFile.errorSize = autoCreateFile(rLogFile.errorFile, path.Join(rLogFile.LogDir, "error", "error.log"))
					}
				}
			case LogLevelPanic:
				if rLogFile.panicFile != nil {
					size, _ := rLogFile.panicFile.WriteString(d.Content)
					rLogFile.panicSize += int64(size)
					if rLogFile.panicSize > rLogFile.FileMaxSize {
						rLogFile.panicFile, rLogFile.panicSize = autoCreateFile(rLogFile.panicFile, path.Join(rLogFile.LogDir, "panic", "panic.log"))
					}
				}
			}
		}
		rLogFile.lock.Unlock()
		atomic.AddInt64(&curPool, -1)
	}
}
