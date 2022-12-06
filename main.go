/*
# The Question

A core concept in blockchains is the ability to store and access a historical record of account balances. Design a structure allows setting the current value and retrieving a historical value the latest value at a certain block height.

**Example interface:**

- `set(account, balance)`
- `get(account, height)`
- `increment_height()`

**Example flow:**

- `initialize` → height=0
- `set(A, 1)`
- `increment_height()`
- `set(B, 2)`
- `increment_height()`
- `set(A, 3)`
- `increment_height()`
- `get(A, 0)` → returns 1
- `get(B, 0)` → returns 0
- `get(B, 1)` → returns 2
- `get(A, 1)` → returns 1
- `get(B, 2)` -> returns 2
- `get(A, 2)` → returns  3
*/

package main

import "log"

// Block ...
type Block struct {
	balances map[string]uint64
	current  uint64
}

// Blockchain ...
type Blockchain struct {
	chain []Block
}

// Balances ..
type Balances interface {
	set(id string, balance uint64)
	get(id string, height uint64) uint64
	incrementHeight()
}

var _ Balances = (*Blockchain)(nil)

func initialize() *Blockchain {
	return &Blockchain{
		chain: []Block{
			Block{
				balances: map[string]uint64{},
				current:  0,
			},
		},
	}
}

func (b *Blockchain) set(id string, balance uint64) {
	last := b.chain[len(b.chain)-1]
	last.balances[id] = balance
}

func (b *Blockchain) get(id string, height uint64) uint64 {
	for _, block := range b.chain {
		// log.Printf("found block: %+v", block)
		if block.current == height {
			log.Printf("found correct height: %+v", block)
			bal, ok := block.balances[id]
			if !ok {
				return 0
			}
			// log.Printf("Found balance: %v", bal)
			return bal
		}
	}
	return 0
}

func (b *Blockchain) incrementHeight() {
	last := b.chain[len(b.chain)-1]
	next := last.current + 1

	var newBalances = map[string]uint64{}
	for k, v := range last.balances {
		newBalances[k] = v
	}

	b.chain = append(b.chain, Block{
		balances: newBalances,
		current:  next,
	})
}

func (b *Blockchain) height() uint64 {
	last := b.chain[len(b.chain)-1]
	return last.current
}
