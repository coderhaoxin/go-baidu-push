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
	if sign != "783qeu/JHLGA/lnzpNEPIw==" {
		t.Fail()
	}
}
