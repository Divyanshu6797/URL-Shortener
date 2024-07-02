package main

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	Originalurl  string    `json:"original_url"`
	Shorturl     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

// hash generator
func generateShortURL(originalUrl string) string {
	hasher := md5.New()               //converts to hash
	hasher.Write([]byte(originalUrl)) // converts to byte
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:8]

}

func main() {

}
