package consts

import (
	"wechatGpt/model"
)

type ModelName string

const (
	ModelNameWenXinYiYan   ModelName = "文心一言"
	ModelNameWeTab         ModelName = "gpt3.5插件版"
	ModelNameGPT4          ModelName = "gpt4"
	ModelNameGPT4O         ModelName = "gpt4o"
	ModelNameImgDallE3     ModelName = "画图"
	ModelNameAdministrator ModelName = "管理员模式"
)

var ModelInfoMap = map[ModelName]*model.ModelDetail{
	ModelNameWenXinYiYan: {
		ModelIntroduce: "",
	},
	ModelNameWeTab: {
		ModelIntroduce: "",
	},
	ModelNameGPT4: {
		ModelIntroduce: "",
	},
	ModelNameAdministrator: {
		ModelIntroduce: "",
	},
	ModelNameImgDallE3: {
		ModelIntroduce: "",
	},
	ModelNameGPT4O: {
		ModelIntroduce: "",
	},
}
