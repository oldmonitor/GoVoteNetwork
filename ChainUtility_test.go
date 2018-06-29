package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"testing"
)

func TestSignatureWithGeneratedKey(t *testing.T) {
	var keyPair RSAKeyPair
	keyPair.generateKey()
	var message = []byte("this is a test")
	var hashed = sha256.Sum256(message)

	var publicKey, _ = x509.ParsePKCS1PublicKey(keyPair.PublicKey)
	var privateKey, _ = x509.ParsePKCS1PrivateKey(keyPair.PrivateKey)
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])

	error := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)

	if error != nil {
		t.Errorf("signature is not valid")
	}
}

func TestVerifySignatureFunction(t *testing.T) {
	var keyPair RSAKeyPair
	keyPair.generateKey()
	var message = []byte("this is a test")
	var hashed = createHash(message)

	var signature = rsaSign(keyPair.PrivateKey, []byte(hashed))

	error := rsaUnsign(keyPair.PublicKey, []byte(hashed), signature)

	if error != true {
		t.Errorf("signature is not valid")
	}
}
func TestByteArrayStringConversion(t *testing.T) {
	sourceData := []byte{1, 2, 3}
	finalData := convertHexStringToByteArray(convertByteArrayToHexString(sourceData))
	if bytes.Equal(sourceData, finalData) == false {
		t.Errorf("array string conversion is not valid")
	}
}
