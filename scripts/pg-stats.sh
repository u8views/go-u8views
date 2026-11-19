#!/usr/bin/env bash
# Read-only statistics of the database, for a quick look over ssh on the server.
#
#   ./scripts/pg-stats.sh [all|size|tables|users|referrals|usernames|views|views-monthly]
#
# The queries run with psql inside the Postgres container, which knows its own POSTGRES_USER and POSTGRES_DB,
# so it works the same in development and in production. The session is read-only: nothing can be changed.
#
# PG_CONTAINER           overrides the container name (default go_u8views_postgres)
# PG_STATS_VIEW_MONTHS   the number of months (the current one included) for views-monthly (default 6)
#
# profile_hourly_views_stats is huge (tens of millions of rows): there is no COUNT(*) over it,
# the number of its rows is the planner estimate. Only the command views-monthly reads it, see the warning there.
set -euo pipefail

PG_CONTAINER="${PG_CONTAINER:-go_u8views_postgres}"
PG_STATS_VIEW_MONTHS="${PG_STATS_VIEW_MONTHS:-6}"

# psql_exec reads SQL from the standard input, the arguments are passed to psql (e.g. -v name=value)
psql_exec() {
	docker exec -i -e PGOPTIONS='-c default_transaction_read_only=on' "$PG_CONTAINER" \
		sh -c 'psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -X -q -v ON_ERROR_STOP=1 -P pager=off "$@"' sh "$@"
}

check_container() {
	if [ "$(docker inspect -f '{{.State.Running}}' "$PG_CONTAINER" 2>/dev/null || true)" != "true" ]; then
		echo "The container $PG_CONTAINER is not running." >&2
		exit 1
	fi
}

stats_size() {
	psql_exec <<'SQL'
\echo
\echo '== Database size'
SELECT current_database()                                   AS database,
       pg_size_pretty(pg_database_size(current_database())) AS size;
SQL
}

stats_tables() {
	# Tables smaller than 64 MB are counted exactly. The bigger ones (profile_hourly_views_stats) are never counted:
	# the number of rows is the planner estimate (pg_class.reltuples, or pg_stat_user_tables.n_live_tup
	# if the table has not been analyzed yet), it can be off by some percent.
	psql_exec <<'SQL'
\echo
\echo '== Tables, the biggest first (rows: exact for the tables < 64 MB, estimate for the bigger ones)'
WITH tables AS (
    SELECT c.oid,
           c.relname,
           pg_relation_size(c.oid) < 64 * 1024 * 1024 AS small,
           c.reltuples,
           COALESCE(s.n_live_tup, 0)                  AS n_live_tup,
           n.nspname
    FROM pg_class c
             JOIN pg_namespace n ON n.oid = c.relnamespace
             LEFT JOIN pg_stat_user_tables s ON s.relid = c.oid
    WHERE c.relkind = 'r'
      AND n.nspname = 'public'
),
counted AS (
    SELECT t.*,
           CASE
               WHEN t.small THEN
                   (xpath('/row/cnt/text()', query_to_xml(FORMAT('SELECT COUNT(*) AS cnt FROM %I.%I', t.nspname, t.relname), FALSE, TRUE, '')))[1]::TEXT::BIGINT
               WHEN t.reltuples >= 0 THEN t.reltuples::BIGINT
               ELSE t.n_live_tup
               END AS row_count
    FROM tables t
)
SELECT relname                                     AS "table",
       row_count                                   AS "rows",
       CASE WHEN small THEN 'exact' ELSE 'estimate' END AS "rows_kind",
       pg_size_pretty(pg_table_size(oid))          AS table_size,
       pg_size_pretty(pg_indexes_size(oid))        AS indexes_size,
       pg_size_pretty(pg_total_relation_size(oid)) AS total_size
FROM counted
ORDER BY pg_total_relation_size(oid) DESC, relname;

\echo
\echo '== Tables, total'
SELECT COUNT(*)                                                   AS tables,
       pg_size_pretty(SUM(pg_total_relation_size(c.oid))::BIGINT) AS total_size
FROM pg_class c
         JOIN pg_namespace n ON n.oid = c.relnamespace
WHERE c.relkind = 'r'
  AND n.nspname = 'public';
SQL
}

