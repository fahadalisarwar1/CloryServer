package main

import (
	"encoding/hex"
	"fmt"
	"github.com/fahadalisarwar1/CloryServer/core"
	"io/ioutil"
	"log"
)

func main(){
	filename := "/home/fahad/go/src/github.com/fahadalisarwar1/CloryServer/myfile.txt"

	data, _ := ioutil.ReadFile(filename)
	key, err := core.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	ciphertext, err := core.Encrypt(key, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ciphertext: %s\n", hex.EncodeToString(ciphertext))
	plaintext, err := core.Decrypt(key, ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("plaintext: %s\n", plaintext)
}
