package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
)

func HashPassword(password string, salt []byte) string {
	passwordBytes := []byte(password)
	passwordBytes = append(passwordBytes, salt...)

	hasher := sha512.New()
	hasher.Write(passwordBytes)
	hashedPassword := hasher.Sum(nil)

	return hex.EncodeToString(hashedPassword) + "." + hex.EncodeToString(salt)
}

func ComparePassword(password string, hashedPassword string) bool {
	v := strings.Split(hashedPassword, ".")

	salt, err := hex.DecodeString(v[1])
	if err != nil {
		fmt.Println("Error decoding salt:", err)
		return false
	}

	fmt.Println("Salt", salt)

	candidatePassword := HashPassword(password, salt)
	fmt.Println("candidatePassword", candidatePassword)

	return hashedPassword == candidatePassword
}

func GenRandomSalt() []byte {
	salt := make([]byte, 30)

	_, err := rand.Read(salt)

	if err != nil {
		panic(err)
	}

	return salt
}
