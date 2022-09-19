package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"os"
	"regexp"
)

func decryptLog(text string) string {

	cipherList := parseLog(&text)

	sizeChanging := 0

	for _, item := range cipherList {

		cipher := text[item[0]-sizeChanging : item[1]-sizeChanging]

		plainText := decryptLogSnippet(cipher[3 : len(cipher)-4])

		text = text[:item[0]-sizeChanging] + plainText + text[item[1]-sizeChanging:]

		sizeChanging = sizeChanging + ((item[1] - item[0]) - len(plainText))
	}

	return text
}

func parseLog(text *string) [][]int {

	r := regexp.MustCompile("<" + regexp.QuoteMeta("?") + ">.*?</" + regexp.QuoteMeta("?") + ">")
	matches := r.FindAllStringIndex(*text, -1)

	return matches
}

func decryptLogSnippet(cipherInBase64 string) string {

	var commonIV = []byte("opswatmetaaccess")

	// aes encryption string
	keyText := sha256.Sum256([]byte("MetaAccess_General_@#113412"))

	// Create the aes encryption algorithm
	c, err := aes.NewCipher(keyText[:])
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(keyText), err)
		os.Exit(-1)
	}

	ciphertext, _ := b64.StdEncoding.DecodeString(cipherInBase64)

	// Decrypt strings
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(ciphertext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)

	return string(plaintextCopy)
}
