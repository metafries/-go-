package dblayer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoDBDataStore struct {
	*mgo.Session
}

func NewMongoDBDataStore(conn string) (*mongoDBDataStore, error) {
	session, err := mgo.Dial(conn)
	if err != nil {
		return nil, err
	}
	return &mongoDBDataStore{Session: session}, nil
}

// In case of mongodb, the id field doesn't auto increment as the case was with mysql.
// So the json string used in the API post reqest body need to supply the id.

func (ms *mongoDBDataStore) AddClub(ci *ClubInfo) error {
	session := ms.Copy()
	defer session.Close()
	clubinfo := session.DB("MF").C("club_info")
	return clubinfo.Insert(ci)
}

func (ms *mongoDBDataStore) FindClub(id int) (ClubInfo, error) {
	session := ms.Copy()
	defer session.Close()
	clubinfo := session.DB("MF").C("club_info")
	ci := ClubInfo{}
	err := clubinfo.Find(bson.M{"id": id}).One(&ci)
	return ci, err
}

func (ms *mongoDBDataStore) AllClubs() (club, error) {
	session := ms.Copy()
	defer session.Close()
	clubinfo := session.DB("MF").C("club_info")
	clubs := club{}
	err := clubinfo.Find(nil).All(&clubs)
	return clubs, err
}
