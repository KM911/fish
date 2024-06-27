package key

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

var ()

type PrivateKey struct {
	key []byte
}

func NewPrivateKey() *PrivateKey {
	return &PrivateKey{key: GetPrivateKey()}
}

func (p *PrivateKey) DecryptByte(ciphertext []byte) ([]byte, error) {
	return PrivateDecryptByte(p.key, ciphertext)
}

func (p *PrivateKey) EncryptByte(plaintext []byte) ([]byte, error) {
	return PrivateEncryByte(p.key, plaintext)
}

func (p *PrivateKey) DecryptString(ciphertext string) (string, error) {
	plaintext, err := p.DecryptByte([]byte(ciphertext))
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func (p *PrivateKey) EncryptString(plaintext string) (string, error) {
	ciphertext, err := p.EncryptByte([]byte(plaintext))
	if err != nil {
		return "", err
	}
	return string(ciphertext), nil
}

// 生成密钥
func GetPrivateKey() []byte {
	key := make([]byte, 32)
	rand.Read(key)
	return key
}

// 加密
func PrivateEncryByte(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 使用CBC模式，需要初始化向量
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 随机生成nonce
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	// 加密
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// 解密
func PrivateDecryptByte(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 使用CBC模式，需要初始化向量
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 获取nonce
	nonceSize := aesGCM.NonceSize()
	nonce := ciphertext[:nonceSize]
	ciphertext = ciphertext[nonceSize:]

	// 解密
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
