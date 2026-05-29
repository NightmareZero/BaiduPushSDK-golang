package push

import (
	"encoding/json"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("test-api-key", "test-secret-key")
	if client == nil {
		t.Fatal("NewClient returned nil")
	}
	if client.apikey != "test-api-key" || client.secretKey != "test-secret-key" {
		t.Fatalf("unexpected client fields: apikey=%q secretKey=%q", client.apikey, client.secretKey)
	}
}

func TestDebug(t *testing.T) {
	client := NewClient("a", "b")
	if client.debug {
		t.Fatal("debug should default to false")
	}
	client.Debug(true)
	if !client.debug {
		t.Fatal("debug should be true after enable")
	}
}

func TestOrderedParams(t *testing.T) {
	params := NewOrderedParams()
	params.AddUnescaped("b", "2")
	params.AddUnescaped("a", "1")
	params.AddUnescaped("c", "3")

	keys := params.Keys()
	if len(keys) != 3 {
		t.Fatalf("expected 3 keys, got %d", len(keys))
	}
	if keys[0] != "a" || keys[1] != "b" || keys[2] != "c" {
		t.Fatalf("keys not sorted: %v", keys)
	}
	if params.Get("a") != "1" || params.Get("b") != "2" || params.Get("c") != "3" {
		t.Fatal("Get returned wrong values")
	}
}

func TestPushMsgToAllResponseUnmarshal(t *testing.T) {
	jsonStr := `{"request_id":12345,"response_params":{"msg_id":"msg123","timer_id":"timer456","send_time":1700000000}}`
	var resp PushMsgToAllJSONResponse
	if err := json.Unmarshal([]byte(jsonStr), &resp); err != nil {
		t.Fatal(err)
	}
	r := resp.PushMsgToAllResponse
	if r.MsgId != "msg123" || r.TimerId != "timer456" || r.SendTime != 1700000000 {
		t.Errorf("PushMsgToAllResponse: got MsgId=%q TimerId=%q SendTime=%d", r.MsgId, r.TimerId, r.SendTime)
	}
}

func TestPushMsgToSingleDeviceResponseUnmarshal(t *testing.T) {
	jsonStr := `{"request_id":12345,"response_params":{"msg_id":"msg123","send_time":1700000000}}`
	var resp PushMsgToSingleDeviceJSONResponse
	if err := json.Unmarshal([]byte(jsonStr), &resp); err != nil {
		t.Fatal(err)
	}
	r := resp.PushMsgToSingleDeviceResponse
	if r.MsgId != "msg123" || r.SendTime != 1700000000 {
		t.Errorf("PushMsgToSingleDeviceResponse: got MsgId=%q SendTime=%d", r.MsgId, r.SendTime)
	}
}

func TestAndroidNotificationMarshaling(t *testing.T) {
	n := AndroidNotification{
		Title:                  "Hello",
		Description:            "World",
		NotificationBuilderId:  0,
		NotificationBasicStyle: 7,
		OpenType:               1,
		Url:                    "https://example.com",
		TargetChannelId:        "default",
	}
	data, err := json.Marshal(n)
	if err != nil {
		t.Fatal(err)
	}
	var decoded AndroidNotification
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.Title != "Hello" || decoded.Description != "World" || decoded.TargetChannelId != "default" {
		t.Errorf("AndroidNotification round-trip failed: %+v", decoded)
	}
}

func TestSignAlgorithm(t *testing.T) {
	// 百度官方示例
	params := NewOrderedParams()
	params.AddUnescaped("apikey", "Ljc710pzAa99GULCo8y48NvB")
	params.AddUnescaped("expires", "1313293565")
	params.AddUnescaped("timestamp", "1427180905")

	reqStr := requestString("POST", "http://api.tuisong.baidu.com/rest/3.0/test/echo",
		"87772555E1C16715EBA5C85341684C58", params)

	expected := "POSThttp://api.tuisong.baidu.com/rest/3.0/test/echoapikey=Ljc710pzAa99GULCo8y48NvBexpires=1313293565timestamp=142718090587772555E1C16715EBA5C85341684C58"
	if reqStr != expected {
		t.Errorf("sign base string mismatch\nexpected: %s\ngot:      %s", expected, reqStr)
	}
}

func TestZeroValueParamsSkipped(t *testing.T) {
	// 验证零值参数不会被发送
	req := PushMsgToAllRequest{
		MsgType:    1,
		Message:    `{"title":"test","description":"hello"}`,
		MsgExpires: 0, // 零值应该跳过
		SendTime:   0, // 零值应该跳过
	}
	params := NewOrderedParams()
	params = req.AddToParams(params)

	if params.Get("msg_expires") != "" {
		t.Error("msg_expires should be empty when zero")
	}
	if params.Get("send_time") != "" {
		t.Error("send_time should be empty when zero")
	}
}
