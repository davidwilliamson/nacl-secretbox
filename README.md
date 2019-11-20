# Secret Box

## What it is

A program to sequrely store data to disk.  Uses golang's implementation of [NaCl](https://godoc.org/golang.org/x/crypto/nacl/secretbox) to encrypt/decrypt.

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
- Create a backup file on update
- Handle multiple records
- Ability to change master password
- Save master password hint in data file
- CLI to support CRUD operations
   - new record
   - find record and display
   - modify record
   - delete record
   - dump all records with obscured passwords
   - sort output (Use `golang sort.Slice(list, sortFunc)`)
- Ability to work with binary files (pdf, jpg) in secret box
- YAML file templates for different record types
- Allow master passphrase be provided in either env or user prompt
- Multiple data files: command line flag for name of datafile
- run in Docker container with datafile as external mount
