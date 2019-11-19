package main

import (
	"crypto/rand"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/scrypt"
	"io"
)

func Encrypt(passPhrase string, plainText []byte) []byte {
	// You must use a different nonce for each message you encrypt with the
	// same key. Since the nonce here is 192 bits long, a random value
	// provides a sufficiently small probability of repeats.
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}
	secretKey := getSecretKey(passPhrase, nonce)

	// This encrypts "hello world" and appends the result to the nonce.
	encrypted := secretbox.Seal(nonce[:], plainText, &nonce, &secretKey)
	return encrypted
}

func Decrypt(passPhrase string, encrypted []byte) []byte {
	// When you decrypt, you must use the same nonce and key you used to
	// encrypt the message. One way to achieve this is to store the nonce
	// alongside the encrypted message. Above, we stored the nonce in the first
	// 24 bytes of the encrypted text.
	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])

	secretKey := getSecretKey(passPhrase, decryptNonce)
	decrypted, ok := secretbox.Open([]byte{}, encrypted[24:], &decryptNonce, &secretKey)
	if !ok {
		panic("decryption error")
	}
	return decrypted
}

func getSecretKey(passPhrase string, nonce [24]byte) [32]byte {
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
