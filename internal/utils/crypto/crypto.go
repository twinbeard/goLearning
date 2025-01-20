package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(key string) string {
	hash := sha256.New()                 // Create a new SHA256 hash
	hash.Write([]byte(key))              // Write the key to the hash
	hashBytes := hash.Sum(nil)           // Get the hash sum
	return hex.EncodeToString(hashBytes) // Return the hash as a hex string
}

// GenerateSalt a random salt
func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)

	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	return hex.EncodeToString(salt), nil
}

// PasswordHash = Hash(password + salt)
func HashPassword(password string, salt string) string {
	// concatenate password and salt
	saltedPassword := password + salt

	// hash the combined string
	hashedPassword := sha256.Sum256([]byte(saltedPassword))

	return hex.EncodeToString(hashedPassword[:])
}

func MatchingPassword(storeHash string, password string, salt string) bool {
	hashPassword := HashPassword(password, salt)
	return storeHash == hashPassword
}
