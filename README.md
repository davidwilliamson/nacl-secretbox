https://godoc.org/golang.org/x/crypto/nacl/secretbox

https://godoc.org/golang.org/x/crypto/nacl/secretbox#example-package

A worked example
https://godoc.org/github.com/andmarios/crypto/nacl/saltsecret

TODO
1. have master passphrase be provided in either env or user prompt
2. CLI to support CRUD operations
   - new record
   - find record and display
   - modify record
   - delete record
   - dump all records with obscured passwords
   - sort output (golang sort.Slice(list, sortFunc) method
3. search using command line flag ./secretbox -f 'website'
4. command line flag for name of datafile
5. create a backup file on update
6. run in container with datafile as external mount
7. Ability to read binary files (pdf, jpg) into secret box;
   ability to write them back out as clear files.
