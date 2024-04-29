package random

import (
	"crypto/rand"
	"math/big"
)

func RandInt() int {
	v, e := rand.Int(rand.Reader, big.NewInt(10000))
	if e != nil {
		return 0
	}
	return int(v.Int64())
}

func RandIntArr(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr = append(arr, RandInt())
	}
	return arr
}

func RandIntn(n int) int {
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

func RandString() string {
	b := make([]rune, RandIntn(maxStringLength))
	for i := range b {
		b[i] = runes[RandIntn(len(runes))]
	}
	return string(b)
}

func RandStringArr(length int) []string {
	arr := make([]string, length)
	for i := 0; i < length; i++ {
		arr = append(arr, RandString())
	}
	return arr
}

func RandBool() bool {
	return RandIntn(2) == 1
}

func RandFloat() float64 {
	v, e := rand.Int(rand.Reader, big.NewInt(10000))
	if e != nil {
		return 0
	}
	r, _ := v.Float64()
	return r
}
