// Package utils contains utility functions that are used throughout the application.
package utils

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
)

// GetCredUnsafe is a function that gets a credential from the environment variables. If the credential is not found, it will log a fatal error.
func GetCredUnsafe(l *zap.SugaredLogger, value string) string {
	cred := os.Getenv(value)
	if cred == "" {
		l.Fatal(value + " is not set")
	}

	l.Info("Found " + value)
	return cred
}

// GetCred is a function that gets a credential from the environment variables. If the credential is not found, it will return an error.
func GetCred(value string) (string, error) {
	l := slog.Default().With("cred", value)
	cred := os.Getenv(value)
	if cred == "" {
		noCredError := &NoCredFoundError{
			CredentialName: value,
		}
		return "", noCredError
	}
	l.Info("Credential found")
	return cred, nil
}
