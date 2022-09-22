/**
 * @Author: chenwei
 * @Description: consts of this package
 * @Date: 2022/09/09 4:05 PM
 */

package dupco

const (
	clientId         = "client_id"
	secretKey        = "secret_key"
	fmtHttpCodeError = "HTTP CODE:%d"
)

const (
	sdkVerKey            = "sdk_ver"
	sdkVerForTest        = "test"
	domain               = "http://data.shuzilm.cn/api"
	httpMethodPost       = "POST"
	contentTypeJson      = "application/json; charset=utf-8"
	zlibInvalidHeaderErr = "zlib: invalid header"
	secretNotMathMsg     = "Secret key and secret value not match"
	pageNotFoundErr      = "404 page not found"
	pageNotFoundMsg      = "Request path error"
	pathErrorCode        = "2"
	otherErrorCode       = "10999"
)

var sdkVer = "v0.0.5"
