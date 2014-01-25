package push

import (
	// "fmt"
	"testing"
)

func TestGenerateSign(t *testing.T) {
	kv := make(map[string]interface{})
	kv["k1"] = "v1"
	kv["k3"] = "v3"
	kv["k2"] = "v2"
	sign, _ := GenerateSign("POST", "duapp.com/rest/2.0/channel", "testSecretKey", kv)
	if sign != "783qeu/JHLGA/lnzpNEPIw==" {
		t.Fail()
	}
}

func assert(result, should, errmsg string, t *testing.T) {
	if result != should {
		t.Error(errmsg)
		t.Error("result: ", result, " should: ", should)
		t.Fail()
	}
}

func TestConvertString(t *testing.T) {
	var in string = "123456"
	result, _ := String(in)
	assert(result, "123456", "convert string error", t)
}

func TestConvertInt(t *testing.T) {
	var in int = 10086
	result, _ := String(in)
	assert(result, "10086", "convert int error", t)
}

func TestConvertInt32(t *testing.T) {
	var in int32 = 12345
	result, _ := String(in)
	assert(result, "12345", "convert int32 error", t)
}

func TestConvertInt64(t *testing.T) {
	var in int64 = 1258096
	result, _ := String(in)
	assert(result, "1258096", "convert int64 error", t)
}
