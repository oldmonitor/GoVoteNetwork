package main

import (
	"bytes"
	"crypto"
	"crypto/sha256"
	"fmt"
	"os"
	"testing"
)

func TestSignature(t *testing.T) {
	var keyPair RSAKeyPair
	keyPair.generateKey()
	var message = []byte("this is a test")

	/*h := sha256.New()
	h.Write(dataToSigned)
	md := h.Sum(nil)
	mdStr := hex.EncodeToString(md)
	*/

	var hashed = sha256.Sum256(message)
	var pKey []byte = keyPair.PrivateKey

	signature, err := SignPKCS1v15(rng, pKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return
	}

	//	var signature = rsaSign(pKey, hashed[:])

	var err = rsaUnsign(keyPair.PublicKey, hashed[:], signature)
	if err == false {
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
