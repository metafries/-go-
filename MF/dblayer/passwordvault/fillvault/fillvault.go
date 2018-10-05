package main

import (
	"-go-/MF/dblayer/passwordvault"
	"crypto/md5"
)

func main() {
	db, err := passwordvault.ConnectPasswordVault()
	if err != nil {
		return
	}
	carlpss := md5.Sum([]byte("carlspass"))
	passwordvault.AddBytesToVault(db, "Carl", carlpss[:])
	db.Close()
}
