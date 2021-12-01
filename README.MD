# Pulsar

### Installing

go get -u github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar

### Running

cd path/to/proto/files

protoc --go-pulsar_out=. --go-pulsar_opt=paths=source_relative --go-pulsar_opt=features=marshal+unmarshal+size -I .
NAME_OF_FILE.proto


## Acknowledgements

Code for marshal, unmarshal, and size implemented by [planetscale/vtprotobuf](https://github.com/planetscale/vtprotobuf) used in our `ProtoMethods` implementation. 