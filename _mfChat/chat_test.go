package _mfChat

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"
	"testing"
	"time"
)

var once sync.Once

func chatServerFunc(t *testing.T) func() {
	return func() {
		t.Log("[INFO] Starting MF Chat Server ...")
		if err := Run(":2300"); err != nil {
			t.Error("[ERROR] Could not start chat server:", err)
			return
		} else {
			t.Log("[INFO] Started MF Chat Server ...")
		}
	}
}

func TestRun(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode ...")
	}
	t.Log("[INFO] Testing MF Chat Send and Receive ...")

	go once.Do(chatServerFunc(t))

	time.Sleep(1 * time.Second) // wait for one second assuming the chat server succeeded

	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("Anonymous%d", rand.Intn(400))

	t.Logf("[INFO] Hello %s, Connecting to the MF Chat System ... \n", name)
	conn, err := net.Dial("tcp", "127.0.0.1:2300")
	if err != nil {
		t.Fatal("[FATAL] Could not connect to MF chat system:", err)
	}
	t.Log("[INFO] Connected to MF Chat System")
	name += ":"
	defer conn.Close()
	msgCh := make(chan string)

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			recvmsg := scanner.Text()
			sentmsg := <-msgCh
			if strings.Compare(recvmsg, sentmsg) != 0 {
				t.Errorf("[ERROR] Chat message does not match: %s | %s", recvmsg, sentmsg)
			}
		}
	}()
	for i := 0; i <= 10; i++ {
		msgBody := fmt.Sprintf("RandomMessage %d", rand.Intn(400))
		msg := name + msgBody
		_, err := fmt.Fprintf(conn, msg+"\n")
		if err != nil {
			t.Error(err)
			return
		}
		msgCh <- msg
	}
}

func TestServerConnection(t *testing.T) {
	t.Log("[INFO] Test MF Chat Receive Messages ...")
	f := chatServerFunc(t)
	go once.Do(f)
	time.Sleep(1 * time.Second) // wait for one second assuming the chat server succeeded

	conn, err := net.Dial("tcp", "127.0.0.1:2300")
	if err != nil {
		t.Fatal("[FATAL] Could not connect to MF chat system:", err)
	}
	conn.Close()
}
