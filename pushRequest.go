// Copyright 2015 Beijing Venusource Tech.Co.Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//Rest Api请求参数构造
package push

import (
	"fmt"
	"strconv"
)

type QueryMsgStatusRequest struct {
	MsgId string // 推送接口返回的msg_id，支持json数组格式
}

func (r *QueryMsgStatusRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("msg_id", r.MsgId)
	return params
}

type PushMsgToSingleDeviceRequest struct {
	ChannelId    string
	MsgType      int
	Message      string
	MsgExpires   int64
	DeployStatus int    // iOS: 1=开发 2=生产
	TopicId      string // 消息topic，可选
	ExtraMsgType int    // 与MsgType互斥，0=消息 1=通知
	ExtraMsg     string // 与extra_msg_type配合
	ChannelType  int    // 0=默认通道 1=私信通道(小米OPPO)
	LongConnThirdList string // 厂商ID json数组，如 "[1,4,5]"
}

func (r *PushMsgToSingleDeviceRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("channel_id", r.ChannelId)
	params.AddUnescaped("msg_type", strconv.Itoa(r.MsgType))
	params.AddUnescaped("msg", r.Message)
	if r.MsgExpires != 0 {
		params.AddUnescaped("msg_expires", fmt.Sprintf("%d", r.MsgExpires))
	}
	if r.DeployStatus != 0 {
		params.AddUnescaped("deploy_status", strconv.Itoa(r.DeployStatus))
	}
	if r.TopicId != "" {
		params.AddUnescaped("topic_id", r.TopicId)
	}
	if r.ExtraMsgType != 0 {
		params.AddUnescaped("extra_msg_type", strconv.Itoa(r.ExtraMsgType))
	}
	if r.ExtraMsg != "" {
		params.AddUnescaped("extra_msg", r.ExtraMsg)
	}
	if r.ChannelType != 0 {
		params.AddUnescaped("channel_type", strconv.Itoa(r.ChannelType))
	}
	if r.LongConnThirdList != "" {
		params.AddUnescaped("long_conn_thirdlist", r.LongConnThirdList)
	}
	return params
}

type PushMsgToAllRequest struct {
	MsgType      int
	Message      string
	MsgExpires   int
	DeployStatus int    // iOS: 1=开发 2=生产
	SendTime     int64  // 定时推送unix时间戳
	ExtraMsgType int    // 与MsgType互斥
	ExtraMsg     string
}

func (r *PushMsgToAllRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("msg_type", strconv.Itoa(r.MsgType))
	params.AddUnescaped("msg", r.Message)
	if r.MsgExpires != 0 {
		params.AddUnescaped("msg_expires", strconv.Itoa(r.MsgExpires))
	}
	if r.DeployStatus != 0 {
		params.AddUnescaped("deploy_status", strconv.Itoa(r.DeployStatus))
	}
	if r.SendTime != 0 {
		params.AddUnescaped("send_time", fmt.Sprintf("%d", r.SendTime))
	}
	if r.ExtraMsgType != 0 {
		params.AddUnescaped("extra_msg_type", strconv.Itoa(r.ExtraMsgType))
	}
	if r.ExtraMsg != "" {
		params.AddUnescaped("extra_msg", r.ExtraMsg)
	}
	return params
}

type PushMsgToTagRequest struct {
	TagName      string
	MsgType      int
	Message      string
	MsgExpires   int
	DeployStatus int    // iOS: 1=开发 2=生产
	SendTime     int64  // 定时推送unix时间戳
	ExtraMsgType int    // 与MsgType互斥
	ExtraMsg     string
}

func (r *PushMsgToTagRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("type", "1")
	params.AddUnescaped("tag", r.TagName)
	params.AddUnescaped("msg_type", strconv.Itoa(r.MsgType))
	params.AddUnescaped("msg", r.Message)
	if r.MsgExpires != 0 {
		params.AddUnescaped("msg_expires", strconv.Itoa(r.MsgExpires))
	}
	if r.DeployStatus != 0 {
		params.AddUnescaped("deploy_status", strconv.Itoa(r.DeployStatus))
	}
	if r.SendTime != 0 {
		params.AddUnescaped("send_time", fmt.Sprintf("%d", r.SendTime))
	}
	if r.ExtraMsgType != 0 {
		params.AddUnescaped("extra_msg_type", strconv.Itoa(r.ExtraMsgType))
	}
	if r.ExtraMsg != "" {
		params.AddUnescaped("extra_msg", r.ExtraMsg)
	}
	return params
}

type PushBatchUniMsgRequest struct {
	ChannelIds string // channel_id的json数组字符串
	MsgType    int
	Message    string
	MsgExpires int
	TopicId    string // 分类主题名称
	ExtraMsgType int  // 与MsgType互斥
	ExtraMsg     string
	DeployStatus int    // iOS: 1=开发 2=生产
	LongConnThirdList string // 厂商ID json数组，如 "[1,4,5]"
}

