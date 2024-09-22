package image

type Image interface {
	// GetImgUrl 获取图片链接
	GetImgUrl(prompt string) (string, error)
	// GetName 获取模型名称
	GetName() string
}