stats_users() {
	psql_exec <<'SQL'
SET TIME ZONE 'UTC';

\echo
\echo '== Users, total'
SELECT COUNT(*) AS total
FROM users;

\echo
\echo '== Users registered by month (UTC)'
WITH months AS (
    SELECT GENERATE_SERIES(DATE_TRUNC('month', MIN(created_at)), DATE_TRUNC('month', NOW()), '1 month') AS month
    FROM users
),
registered AS (
    SELECT DATE_TRUNC('month', created_at) AS month,
           COUNT(*)                        AS new_users
    FROM users
    GROUP BY 1
)
SELECT TO_CHAR(m.month, 'YYYY-MM')                       AS month,
       COALESCE(r.new_users, 0)                          AS new_users,
       SUM(COALESCE(r.new_users, 0)) OVER (ORDER BY m.month) AS total
FROM months m
         LEFT JOIN registered r ON r.month = m.month
ORDER BY m.month;

\echo
\echo '== Users by social provider'
SELECT p::TEXT                   AS provider,
       COUNT(u.id)               AS users
FROM UNNEST(ENUM_RANGE(NULL::SOCIAL_PROVIDER)) AS p
         LEFT JOIN users u ON u.social_provider = p
GROUP BY p
ORDER BY users DESC, provider;

\echo
\echo '== Users activity (last_login_at)'
SELECT COUNT(*) FILTER (WHERE last_login_at >= NOW() - INTERVAL '1 day')  AS last_1_day,
       COUNT(*) FILTER (WHERE last_login_at >= NOW() - INTERVAL '7 days')  AS last_7_days,
       COUNT(*) FILTER (WHERE last_login_at >= NOW() - INTERVAL '30 days') AS last_30_days
FROM users;
SQL
}

stats_referrals() {
	psql_exec <<'SQL'
SET TIME ZONE 'UTC';

\echo
\echo '== Referrals, total'
SELECT (SELECT COUNT(*) FROM referrals)                         AS referrals,
       (SELECT COUNT(DISTINCT referrer_user_id) FROM referrals) AS referrers,
       (SELECT COUNT(*) FROM users)                             AS users;

\echo
\echo '== Referred users by month of their registration (UTC)'
WITH months AS (
    SELECT GENERATE_SERIES(DATE_TRUNC('month', MIN(u.created_at)), DATE_TRUNC('month', NOW()), '1 month') AS month
    FROM referrals r
             JOIN users u ON u.id = r.referee_user_id
),
referred AS (
    SELECT DATE_TRUNC('month', u.created_at) AS month,
           COUNT(*)                          AS referred_users
    FROM referrals r
             JOIN users u ON u.id = r.referee_user_id
    GROUP BY 1
)
SELECT TO_CHAR(m.month, 'YYYY-MM')                          AS month,
       COALESCE(d.referred_users, 0)                        AS referred_users,
       SUM(COALESCE(d.referred_users, 0)) OVER (ORDER BY m.month) AS total
FROM months m
         LEFT JOIN referred d ON d.month = m.month
ORDER BY m.month;

\echo
\echo '== Top 10 referrers'
SELECT u.username,
       u.social_provider::TEXT AS provider,
       COUNT(*)                AS referred_users
FROM referrals r
         JOIN users u ON u.id = r.referrer_user_id
GROUP BY u.id
ORDER BY referred_users DESC, u.id
LIMIT 10;
SQL
}

stats_usernames() {
	psql_exec <<'SQL'
\echo
\echo '== Username history'
WITH per_user AS (
    SELECT user_id,
           COUNT(*) AS names
    FROM username_history
    GROUP BY user_id
)
SELECT COALESCE(SUM(names), 0)                    AS entries,
       COUNT(*)                                   AS users,
       COUNT(*) FILTER (WHERE names > 1)          AS users_changed_username,
       COALESCE(MAX(names), 0)                    AS max_names_per_user
FROM per_user;
SQL
}

