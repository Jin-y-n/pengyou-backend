package security

import (
	"math/rand"
)

const numberBytes = "0123456789"
const characterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const defaultLength = 6 // 验证码长度

func GenerateCaptcha() string {
	// 初始化随机数生成器
	rand.New(rand.NewSource(123))

	// 创建一个长度为length的字节切片
	b := make([]byte, defaultLength)
	for i := range b {
		// 从numberBytes中随机选择一个字符
		b[i] = numberBytes[rand.Intn(len(numberBytes))]
	}
	return string(b)
}

func GenerateCaptchaWithLength(length int) string {
	// 初始化随机数生成器
	rand.New(rand.NewSource(123))

	// 创建一个长度为length的字节切片
	b := make([]byte, length)
	for i := range b {
		// 从numberBytes中随机选择一个字符
		b[i] = numberBytes[rand.Intn(len(numberBytes))]
	}

	return string(b)
}

func GenerateCaptchaWithCharacter(length int) string {
	// 初始化随机数生成器
	rand.New(rand.NewSource(123))

	// 创建一个长度为length的字节切片
	b := make([]byte, length)
	for i := range b {
		// 从characterBytes中随机选择一个字符
		b[i] = characterBytes[rand.Intn(len(characterBytes))]
	}

	return string(b)
}
