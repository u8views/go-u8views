# go-u8views

Profile views counter

### Development

##### Start development session
```bash
make up
make migrate-all-reset
make postgres-fixtures
make postgres-fixtures-count
```

##### Run after code changes
```bash
make local-run
```

##### End development session
```bash
make down
```

##### Benchmark
```bash
BENCHTIME=100x make bench
```
```text
BenchmarkProfileStatsService
BenchmarkProfileStatsService-12    	     100	  93718529 ns/op	    4774 B/op	      77 allocs/op
PASS
ok  	github.com/u8views/go-u8views/internal/tests	9.625s
```
```bash
BENCHTIME=1000x make bench
```
```text
BenchmarkProfileStatsService
BenchmarkProfileStatsService-12    	    1000	  58595528 ns/op	    2267 B/op	      51 allocs/op
PASS
ok  	github.com/u8views/go-u8views/internal/tests	58.879s
```
```bash
BENCHTIME=10000x make bench
```
```text
BenchmarkProfileStatsService
BenchmarkProfileStatsService-12    	   10000	  58260299 ns/op	    2410 B/op	      54 allocs/op
PASS
ok  	github.com/u8views/go-u8views/internal/tests	582.842s
```
