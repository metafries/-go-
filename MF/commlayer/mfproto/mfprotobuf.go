package mfproto

import (
	"errors"
	"io/ioutil"
	"log"
	"net"

	"github.com/golang/protobuf/proto"
)

type ProtoHandler struct{}

// Constructor for the protocol buffer sender
func NewProtoHandler() *ProtoHandler {
	return new(ProtoHandler)
}

func (pSender *ProtoHandler) EncodeAndSend(obj interface{}, destination string) error {
	v, ok := obj.(*Club)
	if !ok {
		return errors.New("Proto: Unknown Message Type")
	}
	data, err := proto.Marshal(v)
	if err != nil {
		return err
	}
	return sendmessage(data, destination)
}

func (pSender *ProtoHandler) DecodeProto(buffer []byte) (*Club, error) {
	pb := new(Club)
	return pb, proto.Unmarshal(buffer, pb)
}

func (pSender *ProtoHandler) ListenAndDecode(listenaddress string) (chan interface{}, error) {
	outChan := make(chan interface{})
	l, err := net.Listen("tcp", listenaddress)
	if err != nil {
		return outChan, err
	}
	log.Println("INFO: Listening to", listenaddress)
	go func() {
		defer l.Close()
		for {
			c, err := l.Accept()
			if err != nil {
				break
			}
			log.Println("INFO: Accepted Connection From ", c.RemoteAddr())
			go func(c net.Conn) {
				defer c.Close()
				for {
					buffer, err := ioutil.ReadAll(c)
					if err != nil {
						break
					}
					if len(buffer) == 0 {
						continue
					}
					obj, err := pSender.DecodeProto(buffer)
					if err != nil {
						continue
					}
					select {
					case outChan <- obj:
					// case <-time.After(1 * time.Second):
					default:
					}
				}
			}(c)
		}
	}()
	return outChan, nil
}

func sendmessage(buffer []byte, destination string) error {
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Printf("Sending %d bytes to %s \n", len(buffer), destination)
	_, err = conn.Write(buffer)
	return err
}
