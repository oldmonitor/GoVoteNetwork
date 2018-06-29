package main

/*
Location of common utility methods
*/
import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

//general method to handle errors
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

// newUUID generates a random UUID according to RFC 4122
// example code from Golang playground: https://play.golang.org/p/4FkNSiUDMg
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

//CreateHash create a has value of given data
func createHash(data []byte) string {
	h := sha256.New()
	h.Write(data)
	md := h.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}

//sign hash value with private key
func rsaSign(privateKey []byte, data []byte) []byte {
	var pKey, _ = x509.ParsePKCS1PrivateKey(privateKey)
	signature, _ := rsa.SignPKCS1v15(rand.Reader, pKey, crypto.SHA256, data)
	return signature
}

//verify the signature with public key
func rsaUnsign(publicKey []byte, message []byte, signature []byte) bool {
	var pKey, _ = x509.ParsePKCS1PublicKey(publicKey)
	error := rsa.VerifyPKCS1v15(pKey, crypto.SHA256, message, signature)

	if error != nil {
		return true
	}

	return false

}

func convertByteArrayToHexString(data []byte) string {
	return hex.EncodeToString(data)
}

func convertHexStringToByteArray(data string) []byte {
	arrayData, _ := hex.DecodeString(data)
	return arrayData
}
