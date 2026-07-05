package secretutil

import "encoding/base64"

// 获取Base64加密内容
func ToBase64(source string) string {
	dataBase64 := []byte(source)
	return base64.StdEncoding.EncodeToString(dataBase64)
}

// 获取Base64解密内容
func FromBase64(dest string) (string, error) {
	// base64解密
	dataFromBase64, err := base64.StdEncoding.DecodeString(dest)
	if err != nil {
		return "", err
	}
	// byte 和 string 类型之间的转换 https://www.cnblogs.com/smallleiit/p/10282371.html
	result := string(dataFromBase64[:])
	return result, err
}
