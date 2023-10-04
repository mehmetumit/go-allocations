# Allocation comparisons of different approaches in go
## Running benchmark
```
go test -benchmem -bench=Allocs -benchtime=3000000x
```
## Results
```
goos: linux
goarch: amd64
pkg: go-allocations
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
BenchmarkAllocs/LotsOfAlloc-8         	 3000000	       391.1 ns/op	     248 B/op	       5 allocs/op
BenchmarkAllocs/AppendAlloc-8         	 3000000	       304.2 ns/op	     320 B/op	       2 allocs/op
BenchmarkAllocs/IndexAlloc-8          	 3000000	       145.3 ns/op	     112 B/op	       1 allocs/op
BenchmarkAllocs/ZeroAlloc-8           	 3000000	        64.29 ns/op	       0 B/op	       0 allocs/op
```
