package push

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/coderhaoxin/go-convert/cov"
)

func GenerateSign(httpMethod string, httpUrl string, secretKey string, keyValues map[string]string) (string, error) {
	var kv []string
	for k, v := range keyValues {
		kv = append(kv, k+"="+v)
	}
	sort.Strings(kv)
	kvString := strings.Join(kv, "")

	h := md5.New()
	h.Write([]byte(url.QueryEscape(httpMethod + httpUrl + kvString + secretKey)))
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func Request(method, httpUrl string, headers, options map[string]string) (map[string]interface{}, error) {
	v := url.Values{}
	v.Add("method", options["method"])
	v.Add("apikey", options["apikey"])
	v.Add("timestamp", options["timestamp"])
	v.Add("sign", options["sign"])

	body := strings.NewReader(v.Encode())
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
		return nil, errors.New("do http request error")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	return cov.JSONmap(resBody)
}
