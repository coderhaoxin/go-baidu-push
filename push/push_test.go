package push

import (
	"fmt"
	"testing"
)

var push = New(map[string]string{
	"apiKey":    "12306",
	"secretKey": "fuck12306",
})

func TestQueryBindList(t *testing.T) {
	options := make(map[string]interface{})
	options["user_id"] = "xxoo10086"
	options["start"] = 0
	options["limit"] = 30
	result, err := push.QueryBindList(options)
	fmt.Println(result, err)
}

func TestPushMsg(t *testing.T) {
	options := make(map[string]interface{})
	options["push_type"] = 1
	options["user_id"] = "xxoo10086"
	messages := make(map[string]interface{})
	messages["hello"] = "baidu push"
	result, err := push.PushMsg(options, messages)
	fmt.Println(result, err)
}
