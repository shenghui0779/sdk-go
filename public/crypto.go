package public

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// AESPaddingMode aes padding mode
type AESPaddingMode string

const (
	// PKCS5 PKCS#5 padding mode
	PKCS5 AESPaddingMode = "PKCS#5"
	// PKCS7 PKCS#7 padding mode
	PKCS7 AESPaddingMode = "PKCS#7"
)

// AESCBCCrypto aes-cbc crypto
type AESCBCCrypto struct {
	key []byte
	iv  []byte
}

// NewAESCBCCrypto returns new aes-cbc crypto
func NewAESCBCCrypto(key, iv []byte) *AESCBCCrypto {
	return &AESCBCCrypto{
		key: key,
		iv:  iv,
	}
}

// Encrypt aes-cbc encrypt with PKCS#7 padding
func (c *AESCBCCrypto) Encrypt(plainText []byte, mode AESPaddingMode) ([]byte, error) {
	block, err := aes.NewCipher(c.key)

	if err != nil {
		return nil, err
	}

	if len(c.iv) != block.BlockSize() {
		return nil, errors.New("yiigo: IV length must equal block size")
	}

	switch mode {
	case PKCS5:
		plainText = c.padding(plainText, block.BlockSize())
	case PKCS7:
		plainText = c.padding(plainText, len(c.key))
	}

	cipherText := make([]byte, len(plainText))

	blockMode := cipher.NewCBCEncrypter(block, c.iv)
	blockMode.CryptBlocks(cipherText, plainText)

	return cipherText, nil
}

// Decrypt aes-cbc decrypt with PKCS#7 padding
func (c *AESCBCCrypto) Decrypt(cipherText []byte, mode AESPaddingMode) ([]byte, error) {
	block, err := aes.NewCipher(c.key)

	if err != nil {
		return nil, err
	}

	if len(c.iv) != block.BlockSize() {
		return nil, errors.New("yiigo: IV length must equal block size")
	}

	plainText := make([]byte, len(cipherText))

	blockMode := cipher.NewCBCDecrypter(block, c.iv)
	blockMode.CryptBlocks(plainText, cipherText)

	switch mode {
	case PKCS5:
		plainText = c.unpadding(plainText, block.BlockSize())
	case PKCS7:
		plainText = c.unpadding(plainText, len(c.key))
	}

	return plainText, nil
}

func (c *AESCBCCrypto) padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize

	if padding == 0 {
		padding = blockSize
	}

	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padText...)
}

func (c *AESCBCCrypto) unpadding(plainText []byte, blockSize int) []byte {
	l := len(plainText)
	unpadding := int(plainText[l-1])

	if unpadding < 1 || unpadding > blockSize {
		unpadding = 0
	}

	return plainText[:(l - unpadding)]
}

// RSAEncrypt rsa encryption with public key
func RSAEncrypt(data, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)

	if block == nil {
		return nil, errors.New("gochat: invalid rsa public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	key, ok := pubKey.(*rsa.PublicKey)

	if !ok {
		return nil, errors.New("gochat: invalid rsa public key")
	}

	return rsa.EncryptPKCS1v15(rand.Reader, key, data)
}

// RSADecrypt rsa decryption with private key
func RSADecrypt(cipherText, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)

	if block == nil {
		return nil, errors.New("gochat: invalid rsa private key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, key, cipherText)
}
