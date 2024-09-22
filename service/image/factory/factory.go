package factory

import (
	"wechatGpt/common/consts"
	"wechatGpt/dao/local_cache"
	"wechatGpt/service/image"
	"wechatGpt/service/image/dall_e_3"
)

// GetImageService 获取具体图片模型服务
func GetImageService(sendId string) image.Image {
	modelName := local_cache.GetCurrentModel(sendId)

	switch modelName {
	case consts.ModelNameImgDallE3:
		return dall_e_3.NewDallE3Service(sendId)
	default:
		// 默认Dall-E-3
		return dall_e_3.NewDallE3Service(sendId)
	}
}
