package wenxin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"wechatGpt/common/consts"
	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/config"
	"wechatGpt/dao/local_cache"
)

type WenXinYiYanService struct {
	accessToken string
	apiKey      string
	senderId    string
	secretKey   string
	cacheList   []message // 上下文缓存内容
}

func NewWenXinYiYanService(sendId string) *WenXinYiYanService {
	w := &WenXinYiYanService{senderId: sendId}
	w.apiKey = config.GetLLMConfig().WenXinAk    // 请替换为您自己的API Key
	w.secretKey = config.GetLLMConfig().WenXinSk // 请替换为您自己的Secret Key
	if accessToken == "" {
		accessToken, _ = w.getAccessToken()
	}
	w.accessToken = accessToken
	return w
}

// url = `https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s`
func (w *WenXinYiYanService) getAccessToken() (string, error) {
	params := url.Values{}
	params.Set("grant_type", "client_credentials")
	params.Set("client_id", w.apiKey)
	params.Set("client_secret", w.secretKey)
	reqURL := fmt.Sprintf("%s?%s", accessTokenBaseURL, params.Encode())
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	logs.Info("Testurl:%v", reqURL)
	resp, err := client.Get(reqURL)
	if err != nil {
		logs.Error("请求失败: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("读取响应失败: %v", err)
		return "", err
	}
	accessTokenResp := &AccessTokenResp{}
	err = json.Unmarshal(bodyBytes, &accessTokenResp)
	if err != nil {
		logs.Error("fail to json.Unmarshal,[(w *WenXinYiYanService) getAccessToken],err: %v", err)
		return "", err
	}
	if accessTokenResp.Error != nil {
		logs.Error("fail to getAccessToken,accessTokenResp:%v", utils.Encode(accessTokenResp))
		return "", errors.New(*accessTokenResp.ErrorDescription)
	}
	return accessTokenResp.AccessToken, nil
}

func (w *WenXinYiYanService) PreQuery() {
	key := consts.RedisKeyWenXinContext + w.senderId
	v, ok := local_cache.Get(key)
	if !ok {
		return
	}
	if cachInfoStr, ok := v.(string); ok {
		err := utils.Decode(cachInfoStr, &w.cacheList)
		if err != nil {
			logs.Error("fail to PreQuery,err:%v")
		}
		return
	}
	return
}

func (w *WenXinYiYanService) Query(content string) (string, error) {
	if w.accessToken == "" {
		return "", errors.New("accessToken is nil")
	}
	requestBody := &wenXinYiYanReqBody{}
	w.cacheList = append(w.cacheList, message{
		Role:    "user",
		Content: content,
	})
	requestBody.Messages = w.cacheList

	requestData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("access_token", w.accessToken)
	reqURL := fmt.Sprintf("%s?%s", wenxin4BaseUrl, params.Encode())
	logs.Info("test request:%v", string(requestData))
	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
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

	logs.Info(string(body))
	wenXinResp := &wenXinYiYanResp{}
	err = json.Unmarshal(body, &wenXinResp)
	if err != nil {
		return "", err
	}
	// 字数超了则清空不存，并提示
	ifSave := w.whetherNeedSave(wenXinResp.Result)
	if !ifSave {
		wenXinResp.Result += "\n上下文超出限制，已自动清空"
	}
	return wenXinResp.Result, nil
}

// 是否需要缓存，是否超出字数
func (w *WenXinYiYanService) whetherNeedSave(result string) bool {
	// 多轮对话
	w.cacheList = append(w.cacheList, message{
		Role:    "assistant",
		Content: result,
	})
	length := 0
	for _, cache := range w.cacheList {
		length += len(cache.Content)
	}
	if length > 1800 {
		w.cacheList = make([]message, 0)
		local_cache.Set(consts.RedisKeyWenXinContext+w.senderId, utils.Encode(w.cacheList))
		return false
	}
	return true
}

func (w *WenXinYiYanService) PostQuery() {
	if len(w.cacheList) == 0 {
		return
	}
	key := consts.RedisKeyWenXinContext + w.senderId
	local_cache.Set(key, utils.Encode(w.cacheList))
}
