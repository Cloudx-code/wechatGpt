package config

import (
	"fmt"
	"testing"

	"wechatGpt/common/utils"
)

func TestName(t *testing.T) {
	Init("config.yaml")
	fmt.Println(utils.Encode(conf))
}
