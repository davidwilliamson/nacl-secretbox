package main

import (
	"fmt"
)

// CreateAccount creates a new account and appends it to the accounts slice
func CreateAccount(accounts []*Account) []*Account {
	accountInfo := GetAccountInfoFromStdin()
	return append(accounts, accountInfo)
}

// FindAccount finds the Account that contains searchStr
func FindAccount(accounts []*Account, searchStr string) (*Account, int, bool) {
	for index, account := range accounts {
		if account.ContainsStr(searchStr) {
			return account, index, true
		}
	}
	return NewAccount(), -1, false
}

// ShowAccount displays the Account info for the account that contains searchStr
func ShowAccount(accounts []*Account, searchStr string) {
	account, _, found := FindAccount(accounts, searchStr)
	if found {
		fmt.Println("---- Account ----")
		account.Display()
		return
	}
	fmt.Printf("No account matches '%s'", searchStr)
}

// RemoveAccount removes the Account for the account that contains searchStr
func RemoveAccount(accounts []*Account, searchStr string) ([]*Account, bool) {
	account, index, found := FindAccount(accounts, searchStr)
	if found {
		fmt.Println("---- Account ----")
		account.Display()
		if YesNo("OK to remove (y/n) ") {
			accounts[len(accounts)-1], accounts[index] = accounts[index], accounts[len(accounts)-1]
			return accounts[:len(accounts)-1], true
		}
	}
	return accounts, false
}

// DumpAllAccounts displays info for all accounts
func DumpAllAccounts(accounts []*Account) {
	// dump so user can see we decrypted OK.
	for _, account := range accounts {
		fmt.Println("---- Account ----")
		account.Display()
	}
}

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
