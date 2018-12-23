package excp

import "github.com/Kamva/nautilus/http"

// IException is an interface for Exceptions
type IException interface {
	// GetCode returns exception code
	GetCode() string

	// GetStatus returns http status code
	GetStatus() http.StatusCode

	// GetMessage returns exception message
	GetMessage() string

	// GetTransKey return the key for translation
	GetTransKey() string
}
