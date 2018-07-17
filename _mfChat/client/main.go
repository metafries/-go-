package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("Anonymous%d", rand.Intn(400))
	fmt.Println("Starting MFChatClient ...")
	fmt.Println("What's your name?")
	fmt.Scanln(&name)

	fmt.Println("Hello %s, connecting to the MF chat system ... \n", name)
	conn, err := net.Dial("tcp", "127.0.0.1:2300")
	if err != nil {
		log.Fatal("Could not connect to MF Chat system: ", err)
	}
	fmt.Println("Connected to MF Chat system")
	name += ":"
	defer conn.Close()
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && err == nil {
		msg := scanner.Text()
		_, err = fmt.Fprintf(conn, name+msg+"\n")
	}
}
