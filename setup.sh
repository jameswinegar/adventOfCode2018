#!/usr/bin/env bash
for i in $(seq 30)
do
  mkdir -p Day${i}
  touch Day${i}/1.go
  touch Day${i}/2.go
  touch Day${i}/input.txt
done
