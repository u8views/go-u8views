#!/usr/bin/env bash
# Updates only the modules that have known, reachable vulnerabilities,
# to the minimal fixed version. Does not commit anything.
# If ./vendor exists, it is regenerated at the end (go-mod-update creates it locally, it is git-ignored).
#
# Usage: GOVULNCHECK="go run golang.org/x/vuln/cmd/govulncheck@vX.Y.Z" scripts/deps/update-security.sh
set -euo pipefail

: "${GOVULNCHECK:?GOVULNCHECK is required}"

# A stale local vendor/ must not be used by the checks
export GOFLAGS="-mod=mod"

report="$(mktemp)"
trap 'rm -f "$report"' EXIT

# Fixing one module can pull in new versions of the others and the fixed versions can change,
# so the scan and the update are repeated until nothing reachable is left.
max_rounds=5
unfixed=""

for round in $(seq 1 "$max_rounds"); do
	# -format json always exits with 0, findings are in the stream
	$GOVULNCHECK -format json ./... > "$report"

	# Only findings with a called function (reachable), same as the text mode reports
	findings="$(jq -r '.finding? | select(. != null and .trace[0].function != null) | "\(.trace[0].module)\t\(.fixed_version // "")"' "$report" | sort -u)"

	if [ -z "$findings" ]; then
		echo "No reachable vulnerabilities found."
		break
	fi

	# Highest fixed version per module (sort -V, the last one wins)
	fixes="$(printf '%s\n' "$findings" | awk -F'\t' '$2 != ""' | sort -t$'\t' -k1,1 -k2,2V | awk -F'\t' '{m[$1]=$2} END {for (k in m) print k "\t" m[k]}' | sort)"
	unfixed="$(printf '%s\n' "$findings" | awk -F'\t' '$2 == "" {print $1}' | sort -u)"

	if [ -z "$fixes" ]; then
		break
	fi

	while IFS=$'\t' read -r module version; do
		if [ "$module" = "stdlib" ]; then
			echo "stdlib: fixed in Go ${version#go}. Update the Go version (go.mod, docker-compose.yml, scripts/ansible/install-go.yml, .github/workflows)."
			continue
		fi

		# An earlier `go get` may have already upgraded this module as a dependency of another one.
		# `go get module@version` would downgrade it back, e.g. x/text@v0.39.0 downgrades x/net, which needs a newer x/text.
		current="$(go list -m -f '{{.Version}}' "$module")"
		if [ "$(printf '%s\n%s\n' "$current" "$version" | sort -V | tail -n 1)" = "$current" ]; then
			echo "Round $round: $module $current already satisfies $version"
			continue
		fi

		echo "Round $round: updating $module to $version"
		go get "$module@$version"
	done <<< "$fixes"

	go mod tidy

	# Only the standard library is left: nothing more to update here
	if ! printf '%s\n' "$fixes" | awk -F'\t' '$1 != "stdlib" {found=1} END {exit !found}'; then
		break
	fi
done

if [ -n "$unfixed" ]; then
	echo "No fixed version is available for (decide manually):"
	printf '  %s\n' $unfixed
fi

echo "Verifying..."
go build ./...
$GOVULNCHECK ./...

# The default build uses ./vendor when it exists, it must match go.mod
if [ -d vendor ]; then
	echo "Regenerating vendor/..."
	go mod vendor
fi
