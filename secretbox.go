package main

import (
	"fmt"
)

func main() {
	dataFile := "./CryptData.dat"

	passPhrase := getMasterPassphrase()
	// Generate account data
	writeAccounts := make([]*Account, 0)
	for cnt := 1; cnt <= 3; cnt++ {
		accountInfo := GetAccountInfoFromStdin()
		writeAccounts = append(writeAccounts, accountInfo)
	}

	// Encrypt and write to disk
	err := WriteDataStore(passPhrase, dataFile, writeAccounts)
	if err != nil {
		panic(err)
	}

	// Decrypt from disk
	readAccounts, err := ReadDataStore(passPhrase, dataFile)
	if err != nil {
		panic(err)
	}

	// dump so user can see we decrypted OK.
	for _, decodedAccountInfo := range readAccounts {
		fmt.Println("---- Account ----")
		decodedAccountInfo.Display()
	}
}
