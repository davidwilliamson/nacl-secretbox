package main

import (
	"fmt"
)

func main() {
	dataFile := "./CryptData.dat"
	dirty := false
	done := false

	// Init
	passPhrase := GetMasterPassphrase()
	accounts, err := ReadDataStore(passPhrase, dataFile)
	if err != nil {
		panic(err)
	}

	for !done {
		fmt.Println("\nc (create) f (find) a (list all) d (delete) q (quit)")
		switch choice := GetLineFromStdin("Select -> "); choice {
		case "c":
			accounts = CreateAccount(accounts)
			dirty = true
		case "f":
			searchStr := GetLineFromStdin("Enter site name to find -> ")
			ShowAccount(accounts, searchStr)
		case "a":
			DumpAllAccounts(accounts)
		case "d":
			var removed bool
			searchStr := GetLineFromStdin("Enter site name to delete-> ")
			accounts, removed = RemoveAccount(accounts, searchStr)
			if removed {
				dirty = true
			}
		case "u":
			newPassPhrase := GetLineFromStdin("Enter new pass phrase-> ")
			fmt.Printf("Changing master pass phrase to: %s\n", newPassPhrase)
			if YesNo("This can not be undone. Are you sure? (y/n) -> ") {
				dirty = true
				passPhrase = newPassPhrase
			}
		case "q":
			done = true
		default:
			fmt.Printf("Unknown choice '%s'\n", choice)
		}
	}

	// Encrypt and write to disk
	if dirty {
		err := WriteDataStore(passPhrase, dataFile, accounts)
		if err != nil {
			panic(err)
		}
	}
}
