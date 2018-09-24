package mfrestapi

import (
	"-go-/MF/dblayer"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type MFClubReqHandler struct {
	dbconn dblayer.DBLayer
}

func NewMFClubReqHandler() *MFClubReqHandler {
	return new(MFClubReqHandler)
}

func (mfcreq *MFClubReqHandler) connect(o, conn string) error {
	dbl, err := dblayer.ConnectDatabase(o, conn)
	if err != nil {
		return err
	}
	mfcreq.dbconn = dbl
	return nil
}

func (mfcreq *MFClubReqHandler) handleMFClubReq(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ids := r.RequestURI[len("/mfclub/"):]
		id, err := strconv.Atoi(ids)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "ERROR: id %s provided is not of a valid number. \n", ids)
			return
		}
		ci, err := mfcreq.dbconn.FindClub(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "ERROR: %s occured while searching for id %d \n", err.Error(), id)
			return
		}
		json.NewEncoder(w).Encode(&ci)
	case "POST":
		ci := new(dblayer.ClubInfo)
		err := json.NewDecoder(r.Body).Decode(ci)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "ERROR: %s occured", err)
			return
		}
		mfcreq.dbconn.AddClub(ci)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "ERROR: %s occured while adding a club info to the MF database", err)
		}
		fmt.Fprintf(w, "INFO: Successfully inserted id %d \n", ci.ID)
	}
}
