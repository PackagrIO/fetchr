package provider

import (
	"fmt"
	"github.com/packagrio/fetchr/pkg/errors"
)

func CreateDefault(providerType string) (Interface, error) {

	var prov Interface
	var err error

	switch providerType {
	//empty/generic package manager. Noop.
	case "github":
		prov, err = DefaultGithubProvider()

	case "local":
		prov, err = DefaultLocalProvider()

	default:
		return nil, errors.NewProviderUnknownError(fmt.Sprintf("Unknown Packager Manager Type: %s", providerType))
	}

	return prov, err
}
