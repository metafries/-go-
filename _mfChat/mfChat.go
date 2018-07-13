package _mfChat

import (
	"-go-/_mfLogger"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var logger = _mfLogger.GetInstance()

// Start MF Chat
func Run(connection string) error {
	l, err := net.Listen("tcp", connection)
	if err != nil {
		logger.Println("Error Connecting to Chat Client ", err)
		return err
	}
	r := CreateRoom("MFChat")
	go func() {
		// Hangdle SIGINT and SIGTERM
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		l.Close()
		fmt.Println("Closing TCP Connection")
		close(r.Quit)
		if r.CLCount() > 0 {
			<-r.Msgch
		}
		os.Exit(0)
	}()
	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Println("Error Accepting Connection From Chat Client ", err)
			break
		}
		go handleConnection(r, conn)
	}
	return err
}

func handleConnection(r *room, c net.Conn) {
	logger.Println("Received Request From Client ", c.RemoteAddr)
	r.AddClient(c)
}
