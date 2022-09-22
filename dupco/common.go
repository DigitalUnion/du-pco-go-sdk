/**
 * @Author: chenwei
 * @Description: A simple way to call DigitalUnion service
 * @Date: 2022/09/08 6:44 PM
 */

package dupco

import (
	"bytes"
	"compress/zlib"
	"github.com/valyala/fasthttp"
	"io/ioutil"
)

func http(reqMethod, url string, reqBody []byte, header map[string]string) (int, []byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetContentType(contentTypeJson)
	req.Header.SetMethod(reqMethod)
	for k, v := range header {
		req.Header.Add(k, v)
	}
	req.Header.Add(sdkVerKey, sdkVer)
	req.SetRequestURI(domain + url)
	req.SetBody(reqBody)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := fasthttp.Do(req, resp); err != nil {
		return resp.StatusCode(), []byte{}, err
	}
	return resp.StatusCode(), resp.Body(), nil
}

func encode(data, sv []byte) []byte {
	compressData := zlibCompress(data)
	return xor(compressData, sv)
}
func decode(bs, sv []byte) ([]byte, error) {
	xorBs := xor(bs, sv)
	return zlibUnCompress(xorBs)
}

func xor(raw, key []byte) []byte {
	rawLen := len(raw)
	keyLen := len(key)
	dst := make([]byte, 0, rawLen)
	for i := 0; i < rawLen; i++ {
		dst = append(dst, raw[i]^key[i%keyLen])
	}
	return dst
}

func zlibCompress(data []byte) []byte {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(data)
	w.Close()
	return b.Bytes()
}

func zlibUnCompress(s []byte) ([]byte, error) {
	b := bytes.NewReader(s)
	r, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
