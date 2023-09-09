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

	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/config"
)

type WenXinYiYanService struct {
	accessToken string
	apiKey      string
	secretKey   string
}

func NewWenXinYiYanService() *WenXinYiYanService {
	w := &WenXinYiYanService{}
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

}

func (w *WenXinYiYanService) Query(content string) (string, error) {
	if w.accessToken == "" {
		return "", errors.New("accessToken is nil")
	}
	requestBody := &wenXinYiYanReqBody{
		Messages: []message{
			{
				Role:    "user",
				Content: content,
			},
		},
	}
	requestData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("access_token", w.accessToken)
	reqURL := fmt.Sprintf("%s?%s", wenxinBaseUrl, params.Encode())
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
	// todo xiongyun 可删
	logs.Info(string(body))
	wenXinResp := &wenXinYiYanResp{}
	err = json.Unmarshal(body, &wenXinResp)
	if err != nil {
		return "", err
	}

	return wenXinResp.Result, nil
}

func (w *WenXinYiYanService) PostQuery() {

}
