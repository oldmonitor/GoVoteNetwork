package main

import (
	"testing"
)

func TestSignature(t *testing.T) {
	var keyPair RSAKeyPair
	keyPair.generateKey()
	var dataToSigned = []byte("this is a test")

	/*h := sha256.New()
	h.Write(dataToSigned)
	md := h.Sum(nil)
	mdStr := hex.EncodeToString(md)
	*/
	var signature = rsaSign(keyPair.PrivateKey, []byte(dataToSigned))

	var err = rsaUnsign(keyPair.PublicKey, []byte(dataToSigned), signature)
	if err == false {
		t.Errorf("signature is not valid")
	}
}
