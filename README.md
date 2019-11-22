# Secret Box

## What it is

A program to securely store data to disk. Uses golang's implementation of [NaCl](https://godoc.org/golang.org/x/crypto/nacl/secretbox) to encrypt/decrypt.

Currently in proof of concept state: proof we can encrypt and decrypt the data
- Read data from STDIN
- Encrypt data, write encrypted data out to file
- Read file and decrypt contents
- Display decrypted data to STDOUT

## Build

`make build`

## Run

`./bin/nacl-secretbox`

## TODO
- Need Unit tests
    - encrypt -> decrypt recreates the original Account struct (all fields match exactly)
    - encrypt -> write-file -> read-file -> decrypt recreates original Account structs
    - Two copies of the same Account encrypt to different strings (nonce is used correctly)
    - Decrypting with wrong masterPassword casues error
    - Decrypting with wrong masterPassword does not corrupt dataStore
- Vendor golang.org/x/crypto/nacl/secretbox
- Create a backup file on update
- Save master password hint in data file
- CLI to support CRUD operations
   - modify record
   - sort output (Use `golang sort.Slice(list, sortFunc)`)
- Ability to work with binary files (pdf, jpg) in secret box
- YAML file templates for different record types
- Multiple data files: command line flag for name of datafile
- run in Docker container with datafile as external mount
