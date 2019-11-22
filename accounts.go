package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

// Account represents an online account
type Account struct {
	Site     string `json:"site"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Comment  string `json:"comment,omitempty"`
}

// NewAccount creates an Account struct and returns a pointer to it
func NewAccount() *Account {
	return &Account{}
}

// GetAccountInfoFromStdin prompts the user for data and returns a ptr to Account
func GetAccountInfoFromStdin() *Account {
	accountInfo := NewAccount()
	accountInfo.Site = GetLineFromStdin("Site name  -> ")
	accountInfo.URL = GetLineFromStdin("URL        -> ")
	accountInfo.Username = GetLineFromStdin("Username   -> ")
	accountInfo.Password = GetLineFromStdin("Password   -> ")
	accountInfo.Comment = GetLineFromStdin("Comment    -> ")
	fmt.Println()
	return accountInfo
}

// Encrypt will encrupt the Account struct
func (acct *Account) Encrypt(passPhrase string) (string, error) {
	message, err := json.Marshal(acct)
	if err != nil {
		return "", err
	}
	encryptedBytes := Encrypt(passPhrase, message)
	encryptedString := hex.EncodeToString(encryptedBytes[:])
	return encryptedString, nil
}

// ContainsStr returns true if string is found in Site or URL field
func (acct *Account) ContainsStr(searchStr string) bool {
	return strings.Contains(acct.Site, searchStr) || strings.Contains(acct.URL, searchStr)
}

// DecryptAccount will take a hex-encoded string and return an Account struct
func DecryptAccount(encryptedString, passPhrase string) (Account, error) {
	var account Account
	encryptedBytes, err := hex.DecodeString(encryptedString)
	if err != nil {
		return account, err
	}
	decryptedString := Decrypt(passPhrase, encryptedBytes)
	err = json.Unmarshal(decryptedString, &account)
	return account, err
}

// Display dumps the account info to STDOUT
func (acct *Account) Display() {
	fmt.Printf("Site: %s\n", acct.Site)
	fmt.Printf("URL : %s\n", acct.URL)
	fmt.Printf("User: %s\n", acct.Username)
	fmt.Printf("Pass: %s\n", acct.Password)
	fmt.Printf("Cmnt: %s\n", acct.Comment)
}
