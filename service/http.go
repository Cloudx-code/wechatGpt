package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetGpt(text string) string {
	url := "https://wetabchat.haohuola.com/api/chat/conversation-v2"

	payload := struct {
		Prompt      string `json:"prompt"`
		AssistantID string `json:"assistantId"`
	}{
		Prompt:      text,
		AssistantID: "",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return ""
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVpZCI6IjY0YTM4OGZlMmE3NWZjNzBkMTU2NTg4MSIsInZlcnNpb24iOjEsImJyYW5jaCI6InpoIiwiY2hhdFZpcEVuZFRpbWUiOjE2OTM2Mzg1MDkzNjh9LCJpYXQiOjE2OTA5NjAxMTMsImV4cCI6MTY5MTEzMjkxM30.4nsk2w-R-xBWTJa2JRksxnoX_bj9YK7WLNRnkAGPNu8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Origin", "chrome-extension://aikflfpejipbpjdlfabpgclhblkpaafo")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("i-app", "hitab")
	req.Header.Set("i-branch", "zh")
	req.Header.Set("i-lang", "zh-CN")
	req.Header.Set("i-platform", "chrome")
	req.Header.Set("i-version", "1.1.2")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)

	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	// 解析结果
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("fail to ioutil.ReadAll，err:", err)
		return ""
	}

	var content string
	jsonStrArray := SplitString(string(body), "_e79218965e_")
	// 解析每个JSON片段并提取所需的数据
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	for _, jsonStr := range jsonStrArray {
		var data map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &data)
		if err == nil {
			fmt.Println("Content:", data["data"].(map[string]interface{})["content"])
			if _, ok := data["data"].(map[string]interface{}); !ok {
				return content
			}
			if _, ok := data["data"].(map[string]interface{})["content"].(string); !ok {
				return content
			}
			content += data["data"].(map[string]interface{})["content"].(string)
		} else {
			fmt.Println("Error:", err)
		}
	}
	return content
}

// 将字符串按照指定的分隔符拆分为数组
func SplitString(s string, sep string) []string {
	var result []string
	length := len(s)
	start := 0
	for i := 0; i < length; i++ {
		if i+len(sep) <= length && s[i:i+len(sep)] == sep {
			if start != i {
				result = append(result, s[start:i])
			}
			start = i + len(sep)
			i += len(sep) - 1
		}
	}
	if start != length {
		result = append(result, s[start:length])
	}
	return result
}
