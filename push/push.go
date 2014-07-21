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
		p.url = "http://channel.api.duapp.com/rest/2.0/channel"
	}

	return p
}

// base api

func (p *Push) QueryBindList(options map[string]string) (map[string]interface{}, error) {
	var httpUrl = p.url + "/channel"

	options["method"] = "query_bindlist"
	addDefaultOptions(options, httpUrl, p.apiKey, p.secretKey)

	res, err := Request("POST", httpUrl, nil, options)

	return res, err
}

func (p *Push) PushMsg(options map[string]string) (map[string]interface{}, error) {
	var httpUrl = p.url + "/channel"

	options["method"] = "push_msg"
	addDefaultOptions(options, httpUrl, p.apiKey, p.secretKey)

	res, err := Request("POST", httpUrl, nil, options)

	return res, err
}

// advanced api

func (p *Push) VerifyBind(options map[string]string) (map[string]interface{}, error) {
	var httpUrl = p.url + "/channel"

	options["method"] = "verify_bind"
	addDefaultOptions(options, httpUrl, p.apiKey, p.secretKey)

	res, err := Request("POST", httpUrl, nil, options)

	return res, err
}

func (p *Push) FetchMsg(options map[string]string) (map[string]interface{}, error) {
	var httpUrl = p.url + "/channel"

	options["method"] = "fetch_msg"
	addDefaultOptions(options, httpUrl, p.apiKey, p.secretKey)

	res, err := Request("POST", httpUrl, nil, options)

	return res, err
}

func (p *Push) FetchMsgCount(options map[string]string) (map[string]interface{}, error) {
	var httpUrl = p.url + "/channel"

	options["method"] = "fetch_msgcount"
	addDefaultOptions(options, httpUrl, p.apiKey, p.secretKey)

	res, err := Request("POST", httpUrl, nil, options)

	return res, err
}

func (p *Push) DeleteMsg(options map[string]string) (map[string]interface{}, error) {
	var httpUrl = p.url + "/channel"

	options["method"] = "delete_msg"
	addDefaultOptions(options, httpUrl, p.apiKey, p.secretKey)

	res, err := Request("POST", httpUrl, nil, options)

	return res, err
}

func (p *Push) SetTag(options map[string]string) (map[string]interface{}, error) {
	var httpUrl = p.url + "/channel"

	options["method"] = "set_tag"
	addDefaultOptions(options, httpUrl, p.apiKey, p.secretKey)

	res, err := Request("POST", httpUrl, nil, options)

	return res, err
}

func (p *Push) FetchTag(options map[string]string) (map[string]interface{}, error) {
	var httpUrl = p.url + "/channel"

	options["method"] = "fetch_tag"
	addDefaultOptions(options, httpUrl, p.apiKey, p.secretKey)

	res, err := Request("POST", httpUrl, nil, options)

	return res, err
}

func (p *Push) DeleteTag(options map[string]string) (map[string]interface{}, error) {
	var httpUrl = p.url + "/channel"

	options["method"] = "delete_tag"
	addDefaultOptions(options, httpUrl, p.apiKey, p.secretKey)

	res, err := Request("POST", httpUrl, nil, options)

	return res, err
}

func (p *Push) QueryUserTags(options map[string]string) (map[string]interface{}, error) {
	var httpUrl = p.url + "/channel"

	options["method"] = "query_user_tags"
	addDefaultOptions(options, httpUrl, p.apiKey, p.secretKey)

	res, err := Request("POST", httpUrl, nil, options)

	return res, err
}

func (p *Push) QueryDeviceType(options map[string]string) (map[string]interface{}, error) {
	var httpUrl = p.url + "/" + options["channel"]

	delete(options, "channel")

	options["method"] = "query_device_type"
	addDefaultOptions(options, httpUrl, p.apiKey, p.secretKey)

	res, err := Request("POST", httpUrl, nil, options)

	return res, err
}

func addDefaultOptions(options map[string]string, httpUrl, apiKey, secretKey string) {
	options["apikey"] = apiKey
	options["timestamp"], _ = cov.String(time.Now().Unix() * 1000)

	sign, _ := GenerateSign("POST", httpUrl, secretKey, options)
	options["sign"] = sign
}
