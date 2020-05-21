package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Hash(i []byte) string {
	hash := md5.Sum(i)
	r := strings.ToLower(hex.EncodeToString(hash[:]))
	return r
}

func MergeBytes(bs ...[]byte) []byte {
	i := 0
	for _, b := range bs {
		i += len(b)
	}
	r := make([]byte, 0, i)
	for _, b := range bs {
		r = append(r, b...)
	}
	return r
}
