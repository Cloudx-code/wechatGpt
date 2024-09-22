package chat_manage

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"wechatGpt/common/logs"
	"wechatGpt/service/image/factory"
)

/*
	@Author: xy
	@Date: 2023/8/28
	@Desc: 聊天功能
*/

type ImgChatService struct {
	senderId string
	content  string
}

func NewImgChatService(senderId string, content string) *ImgChatService {
	return &ImgChatService{senderId: senderId, content: content}
}

// Chat 返回msgInfo, urlPath,
func (i *ImgChatService) Chat() (string, string, error) {
	imgModel := factory.GetImageService(i.senderId)
	// 生成图片url
	urlLink, err := imgModel.GetImgUrl(i.content)
	if err != nil {
		return "", "", err
	}
	// 下载图片并保存到本地
	urlPath, err := i.DownloadImg2Local(urlLink)
	if err != nil {
		return fmt.Sprintf("保存图片有误，url:%v", urlLink), "", err
	}
	return "图片保存成功，生成图片模型为：" + imgModel.GetName(), urlPath, nil
}

// DownloadImg2Local 保存到本地
func (i *ImgChatService) DownloadImg2Local(imgUrl string) (string, error) {
	response, err := http.Get(imgUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", err
	}
	// 创建文件
	urlPath := fmt.Sprintf("downloaded_image:%v.jpg", time.Now().Format("2006-01-02 15:04:05"))
	file, err := os.Create(urlPath) // 保存的文件名
	if err != nil {
		logs.Error("fail to Create,[ImgChatService],err:%v")
		return "", err
	}
	defer file.Close()
	// 将HTTP响应的Body内容写入文件
	_, err = io.Copy(file, response.Body)
	if err != nil {
		logs.Error("fail to copy,[ImgChatService],err:%v")
		return "", err
	}
	return urlPath, nil
}
