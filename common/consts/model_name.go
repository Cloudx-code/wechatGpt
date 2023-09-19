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
	ModelNameAdministrator = "管理员模式"
)

var ModelInfoMap = map[ModelName]*model.ModelDetail{
	ModelNameWenXinYiYan: {
		ModelIntroduce: "",
	},
	ModelNameWeTab: {
		ModelIntroduce: "",
	},
	ModelNameGPT: {
		ModelIntroduce: "",
	},
	ModelNameAdministrator: {
		ModelIntroduce: "",
	},
	//ModelNameImageTest:   nil,
}
