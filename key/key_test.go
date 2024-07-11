package key

import (
	"fmt"
	"testing"
)

func TestKey(t *testing.T) {
	ak, sk := GenerateKeyPair()
	fmt.Println(len(ak), len(sk))
	// 459 1675
	message := "I love you"
	encrypted := Encrypt(message, ak)
	fmt.Println(encrypted)
	decrypted := Decrypt(encrypted, sk)
	if message != decrypted {
		t.Error("Test failed")
	}
}

func BenchmarkStringKey(b *testing.B) {
	// 112779731 ns =
	message := "I love you"
	for i := 0; i < b.N; i++ {
		ak, sk := GenerateKeyPair()
		encrypted := Encrypt(message, ak)
		decrypted := Decrypt(encrypted, sk)
		_ = decrypted
	}
}

// 86554425
// 85.673262ms
func BenchmarkByteKey(b *testing.B) {
	message := []byte("I love you")

	for i := 0; i < b.N; i++ {
		ak, sk := GenerateKeyPairByte()
		encrypted := EncryptByte(message, ak)
		decrypted := DecryptByte(encrypted, sk)
		_ = decrypted
	}
}

func BenchmarkByteToStringKey(b *testing.B) {
	message := []byte("I love you")

	for i := 0; i < b.N; i++ {
		ak, sk := GenerateKeyPairByte()
		encrypted := Encrypt(string(message), string(ak))
		decrypted := Decrypt(encrypted, string(sk))
		_ = decrypted
	}
}
