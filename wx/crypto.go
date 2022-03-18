package wx

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// PaddingMode aes padding mode
type PaddingMode string

const (
	// ZERO zero padding mode
	ZERO PaddingMode = "ZERO"
	// PKCS5 PKCS#5 padding mode
	PKCS5 PaddingMode = "PKCS#5"
	// PKCS7 PKCS#7 padding mode
	PKCS7 PaddingMode = "PKCS#7"
)

// PemBlockType pem block type which taken from the preamble.
type PemBlockType string

const (
	// RSAPKCS1 private key in PKCS#1
	RSAPKCS1 PemBlockType = "RSA PRIVATE KEY"
	// RSAPKCS8 private key in PKCS#8
	RSAPKCS8 PemBlockType = "PRIVATE KEY"
)

// AESCrypto is the interface for aes crypto.
type AESCrypto interface {
	// Encrypt encrypts the plain text.
	Encrypt(plainText []byte) ([]byte, error)

	// Decrypt decrypts the cipher text.
	Decrypt(cipherText []byte) ([]byte, error)
}

type cbccrypto struct {
	key  []byte
	iv   []byte
	mode PaddingMode
}

func (c *cbccrypto) Encrypt(plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.key)

	if err != nil {
		return nil, err
	}

	if len(c.iv) != block.BlockSize() {
		return nil, errors.New("IV length must equal block size")
	}

	switch c.mode {
	case ZERO:
		plainText = ZeroPadding(plainText, block.BlockSize())
	case PKCS5:
		plainText = PKCS5Padding(plainText, block.BlockSize())
	case PKCS7:
		plainText = PKCS5Padding(plainText, len(c.key))
	}

	cipherText := make([]byte, len(plainText))

	blockMode := cipher.NewCBCEncrypter(block, c.iv)
	blockMode.CryptBlocks(cipherText, plainText)

	return cipherText, nil
}

func (c *cbccrypto) Decrypt(cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.key)

	if err != nil {
		return nil, err
	}

	if len(c.iv) != block.BlockSize() {
		return nil, errors.New("IV length must equal block size")
	}

	plainText := make([]byte, len(cipherText))

	blockMode := cipher.NewCBCDecrypter(block, c.iv)
	blockMode.CryptBlocks(plainText, cipherText)

	switch c.mode {
	case ZERO:
		plainText = ZeroUnPadding(plainText)
	case PKCS5:
		plainText = PKCS5Unpadding(plainText, block.BlockSize())
	case PKCS7:
		plainText = PKCS5Unpadding(plainText, len(c.key))
	}

	return plainText, nil
}

// NewCBCCrypto returns a new aes-cbc crypto.
func NewCBCCrypto(key, iv []byte, mode PaddingMode) AESCrypto {
	return &cbccrypto{
		key:  key,
		iv:   iv,
		mode: mode,
	}
}

type ecbcrypto struct {
	key  []byte
	mode PaddingMode
}

func (c *ecbcrypto) Encrypt(plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.key)

	if err != nil {
		return nil, err
	}

	switch c.mode {
	case ZERO:
		plainText = ZeroPadding(plainText, block.BlockSize())
	case PKCS5:
		plainText = PKCS5Padding(plainText, block.BlockSize())
	case PKCS7:
		plainText = PKCS5Padding(plainText, len(c.key))
	}

	cipherText := make([]byte, len(plainText))

	blockMode := NewECBEncrypter(block)
	blockMode.CryptBlocks(cipherText, plainText)

	return cipherText, nil
}

func (c *ecbcrypto) Decrypt(cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.key)

	if err != nil {
		return nil, err
	}

	plainText := make([]byte, len(cipherText))

	blockMode := NewECBDecrypter(block)
	blockMode.CryptBlocks(plainText, cipherText)

	switch c.mode {
	case ZERO:
		plainText = ZeroUnPadding(plainText)
	case PKCS5:
		plainText = PKCS5Unpadding(plainText, block.BlockSize())
	case PKCS7:
		plainText = PKCS5Unpadding(plainText, len(c.key))
	}

	return plainText, nil
}

// NewECBCrypto returns a new aes-ecb crypto.
func NewECBCrypto(key []byte, mode PaddingMode) AESCrypto {
	return &ecbcrypto{
		key:  key,
		mode: mode,
	}
}

func ZeroPadding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{0}, padding)

	return append(cipherText, padText...)
}

func ZeroUnPadding(plainText []byte) []byte {
	return bytes.TrimRightFunc(plainText, func(r rune) bool {
		return r == rune(0)
	})
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize

	if padding == 0 {
		padding = blockSize
	}

	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padText...)
}

func PKCS5Unpadding(plainText []byte, blockSize int) []byte {
	l := len(plainText)
	unpadding := int(plainText[l-1])

	if unpadding < 1 || unpadding > blockSize {
		unpadding = 0
	}

	return plainText[:(l - unpadding)]
}

// RSAEncryptOEAP rsa encrypt with PKCS #1 OEAP.
func RSAEncryptOEAP(plainText, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)

	if block == nil {
		return nil, errors.New("invalid rsa public key for pem.Decode")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	key, ok := pubKey.(*rsa.PublicKey)

	if !ok {
		return nil, errors.New("invalid rsa public key, expects rsa.PublicKey")
	}

	return rsa.EncryptOAEP(sha1.New(), rand.Reader, key, plainText, nil)
}

// RSADecryptOEAP rsa decrypt with PKCS #1 OEAP.
func RSADecryptOEAP(cipherText, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)

	if block == nil {
		return nil, errors.New("invalid rsa private key for pem.Decode")
	}

	var (
		key interface{}
		err error
	)

	switch PemBlockType(block.Type) {
	case RSAPKCS1:
		key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	case RSAPKCS8:
		key, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	}

	if err != nil {
		return nil, err
	}

	rsaKey, ok := key.(*rsa.PrivateKey)

	if !ok {
		return nil, errors.New("invalid rsa private key, expects rsa.PrivateKey")
	}

	return rsa.DecryptOAEP(sha1.New(), rand.Reader, rsaKey, cipherText, nil)
}

// ------------- AES-256-ECB -------------

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}

	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}

	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}

	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}

	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])

		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
