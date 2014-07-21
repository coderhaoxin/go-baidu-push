package push

import (
	"testing"
)

var push = New(map[string]string{
	"apiKey":    "",
	"secretKey": "",
})

var userId = ""
var tagName = "xxooxoxo"

func TestQueryBindList(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = userId
	options["start"] = "0"
	options["limit"] = "30"
	result, err := push.QueryBindList(options)
	if err != nil {
		t.Error(err)
	}
	if _, ok := result["error_code"]; ok {
		t.Error(result)
	}
	t.Log(result)
}

func TestPushMsg(t *testing.T) {
	options := make(map[string]string)
	options["push_type"] = "1"
	options["user_id"] = userId
	options["msg_keys"] = `"["title"]"`
	options["messages"] = `"["hello"]"`
	result, err := push.PushMsg(options)
	if err != nil {
		t.Error(err)
	}
	if _, ok := result["error_code"]; ok {
		t.Error(result)
	}
	t.Log(result)
}

func TestVerifyBind(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = userId
	result, err := push.VerifyBind(options)
	if err != nil {
		t.Error(err)
	}
	if _, ok := result["error_code"]; ok {
		t.Error(result)
	}
	t.Log(result)
}

func TestFetchMsg(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = userId
	result, err := push.FetchMsg(options)
	if err != nil {
		t.Error(err)
	}
	if _, ok := result["error_code"]; ok {
		t.Error(result)
	}
	t.Log(result)
}

func TestFetchMsgCount(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = userId
	result, err := push.FetchMsgCount(options)
	if err != nil {
		t.Error(err)
	}
	if _, ok := result["error_code"]; ok {
		t.Error(result)
	}
	t.Log(result)
}

func TestDeleteMsg(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = userId
	options["msg_ids"] = "123"
	result, err := push.DeleteMsg(options)
	if err != nil {
		t.Error(err)
	}
	if _, ok := result["error_code"]; ok {
		t.Error(result)
	}
	t.Log(result)
}

func TestSetTag(t *testing.T) {
	options := make(map[string]string)
	options["tag"] = tagName
	result, err := push.SetTag(options)
	if err != nil {
		t.Error(err)
	}
	if _, ok := result["error_code"]; ok {
		t.Error(result)
	}
	t.Log(result)
}

func TestFetchTag(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = userId
	result, err := push.FetchTag(options)
	if err != nil {
		t.Error(err)
	}
	if _, ok := result["error_code"]; ok {
		t.Error(result)
	}
	t.Log(result)
}

func TestDeleteTag(t *testing.T) {
	options := make(map[string]string)
	options["tag"] = tagName
	result, err := push.DeleteTag(options)
	if err != nil {
		t.Error(err)
	}
	if _, ok := result["error_code"]; ok {
		t.Error(result)
	}
	t.Log(result)
}

func TestQueryUserTags(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = userId
	result, err := push.QueryUserTags(options)
	if err != nil {
		t.Error(err)
	}
	if _, ok := result["error_code"]; ok {
		t.Error(result)
	}
	t.Log(result)
}

func TestQueryDeviceType(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = userId
	options["channel"] = "10086"
	result, err := push.QueryDeviceType(options)
	if err != nil {
		t.Error(err)
	}
	if _, ok := result["error_code"]; ok {
		t.Error(result)
	}
	t.Log(result)
}
