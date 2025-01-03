package shortener

import (
	"crypto/md5"
	"encoding/binary"
)

func encodeUrl(url string) string {
	const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	hash := md5.Sum([]byte(url))
	firstSevenBytes := hash[:7]
	paddedBytes := append([]byte{0}, firstSevenBytes...)
	dec := binary.BigEndian.Uint64(paddedBytes)

	var result string
	for dec > 0 {
		remainder := dec % 62
		result = string(base62[remainder]) + result
		dec /= 62
	}

	return result
}
