package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"fmt"
	"testing"
)

//keys generated must not be null
func TestRSAGenerateNonNullKeys(t *testing.T) {
	var keyPair RSAKeyPair
	keyPair.generateKey()

	if len(keyPair.PrivateKey) == 0 {
		t.Errorf("private key is empty")
		return
	}

	if len(keyPair.PublicKey) == 0 {
		t.Errorf("Public key is empty")
		return
	}
}

func TestEncryptionDecription(t *testing.T) {
	var keyPair RSAKeyPair
	keyPair.generateKey()
	var dataToEncrypt = "this is a test"
	var encryptedData []byte
	var decryptedData []byte

	//encrypt data
	var publicKey, _ = x509.ParsePKCS1PublicKey(keyPair.PublicKey)
	encryptedData, _ = rsa.EncryptOAEP(sha1.New(), rand.Reader, publicKey, []byte(dataToEncrypt), []byte("test"))

	fmt.Println(string([]byte(dataToEncrypt)))
	if len(encryptedData) == 0 {
		t.Errorf("encrypted byte array is empty")
		return
	}

	var privateKey, _ = x509.ParsePKCS1PrivateKey(keyPair.PrivateKey)
	decryptedData, _ = rsa.DecryptOAEP(sha1.New(), rand.Reader, privateKey, encryptedData, []byte("test"))

	if string(decryptedData) != dataToEncrypt {
		t.Errorf("decrypted data is not the same as original data. Expected:%s. Received:%s", dataToEncrypt, string(decryptedData))
	}

}
