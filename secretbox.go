package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	// Encrypt
	var message []byte
	var err error

	passPhrase := getMasterPassphrase()
	accountInfo := GetAccountInfoFromStdin()
	message = accountInfo.Marshal()
	fmt.Printf("Input   '%s'\n", string(message))

	fmt.Println("Encrypted:")
	encrypted := Encrypt(passPhrase, message)
	// fmt.Printf("encrypted: %s\n", encrypted)
	encryptedString := hex.EncodeToString(encrypted[:])
	fmt.Println(encryptedString)

	// Save to file as ascii text
	err = ioutil.WriteFile("./CryptData.dat", []byte(encryptedString), 0600)
	if err != nil {
		panic(err)
	}

	// Decrypt
	// decrypted := Decrypt(encrypted)
	// fmt.Printf("Decoded '%s'\n", string(decrypted))
	// fileData is []byte
	fileData, err := ioutil.ReadFile("./CryptData.dat")
	if err != nil {
		panic(err)
	}
	encryptedBytes, err := hex.DecodeString(string(fileData))
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
