package antchain

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math/big"
)

const (
	CHAIN_CALL_FOR_BIZ = "/api/contract/chainCallForBiz"
	CHAIN_CALL         = "/api/contract/chainCall"
	SHAKE_HAND         = "/api/contract/shakeHand"
)

// Identity 链账户对应的Identity
type Identity struct {
	Data string `json:"data"`
}

// GetIdentityByName 根据链账户名称获取对应的Identity
func GetIdentityByName(name string) *Identity {
	h := sha256.New()
	h.Write([]byte(name))
	return &Identity{
		Data: base64.StdEncoding.EncodeToString(h.Sum(nil)),
	}
}

// TokenID 链上资产(NFT)的唯一标识
type TokenID *big.Int

// GetTokenID 获取token(如：md5值)对应的tokenID(uint256)
func GetTokenID(token string) TokenID {
	v, _ := big.NewInt(0).SetString(token, 16)
	return TokenID(v)
}

// ParseOutput 解析合约方法返回的output
func ParseOutput(data string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
