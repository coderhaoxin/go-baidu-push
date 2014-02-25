package push

import (
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
		p.url = "http://localhost:3000"
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

	res, err := Request("POST", httpUrl, nil, options)

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

	res, err := Request("POST", httpUrl, nil, nil)

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
