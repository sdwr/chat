package main
import (
	"math/rand"
	"time"
)

func Random(max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max)
}