package weTab

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"wechatGpt/common/consts"
	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/config"
	"wechatGpt/dao/local_cache"
)

type WeTabService struct {
	authorization  string // 用户验证信息
	senderId       string // 发送者Id
	conversationId string // 对话Id
}

func NewWeTabService(senderId string) *WeTabService {
	w := &WeTabService{
		senderId: senderId,
	}
	if len(authorization) == 0 {
		authorization = w.getAuthorization()
	}
	w.authorization = authorization
	return w
}

func (w *WeTabService) GetName() string {
	return "GPT3.5插件版WeTab"
}

func (w *WeTabService) getAuthorization() string {

	loginInfo := &LoginInfo{
		Email:    config.GetLLMConfig().WeTabEmail,
		Password: config.GetLLMConfig().WeTabPwd,
	}
	reqBuffer, _ := json.Marshal(loginInfo)

	req, _ := http.NewRequest("POST", loginBaseUrl, bytes.NewBuffer(reqBuffer))

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Origin", "chrome-extension://aikflfpejipbpjdlfabpgclhblkpaafo")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	req.Header.Add("i-app", "hitab")
	req.Header.Add("i-branch", "zh")
	req.Header.Add("i-lang", "zh-CN")
	req.Header.Add("i-platform", "chrome")
	req.Header.Add("i-version", "1.2.2")
	req.Header.Add("sec-ch-ua", `"Chromium";v="116", "Not)A;Brand";v="24", "Google Chrome";v="116"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", `"macOS"`)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("fail to getAuthorization Do,err:%v", err)
		return ""
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	loginResp := &LoginResp{}
	err = json.Unmarshal(body, &loginResp)
	if err != nil || loginResp.Code != 0 {
		logs.Error("fail to getAuthorization Unmarshal,err:%v", err)
		logs.Error("fail to getAuthorization Unmarshal,info:%v", utils.Encode(loginResp))
		return ""
	}
	return loginResp.Data.Token
}

func (w *WeTabService) PreQuery() {
	w.conversationId = w.getWebTabContext()
}

func (w *WeTabService) Query(text string) (string, error) {
	reqBody := &WebTabReqBody{
		Prompt:      text,
		AssistantID: "",
	}
	if len(w.conversationId) > 0 {
		reqBody.ConversationId = w.conversationId
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		logs.Error("Error marshalling reqBody:%v", err)
		return "", err
	}

	req, _ := http.NewRequest("POST", baseUrl, bytes.NewBuffer(reqBodyBytes))

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Authorization", "Bearer "+w.authorization)
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

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("fail to client.Do,[WeTabService],err:%v", err)
		return "", err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	logs.Info("weTab:%v", string(body))

	return w.parseResp(string(body))
}

// 获取gpt返回内容及conversationId
func (w *WeTabService) parseResp(respContent string) (string, error) {
	// 看看是否出错
	errStruct := &ErrStruct{}
	utils.Decode(respContent, &errStruct)
	if len(errStruct.Message) > 0 {
		// 重新获取Token
		SetAuthorization2Nil()
		return errStruct.Message + "，请重试", nil
	}

	var content string
	jsonStrArray := w.splitString(string(respContent), "_e79218965e_")
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
			logs.Info("Content:%v", data["data"].(map[string]interface{})["content"])
			if _, ok := data["data"].(map[string]interface{}); !ok {
				return content, nil
			}
			if id, ok := data["data"].(map[string]interface{})["conversationId"].(string); ok {
				logs.Info("测试下conversationId:%v", id)
				w.conversationId = id
			}
			if _, ok := data["data"].(map[string]interface{})["content"].(string); !ok {
				return content, nil
			}
			content += data["data"].(map[string]interface{})["content"].(string)
		} else {
			logs.Error("Error:%v", err)
		}
	}
	return content, nil
}

func (w *WeTabService) splitString(s string, sep string) []string {
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

func (w *WeTabService) PostQuery() {
	w.setWebTabContext()
	// 如果继续聊天，继续weTab模式
	local_cache.SetCurrentModel(w.senderId, consts.ModelNameWeTab)
}

// GetWebTabContext  获取WebTab上下文
func (w *WeTabService) getWebTabContext() string {
	key := consts.RedisKeyWebTabContext + w.senderId
	v, ok := local_cache.Get(key)
	if !ok {
		return ""
	}
	if conversationId, ok := v.(string); ok {
		return conversationId
	}
	return ""
}

func (w *WeTabService) setWebTabContext() {
	key := consts.RedisKeyWebTabContext + w.senderId
	local_cache.Set(key, w.conversationId)
}
