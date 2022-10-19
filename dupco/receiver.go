/**
 * @Author: chenwei
 * @Description:
 * @Date: 2022/10/18 2:03 PM
 */

package dupco

import (
	"io/ioutil"
	"net/http"
)

// DecodeData: get decode data from http.Request
func DecodeData(req http.Request, secret []byte) ([]byte, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	data, err := Decode(body, secret)
	if err != nil {
		return nil, err
	}
	return data, nil
}
