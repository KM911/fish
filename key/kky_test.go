package key

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestKKy(t *testing.T) {
	// c740caa51ca607066039af01e8857f4f0000000000000000
	data := []byte("This is the data to be . right alkdfjaslkdfj")

	// 生成密钥
	key := generateKey()
	fmt.Println(string(key), cap(key), len(key))

	// 加密数据
	cipherText, err := encrypt(key, data)
	if err != nil {
		fmt.Println("Error encrypting data:", err)
		return
	}

	fmt.Println("Encrypted data:", hex.EncodeToString(cipherText))

	// 解密数据
	plainText, err := decrypt(key, cipherText)
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}

	fmt.Println("Decrypted data:", string(plainText))
}
