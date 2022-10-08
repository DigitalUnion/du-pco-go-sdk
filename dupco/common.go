/**
 * @Author: chenwei
 * @Description: A simple way to call DigitalUnion service
 * @Date: 2022/09/08 6:44 PM
 */

package dupco

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"github.com/valyala/fasthttp"
	"io/ioutil"
)

// Decode : decode data
func Decode(data, secret []byte) ([]byte, error) {
	xorBs, err := aesDecrypt(data, secret)
	if err != nil {
		return nil, err
	}
	return zlibUnCompress(xorBs)
}

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

// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
func encode(data, key []byte) ([]byte, error) {
	compressData := zlibCompress(data)
	return aesEncrypt(compressData, key)
}

func aesEncrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	encryptBytes := pkcs7Padding(data, blockSize)
	crypted := make([]byte, len(encryptBytes))
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

func aesDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return data, nil
	}
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
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
