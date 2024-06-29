package core

import (
	"crypto/rand"
)

// GenerateSimplePassword creates a password with the given length.
// TODO: use hashicorp vault with policies to generate a password
// look at this: https://github.com/1Password/spg
func GenerateSimplePassword(length int) (string, error) {
	// Definieren Sie die Zeichensätze für das Passwort.
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/"

	// Erstellen Sie einen Slice, um das Passwort zu speichern.
	b := make([]byte, length)

	// Generieren Sie jedes Zeichen des Passworts.
	for i := range b {
		if _, err := rand.Read(b[i : i+1]); err != nil {
			return "", err
		}
		b[i] = charset[b[i]%byte(len(charset))]
	}

	return string(b), nil
}
