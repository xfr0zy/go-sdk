package client

import (
	"github.com/qiniu/go-sdk/v7/conf"
	"net/http"
	"time"
)

const (
	RequestHeaderKeyXQiniuDate = "X-Qiniu-Date"
)

func AddDefaultHeader(headers http.Header) error {
	return addXQiniuDate(headers)
}

func addXQiniuDate(headers http.Header) error {
	if !conf.IsEnableRequestHeaderXQiniuDate() {
		return nil
	}

	timeString := time.Now().UTC().Format("20060102T150405Z")
	headers.Set(RequestHeaderKeyXQiniuDate, timeString)
	return nil
}
