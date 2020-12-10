package test

import (
	"testing"

	"gotest/game"
)

func TestCheckCOunt(t *testing.T) {
	g := game.New()
	g.Check(1)
	g.Check(1)
	g.Check(1)

	count := g.CheckCount()

	if 3 != count {
		t.Error("Expected 3, got", count)
	}
}

func TestCointainsEqual(t *testing.T) {
	g := game.New()

	for i := 0; i < 100; i++ {
		v := g.Check(i)
		if v == game.Equal {
			return
		}
	}

	t.Error("No equal")
}
