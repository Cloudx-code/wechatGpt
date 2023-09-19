package config

import (
	"fmt"
	"time"

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

	err := getDynamicConfig(v, true)
	if err != nil {
		logs.Error("fail to getDynamicConfig,err:%v", err)
		panic(err)
	}
	go func() {
		ticker := time.NewTicker(20 * time.Second)
		for range ticker.C {
			getDynamicConfig(v, false)
		}
	}()
}

func getDynamicConfig(v *viper.Viper, isFirst bool) error {
	if err := v.ReadInConfig(); err != nil {
		logs.Error("viper read in config err: %s", err.Error())
		return err
	}
	temp := &BotDynamicConfig{}
	if err := v.Unmarshal(&temp); err != nil {
		logs.Error("viper decode config err: %s", err.Error())
		return err
	}
	if temp.IsChange || isFirst {
		DynamicConf = *temp
		fmt.Println(utils.Encode(DynamicConf))
	}
	return nil
}
