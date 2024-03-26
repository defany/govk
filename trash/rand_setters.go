package trash

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randInt() int {
	return rand.Intn(10000)
}

func randIntArr(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr = append(arr, randInt())
	}
	return arr
}

func randIntn(n int) int {
	return rand.Intn(n)
}

const (
	maxArrayLength  = 1
	maxStringLength = 25
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randString() string {
	b := make([]rune, rand.Intn(maxStringLength))
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func randStringArr(length int) []string {
	arr := make([]string, length)
	for i := 0; i < length; i++ {
		arr = append(arr, randString())
	}
	return arr
}

func randBool() bool {
	return rand.Intn(2) == 1
}

func randFloat() float64 {
	return rand.NormFloat64()
}
