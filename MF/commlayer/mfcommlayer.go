package commlayer

import (
	"-go-/MF/commlayer/mfproto"
)

// Communication messages types
const (
	Protobuf uint8 = iota
)

type mfConnection interface {
	EncodeAndSend(obj interface{}, destination string) error
	ListenAndDecode(listenaddress string) (chan interface{}, error)
}

func NewConnection(connType uint8) mfConnection {
	switch connType {
	case Protobuf:
		return mfproto.NewProtoHandler()
	}
	return nil
}
