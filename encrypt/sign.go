package encrypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"hash"
)

func GetHashMessage(message []byte) hash.Hash {
	//msg := []byte(message)
	// Hashing
	msgHash := sha256.New()
	_, err := msgHash.Write(message)
	if err != nil {
		panic(err)
	}
	return msgHash
}

func GenerateSignature(privateKey *rsa.PrivateKey, msgHashSum []byte) []byte {
	// Signing
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}
	return signature
}

func CheckSignature(signature []byte, publicKey rsa.PublicKey, msgHashSum []byte) {
	// Verify
	err := rsa.VerifyPSS(&publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		fmt.Println("No se pudo verificar")
		return
	}

	// If there isn't any error
	fmt.Println("signature verified")
}
