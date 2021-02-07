package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

// GetHmac256 !
func GetHmac256(s string) string {
	mac := hmac.New(sha256.New, []byte(os.Getenv("SECRET_KEY")))
	bs := []byte(s)
	mac.Write(bs)
	r := hex.EncodeToString(mac.Sum(nil))
	return r
}
