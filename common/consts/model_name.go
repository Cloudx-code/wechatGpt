package consts

import (
	"wechatGpt/model"
)

type ModelName string

const (
	ModelNameWenXinYiYan ModelName = "文心一言"
	ModelNameWeTab       ModelName = "gpt3.5插件版"
	ModelNameGPT         ModelName = "gpt3.5"
	//ModelNameImageTest   ModelName = "gpt画图"
)

var ModelInfoMap = map[ModelName]*model.ModelDetail{
	ModelNameWenXinYiYan: nil,
	ModelNameWeTab:       nil,
	ModelNameGPT:         nil,
	//ModelNameImageTest:   nil,
}
