package push

import (
	"fmt"
	"testing"
)

var push = New(map[string]string{
	"apiKey":    "xxoo",
	"secretKey": "ooxx",
})

func TestQueryBindList(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = "xxoo10086"
	options["start"] = "0"
	options["limit"] = "30"
	result, err := push.QueryBindList(options)
}

func TestPushMsg(t *testing.T) {
	options := make(map[string]string)
	options["push_type"] = "1"
	options["user_id"] = "xxoo10086"
	messages := make(map[string]string)
	messages["hello"] = "baidu push"
	result, err := push.PushMsg(options, messages)
}

func TestVerifyBind(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = "xxoo10086"
	result, err := push.QueryBindList(options)
}

func TestFetchMsg(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = "xxoo10086"
	result, err := push.QueryBindList(options)
}

func TestFetchMsgCount(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = "xxoo10086"
	result, err := push.QueryBindList(options)
}

func TestDeleteMsg(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = "xxoo10086"
	result, err := push.QueryBindList(options)
}

func TestSetTag(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = "xxoo10086"
	result, err := push.QueryBindList(options)
}

func TestFetchTag(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = "xxoo10086"
	result, err := push.QueryBindList(options)
}

func TestDeleteTag(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = "xxoo10086"
	result, err := push.QueryBindList(options)
}

func TestQueryUserTags(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = "xxoo10086"
	result, err := push.QueryBindList(options)
}

func TestQueryDeviceType(t *testing.T) {
	options := make(map[string]string)
	options["user_id"] = "xxoo10086"
	result, err := push.QueryBindList(options)
}
