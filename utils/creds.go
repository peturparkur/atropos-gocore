package utils

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

// NoCredFoundError represents an error when no credentials are found
type NoCredFoundError struct {
	CredentialName string
}

func (e *NoCredFoundError) Error() string {
	return fmt.Sprintf("no credentials found for %s", e.CredentialName)
}

func getCredUnsafe(l *zap.SugaredLogger, value string) string {
	cred := os.Getenv(value)
	if cred == "" {
		l.Fatal(value + " is not set")
	}

	l.Info("Found " + value)
	return cred
}

func getCred(l *zap.SugaredLogger, value string) (string, error) {
	cred := os.Getenv(value)
	if cred == "" {
		noCredError := &NoCredFoundError{
			CredentialName: value,
		}
		return "", noCredError
	}
	l.Info("Found " + value)
	return cred, nil
}
