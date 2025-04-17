package base

import (
	"os"
)

// AudioToText 语音转文字
func (c *AppClient) AudioToText(filePath string, f *os.File, user string) (text string, err error) {
	if user == "" {
		user = c.GetUser()
	}
	// todo 没有语音转文字模型 暂未实现
	return
}
