/**
 * @Author: chenwei
 * @Description: A example for this package
 * @Date: 2022/09/09 2:34 PM
 */

package main

import (
	"github.com/DigitalUnion/du-pco-go-sdk/dupco"
	"github.com/goccy/go-json"
	"log"
)

func main() {
	ExampleClient()
}

func ExampleClient() {
	api := dupco.NewDataClient("cloud-test", "aa", "yDpDEihpUsF_RyWsCES1H")
	//api.EnableTestMode()
	for i := 0; i < 10; i++ {
		r := api.Call("idmap-query-all", []byte(`{"f":"mac,imei","k":"868862032205613","m":"0"}`))
		j, _ := json.Marshal(r)
		log.Println(string(j))
	}
}
