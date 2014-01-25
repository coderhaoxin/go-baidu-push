package push

import (
	"time"
)

type Push struct {
	host      string
	apiKey    string
	secretKey string
}

func New(options map[string]string) *Push {
	p := &Push{
		apiKey:    options["apiKey"],
		secretKey: options["secretKey"],
	}
	if host, ok := options["host"]; ok {
		p.host = host
	} else {
		// default host
		p.host = "channel.api.duapp.com/rest/2.0/channel/"
	}

	return p
}

// base api

func (p *Push) QueryBindList(options ...map[string]interface{}) (map[string]interface{}, error) {
	var httpMethod = "POST"
	var httpUrl = p.host

	options["method"] = "query_bindlist"
	options["apikey"] = p.apiKey
	options["timestamp"] = time.Now().Unix()

	sign, err := GenerateSign(httpMethod, httpUrl, p.secretKey, options)
	options["sign"] = sign
	return options, err
}

func (p *Push) PushMsg(options map[string]interface{}, messages map[string]string) (map[string]interface{}, error) {
	var httpMethod = "POST"
	var httpUrl = p.host

	options["method"] = "push_msg"
	options["apikey"] = p.apiKey
	options["timestamp"] = time.Now().Unix()

	sign, err := GenerateSign(httpMethod, httpUrl, p.secretKey, options)
	options["sign"] = sign
	return options, err
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
