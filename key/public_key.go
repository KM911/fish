package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

type KeyPair struct {
	PublicKey  string
	PrivateKey string
}

func NewKeyPair() *KeyPair {
	publicKey, privateKey := GenerateKeyPair()
	return &KeyPair{PublicKey: publicKey, PrivateKey: privateKey}
}

func (k *KeyPair) EncryptString(message string) string {
	return Encrypt(message, k.PublicKey)
}

func (k *KeyPair) DecryptString(encryptedMessage string) string {
	return Decrypt(encryptedMessage, k.PrivateKey)
}

func (k *KeyPair) EncryptByte(message []byte) []byte {
	return EncryptByte(message, []byte(k.PublicKey))
}

func GenerateKeyPair() (public_key, private_key string) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	pubKey := &key.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		panic(err)
	}
	pemPublicKey := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: publicKeyBytes})

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(key)
	pemPrivateKey := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes})

	return string(pemPublicKey), string(pemPrivateKey)
}

func GenerateKeyPairByte() ([]byte, []byte) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	pubKey := &key.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		panic(err)
	}
	pemPublicKey := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: publicKeyBytes})

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(key)
	pemPrivateKey := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes})

	return pemPublicKey, pemPrivateKey
}

func Encrypt(message string, ak string) string {
	block, _ := pem.Decode([]byte(ak))
	pubKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, pubKeyInterface.(*rsa.PublicKey), []byte(message))
	if err != nil {
		panic(err)
	}
	return string(encrypted)
}
func EncryptByte(message []byte, ak []byte) []byte {
	block, _ := pem.Decode(ak)
	pubKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, pubKeyInterface.(*rsa.PublicKey), message)
	if err != nil {
		panic(err)
	}
	return encrypted
}

func Decrypt(encryptedMessage string, sk string) string {
	block, _ := pem.Decode([]byte(sk))
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, []byte(encryptedMessage))
	if err != nil {
		panic(err)
	}
	return string(decrypted)
}
func DecryptByte(encryptedMessage []byte, sk []byte) string {
	block, _ := pem.Decode(sk)
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, encryptedMessage)
	if err != nil {
		panic(err)
	}
	return string(decrypted)
}
