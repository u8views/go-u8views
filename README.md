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
BenchmarkProfileStatsService-12    	     100	    627990 ns/op	    5033 B/op	      80 allocs/op
PASS
ok  	github.com/u8views/go-u8views/internal/tests	0.088s
```
```bash
BENCHTIME=1000x make bench
```
```text
BenchmarkProfileStatsService
BenchmarkProfileStatsService-12    	    1000	    449478 ns/op	    4124 B/op	      72 allocs/op
PASS
ok  	github.com/u8views/go-u8views/internal/tests	0.471s
```
```bash
BENCHTIME=10000x make bench
```
```text
BenchmarkProfileStatsService
BenchmarkProfileStatsService-12    	   10000	    546875 ns/op	    4885 B/op	      81 allocs/op
PASS
ok  	github.com/u8views/go-u8views/internal/tests	5.492s
```

### Database schema templates
* [DrawSQL](https://drawsql.app/templates)

### Database schema
![Database schema](https://github.com/u8views/go-u8views/blob/master/database-schema/v001.png?raw=true)

### SQL
```sql
SELECT user_id, SUM(count), COUNT(*)
FROM profile_hourly_views_stats
GROUP BY user_id
ORDER BY SUM(count) DESC
LIMIT 100;
```