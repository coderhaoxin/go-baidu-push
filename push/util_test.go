package push

import (
	"testing"
)

func TestGenerateSign(t *testing.T) {
	kv := make(map[string]string)
	kv["k1"] = "v1"
	kv["k3"] = "v3"
	kv["k2"] = "v2"
	sign, _ := GenerateSign("POST", "duapp.com/rest/2.0/channel", "testSecretKey", kv)
	if sign != "efcdea7aefc91cb180fe59f3a4d10f23" {
		t.Error(sign)
	}
}

func TestRequest(t *testing.T) {
	m, e := Request("POST", "http://channel.api.duapp.com/rest/2.0/channel/channel", nil, nil)

	if e != nil {
		t.Error(e)
	}

	if m["error_msg"].(string) != "request params not valid. " {
		t.Error(m["error_msg"])
	}
}
