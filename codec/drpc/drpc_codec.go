package drpc

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type vtprotoMessage interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

func Marshal(msg interface{}) ([]byte, error) {
	return msg.(vtprotoMessage).Marshal()
}

func Unmarshal(buf []byte, msg interface{}) error {
	return msg.(vtprotoMessage).Unmarshal(buf)
}

func JSONMarshal(msg interface{}) ([]byte, error) {
	return protojson.Marshal(msg.(proto.Message))
}

func JSONUnmarshal(buf []byte, msg interface{}) error {
	return protojson.Unmarshal(buf, msg.(proto.Message))
}
