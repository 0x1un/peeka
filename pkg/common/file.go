package common

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func ComputeFileSHA(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	sha := md5.New()
	_, err = io.Copy(sha, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sha.Sum(nil)), nil
}

func IsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
