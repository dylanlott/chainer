package main

// Storage is the main interface for this layer
type Storage interface {
	Set(id string, balance uint64)
	Get(id string, height uint64) uint64
	IncrementHeight()
}

var _ Storage = (*chainer)(nil)

// tx holds an update to a balance at a given height
type tx struct {
	height  uint64
	balance uint64
}

// block holds a list of transactions and a height.
// NB: this could be made into tx hashes to further decrease size
type block struct {
	height uint64
	txs    []tx
}

// holds the current height and a map of users to tx slices
type chainer struct {
	balances map[string][]tx
	height   uint64

	// NB: Log is only attached to chainer for demonstration.
	// It would be better served to be pulled out into its own
	// collection layer behind an interface.
	log    []block
	buffer []tx
}

// initialize returns a new instantiated blockchain
func initialize() *chainer {
	return &chainer{
		height:   0,
		balances: map[string][]tx{},
		log:      []block{},
	}
}

// Set adds a tx to the list at the current height
func (b *chainer) Set(id string, balance uint64) {
	newtx := tx{height: b.height, balance: balance}

	v, ok := b.balances[id]
	if !ok {
		b.balances[id] = []tx{newtx}
		return
	}

	v = append(v, newtx)
	b.balances[id] = v
	b.buffer = append(b.buffer, newtx)
}

// Get returns the balance at the given height.
// If the height is greater than the block height, it uses
// the current block height.
func (b *chainer) Get(id string, height uint64) uint64 {
	var selected int

	txlist, ok := b.balances[id]
	if !ok {
		return 0
	}

	for idx, item := range txlist {
		if item.height <= height {
			selected = idx
		}
		if item.height > height {
			break
		}
	}

	return txlist[selected].balance
}

// IncrementHeight will tick the height of the chain up by 1
func (b *chainer) IncrementHeight() {
	b.height++
	b.append()
}

// append will add a block to the chain
func (b *chainer) append() {
	txlist := []tx{}
	copy(txlist, b.buffer)
	newblock := block{
		height: b.height,
		txs:    txlist,
	}
	b.log = append(b.log, newblock)
	b.buffer = make([]tx, 0)
}
