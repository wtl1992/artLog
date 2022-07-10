package downloadutils

import (
	"artlog/artlog"
	"artlog/artlog/colors"
	"artlog/artlog/files"
	"bufio"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type DownloadUtils struct {
}

// FileContentLengthType 文件大小类型
type FileContentLengthType struct {
	FileContentLength int64  `json:"fileContentLength"`
	Unit              string `json:"unit"`
}

type DownloadWriter struct {
	reader      io.ReadCloser
	total       int64
	currentSize int64
	startTime   int64
	endTime     int64
}

var (
	DefaultDownloadUtils = &DownloadUtils{}
)

func (du *DownloadUtils) DownloadFileWithProgress(url string) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("User-Agent", "Golang")
	response, _ := http.DefaultClient.Do(request)

	fileContentLength := files.GetNetWorkFileContentLength(url)

	var downloadFileName = files.GetNetWorkFileName(url)

	fileContentLengthType := computeFileSize(fileContentLength)

	artlog.InfoLog("%s\n", "Downloading "+downloadFileName+"  ( "+fileContentLengthType.Unit+" )")
	printDownloadFileProgressWithAllInfo(response, fileContentLength, "/Users/tlzs/Downloads/"+downloadFileName)
}

// 计算文件大小格式
func computeFileSize(fileContentLength int64) *FileContentLengthType {
	if fileContentLength < 1024 {
		return &FileContentLengthType{FileContentLength: fileContentLength, Unit: strconv.FormatInt(fileContentLength, 10) + "B"}
	} else if fileContentLength < 1024*1024 {
		return &FileContentLengthType{FileContentLength: fileContentLength, Unit: strconv.FormatFloat(float64(fileContentLength)/float64(1024), 'f', 6, 64) + "KB"}
	} else if fileContentLength < 1024*1024*1024 {
		return &FileContentLengthType{FileContentLength: fileContentLength, Unit: strconv.FormatFloat(float64(fileContentLength)/float64(1024*1024), 'f', 6, 64) + "MB"}
	} else if fileContentLength < 1024*1024*1024*1024 {
		return &FileContentLengthType{FileContentLength: fileContentLength, Unit: strconv.FormatFloat(float64(fileContentLength)/float64(1024*1024*1024), 'f', 6, 64) + "GB"}
	} else {
		return &FileContentLengthType{FileContentLength: fileContentLength, Unit: strconv.FormatFloat(float64(fileContentLength)/float64(1024*1024*1024*1024), 'f', 6, 64) + "TB"}
	}
}

// 计算当前文件下载速度
func computeDownloadRealSpeed(bytesCount int, usedMillTime uint64) string {
	usedSecs := float64(usedMillTime) / float64(1000)
	speed := float64(bytesCount) / usedSecs
	if speed < 1000 {
		return fmt.Sprintf("%10s  B/S", strconv.FormatFloat(speed, 'f', 2, 64))
	} else if speed < 1000*1000 {
		return fmt.Sprintf("%10s KB/S", strconv.FormatFloat(speed/float64(1000), 'f', 2, 64))
	} else if speed < 1000*1000*1000 {
		return fmt.Sprintf("%10s MB/S", strconv.FormatFloat(speed/float64(1000*1000), 'f', 2, 64))
	} else if speed < 1000*1000*1000*1000 {
		return fmt.Sprintf("%10s GB/S", strconv.FormatFloat(speed/float64(1000*1000*1000), 'f', 2, 64))
	} else {
		return fmt.Sprintf("%10s KB/S", "0")
	}
}

func printALongBackgroundProgressLine(percent float64, charNum uint8) {
	var aLongBackgroundLine string
	var joinNum = int(charNum) - int(math.Round(percent*float64(charNum)))
	for i := 0; i < int(joinNum); i++ {
		aLongBackgroundLine += "─"
	}

	artlog.InfoLog(fmt.Sprintf("  \r%%%ds", charNum), aLongBackgroundLine, colors.Color_Font_Purple)
}

func printFrontProgressLine(percent float64, charNum uint8) {
	var frontLine string
	var currentNum = int(math.Round(percent * float64(charNum)))
	for i := 0; i < currentNum; i++ {
		frontLine += "─"
	}
	artlog.InfoLog("\r%s", frontLine, colors.Color_Font_Blue)

	var leftLine string
	var leftNum = int(charNum) - currentNum
	for i := 0; i < leftNum; i++ {
		leftLine += "─"
	}
	artlog.InfoLog("%s", leftLine, colors.Color_Font_Purple)
}

func printDownloadFileProgressWithAllInfo(response *http.Response, total int64, savePath string) {
	var downloadReader = &DownloadWriter{reader: response.Body, total: total, currentSize: 0}
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		reader := bufio.NewReader(downloadReader.reader)
		file, err := os.OpenFile(savePath, os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			panic(err)
		}

		defer func() {
			_ = file.Close()
		}()

		var timer = time.NewTimer(time.Millisecond * 100)

		go func() {
			for {
				select {
				case <-timer.C:
					downloadReader.endTime = time.Now().UnixMilli()
					percent := float64(downloadReader.currentSize) / float64(downloadReader.total)

					// 打印背景进度条暗线
					printALongBackgroundProgressLine(percent, 50)
					printFrontProgressLine(percent, 50)
					timer.Reset(time.Millisecond * 100)

					artlog.InfoLog("%26s",
						computeFileSize(downloadReader.currentSize).Unit+
							"/"+
							computeFileSize(downloadReader.total).Unit, colors.Attr_HideCursor, colors.Color_Font_Yellow)

					artlog.InfoLog("      %f", percent, colors.Color_Font_Blue)

					artlog.InfoLog("  %s", computeDownloadRealSpeed(int(downloadReader.currentSize), uint64(downloadReader.endTime-downloadReader.startTime)), colors.Color_Font_Blue, colors.Attr_ClearCcontentsFromCursor2Endline)
				}
			}
		}()

		var buffer []byte = make([]byte, 10240)
		downloadReader.startTime = time.Now().UnixMilli()
		for {
			n, err := reader.Read(buffer)
			downloadReader.currentSize += int64(n)
			_, _ = file.Write(buffer[:n])

			if err != nil {
				timer.Reset(time.Millisecond * 100)
				time.Sleep(time.Second)
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
