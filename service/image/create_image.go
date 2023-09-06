package image

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func CreateImage(prompt string) (string, error) {
	// 定义请求的数据结构
	type RequestData struct {
		Prompt string `json:"prompt"`
		N      int    `json:"n"`
		Size   string `json:"size"`
	}

	// 创建一个新的请求数据
	requestData := RequestData{
		Prompt: prompt,
		N:      1,
		Size:   "1024x1024",
	}

	// 将请求数据编码为 JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return "fail to json.Marshal，err:" + err.Error(), err
	}

	// 创建一个新的 HTTP 请求
	apiKey := "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"
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "fail to http.NewRequest，err:" + err.Error(), err
	}

	// 添加 HTTP 头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// 发送 HTTP 请求并获取响应
	client := &http.Client{Timeout: 90 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "fail to client.Do，err:" + err.Error(), err
	}
	defer resp.Body.Close()

	// 读取响应体并输出
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "fail to ioutil.ReadAll，err:" + err.Error(), err
	}
	fmt.Println(string(body))
	respInfo := &ImageResp{}
	err = json.Unmarshal(body, &respInfo)
	if err != nil {
		return "fail to json.Unmarshal，err:" + err.Error(), err
	}
	if len(respInfo.Data) == 0 {
		return "fail to len(respInfo.Data)==0", errors.New("error")
	}
	return respInfo.Data[0].Url, nil
}

type ImageResp struct {
	Data []*UrlInfo `json:"data,omitempty"`
}

type UrlInfo struct {
	Url string `json:"url,omitempty"`
}
