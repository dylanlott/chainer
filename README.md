chainer ⛓️
==========

chainer is an experimental transaction log.

this could power a write ahead log in a database, store transactions for a blockchain, or any number of similar responsibilities.

this library is entirely experimental and not to be used in production without further scrutiny and testing.

it has a set of benchmark tests and some utilities for generating arbitrarily large sets of data to act on and test.

## tests

`go test -v ./...` to run the unit tests.

`go test -bench=.` to run the benchmark tests.

*benchmark test results*
```sh
❯ go test -v -bench=.
=== RUN   TestBlockchain
--- PASS: TestBlockchain (0.00s)
=== RUN   TestAppend
--- PASS: TestAppend (0.00s)
=== RUN   TestSeed
--- PASS: TestSeed (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/dylanlott/blockchain
cpu: VirtualApple @ 2.50GHz
BenchmarkBlockchainGet
BenchmarkBlockchainGet-10    	16567231	        69.93 ns/op
BenchmarkBlockchainSet
BenchmarkBlockchainSet-10    	   85296	     14613 ns/op
PASS
ok  	github.com/dylanlott/blockchain	3.873s
```
