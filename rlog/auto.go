package rlog

import (
	"fmt"
	"os"
	"path"
	"time"
)

// autoRemoveFile 自动删除旧日志
func autoRemoveFile() {
	ticker := time.NewTicker(time.Minute)
	for {
		if rCfg.AutoRemoveFile > 0 {
			removeFile(path.Dir(rLogFile.LogDir), int(rCfg.AutoRemoveFile))
		}
		<-ticker.C
	}
}

// removeFile 删除指定文件下最后修改日期为 day 天前的文件和文件夹。
func removeFile(dir string, day int) {
	now := time.Now()
	readDir, err := os.ReadDir(dir)
	if err == nil {
		for _, d := range readDir {
			info, err := d.Info()
			if err == nil {
				if info.ModTime().Sub(now.AddDate(0, 0, -day)) < 0 {
					if d.IsDir() {
						removeFile(path.Join(dir, d.Name()), day) //去遍历下级文件夹
					} else {
						_ = os.Remove(path.Join(dir, d.Name()))
					}
				}
			}
		}
	}
}

// autoCreateFile 自动创建日志文件
func autoCreateFile(old *os.File, filepath string) (*os.File, int64) {
	dir := path.Dir(filepath)  //提取出文件夹
	_ = os.MkdirAll(dir, 0777) //创建文件夹

	stat, err := os.Stat(filepath) //判断文件是否存在
	if err != nil {
		//不存在则创建
		file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			fmt.Println("创建日志文件失败：", err)
			return old, 0
		} else {
			if old != nil {
				_ = old.Close()
			}
			return file, 0
		}
	} else {
		//存在则先关闭文件
		if old != nil {
			_ = old.Close()
		}

		if stat.Size() < rLogFile.FileMaxSize { //文件大小没有超出限制，那就读取就文件，进行日志追加。
			file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0777)
			if err == nil {
				return file, stat.Size()
			}
		}
		base := path.Base(filepath)
		ext := path.Ext(filepath)
		if len(ext) > 0 {
			base = base[:len(base)-len(ext)] //去除扩展名
		}
		//重命名旧文件名
		err := os.Rename(filepath, path.Join(dir, fmt.Sprint(base, "_", time.Now().Format("20060102150405_-0700"), ext)))
		if err != nil {
			return nil, 0
		}
		file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			return nil, 0
		} else {
			return file, 0
		}
	}
}
