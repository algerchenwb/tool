package rlog

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// parseLog 生成日志文本内容
func parseLog(model PrintModel, tag string, v ...any) string {
	var str = make([]string, 0, 4)
	if model != ModelNoPathTime {
		if model != ModelNoTime {
			str = append(str, time.Now().Format("2006/01/02·15:04:05-0700 "))
		}
		str = append(str, tag+"\t")
		if model != ModelNoPath {
			if rCfg.PrintCodeFilePath {
				_, file, line, _ := runtime.Caller(2)
				if rCfg.CropPathLength > 0 && len(file) > int(rCfg.CropPathLength) {
					file = file[rCfg.CropPathLength:]
				}
				str = append(str, fmt.Sprint(file, ":", line, " "))
			}
		}
	} else {
		str = append(str, tag+"\t")
	}
	return strings.Join(str, "") + "：" + fmt.Sprintln(v...)
}
