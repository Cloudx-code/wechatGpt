package config

import (
	"fmt"
	"testing"

	"wechatGpt/common/utils"
)

func TestStatic(t *testing.T) {
	InitStaticConfig("static_config.json")
	fmt.Println(utils.Encode(StaticConf))
}

func TestDynamic(t *testing.T) {
	InitDynamicConfig("test.json")
	fmt.Println(utils.Encode(DynamicConf))
	//	str := `{
	//  "GroupInfoMap":{
	//    "俞越机器人内部测试": {
	//
	//    }
	//  }
	//}`
	//	err := json.Unmarshal([]byte(str), &DynamicConf)
	//	if err != nil {
	//		fmt.Println("fail to ", err)
	//		return
	//	}
	//	fmt.Println(utils.Encode(DynamicConf))
}
