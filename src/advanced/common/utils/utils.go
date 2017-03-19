package utils

import (
	"math/rand"
	"time"
)

var defaultCharset = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-")

func RandomString(length int, charsets ...[]byte) string {
	if length < 1 {
		return ""
	}

	var charset []byte
	if len(charsets) == 1 && len(charsets[0]) > 0 {
		charset = charsets[0]
	} else {
		charset = defaultCharset
	}

	data := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		data[i] = charset[r.Intn(len(charset)-1)]
	}

	return string(data)
}
