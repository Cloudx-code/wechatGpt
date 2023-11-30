package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"wechatGpt/common/consts"
	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/config"
	"wechatGpt/dao/local_cache"
)

// 做代理
const BASEURL = "https://ng.dawu.world/v1/chat/completions"
const BASEURL2 = "https://api.openai.com/v1/chat/completions"

type GPT4VService struct {
	ak        string
	senderId  string
	cacheList []Message
}

func NewGPT4VService(senderId string) *GPT4VService {
	return &GPT4VService{
		ak:       "Bearer " + config.GetLLMConfig().GPTAk,
		senderId: senderId,
	}
}

func (g *GPT4VService) PreQuery() {
	key := consts.RedisKeyGPT4VContext + g.senderId
	v, ok := local_cache.Get(key)
	if !ok {
		return
	}
	if cachInfoStr, ok := v.(string); ok {
		err := utils.Decode(cachInfoStr, &g.cacheList)
		if err != nil {
			logs.Error("fail to PreQuery,err:%v")
		}
		return
	}
	return
}

func (g *GPT4VService) Query(content string) (string, error) {
	if content == "清空" {
		key := consts.RedisKeyGPT4VContext + g.senderId
		local_cache.Set(key, "")
		return "已清空内容", nil
	}
	buf, err := g.getGPTReqBuf(content)
	if err != nil {
		logs.Error("fail to getGPTReqBuf,[GPT4VService],err:%v", err)
		return "", err
	}
	// 2. 调用HTTP
	req, err := http.NewRequest("POST", BASEURL, bytes.NewBuffer(buf))
	if err != nil {
		logs.Error("fail to NewRequest,[GPT4VService],err:%v", err)
		return "", err
	}
	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", g.ak)
	// Send the request
	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("fail to client.Do,[GPT4VService],err:%v", err)
		return "", err
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("fail to ReadAll,[GPT4VService],err:%v", err)
		return "", err
	}
	logs.Info("success to get resp", string(body))
	chatGptResp := &ChatGPTResponseBody{}
	err = json.Unmarshal(body, &chatGptResp)
	if err != nil || len(chatGptResp.Choices) == 0 {
		logs.Error("fail to json.Unmarshal,[GPT4VService],err:%v", err)
		return "", err
	}
	cost := float64(chatGptResp.Usage.PromptTokens)*0.001 + float64(chatGptResp.Usage.CompletionTokens)*0.003

	logs.Info("success to content:%v", chatGptResp.Choices[0].Message.Content)
	g.setCacheMsg(content, chatGptResp.Choices[0].Message.Content)
	return chatGptResp.Choices[0].Message.Content + fmt.Sprintf("\n\nGPT4本次花费：%v美元\n 清空上下文可输入：清空", cost), nil
}

func (g *GPT4VService) setCacheMsg(content, reply string) {
	if len(g.cacheList) == 0 {
		g.cacheList = make([]Message, 0)
	}
	g.cacheList = append(g.cacheList, []Message{
		{
			Role: "user",
			Content: []Content{
				{
					Type: "text",
					Text: content,
				},
			},
		},
		{
			Role: "assistant",
			Content: []Content{
				{
					Type: "text",
					Text: reply,
				},
			},
		},
	}...)
}

func (g *GPT4VService) getGPTReqBuf(content string) ([]byte, error) {
	requestBody := RequestBody{
		Model:     "gpt-4-vision-preview",
		Messages:  []Message{},
		MaxTokens: 1000,
	}
	if len(g.cacheList) != 0 {
		requestBody.Messages = append(requestBody.Messages, g.cacheList...)
	}

	requestBody.Messages = append(requestBody.Messages, Message{
		Role: "user",
		Content: []Content{
			{
				Type: "text",
				Text: content,
			},
		},
	})
	return json.Marshal(requestBody)
}

func (g *GPT4VService) PostQuery() {
	if len(g.cacheList) == 0 {
		return
	}
	key := consts.RedisKeyGPT4VContext + g.senderId
	local_cache.Set(key, utils.Encode(g.cacheList))
	return
}
