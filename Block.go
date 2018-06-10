package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

//Block is struct class for Block
type Block struct {
	timestamp time.Time
	lasthash  string
	hash      string
	data      []byte
}

//ToString return string representation of the block
func (b Block) ToString() string {
	output := fmt.Sprintf(`Block-
		Timestamp: %s
		Last Hash: %s
		Hash: %s
		Data: %d`,
		b.timestamp.Format(time.RFC3339),
		string(b.lasthash[0:9]),
		string(b.hash[0:9]),
		b.data[0:9])
	return output
}

//EncryptData encrypt data and store the hash value in hash property of block
func (b *Block) EncryptData() {
	rawData := append([]byte(b.timestamp.String()+b.lasthash+b.hash), b.data...)
	hashValue := CreateHash(rawData)
	b.hash = hashValue
}

//CreateHash create a has value of given data
func CreateHash(data []byte) string {
	h := sha256.New()
	h.Write(data)
	md := h.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}
