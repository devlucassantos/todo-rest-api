package crypto

import (
	"encoding/hex"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

func HashPassword(password string) (string, string, error) {
	salt := generateSalt()
	bytes := applySaltToPass(password, salt)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)

	if err != nil {
		log.Error(err.Error())
		return "", "", err
	}

	return encodeHashToString(hashedPassword), encodeHashToString(salt), nil
}

func ComparePassword(password string, salt string, hashedPwd string) bool {
	hashSalt, err := decodedHashFromString(salt)
	if err != nil {
		return false
	}

	hashPwd, err := decodedHashFromString(hashedPwd)
	if err != nil {
		return false
	}

	bytes := applySaltToPass(password, hashSalt)
	if err = bcrypt.CompareHashAndPassword(hashPwd, bytes); err != nil {
		return false
	}

	return true
}

func generateSalt() []byte {
	size := 256
	bytes := make([]byte, size)

	_, err := rand.Read(bytes)
	if err != nil {
		log.Error("Error while generating salt: " + err.Error())
	}

	return bytes
}

func applySaltToPass(password string, salt []byte) []byte {
	return append([]byte(password), salt[:]...)
}

func encodeHashToString(hash []byte) string {
	return hex.EncodeToString(hash)
}

func decodedHashFromString(hash string) ([]byte, error) {
	if decoded, err := hex.DecodeString(hash); err != nil {
		log.Error("Error while decoding password: " + err.Error())
		return nil, err
	} else {
		return decoded, nil
	}
}
