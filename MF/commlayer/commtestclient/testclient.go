package main

import (
	"-go-/MF/commlayer"
	"-go-/MF/commlayer/mfproto"
	"flag"
	"log"
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

func runServer(dest string) {
	c := commlayer.NewConnection(commlayer.Protobuf)
	recvChan, err := c.ListenAndDecode(dest)
	if err != nil {
		log.Fatal("FATAL: ", err)
	}
	for msg := range recvChan {
		log.Println("INFO: Received ", msg)
	}
}

func runClient(dest string) {
	c := commlayer.NewConnection(commlayer.Protobuf)
	club := &mfproto.Club{
		Clubname:   "Spurs",
		LeagueName: "Premier League",
		CI: []*mfproto.Club_ClubInfo{
			&mfproto.Club_ClubInfo{1, "Tottenham Hotspur", "3_2017-18", "Premier League"},
			&mfproto.Club_ClubInfo{2, "Tottenham Hotspur", "3_2015-16", "Premier League"},
		},
	}

	if err := c.EncodeAndSend(club, dest); err != nil {
		log.Println("ERROR: Error occured while sending message", err)
	} else {
		log.Println("INFO: Send operation successful")
	}

}
