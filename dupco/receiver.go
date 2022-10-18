/**
 * @Author: chenwei
 * @Description:
 * @Date: 2022/10/18 2:03 PM
 */

package dupco

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// RegisterDataHandler: register a function that handle data
func RegisterDataHandler(req http.Request, secret []byte, handleFunc func(bs []byte)) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	data, err := Decode(body, secret)
	if err != nil {
		return err
	}
	if handleFunc != nil {
		handleFunc(data)
		return nil
	} else {
		return errors.New("Unhandled data:" + string(data))
	}
}
