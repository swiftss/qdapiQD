package sign

import (
	"strings"
)

func EncryptSDKSign(input string) (string, error) {
	encoded, err := Encrypt3DES(input, SDKSignPass, SDKSignIV)
	if err != nil {
		return "", err
	}
	// Split into chunks of 60 characters
	var chunks []string
	for i := 0; i < len(encoded); i += 60 {
		end := i + 60
		if end > len(encoded) {
			end = len(encoded)
		}
		chunks = append(chunks, encoded[i:end])
	}
	return strings.Join(chunks, " "), nil
}

func DecryptSDKSign(encodedSign string) (string, error) {
	// 移除所有空白字符
	encodedSign = strings.ReplaceAll(encodedSign, " ", "")
	return Decrypt3DES(encodedSign, SDKSignPass, SDKSignIV)
}

// EncryptQDInfo encrypts the QDInfo string using DES-EDE3-CBC.
func EncryptQDInfo(info string) (string, error) {
	return Encrypt3DES(info, QDInfoPass, QDInfoIV)
}

// DecryptQDInfo decrypts the given input string using DES-EDE3-CBC.
func DecryptQDInfo(input string) (string, error) {
	return Decrypt3DES(input, QDInfoPass, QDInfoIV)
}
