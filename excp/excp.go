package excp

import "github.com/kamva/nautilus/http"

// Exception is an exception for fatal error happened in libraries
type Exception struct {
	Status   http.StatusCode
	Message  string
	TransKey string
}

// GetCode returns exception code
func (e Exception) GetCode() string {
	return "LIB_ERR"
}

// GetStatus returns http status code
func (e Exception) GetStatus() http.StatusCode {
	return e.Status
}

// GetMessage returns exception message
func (e Exception) GetMessage() string {
	return e.Message
}

// GetTransKey return the key for translation
func (e Exception) GetTransKey() string {
	return e.TransKey
}
