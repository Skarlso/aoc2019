#!/bin/zsh

for i in {1..25}; do
    mkdir -p day$i && mkdir -p day$i/part1 day$i/part2
    (
        cd day$i/part1
        go mod init aoc0219/day$i-part1
        cd ../part2
        go mod init aoc0219/day$i-part2
    )
done
