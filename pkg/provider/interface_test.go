package provider_test

import (
	"github.com/packagrio/fetchr/pkg/provider"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGithubProvider(t *testing.T) {
	t.Parallel()
	prov := new(provider.GithubProvider)
	require.Implements(t, (*provider.Interface)(nil), prov, "should implement the Provider interface")
}

func TestLocalProvider(t *testing.T) {
	t.Parallel()
	prov := new(provider.LocalProvider)
	require.Implements(t, (*provider.Interface)(nil), prov, "should implement the Provider interface")
}
