package main

import (
	"fmt"
	"strconv"
	"time"
)

//Block is struct class for Block
type Block struct {
	Timestamp  time.Time
	Lasthash   string
	Hash       string
	Data       []byte
	Difficulty int
	Nonce      int
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
	rawData := append([]byte(b.Timestamp.String()+b.Lasthash+strconv.Itoa(b.Nonce)), b.Data...)
	hashValue := createHash(rawData)
	b.Hash = hashValue
}
