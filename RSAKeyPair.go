package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

//RSAKeyPair struct for private & public key pair
type RSAKeyPair struct {
	PublicKey  []byte
	PrivateKey []byte
}

//generate public private key pair
func (kp *RSAKeyPair) generateKey() {
	reader := rand.Reader
	bitSize := 2048

	//generate key pair
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	//save public key
	asn1Bytes := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	kp.PublicKey = asn1Bytes
	checkError(err)

	//save private key
	privDER := x509.MarshalPKCS1PrivateKey(key)
	kp.PrivateKey = privDER

}
