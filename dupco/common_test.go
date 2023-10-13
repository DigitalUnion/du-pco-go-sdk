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
	//origin := []byte(`[{"device_id":"DUoRKigM6eE0FschEUx6zKyBelI9_1ZaPwea","application_name":"com.xunwangvip.dlspapp","app_package_name":"","app_version":"1.0.2","scene_type":"MAT2","os":"android","ts":"2023-10-13 15:46:49"}]`)
	//dd, _ := encode(origin, []byte(`QrmsG_IMiA5_oODjJq4OQ`))
	//log.Println(string(dd))
	//ee, err := Decode(dd, []byte("QrmsG_IMiA5_oODjJq4OQ"))
	ee, err := DecodePushData([]byte("eYuCObqPdvByBq0eAWAKhrbtqEfkg1u+b47Bb+Jd8nS9QkxgBPxPiVVlSA0ePvxxRgfZlR0ac/p95cyrpeTIjTr95b0NoeElc4iTeit4xnp6ppsQDXwfDTxLpjQPVMLjb7/Sa6gVV+DVKufYzJ1E2SfG9DMXjc9fNoxGSIa3SZSn99dZQhYrZQZjtPi/r2keo5/cbRcFGeRoZMeYj1ITZKuvJ/QjyNtDbiwYvQudx1HkA+Pg1PCmfljCaIpDqGFZ"), []byte("QrmsG_IMiA5_oODjJq4OQ"))
	log.Println(string(ee), err)
}

func Test_aesEncrypt(t *testing.T) {
	res, _ := aesEncrypt([]byte("hello"), []byte("world"))
	log.Println(res)
}
