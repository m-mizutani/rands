package rands_test

import (
	"fmt"
	"testing"

	"github.com/m-mizutani/rands"
)

func FuzzRandsCharSet(f *testing.F) {
	f.Add(rands.DefaultCharSet)
	f.Add(";l.,'po0c][")
	f.Add("ouopiirpogiusdfoigauop")

	f.Fuzz(func(t *testing.T, chars string) {
		r := rands.New(
			rands.WithSeed(1001), // to prevent massive reading from crypto/rand.Reader
			rands.WithCharSet(chars),
		)

		s1 := r.NewString(128)
		s2 := r.NewString(128)
		if s1 == s2 && len(r.Runes()) > 4 {
			fmt.Println("chars:", len([]rune(chars)), []rune(chars))
			fmt.Println("string:", []rune(s1), []rune(s2))
			t.Errorf("do not match with another NewString")
		}
	})
}
