package config

import (
	"wechatGpt/common/logs"

	"github.com/spf13/viper"
)

func Init(fileName string) {
	v := viper.New()
	v.SetConfigFile(fileName)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		logs.Error("viper read in config err: %s", err.Error())
		panic(err)
	}

	if err := v.Unmarshal(&conf); err != nil {
		logs.Error("viper decode config err: %s", err.Error())
		panic(err)
	}
}
