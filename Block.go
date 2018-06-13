package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

//Block is struct class for Block
type Block struct {
	Timestamp time.Time
	Lasthash  string
	Hash      string
	Data      []byte
}

//ToString return string representation of the block
func (b Block) toString() string {
	output := fmt.Sprintf(`Block-
		Timestamp: %s
		Last Hash: %s
		Hash: %s
		Data: %d`,
		b.Timestamp.Format(time.RFC3339),
		string(b.Lasthash[0:9]),
		string(b.Hash[0:9]),
		b.Data[0:9])
	return output
}

//EncryptData encrypt data and store the hash value in hash property of block
func (b *Block) encryptData() {
	rawData := append([]byte(b.Timestamp.String()+b.Lasthash), b.Data...)
	hashValue := createHash(rawData)
	b.Hash = hashValue
}

//CreateHash create a has value of given data
func createHash(data []byte) string {
	h := sha256.New()
	h.Write(data)
	md := h.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}
