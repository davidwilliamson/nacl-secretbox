package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var message []byte
	var err error

	passPhrase := getMasterPassphrase()
	accountInfo := GetAccountInfoFromStdin()
	message = accountInfo.Marshal()
	fmt.Printf("Input   '%s'\n", string(message))

	// Encrypt
	fmt.Println("Encrypted:")
	encrypted := Encrypt(passPhrase, message)
	err = saveFile(encrypted, "./CryptData.dat", false)
	if err != nil {
		panic(err)
	}
	// Decrypt
	encryptedBytes, err := readFile("./CryptData.dat")
	if err != nil {
		panic(err)
	}
	decrypted := Decrypt(passPhrase, encryptedBytes)
	var decodedAccountInfo Account
	err = json.Unmarshal(decrypted, &decodedAccountInfo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Site: %s\n", decodedAccountInfo.Site)
	fmt.Printf("User: %s\n", decodedAccountInfo.Username)
	fmt.Printf("Pass: %s\n", decodedAccountInfo.Password)
	fmt.Printf("Cmnt: %s\n", decodedAccountInfo.Comment)
}
