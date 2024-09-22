package dall_e_3

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/config"
)

type DallE3Service struct {
	ak          string
	senderId    string
	downloadUrl string
}

const (
	BaseUrl  = "https://api.openai.com/v1/images/generations"
	BaseUrl2 = "https://ng.dwai.world/v1/images/generations"
)

func NewDallE3Service(senderId string) *DallE3Service {
	return &DallE3Service{
		ak:       "Bearer " + config.GetLLMConfig().GPTAk,
		senderId: senderId,
	}
}
func (d *DallE3Service) GetImgUrl(prompt string) (string, error) {
	buf, err := d.getReqBuf(prompt)
	if err != nil {
		logs.Error("fail to getReqBuf,[DallE3Service],err:%v", err)
		return "", err
	}
	req, err := http.NewRequest("POST", BaseUrl2, bytes.NewBuffer(buf))
	if err != nil {
		logs.Error("fail to NewRequest,[DallE3Service],err:%v", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", d.ak)
	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("fail to client.Do,[DallE3Service],err:%v", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error("fail to ReadAll,[DallE3Service],err:%v", err)
		return "", err
	}
	gptResp := &DallE3Resp{}
	err = json.Unmarshal(body, &gptResp)
	if err != nil {
		logs.Error("fail to json.Unmarshal,[DallE3Service],err:%v", err)
		return "", err
	}
	if len(gptResp.Data) != 1 || gptResp.Data[0].Url == "" {
		logs.Error("fail to check len,[DallE3Service],detail:%v", utils.Encode(gptResp))
		return "", errors.New("生成url失败")
	}
	logs.Info("check url : %v", gptResp.Data[0].Url)
	return gptResp.Data[0].Url, nil
}

func (d *DallE3Service) getReqBuf(content string) ([]byte, error) {
	payload := map[string]interface{}{
		"model":  "dall-e-3",
		"prompt": content,
		"n":      1,
		"size":   "1024x1024",
	}
	return json.Marshal(payload)
}

func (d *DallE3Service) GetName() string {
	return "Dall-E-3"
}
