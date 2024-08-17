package check

import (
	"regexp"
)

// 用于名字格式校验的正则表达式模式
// 名字可能由字母（包括大小写）、数字、下划线和连字符组成，且可以包含空格，长度在4 - 20之间
var nameRegex = regexp.MustCompile(`^(?=.{4,20}$)[\p{L}\d*\p{L}-_]+$`)

// 加强密码强度要求的正则表达式模式
// 要求必须包含至少一个数字、一个字母，长度在8到20个字符之间
var passwordRegex = regexp.MustCompile(`^(?=.*[a-zA-Z])(?=.*\d).{8,20}$`)

// 电话号码校验的正则表达式模式，仅适用于中国手机号码
var phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)

// 电子邮件格式校验的正则表达式模式
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// 验证提供的密码是否符合要求
func checkPassword(password string) bool {
	if password == "" {
		return false
	}
	return passwordRegex.MatchString(password)
}

// 验证提供的电话号码是否有效
func checkPhone(phone string) bool {
	if phone == "" {
		return false
	}
	return phoneRegex.MatchString(phone)
}

// 验证提供的电子邮件是否有效
func checkEmail(email string) bool {
	if email == "" {
		return false
	}
	return emailRegex.MatchString(email)
}

// 验证提供的名字是否有效
func checkName(name string) bool {
	if name == "" {
		return false
	}
	return nameRegex.MatchString(name)
}
