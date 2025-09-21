# [u8views.com](https://u8views.com)
[![u8views profile views](https://u8views.com/api/v1/github/profiles/121827373/views/day-week-month-total-count.svg)](https://u8views.com/github/u8views)

### Profile views counter
[![Yaroslav Podorvanov profile views](https://github.com/u8views/go-u8views/blob/main/public/assets/images/yaroslav-podorvanov-developer.jpg?raw=true)](https://u8views.com/github/YaroslavPodorvanov)

# Articles
- Reddit | [I built a GitHub profile view counter](https://www.reddit.com/r/webdev/comments/1n9zdel/i_built_a_github_profile_view_counter/)
- Peerlist | [GitHub profile views counter](https://peerlist.io/podorvanov/project/github-profile-views-counter)

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
![Database schema](https://github.com/u8views/go-u8views/blob/main/database-schema/v003.png?raw=true)

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

### Stats
```sql
SELECT DATE_TRUNC('MONTH', time) AS month,
       COUNT(*)                  AS views,
       COUNT(DISTINCT (user_id)) AS users,
       SUM("count")              AS total
FROM profile_hourly_views_stats
GROUP BY 1
ORDER BY 1;
```
| Month      | Views | Users | Total   |
|------------|-------|-------|---------|
| 2023-01-01 | 15    | 3     | 78      |
| 2023-02-01 | 438   | 18    | 2700    |
| 2023-03-01 | 951   | 32    | 10241   |
| 2023-04-01 | 1110  | 36    | 3441    |
| 2023-05-01 | 2191  | 43    | 9032    |
| 2023-06-01 | 3433  | 57    | 16866   |
| 2023-07-01 | 3331  | 54    | 14233   |
| 2023-08-01 | 4539  | 69    | 18017   |
| 2023-09-01 | 4519  | 77    | 17053   |
| 2023-10-01 | 4473  | 78    | 15771   |
| 2023-11-01 | 4919  | 96    | 17567   |
| 2023-12-01 | 5525  | 115   | 19882   |
| 2024-01-01 | 11185 | 232   | 40202   |
| 2024-02-01 | 11348 | 245   | 39586   |
| 2024-03-01 | 13581 | 273   | 48629   |
| 2024-04-01 | 13613 | 291   | 47198   |
| 2024-05-01 | 14655 | 306   | 50581   |
| 2024-06-01 | 13924 | 319   | 57695   |
| 2024-07-01 | 14589 | 338   | 55751   |
| 2024-08-01 | 15593 | 380   | 51997   |
| 2024-09-01 | 16034 | 400   | 54388   |
| 2024-10-01 | 16599 | 455   | 56601   |
| 2024-11-01 | 16500 | 474   | 56818   |
| 2024-12-01 | 16511 | 483   | 58579   |
| 2025-01-01 | 19778 | 534   | 67856   |
| 2025-02-01 | 18914 | 597   | 112192  |
| 2025-03-01 | 20892 | 611   | 70920   | 
| 2025-04-01 | 20738 | 630   | 67782   | 
| 2025-05-01 | 21916 | 638   | 125460  | 
| 2025-06-01 | 20758 | 689   | 62625   | 
| 2025-07-01 | 23749 | 688   | 2407710 | 
| 2025-08-01 | 23726 | 686   | 78508   |

```sql
SELECT                                    
    DATE_TRUNC('month', created_at),
    COUNT(*)                
FROM users                               
GROUP BY 1                     
ORDER BY 1;
```
| Month      | Users |
|------------|-------|
| 2023-01-01 | 3     |
| 2023-02-01 | 24    |
| 2023-03-01 | 29    |
| 2023-04-01 | 15    |
| 2023-05-01 | 23    |
| 2023-06-01 | 37    |
| 2023-07-01 | 7     |
| 2023-08-01 | 33    |
| 2023-09-01 | 27    |
| 2023-10-01 | 23    |
| 2023-11-01 | 37    |
| 2023-12-01 | 53    |
| 2024-01-01 | 189   |
| 2024-02-01 | 65    |
| 2024-03-01 | 62    |
| 2024-04-01 | 65    |
| 2024-05-01 | 49    |
| 2024-06-01 | 50    |
| 2024-07-01 | 52    |
| 2024-08-01 | 74    |
| 2024-09-01 | 87    |
| 2024-10-01 | 107   |
| 2024-11-01 | 91    |
| 2024-12-01 | 88    |
| 2025-01-01 | 98    |
| 2025-02-01 | 133   |
| 2025-03-01 | 90    |
| 2025-04-01 | 107   |
| 2025-05-01 | 76    |
| 2025-06-01 | 88    |
| 2025-07-01 | 105   |
| 2025-08-01 | 67    |
