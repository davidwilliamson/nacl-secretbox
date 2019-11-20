package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func saveFile(data []byte, filePath string, makeBackup bool) error {
	// fmt.Printf("encrypted: %s\n", encrypted)
	encryptedString := hex.EncodeToString(data[:])
	fmt.Println(encryptedString)

	// Save to file as ascii text
	return ioutil.WriteFile(filePath, []byte(encryptedString), 0600)
}

func readFile(filePath string) ([]byte, error) {
	// Decrypt
	// decrypted := Decrypt(encrypted)
	// fmt.Printf("Decoded '%s'\n", string(decrypted))
	// fileData is []byte
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		var b []byte
		return b, err
	}
	return hex.DecodeString(string(fileData))
}
