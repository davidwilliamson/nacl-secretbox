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
