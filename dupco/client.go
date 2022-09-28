/**
 * @Author: chenwei
 * @Description: A simple way to call DigitalUnion service
 * @Date: 2022/09/07 4:09 PM
 */

package dupco

import (
	"fmt"
	"github.com/goccy/go-json"
	"strings"
)

// Client is a dupco representing a instance
type Client struct {
	// clientId identify of customer
	clientId string

	// secretKey key of secret
	// secreyKey and secretVal are use in pairs
	secretKey string

	// secretVal value of secret
	secretVal []byte
}

// Response is the response of request
type Response struct {
	// Code response code
	//
	// |code  | describe                        |
	// |------|---------------------------------|
	// |0     | success                         |
	// |10000 | IP not in whitelist             |
	// |10001 | Request path error              |
	// |10002 | Internal server error           |
	// |10100 | Param cilent_id required        |
	// |10101 | Param client_id not found       |
	// |10102 | This service is not activated   |
	// |10200 | Secret key required             |
	// |10201 | Secret not found                |
	// |10202 | Decode failed                   |
	// |10203 | Get request body failed         |
	// |10300 | Service not found               |
	// |10999 | Other error                     |
	//
	// code 0 means success, others means errors
	Code int `json:"code,omitempty"`

	// Msg the message of response
	// if code not 0, the msg will tell you the reason
	Msg string `json:"msg,omitempty"`

	// Data the data of response, The responses are different depending on the service
	// if code not 0, the data will be nil
	Data interface{} `json:"data,omitempty"`
}

// NewClient: create and return a new dupco
//
// clientId identify of dupco
// secretKey key of secret
// secretVal value of secret
func NewClient(clientId, secretKey, secretVal string) *Client {
	return &Client{
		clientId:  clientId,
		secretKey: secretKey,
		secretVal: []byte(secretVal),
	}
}

// EnableTestMode: enable test mode
//
// Warnning: if you call method with test mode ,the reponse will be different. So DO NOT use it on production mode
func (p *Client) EnableTestMode() {
	sdkVer = sdkVerForTest
}

// Call: call remote function
//
// Response response from remote server
// response if valid if error is nil
func (p *Client) Call(method string, data []byte) (res Response) {
	header := map[string]string{
		clientId:  p.clientId,
		secretKey: p.secretKey,
	}

	if len(data) != 0 {
		data = encode(data, p.secretVal)
	}
	resCode, respBody, err := http(httpMethodPost, method, data, header)
	if err != nil {
		res.Code = otherErrorCode
		res.Msg = err.Error()
		return
	}
	if resCode > 400 {
		res.Code = otherErrorCode
		res.Msg = fmt.Sprintf(fmtHttpCodeError, resCode)
		return
	}
	if resCode == 200 {
		if len(respBody) != 0 {
			respBody, err = decode(respBody, p.secretVal)
			if err != nil {
				res.Code = otherErrorCode
				errStr := err.Error()
				// warp error
				if strings.HasPrefix(errStr, "flate:") {
					res.Msg = secretNotMathMsg
				} else if errStr == zlibInvalidHeaderErr {
					res.Msg = secretNotMathMsg
				} else {
					res.Msg = err.Error()
				}
				return
			}
		}
	}
	err = json.Unmarshal(respBody, &res)
	if err != nil {
		switch string(respBody) {
		case pageNotFoundErr:
			res.Code = pathErrorCode
			res.Msg = pageNotFoundMsg
		default:
			res.Code = otherErrorCode
			res.Msg = err.Error()
		}
		return
	}
	return
}
