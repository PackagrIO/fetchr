package errors

import (
	"fmt"
	"log"
	"strings"
)

// Raised when config file is missing
type ConfigFileMissingError error

func NewConfigFileMissingError(a ...interface{}) ConfigFileMissingError {
	return ConfigFileMissingError(argsToError("ConfigFileMissingError", a...))
}

// Raised when the config file doesnt match schema
type ConfigValidationError error

func NewConfigValidationError(a ...interface{}) ConfigValidationError {
	return ConfigValidationError(argsToError("ConfigValidationError", a...))
}

// raised when the provider is unsupported/unknown
type ProviderUnknownError error

func NewProviderUnknownError(a ...interface{}) ProviderUnknownError {
	return ProviderUnknownError(argsToError("ProviderUnknownError", a...))
}

//raised when the provider doesnt support the specified action
type ProviderUnsupportedActionError error

func NewProviderUnsupportedActionError(a ...interface{}) ProviderUnsupportedActionError {
	return ProviderUnsupportedActionError(argsToError("ProviderUnsupportedActionError", a...))
}

type FileValidationError error

func NewFileValidationError(a ...interface{}) FileValidationError {
	return FileValidationError(argsToError("FileValidationError", a...))
}

// helpers
// https://stackoverflow.com/a/59095385
func argsToError(errorName string, args ...interface{}) error {
	ph := make([]string, len(args))
	for i, v := range args {
		_, isErr := v.(error)
		if isErr {
			log.Println("ERROR DETECTED")
			ph[i] = "%w"
		} else {
			ph[i] = "%v"
		}
	}
	prefix := fmt.Sprintf("%s: ", errorName)
	log.Printf("error fmt string: %s", strings.Join(ph, ", "))

	return fmt.Errorf(prefix+strings.Join(ph, ", "), args...)
}
