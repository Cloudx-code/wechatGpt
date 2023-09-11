package config

import (
	"fmt"

	"wechatGpt/common/logs"
	"wechatGpt/common/utils"

	"github.com/spf13/viper"
)

func InitStaticConfig(fileName string) {
	v := viper.New()
	v.SetConfigFile(fileName)
	v.SetConfigType("json")

	if err := v.ReadInConfig(); err != nil {
		logs.Error("[InitStaticConfig],viper read in config err: %s", err.Error())
		panic(err)
	}
	if err := v.Unmarshal(&StaticConf); err != nil {
		logs.Error("[InitStaticConfig],viper decode config err: %s", err.Error())
		panic(err)
	}
}

func InitDynamicConfig(fileName string) {
	v := viper.New()
	v.SetConfigFile(fileName)
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		logs.Error("[InitStaticConfig],viper read in config err: %s", err.Error())
		panic(err)
	}
	var result map[string]interface{}
	if err := viper.UnmarshalKey("GroupInfoMap", &result); err != nil {
		panic(err)
	}
	type EmptyStruct struct{}
	// 将map[string]interface{}转换为map[string]struct{}
	finalResult := make(map[string]EmptyStruct)
	for k := range result {
		finalResult[k] = EmptyStruct{}
	}

	fmt.Println(finalResult)

	if err := v.Unmarshal(&DynamicConf); err != nil {
		logs.Error("[InitStaticConfig],viper decode config err: %s", err.Error())
		panic(err)
	}
	//err := getDynamicConfig(v)
	//if err != nil {
	//	logs.Error("fail to getDynamicConfig,err:%v", err)
	//	panic(err)
	//}
	//go func() {
	//	ticker := time.NewTicker(20 * time.Second)
	//	for range ticker.C {
	//		getDynamicConfig(v)
	//		fmt.Println("更新测试", utils.Encode(DynamicConf))
	//	}
	//}()
}

func getDynamicConfig(v *viper.Viper) error {
	if err := v.ReadInConfig(); err != nil {
		logs.Error("viper read in config err: %s", err.Error())
		return err
	}
	if err := v.Unmarshal(&DynamicConf); err != nil {
		logs.Error("viper decode config err: %s", err.Error())
		return err
	}
	fmt.Println(utils.Encode(DynamicConf))
	return nil
}
