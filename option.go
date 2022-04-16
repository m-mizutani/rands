package rands

import "math/rand"

func WithSeed(seed int64) Option {
	return func(r *Rands) {
		r.random = rand.New(rand.NewSource(seed))
	}
}

func WithCharSet(chars string) Option {
	return func(r *Rands) {
		r.chars = unique([]rune(chars))
	}
}

func WithRuneSet(runes []rune) Option {
	return func(r *Rands) {
		r.chars = runes
	}
}
