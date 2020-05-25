package strings2

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func ByteToInt64(buf []byte) (int64, error) {
	value, err := strconv.ParseInt(string(buf), 10, 64)
	return value, err
}
