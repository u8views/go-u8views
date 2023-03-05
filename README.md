# [u8views.com](https://u8views.com)
[![Yaroslav Podorvanov profile views](https://u8views.com/api/v1/github/profiles/63663261/views/day-week-month-total-count.svg)](https://u8views.com/github/YaroslavPodorvanov)

### Profile views counter
[![Yaroslav Podorvanov profile views](https://github.com/u8views/go-u8views/blob/main/public/assets/images/yaroslav-podorvanov-developer.jpg?raw=true)](https://u8views.com/github/YaroslavPodorvanov)

### Development

##### Start local development session
```bash
cp .local.env .env
make env-up
make migrate-all-reset
make postgres-fixtures
make postgres-fixtures-count
# and
# make postgres-fixtures-clear
```

##### Run after code changes
```bash
make env-up
```

##### End development session
```bash
make env-down
# or
# make env-down-with-clear
```

##### Benchmark (PC) Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
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

##### Benchmark ([vultr.com](https://www.vultr.com/?ref=8741375) VPS 1024.00 MB High Frequency) Intel Core Processor (Skylake, IBRS)
```bash
BENCHTIME=100x make bench
```
```text
BenchmarkProfileStatsService
BenchmarkProfileStatsService 	     100	   2275173 ns/op	    1562 B/op	      43 allocs/op
PASS
ok  	github.com/u8views/go-u8views/internal/tests	0.360s
```
```bash
BENCHTIME=1000x make bench
```
```text
BenchmarkProfileStatsService
BenchmarkProfileStatsService 	    1000	   2121516 ns/op	    1571 B/op	      44 allocs/op
PASS
ok  	github.com/u8views/go-u8views/internal/tests	2.153s
```
```bash
BENCHTIME=10000x make bench
```
```text
BenchmarkProfileStatsService
BenchmarkProfileStatsService 	   10000	   2153319 ns/op	    1574 B/op	      44 allocs/op
PASS
ok  	github.com/u8views/go-u8views/internal/tests	21.566s
```

### Database schema templates
* [DrawSQL](https://drawsql.app/templates)

### Database schema
![Database schema](https://github.com/u8views/go-u8views/blob/main/database-schema/v002.png?raw=true)

### SQL
```sql
SELECT user_id, SUM(count), COUNT(*)
FROM profile_hourly_views_stats
GROUP BY user_id
ORDER BY SUM(count) DESC
LIMIT 100;
```
```sql
SELECT g.time::TIMESTAMP
FROM (
    SELECT time::TIMESTAMP
    FROM generate_series(
        (DATE_TRUNC('DAY', NOW()) - INTERVAL '1 MONTH')::TIMESTAMP,
        (DATE_TRUNC('DAY', NOW()))::TIMESTAMP,
        '1 DAY'::INTERVAL
    ) AS time
) AS g;
```
