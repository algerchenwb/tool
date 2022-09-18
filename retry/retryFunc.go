package rtool

import (
	"fmt"
	"runtime"
	"time"
	"tools/rlog"
)

// RetryFunc 如果函数运行失败则重试，如果函数运行成功则结束
// maxBout：最大重试次数
// sleepTime：重试间隔时间
// ptLog：是否打印错误日志
// fx：要执行的函数，函数返回值1：是否继续重试至最大次数；函数返回值2：错误信息
// 返回值：错误信息
func RetryFunc(maxBout int, sleepTime time.Duration, ptLog bool, fx func() (bool, error)) (err error) {
	_, file, line, _ := runtime.Caller(1)
	prev := fmt.Sprint(file[rlog.GetCropPathLength():], ":", line, ":")

	if maxBout <= 0 {
		maxBout = 1
	}
	for i := 0; i < maxBout; i++ {
		ct, er := fx()
		if false == ct {
			return er
		}
		if er == nil {
			return nil
		}
		if ptLog {
			rlog.WarnP(rlog.ModelNoPath, prev, "函数运行错误：", i+1, "/", maxBout, "：", er)
		}
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}
	return
}

// RetryFuncErr 如果函数运行失败则重试，如果函数运行成功则结束
// maxBout：最大重试次数
// sleepTime：重试间隔时间
// ptLog：是否打印错误日志
// fx：要执行的函数，函数返回值1：是否继续重试至最大次数；函数返回值2：错误信息
// 返回值：错误信息
func RetryFuncErr(maxBout int, sleepTime time.Duration, ptLog bool, fx func() error) (err error) {
	_, file, line, _ := runtime.Caller(1)
	prev := fmt.Sprint(file[rlog.GetCropPathLength():], ":", line, ":")

	if maxBout <= 0 {
		maxBout = 1
	}
	for i := 0; i < maxBout; i++ {
		err = fx()
		if err == nil {
			return
		}
		if ptLog {
			rlog.WarnP(rlog.ModelNoPath, prev, "函数运行错误：", i+1, "/", maxBout, "：", err)
		}
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}
	return
}
