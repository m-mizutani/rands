package rands_test

import (
	"testing"

	"github.com/m-mizutani/rands"
)

func TestSeed(t *testing.T) {
	t.Run("not equal by random seed", func(t *testing.T) {
		r1 := rands.New()
		r2 := rands.New()
		s1, s2 := r1.NewString(10), r2.NewString(10)
		if s1 == s2 {
			t.Errorf("should not be matched by other seed: %s, %s", s1, s2)
		}
	})

	t.Run("equal by same seed", func(t *testing.T) {
		r1 := rands.New(rands.WithSeed(666))
		r2 := rands.New(rands.WithSeed(666))
		s1, s2 := r1.NewString(10), r2.NewString(10)
		if s1 != s2 {
			t.Errorf("should not be matched by other seed: %s, %s", s1, s2)
		}
	})
}

func TestCharSet(t *testing.T) {
	const loop = 128

	t.Run("set char set", func(t *testing.T) {
		r := rands.New(
			rands.WithCharSet("ABC"),
		)
		for i := 0; i < loop; i++ {
			s := r.NewString(1)
			if s != "A" && s != "B" && s != "C" {
				t.Errorf("including charactor not in charSet: %s", s)
			}
		}
	})

	t.Run("split unicode charactor", func(t *testing.T) {
		r := rands.New(
			rands.WithCharSet("秩序"),
		)
		for i := 0; i < loop; i++ {
			s := r.NewString(1)
			if s != "秩" && s != "序" {
				t.Errorf("including charactor not in charSet: %s", s)
			}
		}
	})

}
