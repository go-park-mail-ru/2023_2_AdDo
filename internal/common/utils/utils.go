package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"io"
	"regexp"
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

func IsRussianLatinDigitUnderscore(value string) bool {
	// Regular expression pattern for matching Russian letters, Latin letters, digits, and underscore
	pattern := "^[А-Яа-яA-Za-z0-9_]+$"

	// Create a regex object using the compiled pattern
	regex := regexp.MustCompile(pattern)

	// Check if the value matches the regex pattern
	return regex.MatchString(value)
}
