package random

import (
	"crypto/rand"
	"math/big"
)

func Int() int {
	return IntDiapason(10000)
}

func IntArr(length int) []int {
	arr := make([]int, 0, length)
	for i := 0; i < length; i++ {
		arr = append(arr, Int())
	}
	return arr
}

func IntDiapason(n int) int {
	v, e := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if e != nil {
		return 0
	}
	return int(v.Int64())
}

const (
	MaxArrayLength  = 1
	maxStringLength = 25
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func String() string {
	b := make([]rune, 0, IntDiapason(maxStringLength))
	for i := range b {
		b[i] = runes[IntDiapason(len(runes))]
	}
	return string(b)
}

func StringArr(length int) []string {
	arr := make([]string, 0, length)
	for i := 0; i < length; i++ {
		arr = append(arr, String())
	}
	return arr
}

func Bool() bool {
	return IntDiapason(2) == 1
}

func MustFloat() float64 {
	v, _ := rand.Int(rand.Reader, big.NewInt(10000))
	r, _ := v.Float64()
	return r
}
