package main

import "log"

// Storage is the main interface for this layer
type Storage interface {
	Set(id string, balance uint64)
	Get(id string, height uint64) uint64
	IncrementHeight()
}

var _ Storage = (*blockchain)(nil)

// block ...
type block struct {
	height uint64
}

// tx holds an update to a balance at a given height
type tx struct {
	height  uint64
	balance uint64
}

// holds the current height and a map of users to tx slices
type blockchain struct {
	balances map[string][]tx
	height   uint64
}

// initialize returns a new instantiated blockchain
func initialize() *blockchain {
	return &blockchain{
		balances: map[string][]tx{},
		height:   0,
	}
}

// Set adds a tx to the list at the current height
func (b *blockchain) Set(id string, balance uint64) {
	newtx := tx{height: b.height, balance: balance}

	v, ok := b.balances[id]
	if !ok {
		b.balances[id] = []tx{newtx}
		return
	}

	v = append(v, newtx)
	b.balances[id] = v
}

// Get returns the balance at the given height.
func (b *blockchain) Get(id string, height uint64) uint64 {
	var selected int

	txlist, ok := b.balances[id]
	if !ok {
		return 0
	}

	for idx, item := range txlist {
		if item.height <= height {
			selected = idx
			log.Printf("idx: %+v", idx)
			log.Printf("tx %+v - selected: %+v", item, selected)
		}
		if item.height > height {
			break
		}
	}

	return txlist[selected].balance
}

// IncrementHeight will tick the height of the chain up by 1
func (b *blockchain) IncrementHeight() {
	b.height++
}
