package game

import (
	"math/rand"
	"time"
)

type CheckResult int

const (
	_ CheckResult = iota
	Equal
	More
	Less
)

type Game interface {
	Check(value int) CheckResult
	CheckCount() int
}

type game struct {
	number     int
	checkCount int
}

func New() *game {
	rand.Seed(time.Now().UnixNano())
	return &game{
		number: (rand.Intn(100) + 1),
	}
}

func (r *game) Check(value int) CheckResult {
	r.checkCount++
	if value > r.number {
		return Less
	} else if value < r.number {
		return More
	} else {
		return Equal
	}
}

func (r *game) CheckCount() int {
	return r.checkCount
}
