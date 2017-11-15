package core

import (
	"math/rand"
	"time"
)

const (
	TestTrials = 100
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
