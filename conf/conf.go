package conf

const Version = "7.12.1"

const (
	CONTENT_TYPE_JSON      = "application/json"
	CONTENT_TYPE_FORM      = "application/x-www-form-urlencoded"
	CONTENT_TYPE_OCTET     = "application/octet-stream"
	CONTENT_TYPE_MULTIPART = "multipart/form-data"
)

var enableRequestHeaderXQiniuDate = true

func IsEnableRequestHeaderXQiniuDate() bool {
	return enableRequestHeaderXQiniuDate
}

func SetEnableRequestHeaderXQiniuDate(enable bool) {
	enableRequestHeaderXQiniuDate = enable
}
