package consts

type ChatStatus int64

const (
	NormalChat  ChatStatus = iota // 正常聊天模式
	ModelChoose                   // 模型选择中
)
