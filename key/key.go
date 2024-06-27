package key

type Key interface {
	DecryptByte([]byte) ([]byte, error)
	EncryptByte([]byte) ([]byte, error)

	DecryptString(string) (string, error)
	EncryptString(string) (string, error)
}
