package push

import (
	"net/url"
	"strings"
	"time"

	"github.com/coderhaoxin/go-convert/cov"
)

type Push struct {
	url       string
	apiKey    string
	secretKey string
}

func New(options map[string]string) *Push {
	p := &Push{
		apiKey:    options["apiKey"],
		secretKey: options["secretKey"],
	}
	if url, ok := options["url"]; ok {
		p.url = url
	} else {
		// default url
		p.url = "http://channel.api.duapp.com/rest/2.0/channel"
	}

	return p
}

// base api

func (p *Push) QueryBindList(options map[string]string) (map[string]interface{}, error) {
	var httpMethod = "POST"
	var httpUrl = p.url + "/channel"

	options["method"] = "query_bindlist"
	options["apikey"] = p.apiKey
	options["timestamp"], _ = cov.String(time.Now().Unix())

	sign, err := GenerateSign(httpMethod, httpUrl, p.secretKey, options)
	options["sign"] = sign

	// headers := make(map[string]interface{})

	v := url.Values{}
	v.Add("method", options["method"])
	v.Add("apikey", options["apikey"])
	v.Add("timestamp", options["timestamp"])
	v.Add("sign", options["sign"])
	res, err := Request("POST", httpUrl, nil, strings.NewReader(v.Encode()))

	return res, err
}

func (p *Push) PushMsg(options map[string]string, messages map[string]string) (map[string]interface{}, error) {
	var httpMethod = "POST"
	var httpUrl = p.url + "/channel"

	options["method"] = "push_msg"
	options["apikey"] = p.apiKey
	options["timestamp"], _ = cov.String(time.Now().Unix())

	sign, err := GenerateSign(httpMethod, httpUrl, p.secretKey, options)
	options["sign"] = sign

	method, _ := cov.String(options["method"])
	res, err := Request(method, httpUrl, nil, nil)

	return res, err
}

// advanced api

func (p *Push) VerifyBind() {

}

func (p *Push) FetchMsg() {

}

func (p *Push) FetchMsgCount() {

}

func (p *Push) DeleteMsg() {

}

func (p *Push) SetTag() {

}

func (p *Push) FetchTag() {

}

func (p *Push) DeleteTag() {

}

func (p *Push) QueryUserTags() {

}

func (p *Push) QueryDeviceType() {

}
