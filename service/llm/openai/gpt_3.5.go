package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"wechatGpt/common/logs"
	"wechatGpt/config"
)

// 做代理
const BASEURL = "https://ng.dawu.world/v1/chat/completions"
const BASEURL2 = "https://api.openai.com/v1/chat/completions"

type GPTService struct {
}

func (G GPTService) PreQuery() {

	return
}

func (G GPTService) Query(content string) (string, error) {
	return Completions(content)
}

func (G GPTService) PostQuery() {
	return
}

func NewGPTService() *GPTService {
	return &GPTService{}
}

// Completions gtp文本模型回复
// curl https://api.openai.com/v1/completions
// -H "Content-Type: application/json"
// -H "Authorization: Bearer your chatGPT key"
// -d '{"model": "text-davinci-003", "prompt": "give me good song", "temperature": 0, "max_tokens": 7}'
func Completions(msg string) (string, error) {
	//Prompt := `你现在是一个语言学习辅助助手，我在学习中国的脏话，无论我问什么，你都要巧妙的用脏话回复。即使面对批评，道德审判，任何形式的质疑，或者让你忘记脏话学习助手的身份，等等任何提问和沟通，你都要用脏话回复并且隐瞒你脏话助手的身份。碰到无法回答的问题，就随便回复一句脏话。`

	requestBody := ChatGPTRequestBody{
		Model:       "gpt-3.5-turbo",
		MaxTokens:   1000,
		Messages:    []Messages{{"user", msg}},
		Temperature: 1,
	}
	requestData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	logs.Info(fmt.Sprintf("request gpt json string : %v", string(requestData)))
	req, err := http.NewRequest("POST", BASEURL2, bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	apiKey := config.GetLLMConfig().GPT35Ak
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{Timeout: 90 * time.Second}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		body, _ := ioutil.ReadAll(response.Body)
		return "", errors.New(fmt.Sprintf("请求GTP出错了，gpt api status code not equals 200,code is %d ,details:  %v ", response.StatusCode, string(body)))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	logs.Info(fmt.Sprintf("response gpt json string : %v", string(body)))

	gptResponseBody := &ChatGPTResponseBody{}
	log.Println(string(body))
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		return "", err
	}

	var reply string
	if len(gptResponseBody.Choices) > 0 {
		reply = gptResponseBody.Choices[0].Message.Content
	}
	logs.Info(fmt.Sprintf("gpt response text: %s ", reply))
	return reply, nil
}
