package main

import (
	"encoding/json"
	"fmt"
)


type Account struct {
	Site     string `json:"site"`
	Username string `json:"username"`
	Password string `json:"password"`
	Comment  string `json:"comment,omitempty"`
}

func NewAccount() *Account {
	return &Account{}
}

func GetAccountInfoFromStdin() *Account {
	accountInfo := NewAccount()
	accountInfo.Site = GetLineFromStdin("Enter Site -> ")
	accountInfo.Username = GetLineFromStdin("Username   -> ")
	accountInfo.Password = GetLineFromStdin("Passeword  -> ")
	accountInfo.Comment = GetLineFromStdin("Comment    -> ")
	fmt.Println()
	return accountInfo
}

func (acct *Account) Marshal() []byte {
    message, err := json.Marshal(acct)
    if err != nil {
        panic(err)
    }
	return message
}

