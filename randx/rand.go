package randx

import (
	"math/rand"
	"time"
)

const (
	Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits  = "0123456789"
)

// S 生成包含数字和字符串的长度为 n 的随机字符串
func S(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	list := []byte(Letters + Digits)
	res := make([]byte, n)
	for i := range res {
		res[i] = list[r.Intn(len(list))]
	}
	return string(res)
}
