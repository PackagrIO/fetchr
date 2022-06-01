package provider

import (
	"github.com/analogj/go-util/utils"
	"github.com/package-url/packageurl-go"
	fErrors "github.com/packagrio/fetchr/pkg/errors"
	"github.com/packagrio/fetchr/pkg/models"
	"log"
	"os"
	"path/filepath"
)

type LocalProvider struct {
	*BaseProvider
	Name string
}

func DefaultLocalProvider() (*LocalProvider, error) {
	return NewLocalProvider("default")
}

func NewLocalProvider(providerName string) (*LocalProvider, error) {
	p := LocalProvider{
		Name: providerName,
	}

	//this must be "default"
	p.Name = providerName

	return &p, p.ValidateAuth()
}

func (p *LocalProvider) Close() {}

func (p *LocalProvider) ValidateAuth() error {
	return nil
}

func (p *LocalProvider) IsAuthenticatedProvider() bool {
	return false
}

func (p *LocalProvider) IsSupportedArtifactType(actionType string, artifactType string) bool {
	if actionType == ActionSetMetadata {
		return false
	}
	return artifactType == "local"
}

func (p *LocalProvider) ArtifactSearch(artifactPurl *packageurl.PackageURL) ([]*models.QueryResult, error) {
	if !p.IsSupportedArtifactType(ActionSearch, artifactPurl.Type) {
		return nil, fErrors.NewProviderUnsupportedActionError(ActionSearch, artifactPurl.Type)
	}

	artifactPath := artifactPurl.Subpath

	//check if destination is a path or a file.
	if artifactPath[len(artifactPath)-1:] == "/" {
		log.Printf("source is a directory, listing recursively ")
		queryResults := []*models.QueryResult{}
		err := filepath.Walk(artifactPath,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					newArtifactPurl := NewLocalPackageUrl(path)

					//todo calculate checksums and store as qualifiers or in the queryResults
					queryResults = append(queryResults, &models.QueryResult{
						ArtifactPurl: newArtifactPurl.String(),
						Size:         info.Size(),
					})
				}
				return nil
			})
		if err != nil {
			return nil, err
		}
		return queryResults, nil
	} else {
		log.Print("source is a file, searching")
		if utils.FileExists(artifactPath) {

			return []*models.QueryResult{{ArtifactPurl: artifactPurl.String()}}, nil

		} else {
			return nil, fErrors.NewFileValidationError("file does not exist at path", artifactPath)
		}
	}
}

func (p *LocalProvider) ArtifactDownload(downloadFolderPath string, artifactPurl *packageurl.PackageURL) ([]string, []*packageurl.PackageURL, error) {
	if !p.IsSupportedArtifactType(ActionDownload, artifactPurl.Type) {
		return nil, nil, fErrors.NewProviderUnsupportedActionError(ActionSearch, artifactPurl.Type)
	}
	return nil, nil, nil
}

func (p *LocalProvider) ArtifactUpload(artifactCachePath string, sourceArtifactPurl *packageurl.PackageURL, destArtifactPurl *packageurl.PackageURL) error {
	if !p.IsSupportedArtifactType(ActionUpload, destArtifactPurl.Type) {
		return fErrors.NewProviderUnsupportedActionError(ActionSearch, destArtifactPurl.Type)
	}
	return nil
}

func (p *LocalProvider) ArtifactSetMetadata(metadata map[string]string, destArtifactPurl *packageurl.PackageURL) error {
	if !p.IsSupportedArtifactType(ActionSetMetadata, destArtifactPurl.Type) {
		return fErrors.NewProviderUnsupportedActionError(ActionSetMetadata, destArtifactPurl.Type)
	}
	return nil
}

func NewLocalPackageUrl(localPath string) *packageurl.PackageURL {
	return packageurl.NewPackageURL(
		"local",
		"",
		"",
		"",
		nil,
		localPath,
	)
}
