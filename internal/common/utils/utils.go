package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Sum(s string) string {
	hash := md5.Sum([]byte(s))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func GenReqId(s string) string {
	return GetMD5Sum(s)
}
