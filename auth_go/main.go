package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, 12)
	if err != nil {
		fmt.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func main() {
	var password string
	var hashedpw string
	flag.StringVar(&password, "password", "", "the raw password")
	flag.StringVar(&hashedpw, "hashedpw", "", "the hashed (salted) password (use single quotes ' )")
	flag.Parse()

	if password != "" && hashedpw != "" {
		err := bcrypt.CompareHashAndPassword([]byte(hashedpw), []byte(password))
		if err != nil {
			fmt.Println("Password and hashed password do NOT match!", err)
		} else {
			fmt.Println("Password and hashed password do match!")
		}
	} else {
		if password == "" {
			flag.Usage()
			os.Exit(1)
		}
		hash := hashAndSalt([]byte(password))
		fmt.Println("Original:    ", password)
		fmt.Println("Salted hash: ", hash)
	}
}
