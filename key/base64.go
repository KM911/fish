package key

import "encoding/base64"

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode decodes a base64 encoded string
func Base64Decode(str string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(decodedBytes)
}
