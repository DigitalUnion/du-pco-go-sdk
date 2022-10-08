/**
 * @Author: chenwei
 * @Description: A example for this package
 * @Date: 2022/09/09 2:34 PM
 */

package example

import (
	"github.com/DigitalUnion/du-pco-go-sdk/dupco"
	"github.com/goccy/go-json"
	"log"
)

func ExampleClient() {
	api, err := dupco.NewClient("aaaaaa", "test", "aaaaaaaa")
	if err != nil {
		log.Println(err.Error())
		return
	}
	api.EnableTestMode()
	for i := 0; i < 10; i++ {
		r := api.Call("/geofence/v1/list_fence", nil)
		j, _ := json.Marshal(r)
		log.Println(string(j))
	}
}
