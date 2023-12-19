package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/google/uuid"
)

func GetImageName(extension string) string {
	return uuid.New().String() + "." + extension
}

func GetMD5Sum(s string) string {
	hash := md5.Sum([]byte(s))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func GenReqId(s string) string {
	return GetMD5Sum(s)
}

func GetReaderFromBytes(in []byte) io.Reader {
	result := bytes.NewReader(in)
	return result
}

const SecondInMinute = 60

func CastTimeToString(duration int) string {
	minutes := duration / SecondInMinute
	seconds := duration % SecondInMinute

	return fmt.Sprintf("%d:%02d", minutes, seconds)
}

// func Encode(s string) string {
// 	data := base64.StdEncoding.EncodeToString([]byte(s))
// 	return string(data)
// }
//
// func Decode(s string) (string, error) {
// 	data, err := base64.StdEncoding.DecodeString(s)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	return string(data), nil
// }
