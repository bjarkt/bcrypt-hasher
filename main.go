package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	plainTextPassword := flag.String("pass", "", "Password to hash")
	hashedPassword := flag.String("hash", "", "Hashed password to compare")
	shouldBase64Encode := flag.Bool("b64", false, "Should the output hash be base 64 encoded?")
	flag.Parse()

	if len(*plainTextPassword) == 0 {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	if len(*hashedPassword) > 0 {
		match := checkPasswordHash(*plainTextPassword, *hashedPassword)
		if match {
			fmt.Println("Match")
		} else {
			fmt.Println("No match")
		}
	} else {
		hash, err := hashPassword(*plainTextPassword)
		if err != nil {
			log.Fatalf("Could not hash password: %v", err)
		}

		if *shouldBase64Encode {
			b64EncodedHash := base64.StdEncoding.EncodeToString([]byte(hash))
			fmt.Println(b64EncodedHash)
		} else {
			fmt.Println(hash)
		}
	}

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
