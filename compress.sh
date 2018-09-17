#!/usr/bin/env bash
TYPE=$2
[ "$2" == "" ] && TYPE="zip"

filename=$(basename -- "$(realpath $1)")
extension="${filename##*.}"
filename="${filename%.*}"

make

./client-encode --input $1 --type $TYPE
./server-compress --input $filename.pb
./client-decode --input $filename.compressed.pb