package generatetoken

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

type ProviderToken struct {
	Core   Payload
	secret string
}

type GenerateToken interface {
	Encrypt() (string, error)
	Decrypt() (string, error)
}

func NewProviderToken(core Payload, secret string) *ProviderToken {
	return &ProviderToken{core, secret}

}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (p *ProviderToken) Encrypt() (string, error) {
	block, err := aes.NewCipher([]byte(p.secret))

	if err != nil {
		return "", err
	}

	plainText := []byte(p.Core.UserName)

	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *ProviderToken) Decrypt(text string) error {
	block, err := aes.NewCipher([]byte(string(p.secret)))
	if err != nil {
		return err
	}

	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	p.Core.UserName = string(plainText)
	return nil
}
