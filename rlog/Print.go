package rlog

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func Debug(v ...any) {
	if rCfg.LogLevelConsole <= LogLevelDebug || rCfg.LogLevelFile <= LogLevelDebug {
		var content = parseLog(0, "debug", v...)
		if rCfg.LogLevelConsole <= LogLevelDebug {
			_, _ = fmt.Fprint(output, aurora.Index(12, content))
		}
		if rCfg.LogLevelFile <= LogLevelDebug {
			chContent <- logContent{
				LogLevel: LogLevelDebug,
				Content:  content,
			}
		}
	}
}

func DebugP(model PrintModel, v ...any) {
	if rCfg.LogLevelConsole <= LogLevelDebug || rCfg.LogLevelFile <= LogLevelDebug {
		var content = parseLog(model, "debug", v...)
		if rCfg.LogLevelConsole <= LogLevelDebug {
			_, _ = fmt.Fprint(output, aurora.Index(12, content))
		}
		if rCfg.LogLevelFile <= LogLevelDebug {
			chContent <- logContent{
				LogLevel: LogLevelDebug,
				Content:  content,
			}
		}
	}
}

func Info(v ...any) {
	if rCfg.LogLevelConsole <= LogLevelInfo || rCfg.LogLevelFile <= LogLevelInfo {
		var content = parseLog(0, "info", v...)
		if rCfg.LogLevelConsole <= LogLevelInfo {
			_, _ = fmt.Fprint(output, aurora.Index(51, content))
		}
		if rCfg.LogLevelFile <= LogLevelInfo {
			chContent <- logContent{
				LogLevel: LogLevelInfo,
				Content:  content,
			}
		}
	}
}

func InfoP(model PrintModel, v ...any) {
	if rCfg.LogLevelConsole <= LogLevelInfo || rCfg.LogLevelFile <= LogLevelInfo {
		var content = parseLog(model, "info", v...)
		if rCfg.LogLevelConsole <= LogLevelInfo {
			_, _ = fmt.Fprint(output, aurora.Index(51, content))
		}
		if rCfg.LogLevelFile <= LogLevelInfo {
			chContent <- logContent{
				LogLevel: LogLevelInfo,
				Content:  content,
			}
		}
	}
}

func Success(v ...any) {
	if rCfg.LogLevelConsole <= LogLeveSuccess || rCfg.LogLevelFile <= LogLeveSuccess {
		var content = parseLog(0, "succ", v...)
		if rCfg.LogLevelConsole <= LogLeveSuccess {
			_, _ = fmt.Fprint(output, aurora.Index(46, content))
		}
		if rCfg.LogLevelFile <= LogLeveSuccess {
			chContent <- logContent{
				LogLevel: LogLeveSuccess,
				Content:  content,
			}
		}
	}
}

func SuccessP(model PrintModel, v ...any) {
	if rCfg.LogLevelConsole <= LogLeveSuccess || rCfg.LogLevelFile <= LogLeveSuccess {
		var content = parseLog(model, "succ", v...)
		if rCfg.LogLevelConsole <= LogLeveSuccess {
			_, _ = fmt.Fprint(output, aurora.Index(46, content))
		}
		if rCfg.LogLevelFile <= LogLeveSuccess {
			chContent <- logContent{
				LogLevel: LogLeveSuccess,
				Content:  content,
			}
		}
	}
}

func Warn(v ...any) {
	if rCfg.LogLevelConsole <= LogLevelWarn || rCfg.LogLevelFile <= LogLevelWarn {
		var content = parseLog(0, "warn", v...)
		if rCfg.LogLevelConsole <= LogLevelWarn {
			_, _ = fmt.Fprint(output, aurora.Index(226, content))
		}
		if rCfg.LogLevelFile <= LogLevelWarn {
			chContent <- logContent{
				LogLevel: LogLevelWarn,
				Content:  content,
			}
		}
	}
}

func WarnP(model PrintModel, v ...any) {
	if rCfg.LogLevelConsole <= LogLevelWarn || rCfg.LogLevelFile <= LogLevelWarn {
		var content = parseLog(model, "warn", v...)
		if rCfg.LogLevelConsole <= LogLevelWarn {
			_, _ = fmt.Fprint(output, aurora.Index(226, content))
		}
		if rCfg.LogLevelFile <= LogLevelWarn {
			chContent <- logContent{
				LogLevel: LogLevelWarn,
				Content:  content,
			}
		}
	}
}

func Error(v ...any) {
	if rCfg.LogLevelConsole <= LogLevelError || rCfg.LogLevelFile <= LogLevelError {
		var content = parseLog(0, "error", v...)
		if rCfg.LogLevelConsole <= LogLevelError {
			_, _ = fmt.Fprint(output, aurora.Index(196, content))
		}
		if rCfg.LogLevelFile <= LogLevelError {
			chContent <- logContent{
				LogLevel: LogLevelError,
				Content:  content,
			}
		}
	}
}

func ErrorP(model PrintModel, v ...any) {
	if rCfg.LogLevelConsole <= LogLevelError || rCfg.LogLevelFile <= LogLevelError {
		var content = parseLog(model, "error", v...)
		if rCfg.LogLevelConsole <= LogLevelError {
			_, _ = fmt.Fprint(output, aurora.Index(196, content))
		}
		if rCfg.LogLevelFile <= LogLevelError {
			chContent <- logContent{
				LogLevel: LogLevelError,
				Content:  content,
			}
		}
	}
}

func Panic(v ...any) {
	if rCfg.LogLevelConsole <= LogLevelPanic || rCfg.LogLevelFile <= LogLevelPanic {
		var content = parseLog(0, "panic", v...)
		if rCfg.LogLevelConsole <= LogLevelPanic {
			_, _ = fmt.Fprint(output, aurora.Index(207, content))
		}
		if rCfg.LogLevelFile <= LogLevelPanic {
			chContent <- logContent{
				LogLevel: LogLevelPanic,
				Content:  content,
			}
		}
	}
}

func PanicP(model PrintModel, v ...any) {
	if rCfg.LogLevelConsole <= LogLevelPanic || rCfg.LogLevelFile <= LogLevelPanic {
		var content = parseLog(model, "panic", v...)
		if rCfg.LogLevelConsole <= LogLevelPanic {
			_, _ = fmt.Fprint(output, aurora.Index(207, content))
		}
		if rCfg.LogLevelFile <= LogLevelPanic {
			chContent <- logContent{
				LogLevel: LogLevelPanic,
				Content:  content,
			}
		}
	}
}