stats_views() {
	# profile_total_views has one row per user, it is small. profile_hourly_views_stats is NOT read here.
	psql_exec <<'SQL'
\echo
\echo '== Profile views, total (profile_total_views)'
SELECT COALESCE(SUM(count), 0)                                        AS total_views,
       COUNT(*)                                                       AS users_with_a_row,
       COUNT(*) FILTER (WHERE count > 0)                              AS users_with_views,
       (SELECT COUNT(*) FROM users) - COUNT(*)                        AS users_without_a_row
FROM profile_total_views;

\echo
\echo '== Top 10 profiles by views'
SELECT u.username,
       u.social_provider::TEXT AS provider,
       v.count                 AS views
FROM profile_total_views v
         JOIN users u ON u.id = v.user_id
ORDER BY v.count DESC, u.id
LIMIT 10;
SQL
}

stats_views_monthly() {
	case "$PG_STATS_VIEW_MONTHS" in
	'' | *[!0-9]* | 0)
		echo "PG_STATS_VIEW_MONTHS must be a positive integer, got: $PG_STATS_VIEW_MONTHS" >&2
		exit 2
		;;
	esac

	echo >&2
	echo "WARNING: views-monthly reads profile_hourly_views_stats for the last $PG_STATS_VIEW_MONTHS month(s)." >&2
	echo "The primary key is (user_id, time), so a filter by time scans the table: it can take minutes on a big database." >&2

	psql_exec -v months="$PG_STATS_VIEW_MONTHS" <<'SQL'
SET TIME ZONE 'UTC';

\echo
\echo '== Profile views by month (UTC), the last N months of profile_hourly_views_stats (N = PG_STATS_VIEW_MONTHS)'
-- time is TIMESTAMP (without time zone) and holds UTC
WITH months AS (
    SELECT GENERATE_SERIES(DATE_TRUNC('month', NOW() AT TIME ZONE 'UTC') - (:months - 1) * INTERVAL '1 month',
                           DATE_TRUNC('month', NOW() AT TIME ZONE 'UTC'),
                           '1 month') AS month
),
views AS (
    SELECT DATE_TRUNC('month', time) AS month,
           SUM(count)                AS views,
           COUNT(DISTINCT user_id)   AS profiles
    FROM profile_hourly_views_stats
    WHERE time >= DATE_TRUNC('month', NOW() AT TIME ZONE 'UTC') - (:months - 1) * INTERVAL '1 month'
    GROUP BY 1
)
SELECT TO_CHAR(m.month, 'YYYY-MM') AS month,
       COALESCE(v.views, 0)        AS views,
       COALESCE(v.profiles, 0)     AS profiles_viewed
FROM months m
         LEFT JOIN views v ON v.month = m.month
ORDER BY m.month;
SQL
}

usage() {
	cat >&2 <<'USAGE'
Usage: scripts/pg-stats.sh [all|size|tables|users|referrals|usernames|views|views-monthly]

  all            size, tables, users, referrals, usernames, views (default)
  size           the size of the database
  tables         the tables: rows (exact for the small ones, estimate for the big ones) and sizes
  users          the registered users: total, by month, by provider, activity (last_login_at)
  referrals      the referrals: total, by month of the registration, top referrers
  usernames      the username history: how many users changed the username
  views          the profile views from profile_total_views: total, top profiles
  views-monthly  the views by month from profile_hourly_views_stats (SLOW, not in "all"),
                 the number of months: PG_STATS_VIEW_MONTHS (default 6)

PG_CONTAINER overrides the container name (default go_u8views_postgres).
USAGE
}

main() {
	local command="${1:-all}"

	if [ "$#" -gt 1 ]; then
		usage
		exit 2
	fi

	case "$command" in
	all | size | tables | users | referrals | usernames | views | views-monthly) ;;
	-h | --help | help)
		usage
		exit 0
		;;
	*)
		usage
		exit 2
		;;
	esac

	check_container

	case "$command" in
	size) stats_size ;;
	tables) stats_tables ;;
	users) stats_users ;;
	referrals) stats_referrals ;;
	usernames) stats_usernames ;;
	views) stats_views ;;
	views-monthly) stats_views_monthly ;;
	all)
		stats_size
		stats_tables
		stats_users
		stats_referrals
		stats_usernames
		stats_views
		;;
	esac
}

# Not run when the file is sourced (for the tests)
if [ "${BASH_SOURCE[0]}" = "$0" ]; then
	main "$@"
fi
