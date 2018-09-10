package main

import (
	"bufio"
	"flag"
	"fmt"
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

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What message would you like to send?")
	for scanner.Scan() && err == nil {
		fmt.Println("Writing ", scanner.Text())
		conn.Write(scanner.Bytes())
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			log.Fatal("FATAL -> ", err)
		}
		fmt.Println(string(buffer))
		fmt.Println("What message would you like to send?")
	}
	return nil
}

func runUDPServer(address string) error {
	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		log.Fatal("FATAL -> ", err)
	}
	defer pc.Close()

	buffer := make([]byte, 1024)
	log.Println("Listening..... ")
	for {
		_, addr, _ := pc.ReadFrom(buffer)
		fmt.Printf("Received %s from address %s \n", string(buffer), addr)
		_, err := pc.WriteTo([]byte("Message Received"), addr)
		if err != nil {
			log.Fatal("FATAL -> Could not write back on connection: ", err)
		}
	}
}
