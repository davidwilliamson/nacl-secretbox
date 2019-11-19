package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/scrypt"
	"io"
	"io/ioutil"
	"os"
)

func Encrypt(plainText []byte) []byte {
	// You must use a different nonce for each message you encrypt with the
	// same key. Since the nonce here is 192 bits long, a random value
	// provides a sufficiently small probability of repeats.
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}
	secretKey := getSecretKey(nonce)

	// This encrypts "hello world" and appends the result to the nonce.
	encrypted := secretbox.Seal(nonce[:], plainText, &nonce, &secretKey)
	return encrypted
}

func Decrypt(encrypted []byte) []byte {
	// When you decrypt, you must use the same nonce and key you used to
	// encrypt the message. One way to achieve this is to store the nonce
	// alongside the encrypted message. Above, we stored the nonce in the first
	// 24 bytes of the encrypted text.
	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])

	secretKey := getSecretKey(decryptNonce)
	decrypted, ok := secretbox.Open([]byte{}, encrypted[24:], &decryptNonce, &secretKey)
	if !ok {
		panic("decryption error")
	}
	return decrypted
}

func getSecretKey(nonce [24]byte) [32]byte {
	passPhrase := os.Getenv("SECRET_BOX")
	if passPhrase == "" {
		fmt.Println("missing SECRET_BOX env var")
		passPhrase = GetLineFromStdin("Enter secret box passphrase -> ")
		fmt.Println()
	}
	nonceCopy := make([]byte, 24)
	copy(nonceCopy[:], nonce[:24])
    // https://godoc.org/golang.org/x/crypto/scrypt
    // func Key(password, salt []byte, N, r, p, keyLen int) ([]byte, error)
    // Key derives a key from the password, salt, and cost parameters, returning a byte slice of
    // length keyLen that can be used as cryptographic key.
    // N is a CPU/memory cost parameter, which must be a power of two greater than 1. r and p
    // must satisfy r * p < 2³⁰. If the parameters do not satisfy the limits, the function returns
    // a nil byte slice and an error.
    // For example, you can get a derived key for e.g. AES-256 (which needs a 32-byte key) by doing:
    // dk, err := scrypt.Key([]byte("some password"), salt, 32768, 8, 1, 32)
    // The recommended parameters for interactive logins as of 2017 are N=32768, r=8 and p=1.
    // The parameters N, r, and p should be increased as memory latency and CPU parallelism
    // increases; consider setting N to the highest power of 2 you can derive within 100
    // milliseconds. Remember to get a good random salt.
	secretKeyBytes, err := scrypt.Key([]byte(passPhrase), nonceCopy, 32768, 8, 1, 32)
	if err != nil {
		panic(err)
	}

	// This will truncate our secretKeyBytes at 32 bytes
	var secretKey [32]byte
	copy(secretKey[:], secretKeyBytes)
	return secretKey
}

type Account struct {
	Site     string `json:"site"`
	Username string `json:"username"`
	Password string `json:"password"`
	Comment  string `json:"comment,omitempty"`
}

func GetAccountInfoFromStdin() *Account {
	accountInfo := &Account{}
	accountInfo.Site = GetLineFromStdin("Enter Site -> ")
	accountInfo.Username = GetLineFromStdin("Username   -> ")
	accountInfo.Password = GetLineFromStdin("Passeword  -> ")
	accountInfo.Comment = GetLineFromStdin("Comment    -> ")
	fmt.Println()
	return accountInfo
}

func GetLineFromStdin(prompt string) string {
	// fmt.Scanln breaks on any whitespace, so use a bufio.Scanner instead
	// This correctly handles backspace so user can edit inline before hitting enter
	fmt.Printf(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	if scanner.Scan() {
		line = scanner.Text()
		fmt.Printf("Input was: %q\n", line)
	}
	return line
}

func main() {
	// Encrypt
	var message []byte
	accountInfo := GetAccountInfoFromStdin()
	message, err := json.Marshal(accountInfo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Input   '%s'\n", string(message))

	fmt.Println("Encrypted:")
	encrypted := Encrypt(message)
	// fmt.Printf("encrypted: %s\n", encrypted)
	encryptedString := hex.EncodeToString(encrypted[:])
	fmt.Println(encryptedString)

	// Save to file as ascii text
	err = ioutil.WriteFile("./CryptData.dat", []byte(encryptedString), 0600)
	if err != nil {
		panic(err)
	}

	// Decrypt
	// decrypted := Decrypt(encrypted)
	// fmt.Printf("Decoded '%s'\n", string(decrypted))
	// fileData is []byte
	fileData, err := ioutil.ReadFile("./CryptData.dat")
	if err != nil {
		panic(err)
	}
	encryptedBytes, err := hex.DecodeString(string(fileData))
	if err != nil {
		panic(err)
	}
	decrypted := Decrypt(encryptedBytes)
	var decodedAccountInfo Account
	err = json.Unmarshal(decrypted, &decodedAccountInfo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Site: %s\n", decodedAccountInfo.Site)
	fmt.Printf("User: %s\n", decodedAccountInfo.Username)
	fmt.Printf("Pass: %s\n", decodedAccountInfo.Password)
	fmt.Printf("Cmnt: %s\n", decodedAccountInfo.Comment)
}
