package base83

import (
	"fmt"
	"math"
	"strings"
)

const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz#$%*+,-.:;=?@[]^_{|}~"
const base = len(chars)

// Encode encodes an integer value into a base83 string of the given length.
func Encode(val int, length int) (str string) {
	carry := int(math.Pow(float64(base), float64(length-1)))
	for i := 0; i < length; i++ {
		idx := val / carry % base
		str += string(chars[idx])
		carry /= base
	}
	return str
}

// Decode decodes a base83 string into an integer value.
func Decode(str string) (val int, err error) {
	for _, r := range str {
		idx := strings.IndexRune(chars, r)
		if idx == -1 {
			return 0, fmt.Errorf("base83: invalid input %v", r)
		}
		val = val*base + idx
	}
	return val, nil
}