func (r *PushBatchUniMsgRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("channel_ids", r.ChannelIds)
	params.AddUnescaped("msg_type", strconv.Itoa(r.MsgType))
	params.AddUnescaped("msg", r.Message)
	if r.MsgExpires != 0 {
		params.AddUnescaped("msg_expires", strconv.Itoa(r.MsgExpires))
	}
	if r.TopicId != "" {
		params.AddUnescaped("topic_id", r.TopicId)
	}
	if r.ExtraMsgType != 0 {
		params.AddUnescaped("extra_msg_type", strconv.Itoa(r.ExtraMsgType))
	}
	if r.ExtraMsg != "" {
		params.AddUnescaped("extra_msg", r.ExtraMsg)
	}
	if r.DeployStatus != 0 {
		params.AddUnescaped("deploy_status", strconv.Itoa(r.DeployStatus))
	}
	if r.LongConnThirdList != "" {
		params.AddUnescaped("long_conn_thirdlist", r.LongConnThirdList)
	}
	return params
}

type QueryTimerRecordsRequest struct {
	TimerId    string
	Start      int
	Limit      int
	RangeStart int64
	RangeEnd   int64
}

func (r *QueryTimerRecordsRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("timer_id", r.TimerId)
	if r.Start != 0 {
		params.AddUnescaped("start", strconv.Itoa(r.Start))
	}
	if r.Limit != 0 {
		params.AddUnescaped("limit", strconv.Itoa(r.Limit))
	}
	if r.RangeStart != 0 {
		params.AddUnescaped("range_start", fmt.Sprintf("%d", r.RangeStart))
	}
	if r.RangeEnd != 0 {
		params.AddUnescaped("range_end", fmt.Sprintf("%d", r.RangeEnd))
	}
	return params
}

type QueryTopicRecordsRequest struct {
	TopicId    string
	Start      int
	Limit      int
	RangeStart int64
	RangeEnd   int64
}

func (r *QueryTopicRecordsRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("topic_id", r.TopicId)
	if r.Start != 0 {
		params.AddUnescaped("start", strconv.Itoa(r.Start))
	}
	if r.Limit != 0 {
		params.AddUnescaped("limit", strconv.Itoa(r.Limit))
	}
	if r.RangeStart != 0 {
		params.AddUnescaped("range_start", fmt.Sprintf("%d", r.RangeStart))
	}
	if r.RangeEnd != 0 {
		params.AddUnescaped("range_end", fmt.Sprintf("%d", r.RangeEnd))
	}
	return params
}

type QueryTimerListRequest struct {
	TimerId string
	Start   int
	Limit   int
}

func (r *QueryTimerListRequest) AddToParams(params *OrderedParams) *OrderedParams {
	if r.TimerId != "" {
		params.AddUnescaped("timer_id", r.TimerId)
	}
	if r.Start != 0 {
		params.AddUnescaped("start", strconv.Itoa(r.Start))
	}
	if r.Limit != 0 {
		params.AddUnescaped("limit", strconv.Itoa(r.Limit))
	}
	return params
}

type QueryTopicListRequest struct {
	Start int
	Limit int
}

func (r *QueryTopicListRequest) AddToParams(params *OrderedParams) *OrderedParams {
	if r.Start != 0 {
		params.AddUnescaped("start", strconv.Itoa(r.Start))
	}
	if r.Limit != 0 {
		params.AddUnescaped("limit", strconv.Itoa(r.Limit))
	}
	return params
}

type QueryTagsRequest struct {
	TagName string
	Start   int
	Limit   int
}

func (r *QueryTagsRequest) AddToParams(params *OrderedParams) *OrderedParams {
	if r.TagName != "" {
		params.AddUnescaped("tag", r.TagName)
	}
	if r.Start != 0 {
		params.AddUnescaped("start", strconv.Itoa(r.Start))
	}
	if r.Limit != 0 {
		params.AddUnescaped("limit", strconv.Itoa(r.Limit))
	}
	return params
}

type CreateTagRequest struct {
	TagName string
}

func (r *CreateTagRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("tag", r.TagName)
	return params
}

type DeleteTagRequest struct {
	TagName string
}

func (r *DeleteTagRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("tag", r.TagName)
	return params
}

type AddDevicesToTagRequest struct {
	TagName    string
	ChannelIds string
}

func (r *AddDevicesToTagRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("tag", r.TagName)
	params.AddUnescaped("channel_ids", r.ChannelIds)
	return params
}

type DeleteDevicesFromTagRequest struct {
	TagName    string
	ChannelIds string
}

func (r *DeleteDevicesFromTagRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("tag", r.TagName)
	params.AddUnescaped("channel_ids", r.ChannelIds)
	return params
}

type QueryDeviceNumInTagRequest struct {
	TagName string
}

func (r *QueryDeviceNumInTagRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("tag", r.TagName)
	return params
}

type QueryStatisticTopicRequest struct {
	TopicId string
}

func (r *QueryStatisticTopicRequest) AddToParams(params *OrderedParams) *OrderedParams {
	params.AddUnescaped("topic_id", r.TopicId)
	return params
}
