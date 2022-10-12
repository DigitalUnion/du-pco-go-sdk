/**
 * @Author: chenwei
 * @Description:
 * @Date: 2022/09/30 3:31 PM
 */

package dupco

import (
	"log"
	"testing"
)

func Test_encode(t *testing.T) {
	origin := []byte("hello world")
	dd, _ := encode(origin, []byte("keyaaaaakeyaaaaakeyaaaaakeyaaaaa"))
	ee, err := Decode(dd, []byte("keyaaaaakeyaaaaakeyaaaaakeyaaaaa"))
	log.Println(string(ee), err)
}

func Test_aesEncrypt(t *testing.T) {
	res, _ := aesEncrypt([]byte("hello"), []byte("world"))
	log.Println(res)
}
