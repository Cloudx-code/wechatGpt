package weTab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"wechatGpt/common/utils"
)

func GetGpt(text, conversationId string) (string, string) {
	url := "https://wetabchatpro.haohuola.com/api/chat/conversation-v2"

	payload := struct {
		Prompt         string `json:"prompt"`
		AssistantID    string `json:"assistantId"`
		ConversationId string `json:"conversationId"`
	}{
		Prompt:      text,
		AssistantID: "",
	}
	if len(conversationId) > 0 {
		payload.ConversationId = conversationId
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return "", ""
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVpZCI6IjY0YTM4OGZlMmE3NWZjNzBkMTU2NTg4MSIsInZlcnNpb24iOjEsImJyYW5jaCI6InpoIiwiY2hhdFZpcEVuZFRpbWUiOjE2OTM2Mzg1MDkzNjh9LCJpYXQiOjE2OTE4MTE3OTAsImV4cCI6MTY5MTk4NDU5MH0.teBMz-uKBwmCNE6qpq2VRBfkuXjLEvCZ2ic4ggS-hKo")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Origin", "chrome-extension://aikflfpejipbpjdlfabpgclhblkpaafo")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Add("i-app", "hitab")
	req.Header.Add("i-branch", "zh")
	req.Header.Add("i-lang", "zh-CN")
	req.Header.Add("i-platform", "chrome")
	req.Header.Add("i-version", "1.1.2")
	req.Header.Add("sec-ch-ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", `"macOS"`)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println("end GetGpt")
	fmt.Println("GetGpt body:", string(body))

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
				return content, conversationId
			}
			if id, ok := data["data"].(map[string]interface{})["conversationId"].(string); ok {
				fmt.Println("测试下conversationId:", id)
				conversationId = id
			}
			if _, ok := data["data"].(map[string]interface{})["content"].(string); !ok {
				return content, conversationId
			}
			content += data["data"].(map[string]interface{})["content"].(string)
		} else {
			fmt.Println("Error:", err)
		}
	}
	return content, conversationId
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
	fmt.Println("result:")
	fmt.Println(utils.Encode(result))
	return result
}
