chainer ⛓️
==========

chainer is an experimental transaction log.

this could power a write ahead log in a database, store transactions for a blockchain, or any number of similar responsibilities.

this library is entirely experimental and not to be used in production without further scrutiny and testing.

it has a set of benchmark tests and some utilities for generating arbitrarily large sets of data to act on and test.

## tests

`go test -v ./...` to run the unit tests.

`go test -bench=.` to run the benchmark tests.
