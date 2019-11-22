package main

import (
	"io/ioutil"
	"strings"
)

func saveFile(data, filePath string) error {
	// Save to file as ascii text
	return ioutil.WriteFile(filePath, []byte(data), 0600)
}

func readFile(filePath string) (string, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(fileData), nil
}

// ReadDataStore reads the datastore file, decrypts, and returns a slice of Accounts
func ReadDataStore(passPhrase, dataFile string) ([]*Account, error) {
	accounts := make([]*Account, 0)
	encryptedString, err := readFile(dataFile)
	if err != nil {
		panic(err)
	}
	// Empty file
	if encryptedString == "" {
		return accounts, nil
	}
	// The file format is:
	// <encrypted-string-for-account1>\n
	// <encrypted-string-for-account2>\n
	// ...
	// <encrypted-string-for-accountN>
	// So we split on '\n' and each line is the encrypted hex string for an Account
	for _, line := range strings.Split(encryptedString, "\n") {
		account, err := DecryptAccount(line, passPhrase)
		if err != nil {
			panic(err)
		}
		accounts = append(accounts, &account)
	}
	return accounts, err
}

// WriteDataStore accepts a slice of Accounts, encrypts and writes to disk
func WriteDataStore(passPhrase, dataFile string, accounts []*Account) error {
	messages := make([]string, 0)
	for _, account := range accounts {
		encryptedStr, err := account.Encrypt(passPhrase)
		if err != nil {
			panic(err)
		}
		messages = append(messages, string(encryptedStr))
	}
	// The file format is:
	// <encrypted-string-for-account1>\n
	// <encrypted-string-for-account2>\n
	// ...
	// <encrypted-string-for-accountN>
	messageStr := strings.Join(messages, "\n")
	return saveFile(messageStr, dataFile)
}
