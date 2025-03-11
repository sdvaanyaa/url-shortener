package random

import (
	"math/rand"
	"strings"
)

// NewRandomString generates random string with given size.
func NewRandomString(length int) string {
	str := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM123456789"

	var sb strings.Builder
	sb.Grow(length)

	for i := 0; i < length; i++ {
		sb.WriteByte(str[rand.Intn(len(str))])
	}

	return sb.String()
}
