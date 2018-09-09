package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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
		runServer(*address)
	case "C":
		runClient(*address)
	}
}

func runClient(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What message would you like to send?")
	for scanner.Scan() {
		fmt.Println("Writing ", scanner.Text())
		conn.Write(append(scanner.Bytes(), '\r'))

		fmt.Println("What message would you like to send?")
		buffer := make([]byte, 1024)
		// conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		_, err := conn.Read(buffer)

		if err != nil && err != io.EOF {
			log.Fatal("FATAL -> ", err)
		} else if err == io.EOF {
			log.Println("Connection is closed")
			return nil
		}
		fmt.Println(string(buffer))
	}
	return scanner.Err()
}
