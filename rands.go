package rands

import (
	cryptoRand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sort"
)

const (
	LowerSet  = "abcdefghijklmnopqrstuvwxyz"
	UpperSet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumberSet = "0123456789"
	MarkSet   = "!@#$%^&*_-=+;:,.?"

	DefaultCharSet = LowerSet + UpperSet + NumberSet
)

type Rands struct {
	chars  []rune
	random *rand.Rand
}

var (
	global *Rands
)

func init() {
	global = New()
}

type Option func(f *Rands)

func New(options ...Option) *Rands {
	rands, err := NewWithError(options...)
	if err != nil {
		panic("rands error: " + err.Error())
	}

	return rands
}

func NewWithError(options ...Option) (*Rands, error) {
	rands := &Rands{
		chars: unique([]rune(DefaultCharSet)),
	}

	for _, opt := range options {
		opt(rands)
	}

	if rands.random == nil {
		seed, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(math.MaxInt64))
		if err != nil {
			return nil, fmt.Errorf("failed to init crypto/rand: %w", err)
		}

		// #nosec: using crypto/rand for math/rand.seed
		rands.random = rand.New(rand.NewSource(seed.Int64()))
	}

	return rands, nil
}

func unique(runes []rune) []rune {
	m := map[rune]struct{}{}
	for _, r := range runes {
		m[r] = struct{}{}
	}

	i := 0
	resp := make([]rune, len(m))
	for r := range m {
		resp[i] = r
		i++
	}

	sort.Slice(resp, func(i, j int) bool {
		return resp[i] < resp[j]
	})

	return resp
}

func NewString(length int) string {
	return global.NewString(length)
}

func (x *Rands) NewString(length int) string {
	if len(x.chars) == 0 {
		return "" // nothing to be created
	}

	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = x.chars[x.random.Int()%len(x.chars)]
	}

	return string(runes)
}

func (x *Rands) Runes() []rune {
	return x.chars[:]
}
