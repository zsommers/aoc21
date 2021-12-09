#!/usr/bin/env bash
set -euxo pipefail

if [[ $# -ne 1 ]]; then
    echo "Usage: ./new_day.sh 7"
    exit 1
fi

d="./day$1"
f="$d/day$1.go"
tf="$d/day$1_test.go"
import=$(cat <<-IMPORT
	\"github.com/zsommers/aoc21/day${1}\"\n\
	\"github.com/zsommers/aoc21/util\"
IMPORT
)
case=$(cat <<-CASE
	case "${1}a":\n\
		result = day${1}.A(input)\n\
	case "${1}b":\n\
		result = day${1}.B(input)\n\
	default
CASE
)

mkdir -p "$d"

cp template/dayX.go "$f"
cp template/dayX_test.go "$tf"

gsed -i "s/X/$1/g" "$f"
gsed -i "s/X/$1/g" "$tf"
gsed -i "s|\"github.com/zsommers/aoc21/util\"|$import|" "./main.go"
gsed -i "s/default/$case/" "./main.go"

wget "https://adventofcode.com/2021/day/$1/input"\
    -q --no-cookies\
    --header "Cookie: session=$(cat ./session.cookie)"\
     -O "$d/input.txt"