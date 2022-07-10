package main

import (
	"artlog/artlog/downloadutils"
)

func main() {
	downloadutils.DefaultDownloadUtils.DownloadFileWithProgress("https://mirrors.tuna.tsinghua.edu.cn/centos/7.9.2009/isos/x86_64/CentOS-7-x86_64-NetInstall-2009.iso")
}
