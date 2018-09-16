package protobuff

import (
	"errors"

	"github.com/golang/protobuf/proto"
)

func EncodeProto(obj interface{}) ([]byte, error) {
	if v, ok := obj.(*Club); ok {
		return proto.Marshal(v)
	}
	return nil, errors.New("Proto: Unknown Message Type")
}

func DecodeProto(buffer []byte) (*Club, error) {
	pb := new(Club)
	return pb, proto.Unmarshal(buffer, pb)
}
