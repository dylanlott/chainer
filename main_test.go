package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestBlockchain(t *testing.T) {
	is := is.New(t)

	chain := initialize()
	is.Equal(chain.height(), uint64(0))
	t.Logf("chain first: %+v", chain)

	chain.set("A", 1)
	t.Logf("chain 2: %+v", chain)

	chain.incrementHeight()
	t.Logf("chain 3 %+v", chain)

	chain.set("B", 2)

	chain.incrementHeight()
	t.Logf("chain 4 %+v", chain)

	chain.set("A", 3)

	chain.incrementHeight()

	got := chain.get("A", 0)
	is.Equal(got, uint64(1))

	got = chain.get("B", 0)
	t.Logf("B at 0: %+v", got)

	got = chain.get("B", 1)
	t.Logf("B at 1: %+v", got)

	got = chain.get("A", 3)
	is.Equal(got, uint64(3))

	got = chain.get("A", 1)
	is.Equal(got, uint64(1))

	got = chain.get("B", 1)
	is.Equal(got, uint64(2))

	got = chain.get("B", 2)
	is.Equal(got, uint64(2))
}
