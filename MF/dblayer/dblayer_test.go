package dblayer

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkMySQLDBReads(b *testing.B) {
	dblayer, err := ConnectDatabase("mysql", "root:KLin#180812@/MF")
	if err != nil {
		b.Fatal("FATAL: Could not connect to MF chat system -> ", err)
	}
	findCLubsBM(b, dblayer)
}

func BenchmarkMongoDBReads(b *testing.B) {
	dblayer, err := ConnectDatabase("mongodb", "mongodb://127.0.0.1")
	if err != nil {
		b.Error("ERROR: Could not connect to MF chat system -> ", err)
		return
	}
	findCLubsBM(b, dblayer)
}

func findCLubsBM(b *testing.B, dblayer DBLayer) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		_, err := dblayer.FindClub(rand.Intn(2) + 1)
		if err != nil {
			b.Error("ERROR: Query failed -> ", err)
			return
		}
	}
}
