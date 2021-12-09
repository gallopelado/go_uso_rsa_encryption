package main

import (
	"fmt"
	"uso_rsa_encryption/encrypt"
)

func main() {
	prvkey := encrypt.GetPrivateKey()
	pubkey := prvkey.PublicKey
	mensaje := "Caldito de pescado"

	// Using only sha256
	encryptedBytes := encrypt.EncryptMessage(mensaje, pubkey)
	// hacer hash
	msgHash := encrypt.GetHashMessage(encryptedBytes)
	msgHashSum := msgHash.Sum(nil)

	// sign on the hash
	signature := encrypt.GenerateSignature(prvkey, msgHashSum)

	// verify signature
	encrypt.CheckSignature(signature, pubkey, msgHashSum)

	// decrypt
	decrypted := encrypt.DecryptionMessage(prvkey, encryptedBytes)
	fmt.Println(decrypted)
}
