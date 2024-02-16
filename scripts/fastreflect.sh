#!/bin/sh

set -e

PULSAR_EXE=$(which  protoc-gen-go-pulsar)

build() {
    echo finding protobuf files in "$1"
    proto_files=$(find "$1" -name "*.proto")
    for file in $proto_files; do
      echo "building proto file $file"
      protoc -I=. --plugin "$PULSAR_EXE" --go-pulsar_out=. --go-pulsar_opt=features=protoc+fast+zeropb "$file"
    done
}

for dir in "$@"
do
  build "$dir"
done

cp -r github.com/cosmos/cosmos-proto/* ./
rm -rf github.com
