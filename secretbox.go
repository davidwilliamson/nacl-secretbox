package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ReadDataStore reads the datastore file, decrypts, and returns account info
func ReadDataStore(passPhrase, dataFile string) ([]*Account, error) {
	var decodedAccountInfo Account
	accounts := make([]*Account, 0)
	encryptedBytes, err := readFile(dataFile)
	if err != nil {
		panic(err)
	}
	encryptedString := string(encryptedBytes)
	// TODO change readFile to return a list of encrypted hex decoded bytes
	for _, line := range strings.Split(encryptedString, "\n") {
		decrypted := Decrypt(passPhrase, []byte(line))
		err = json.Unmarshal(decrypted, &decodedAccountInfo)
		accounts = append(accounts, &decodedAccountInfo)
	}
	return accounts, err
}

// WriteDataStore encrypts the account info and writes to disk
func WriteDataStore(passPhrase, dataFile string, accounts []*Account) error {
	var message []byte
	messages := make([]string, 0)
	// TODO this seems wrong. the newline could be the same hex as part of
	// the encrypted bytes.
	// We want to build a list of hex encoded encrypted bytes,
	// and write that list to disk, one line per entry in the list
	// so the newlines in the file are unambiguously delimiters
	//
	// var accounts []*Account
	// var account *Account
	// var bytes []byte
	// lines := make([]string, 0)
	// hex encoded encrypted data, with newlines as delimiter
	// var messageStr string
	// for account in accounts:
	//    bytes = Encrypt(account.Marshal)
	//    lines = append(lines, hex.EncodeToString(bytes[:]))
	// messageStr = strings.Join(lines, "\n")
	// TODO change datastore.go saveFile() to accept a string and assume it
	//       is already hex encoded with newlines already embedded.
	// saveFile(messageStr)
	for _, account := range accounts {
		message = account.Marshal()
		fmt.Printf("Input   '%s'\n", string(message))
		encryptedBytes := Encrypt(passPhrase, message)
		messages = append(messages, string(encryptedBytes))
	}
	messageStr := strings.Join(messages, "\n")
	return saveFile([]byte(messageStr), dataFile, false)
}

func main() {
	var err error
	dataFile := "./CryptData.dat"

	passPhrase := getMasterPassphrase()
	writeAccounts := make([]*Account, 0)
	for cnt := 1; cnt <= 3; cnt ++ {
		accountInfo := GetAccountInfoFromStdin()
		writeAccounts = append(writeAccounts, accountInfo)
	}
	err = WriteDataStore(passPhrase, dataFile, writeAccounts)
	if err != nil {
		panic(err)
	}

	// Decrypt
	readAccounts := make([]*Account, 0)
	readAccounts, err = ReadDataStore(passPhrase, dataFile)
	if err != nil {
		panic(err)
	}
	for _,  decodedAccountInfo := range(readAccounts) {
		fmt.Println("---- Account ----")
		decodedAccountInfo.Display()
	}
}
