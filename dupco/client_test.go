/**
 * @Author: chenwei
 * @Description: The test for this dupco
 * @Date: 2022/09/07 4:09 PM
 */

package dupco

import (
	"github.com/goccy/go-json"
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	sdkVer = "test"
	api := NewClient("aaaaaa", "test", "aaaaaaaa")
	for i := 0; i < 10; i++ {
		r := api.Call("/geofence/v1/list_fence", nil)
		j, _ := json.Marshal(r)
		log.Println(string(j))
	}
}
