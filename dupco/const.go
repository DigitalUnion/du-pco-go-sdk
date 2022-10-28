/**
 * @Author: chenwei
 * @Description: consts of this package
 * @Date: 2022/09/09 4:05 PM
 */

package dupco

const (
	apiIdKey         = "api_id"
	clientId         = "client_id"
	secretKey        = "secret_key"
	fmtHttpCodeError = "HTTP CODE:%d"
)

const (
	sdkVerKey            = "sdk_ver"
	sdkVerForTest        = "test"
	dataDomain           = "https://data.shuzijz.cn/pco/data/sdk"
	baseDomain           = "https://data.shuzilm.cn/pco/base/sdk"
	httpMethodPost       = "POST"
	contentTypeJson      = "application/json; charset=utf-8"
	zlibInvalidHeaderErr = "zlib: invalid header"
	secretNotMathMsg     = "Secret key and secret value not match"
	pageNotFoundErr      = "404 page not found"
	pageNotFoundMsg      = "Request path error"
	pathErrorCode        = 2
	otherErrorCode       = 10999
)

var sdkVer = "v1.1.2"
