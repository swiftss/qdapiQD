package sign

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

func Encrypt3DES(plaintext string, key []byte, iv []byte) (string, error) {
	// 创建 3DES 密钥
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}

	// 填充明文数据，确保长度是块大小的倍数
	plaintextBytes := []byte(plaintext)
	plaintextBytes = pad(plaintextBytes, block.BlockSize())

	// 创建 CBC 加密器
	mode := cipher.NewCBCEncrypter(block, iv)

	// 加密
	ciphertext := make([]byte, len(plaintextBytes))
	mode.CryptBlocks(ciphertext, plaintextBytes)

	// 返回加密后的数据，使用 Base64 编码
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// 解密函数
func Decrypt3DES(encryptedBase64 string, key []byte, iv []byte) (string, error) {
	// Base64 解码密文
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", err
	}

	// 创建 3DES 密钥
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}

	// 创建 CBC 解密器
	mode := cipher.NewCBCDecrypter(block, iv)

	// 解密
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// 去除填充并返回解密后的明文
	plaintext = unpad(plaintext, block.BlockSize())
	return string(plaintext), nil
}

// 填充函数，确保数据是块大小的倍数
func pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := make([]byte, padding)
	for i := range padText {
		padText[i] = byte(padding)
	}
	return append(data, padText...)
}

// 去除填充
func unpad(data []byte, blockSize int) []byte {
	length := len(data)
	padding := int(data[length-1])
	return data[:length-padding]
}
