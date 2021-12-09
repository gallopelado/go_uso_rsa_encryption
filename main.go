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
	//fmt.Println(len(signature))

	// verify signature
	encrypt.CheckSignature(signature, pubkey, msgHashSum)

	// decrypt
	decrypted := encrypt.DecryptionMessage(prvkey, encryptedBytes)
	fmt.Println(decrypted)

	// Playing with message
	m := &MetadataServer{
		CodigoServidor: "000001",
		Signature:      signature,
	}
	res, err := m.ToJson()
	if err != nil {
		fmt.Println("Problems")
	}
	fmt.Println(string(res))
	maembo := string(m.Signature)
	//fmt.Printf("Payload mambeado %v\n", string(m.Signature))
	mamberto := []byte(maembo)
	fmt.Printf("Payload mambeado y masajeado %v\n", mamberto)

	fmt.Println("Re-doing the verification")
	encrypt.CheckSignature(signature, pubkey, msgHashSum)

}
