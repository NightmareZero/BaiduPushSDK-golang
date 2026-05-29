// Copyright 2015 Beijing Venusource Tech.Co.Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
//封装了http调用相关的一些方法
package push

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"sort"
	"strings"
)

// phpURLEncode 实现与 PHP urlencode 兼容的编码。
// 百度官方签名算法要求以 PHP urlencode 为准。
func phpURLEncode(s string) string {
	encoded := url.QueryEscape(s)
	// PHP urlencode: ~ 会被编码为 %7E
	encoded = strings.ReplaceAll(encoded, "~", "%7E")
	// PHP urlencode: 空格编码为 + 而非 %20
	encoded = strings.ReplaceAll(encoded, "%20", "+")
	return encoded
}

// userAgent 生成符合百度要求的 User-Agent 格式:
//
//	BCCS_SDK/3.0 (操作系统) 开发语言/版本 (SDK名称及版本)
func userAgent() string {
	return fmt.Sprintf("BCCS_SDK/3.0 (%s) %s (%s/%s)",
		runtime.GOOS, runtime.Version(), "BaiduPushSDK-golang", "1.0.0")
}

//执行http请求
func httpExecute(method string, urlStr string, params *OrderedParams, debug bool) (*http.Response, error) {
	v := url.Values{}
	for _, key := range params.Keys() {
		v.Add(key, params.Get(key))
	}
	body := v.Encode()

	req, err := http.NewRequest(method, urlStr, strings.NewReader(body))
	if err != nil {
		return nil, errors.New("NewRequest failed: " + err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Header.Set("User-Agent", userAgent())

	if debug {
		fmt.Println("[DEBUG] Request URL:", urlStr)
		fmt.Println("[DEBUG] Request Body:", body)
		fmt.Println("[DEBUG] User-Agent:", req.Header.Get("User-Agent"))
	}

	HttpClient := &http.Client{}
	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, errors.New("Do: " + err.Error())
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		defer resp.Body.Close()
		bytes, _ := io.ReadAll(resp.Body)
		if debug {
			fmt.Println("[DEBUG] Response Status:", resp.StatusCode)
			fmt.Println("[DEBUG] Response Body:", string(bytes))
		}
		return resp, errors.New(string(bytes))
	}
	return resp, err
}

//获得HTTP请求的body部分内容
func getBody(method, url string, params *OrderedParams, debug bool) (*string, error) {
	resp, err := httpExecute(method, url, params, debug)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	bodyStr := string(bodyBytes)
	if debug {
		fmt.Println("[DEBUG] Status:", resp.StatusCode, resp.Status)
		fmt.Println("[DEBUG] Body Response:", bodyStr)
	}
	return &bodyStr, nil
}

//计算签名的字符串
func requestString(method string, urlPath string, secretkey string, params *OrderedParams) string {
	result := method + urlPath
	for _, key := range params.Keys() {
		result += fmt.Sprintf("%s=%s", key, params.Get(key))
	}
	return result + secretkey
}

//调用API
func CallApiServer(httpMethod string, server string, class string, method string, params *OrderedParams, secretkey string, debug bool, i interface{}) error {
	reqString := requestString(httpMethod, server+class+method, secretkey, params)
	h := md5.New()
	h.Write([]byte(phpURLEncode(reqString)))
	signature := hex.EncodeToString(h.Sum(nil))
	params.Add("sign", signature)
	if debug {
		fmt.Println("[DEBUG] Sign Base String:", reqString)
		fmt.Println("[DEBUG] Signature:", signature)
	}
	result, err := getBody("POST", server+class+method, params, debug)
	if err == nil {
		json.Unmarshal([]byte(*result), &i)
		return nil
	}
	return err
}

//排序后的参数列表
type OrderedParams struct {
	allParams   map[string]string
	keyOrdering []string
}

func NewOrderedParams() *OrderedParams {
	return &OrderedParams{
		allParams:   make(map[string]string),
		keyOrdering: make([]string, 0),
	}
}

func (o *OrderedParams) Get(key string) string {
	return o.allParams[key]
}

func (o *OrderedParams) Keys() []string {
	sort.Sort(o)
	return o.keyOrdering
}

func (o *OrderedParams) Add(key, value string) {
	o.AddUnescaped(key, value)
}

func (o *OrderedParams) AddUnescaped(key, value string) {
	o.allParams[key] = value
	o.keyOrdering = append(o.keyOrdering, key)
}

func (o *OrderedParams) Len() int {
	return len(o.keyOrdering)
}

func (o *OrderedParams) Less(i int, j int) bool {
	return o.keyOrdering[i] < o.keyOrdering[j]
}

func (o *OrderedParams) Swap(i int, j int) {
	o.keyOrdering[i], o.keyOrdering[j] = o.keyOrdering[j], o.keyOrdering[i]
}

func (o *OrderedParams) Clone() *OrderedParams {
	clone := NewOrderedParams()
	for _, key := range o.Keys() {
		clone.AddUnescaped(key, o.Get(key))
	}
	return clone
}
