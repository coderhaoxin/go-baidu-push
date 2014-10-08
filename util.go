package push

import "github.com/coderhaoxin/go-convert"
import "encoding/hex"
import "crypto/md5"
import "io/ioutil"
import "net/http"
import "strings"
import "net/url"
import "errors"
import "sort"

func GenerateSign(httpMethod string, httpUrl string, secretKey string, keyValues map[string]string) (string, error) {
	var kv []string
	for k, v := range keyValues {
		kv = append(kv, k+"="+v)
	}
	sort.Strings(kv)
	kvString := strings.Join(kv, "")

	h := md5.New()
	h.Write([]byte(url.QueryEscape(httpMethod + httpUrl + kvString + secretKey)))
	return hex.EncodeToString(h.Sum(nil)), nil
}

func Request(method, httpUrl string, headers, options map[string]string) (map[string]interface{}, error) {
	data := url.Values{}

	for k, v := range options {
		data.Add(k, v)
	}

	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest(method, httpUrl, body)
	if err != nil {
		return nil, errors.New("new http request error")
	}

	// set http headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// do http request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	return convert.JSONmap(resBody)
}
