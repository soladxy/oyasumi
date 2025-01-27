package util

import "unicode"

func IsServiceNameValid(input string) bool {
	for _, r := range input {
		if !unicode.IsLetter(r) && r != '_' {
			return false
		}
	}
	return true
}

func CapitalizeFirstLetter(input string) string {
	if len(input) == 0 {
		return input // 如果字符串为空，直接返回
	}

	runes := []rune(input)               // 将字符串转换为 rune 切片，支持 Unicode
	runes[0] = unicode.ToUpper(runes[0]) // 将首字符转换为大写
	return string(runes)                 // 将 rune 切片转换回字符串
}
