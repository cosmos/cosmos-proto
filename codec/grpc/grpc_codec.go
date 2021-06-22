package grpc

import "fmt"

// Name is the name registered for the proto compressor.
const Name = "proto"

type Codec struct{}

type vtprotoMessage interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

func (Codec) Marshal(v interface{}) ([]byte, error) {
	vt, ok := v.(vtprotoMessage)
	if !ok {
		return nil, fmt.Errorf("failed to marshal, message is %T (missing vtprotobuf helpers)", v)
	}
	return vt.Marshal()
}

func (Codec) Unmarshal(data []byte, v interface{}) error {
	vt, ok := v.(vtprotoMessage)
	if !ok {
		return fmt.Errorf("failed to unmarshal, message is %T (missing vtprotobuf helpers)", v)
	}
	return vt.Unmarshal(data)
}

func (Codec) Name() string {
	return Name
}
