package main

import (
	"encoding/base64"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	op := flag.String("type", "", "Server (s) or Client (c) ?")
	address := flag.String("addr", ":8000", "address? host:port ")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runUDPServer(*address)
	case "C":
		runUDPClient(*address)
	}
}

func runUDPClient(address string) error {
	conn, err := net.Dial("udp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	filebytes, err := ioutil.ReadFile("inputfile.csv")
	if err != nil {
		log.Fatal("FATAL -> ", err)
	}
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(filebytes)))
	base64.StdEncoding.Encode(dst, filebytes)
	log.Println("Sending ", len(dst), " bytes")
	_, err = conn.Write(dst)
	return err
}

func runUDPServer(address string) error {
	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		log.Fatal("FATAL -> ", err)
	}
	defer pc.Close()

	buffer := make([]byte, 4096)
	log.Println("Listening..... ")

	n, _, err := pc.ReadFrom(buffer)
	if err != nil {
		log.Fatal("FATAL -> ", err)
	}
	dst := make([]byte, base64.StdEncoding.DecodedLen(n))
	_, err = base64.StdEncoding.Decode(dst, buffer[:n])
	if err != nil {
		log.Fatal("FATAL -> ", err)
	}
	file, _ := os.Create("outputfile.csv")
	file.Write(dst)
	err = file.Close()
	return err
}
