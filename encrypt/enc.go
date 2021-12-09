package encrypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func GetPrivateKey() *rsa.PrivateKey {
	// Generate a key
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	return privateKey
}

// Encrypt to EncryptOAEP using sha256
func EncryptMessage(jsonstring string, publicKey rsa.PublicKey) []byte {
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		[]byte(jsonstring),
		nil)
	if err != nil {
		panic(err)
	}
	return encryptedBytes
}

// Decrypt the beauty to string
func DecryptionMessage(privateKey *rsa.PrivateKey, encryptedBytes []byte) string {
	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	return string(decryptedBytes)
}
