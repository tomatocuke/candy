package candy

import (
	"math/rand"
	"time"
)

var (
	CharsetLowerLetter = []rune("abcdefghijklmnopqrstuvwxyz")
	CharsetUpperLetter = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	CharsetLetter      = append(CharsetLowerLetter, CharsetUpperLetter...)
	CharsetNumber      = []rune("0123456789")
	CharsetAlpha       = append(CharsetLetter, CharsetNumber...)
	CharsetSpecial     = []rune("!@#$%^&*()_+-=[]{}|;':\",./<>?")
	CharsetAll         = append(CharsetAlpha, CharsetSpecial...)
)

var rd = rand.New(rand.NewSource(time.Now().UnixNano()))

// 随机数
func RandInt(max int) int {
	return rd.Intn(max)
}

// 随机字符串，自定义字符集
func RandString(size int, charset []rune) string {
	if size <= 0 {
		panic("size must be greater than 0")
	}
	if len(charset) == 0 {
		panic("charset must not be empty")
	}

	b := make([]rune, size)
	charsetLen := len(charset)
	for i := range b {
		b[i] = charset[rd.Intn(charsetLen)]
	}

	return string(b)
}
