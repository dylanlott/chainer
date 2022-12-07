package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
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

	chain.Set("A", 5)
	chain.IncrementHeight()

	is.Equal(chain.height, uint64(4))

	acctA := chain.balances["A"]
	acctB := chain.balances["B"]

	is.Equal(len(acctA), 3)
	is.Equal(len(acctB), 1)

	got := chain.Get("A", 0)
	is.Equal(got, uint64(1))

	got = chain.Get("B", 0)
	is.Equal(got, uint64(2))

	got = chain.Get("B", 1)

	got = chain.Get("A", 3)
	is.Equal(got, uint64(5))

	got = chain.Get("A", 1)
	is.Equal(got, uint64(1))

	got = chain.Get("B", 1)
	is.Equal(got, uint64(2))

	got = chain.Get("B", 2)
	is.Equal(got, uint64(2))

	got = chain.Get("A", 10)
	is.Equal(got, uint64(5))
}

func TestAppend(t *testing.T) {
	is := is.New(t)

	chain, ids := seedBlockchain(10)
	chain.IncrementHeight()

	is.Equal(len(chain.log), 1)
	is.Equal(len(chain.buffer), 0)

	amount := randomAmount()
	chain.Set(ids[0], amount)
	is.Equal(chain.buffer[0], tx{height: chain.height, balance: amount})
}

func TestSeed(t *testing.T) {
	is := is.New(t)

	_, ids := seedBlockchain(1)
	is.Equal(len(ids), 1)

	_, ids = seedBlockchain(100)
	is.Equal(len(ids), 100)
}

func BenchmarkBlockchainGet(b *testing.B) {
	chain, ids := seedBlockchain(1000)
	rand.Seed(time.Now().Unix())

	for n := 0; n < b.N; n++ {
		rID := ids[rand.Intn(len(ids))]
		chain.Get(rID, 0)
	}
}

func BenchmarkBlockchainSet(b *testing.B) {
	chain := initialize()

	for n := 0; n < b.N; n++ {
		chain.Set(randomID(), randomAmount())
		chain.IncrementHeight()
	}
}

func seedBlockchain(n uint64) (*chainer, []string) {
	b := &chainer{
		balances: map[string][]tx{},
		height:   0,
	}

	ids := make([]string, n)
	for v := range ids {
		id := randomID()
		ids[v] = id
		b.Set(id, randomAmount())
	}

	// TODO: add random transactions to random accounts in the
	// already seeded map to vary lengths of txlists

	return b, ids
}

func randomID() string {
	randID := uuid.New()
	return randID.String()
}

func randomAmount() uint64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Uint64()
}
