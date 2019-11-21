package main

import (
	"encoding/json"
	"fmt"
)

// ReadDataStore reads the datastore file, decrypts, and returns account info
func ReadDataStore(passPhrase, dataFile string) (*Account, error) {
	encryptedBytes, err := readFile(dataFile)
	if err != nil {
		panic(err)
	}
	decrypted := Decrypt(passPhrase, encryptedBytes)
	var decodedAccountInfo Account
	err = json.Unmarshal(decrypted, &decodedAccountInfo)
	return &decodedAccountInfo, err
}

// WriteDataStore encrypts the account info and writes to disk
func WriteDataStore(passPhrase, dataFile string, accountInfo *Account) error {
	var message []byte
	message = accountInfo.Marshal()
	fmt.Printf("Input   '%s'\n", string(message))

	// Encrypt
	fmt.Println("Encrypted:")
	encrypted := Encrypt(passPhrase, message)
	return saveFile(encrypted, dataFile, false)
}

func main() {
	var err error
	dataFile := "./CryptData.dat"

	passPhrase := getMasterPassphrase()
	accountInfo := GetAccountInfoFromStdin()
	err = WriteDataStore(passPhrase, dataFile, accountInfo)
	if err != nil {
		panic(err)
	}

	// Decrypt
	decodedAccountInfo, err := ReadDataStore(passPhrase, dataFile)
	if err != nil {
		panic(err)
	}
	decodedAccountInfo.Display()
}
