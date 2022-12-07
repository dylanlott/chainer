package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestBlockchain(t *testing.T) {
	is := is.New(t)

	chain := initialize()
	is.Equal(chain.height, uint64(0))

	chain.Set("A", 1)
	chain.IncrementHeight() // 1

	chain.Set("B", 2)
	chain.IncrementHeight() // 2

	chain.Set("A", 3)
	chain.IncrementHeight() // 3

	got := chain.Get("A", 0)
	is.Equal(got, uint64(1))

	got = chain.Get("B", 0)
	is.Equal(got, uint64(2))

	got = chain.Get("B", 1)

	got = chain.Get("A", 3)
	is.Equal(got, uint64(3))

	got = chain.Get("A", 1)
	is.Equal(got, uint64(1))

	got = chain.Get("B", 1)
	is.Equal(got, uint64(2))

	got = chain.Get("B", 2)
	is.Equal(got, uint64(2))
}
