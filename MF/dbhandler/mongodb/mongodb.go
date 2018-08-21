package main

import (
	"log"
	"sync"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type clubInfo struct {
	ID      int    `bson:"id"`
	Name    string `bson:"name"`
	Ranking string `bson:"ranking(2017-2018)"`
	League  string `bson:"league"`
}

type Club []clubInfo

func main() {

	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Fatal("[FATAL] mgo - [Dial]: ", err)
	}
	defer session.Close()

	// Get the named collection.
	clubinfo := session.DB("MF").C("club_info")

	// Get the total number of documents in the collection.
	n, _ := clubinfo.Count()
	log.Println("[INFO] mgo - [Collection.Count]: Number of docs of club_info:", n)

	// Perform simple query.
	ci := clubInfo{}
	clubinfo.Find(bson.M{"id": 1}).One(&ci)
	log.Println("[INFO] mgo - [Collection.Find.One]: Docs of club_info at id = 1:", ci)

	// // Insert
	// newcl := clubInfo{ID: 3, Name: "Tottenham Hotspur", Ranking: "3", League: "Premier League"}
	// if err := clubinfo.Insert(newcl); err != nil {
	// 	log.Fatal("[FATAL] mgo - [Collection.Insert]: ", err)
	// }
	// log.Println("[INFO] mgo - [Collection.Insert]: <<<< SUCCESSFULLY INSERTED >>>>")

	// Update
	err = clubinfo.Update(bson.M{"id": 2}, bson.M{"$set": bson.M{"ranking(2017-2018)": "1"}})
	if err != nil {
		log.Fatal("[FATAL] mgo - [Collection.Update]: ", err)
	}
	log.Println("[INFO] mgo - [Collection.Update]: <<<< SUCCESSFULLY UPDATED >>>>")

	// // Remove
	// if err := clubinfo.Remove(bson.M{"id": 3}); err != nil {
	// 	log.Fatal("[FATAL] mgo - [Collection.Remove]: ", err)
	// }
	// log.Println("[INFO] mgo - [Collection.Remove]: <<<< SUCCESSFULLY REMOVED >>>>")

	// Query with expression.
	query := bson.M{
		"id": bson.M{
			"$gt": 0,
		},
		"league": bson.M{
			"$in": []string{"Premier League", "Lega Serie A"},
		},
	}

	var club Club
	err = clubinfo.Find(query).All(&club)
	if err != nil {
		log.Fatal("[FATAL] mgo - [Collection.Find.All]: ", err)
	}
	log.Println("[INFO] mgo - [Collection.Find.All]: Query Results:", club)

	// Use select to get names only
	names := []struct {
		Name string `bson:"name"`
	}{}

	err = clubinfo.Find(query).Select(bson.M{"name": 1}).All(&names)
	if err != nil {
		log.Fatal("[FATAL] mgo - [Collection.Find.Select.All]: ", err)
	}
	log.Println("[INFO] mgo - [Collection.Find.Select.All]: Query Results(names):", names)

	// Concurrent access
	var wg sync.WaitGroup
	count, _ := clubinfo.Count()
	wg.Add(count)
	for i := 1; i <= count; i++ {
		go readId(i, session.Copy(), &wg)
	}
	wg.Wait()
}

func readId(id int, sessionCopy *mgo.Session, wg *sync.WaitGroup) {
	defer func() {
		sessionCopy.Close()
		wg.Done()
	}()
	p := sessionCopy.DB("MF").C("club_info")
	ci := clubInfo{}
	err := p.Find(bson.M{"id": id}).One(&ci)
	if err != nil {
		log.Printf("[ERROR] * - [readId]: Could not retrieve id %d, error %s\n", id, err.Error())
		return
	}
	log.Println("[INFO] * - [readId]: Docs of club_info at id =", id, "=>", ci)
}
