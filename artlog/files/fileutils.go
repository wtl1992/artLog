package files

import (
	"artlog/artlog/constants"
	"net/http"
	"strings"
	"time"
)

func GetNetWorkFileName(netRemoteFileUrl string) string {
	response, err := http.Get(netRemoteFileUrl)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = response.Body.Close()
	}()

	contentDisposition := response.Header.Get("Content-Disposition")
	if contentDisposition == constants.EMPTY_STRING {
		split := strings.Split(netRemoteFileUrl, "/")
		if len(split) > 0 {
			name := split[len(split)-1]
			if strings.Contains(name, ".") {
				return name
			}
		}
	}

	return time.Now().In(constants.TIME_LOCATION).Format("2006-01-02--13:04:05")
}

func GetNetWorkFileContentLength(netRemoteFileUrl string) int64 {
	request, err := http.NewRequest(http.MethodGet, netRemoteFileUrl, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("User-Agent", "Golang")
	response, _ := http.DefaultClient.Do(request)
	return response.ContentLength
}
