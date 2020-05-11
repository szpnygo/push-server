package base

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetSha256(str []byte) string {
	hash := sha256.New()
	hash.Write(str)
	bytes := hash.Sum(nil)
	return hex.EncodeToString(bytes)
}
