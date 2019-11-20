package main

import (
	"encoding/json"
	"fmt"
)

// Account represents an online account
type Account struct {
	Site     string `json:"site"`
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
	accountInfo.Site = GetLineFromStdin("Enter Site -> ")
	accountInfo.Username = GetLineFromStdin("Username   -> ")
	accountInfo.Password = GetLineFromStdin("Passeword  -> ")
	accountInfo.Comment = GetLineFromStdin("Comment    -> ")
	fmt.Println()
	return accountInfo
}

// Marshal will json serialze an Account struct
func (acct *Account) Marshal() []byte {
	message, err := json.Marshal(acct)
	if err != nil {
		panic(err)
	}
	return message
}
