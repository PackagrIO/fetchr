package provider

import (
	"github.com/package-url/packageurl-go"
	"github.com/packagrio/fetchr/pkg/models"
	"github.com/spf13/viper"
	"io"
)

const (
	ActionDownload    = "ActionDownload"
	ActionSearch      = "ActionSearch"
	ActionSetMetadata = "ActionSetMetadata"
	ActionUpload      = "ActionUpload"
)

// Additional provider types should be added, matching those from this list:
// https://github.com/package-url/purl-spec/blob/master/PURL-TYPES.rst
const (
	ProviderTypeDocker = "docker"
	ProviderTypeGithub = "github"

	// special provider type for storing files locally.
	ProviderTypeLocal = "local"
)

// Create mock using:
// mockgen -source=pkg/provider/interface.go -destination=pkg/provider/mock/mock_config.go
type ConfigInterface interface {
	Get(key string) interface{}
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetUint(key string) uint
	GetUint32(key string) uint32
	GetUint64(key string) uint64
	GetFloat64(key string) float64
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetSizeInBytes(key string) uint
	UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error
	Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error
	Set(key string, value interface{})
	SetDefault(key string, value interface{})
	MergeConfig(in io.Reader) error
	AllKeys() []string
}

type Interface interface {
	ValidateAuth() error
	Close()

	IsAuthenticatedProvider() bool
	IsSupportedArtifactType(actionType string, artifactType string) bool

	//given a single purl, may return multiple purl's (one for each sub artifact)
	ArtifactSearch(artifactPurl *packageurl.PackageURL) ([]*models.QueryResult, error)

	ArtifactUpload(cachePath string, sourceArtifactPurl *packageurl.PackageURL, destArtifactPurl *packageurl.PackageURL) error
	ArtifactDownload(basePath string, artifactPurl *packageurl.PackageURL) ([]string, []*packageurl.PackageURL, error)

	ArtifactSetMetadata(metadata map[string]string, destArtifactPurl *packageurl.PackageURL) error
}
