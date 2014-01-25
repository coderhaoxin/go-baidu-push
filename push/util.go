package push

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	// "fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type ResErr struct {
	RequestId int64
	ErrorCode int64
	ErrorMsg  string
}

func (e *ResErr) Error() string {
	return e.ErrorMsg
}

func String(value interface{}) (string, error) {
	switch reflect.TypeOf(value).String() {
	case "string":
		v, _ := value.(string)
		return v, nil
	case "int":
		v, _ := value.(int)
		return strconv.Itoa(v), nil
	case "int32":
		v, _ := value.(int32)
		return strconv.FormatInt(int64(v), 10), nil
	case "int64":
		v, _ := value.(int64)
		return strconv.FormatInt(v, 10), nil
	default:
		return "", errors.New("unknown type")
	}
}

func GenerateSign(httpMethod string, httpUrl string, secretKey string, keyValues map[string]interface{}) (string, error) {
	var kv []string
	for k, v := range keyValues {
		str, err := String(v)
		if err != nil {
			return "", errors.New("convert error")
		}
		kv = append(kv, k+"="+str)
	}
	sort.Strings(kv)
	kvString := strings.Join(kv, "")

	h := md5.New()
	h.Write([]byte(url.QueryEscape(httpMethod + httpUrl + kvString + secretKey)))
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func Request(method, url string, headers map[string]interface{}, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.New("new http request error")
	}

	// set http headers
	for k, v := range headers {
		str, err := String(v)
		if err != nil {
			return nil, err
		}
		req.Header.Add(k, str)
	}

	// do http request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("do http request error")
	}

	if res.StatusCode >= 400 {
		resBody, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		var resErr ResErr
		err = json.Unmarshal(resBody, resErr)
		return nil, err
	} else {
		return res, nil
	}
}
