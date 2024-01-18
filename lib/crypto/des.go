package crypto

import (
	"crypto/des"
	"errors"
)

// DESEncryptECB DES-ECB 加密
func DESEncryptECB(key, data []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	data = pkcs7padding(data, block.BlockSize())

	bm := NewECBEncrypter(block)
	if len(data)%bm.BlockSize() != 0 {
		return nil, errors.New("input not full blocks")
	}

	out := make([]byte, len(data))
	bm.CryptBlocks(out, data)

	return out, nil
}

// DESDecryptECB DES-ECB 解密
func DESDecryptECB(key, data []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bm := NewECBDecrypter(block)
	if len(data)%bm.BlockSize() != 0 {
		return nil, errors.New("input not full blocks")
	}

	out := make([]byte, len(data))
	bm.CryptBlocks(out, data)

	return pkcs7unpadding(out), nil
}
