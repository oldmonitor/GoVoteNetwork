package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
)

type RSAKeyPair struct {
	PublicKey  []byte
	PrivateKey []byte
}

func (kp *RSAKeyPair) generateKey() {
	reader := rand.Reader
	bitSize := 2048

	//generate key pair
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	//save public key
	asn1Bytes, err := asn1.Marshal(key.PublicKey)
	kp.PublicKey = asn1Bytes
	checkError(err)

	//save private key
	privDER := x509.MarshalPKCS1PrivateKey(key)
	kp.PrivateKey = privDER
}
