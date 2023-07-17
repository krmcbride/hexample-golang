package core

import "math/rand"

const Bar = 123

func GetRandomInt() int {
	return rand.Intn(1000)
}
