package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"github.com/fahadalisarwar1/CloryServer/core"
)

type person struct {
	name string
}

func(p *person) printName(){
	p.name = "fahad"

}
func main(){
	data := []byte("our super secret text")
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