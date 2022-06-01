package provider

import (
	"context"
	"fmt"
	"github.com/google/go-github/v33/github"
	"github.com/package-url/packageurl-go"
	fErrors "github.com/packagrio/fetchr/pkg/errors"
	"github.com/packagrio/fetchr/pkg/models"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
)

type GithubProvider struct {
	*BaseProvider
	Name   string
	Client *github.Client
	Config *GithubProviderConfig
}

type GithubProviderConfig struct {
	Host     string `mapstructure:"hostname"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	Token    string `mapstructure:"token"`
	UseSSL   bool   `mapstructure:"use_ssl"`
	Username string `mapstructure:"username"`

	httpClient *http.Client
}

var (
	// DefaultGithubProviderConfig can be used if you don't want to think about options.
	DefaultGithubProviderConfig = &GithubProviderConfig{
		Host:   "",
		UseSSL: true,
	}
)

func DefaultGithubProvider() (*GithubProvider, error) {

	return NewGithubProvider("default", DefaultGithubProviderConfig)
}

func NewGithubProvider(providerName string, cfg *GithubProviderConfig) (*GithubProvider, error) {
	p := GithubProvider{
		Name:   providerName,
		Config: cfg,
	}

	//httpClient will only be set during testing.
	if p.Config.httpClient == nil && p.IsAuthenticatedProvider() {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: p.Config.Token},
		)
		p.Config.httpClient = oauth2.NewClient(ctx, ts)
	}

	var ghClient *github.Client
	if len(p.Config.Host) > 0 {
		uploadUrl := url.URL{Host: p.Config.Host, Path: "/api/uploads/", Scheme: "https"}
		baseUrl := url.URL{Host: p.Config.Host, Path: "/api/v3/", Scheme: "https"}
		var err error

		ghClient, err = github.NewEnterpriseClient(baseUrl.String(), uploadUrl.String(), p.Config.httpClient)
		if err != nil {
			return nil, err
		}
	} else {
		ghClient = github.NewClient(p.Config.httpClient)
	}

	p.Client = ghClient

	return &p, p.ValidateAuth()
}

func (p *GithubProvider) Close() {}

func (p *GithubProvider) ValidateAuth() error {
	if p.IsAuthenticatedProvider() {
		_, _, err := p.Client.Users.Get(context.Background(), "")
		return err
	}
	return nil
}

func (p *GithubProvider) IsAuthenticatedProvider() bool {
	return len(p.Config.Token) > 0 || len(p.Config.Username) > 0 && len(p.Config.Password) > 0
}

func (p *GithubProvider) IsSupportedArtifactType(actionType string, artifactType string) bool {
	if actionType == ActionDownload {
		return artifactType == packageurl.TypeGithub

	} else if actionType == ActionSearch {
		return artifactType == packageurl.TypeGithub

	} else {
		return false
	}
}

func (p *GithubProvider) ArtifactSearch(artifactPurl *packageurl.PackageURL) ([]*models.QueryResult, error) {
	if !p.IsSupportedArtifactType(ActionSearch, artifactPurl.Type) {
		return nil, fErrors.ProviderUnsupportedActionError(fmt.Sprintf("%s %s", ActionSearch, artifactPurl.Type))
	}

	releaseInfo, _, err := p.Client.Repositories.GetReleaseByTag(context.Background(), artifactPurl.Namespace, artifactPurl.Name, artifactPurl.Version)
	if err != nil {
		return nil, err
	}

	assetName, assetNameFilterExists := artifactPurl.Qualifiers.Map()["release_asset"]

	results := []*models.QueryResult{}
	for _, releaseAsset := range releaseInfo.Assets {
		if assetNameFilterExists {
			if releaseAsset.GetName() == assetName {
				results = append(results, &models.QueryResult{
					ArtifactPurl: artifactPurl.String(),
					Size:         int64(releaseAsset.GetSize()),
				})
			}
		} else {
			//no asset name provided, so just list all assets
			//newReleaseAssetPurl := packageurl.PackageURL{}
			//err = copier.CopyWithOption(&newReleaseAssetPurl, &artifactPurl, copier.Option{IgnoreEmpty: true, DeepCopy: true})
			//if err != nil{
			//	return nil, err
			//}
			qualifiersMap := artifactPurl.Qualifiers.Map()
			qualifiersMap["release_asset"] = releaseAsset.GetName()

			newReleaseAssetPurl := packageurl.NewPackageURL(
				artifactPurl.Type,
				artifactPurl.Namespace,
				artifactPurl.Name,
				artifactPurl.Version,
				packageurl.QualifiersFromMap(qualifiersMap),
				artifactPurl.Subpath,
			)

			results = append(results, &models.QueryResult{
				ArtifactPurl: newReleaseAssetPurl.String(),
				Size:         int64(releaseAsset.GetSize()),
			})
		}
	}

	return results, nil
}

func (p *GithubProvider) ArtifactDownload(downloadFolderPath string, artifactPurl *packageurl.PackageURL) ([]string, []*packageurl.PackageURL, error) {
	if !p.IsSupportedArtifactType(ActionDownload, artifactPurl.Type) {
		return nil, nil, fErrors.ProviderUnsupportedActionError(fmt.Sprintf("%s %s", ActionSearch, artifactPurl.Type))
	}
	return nil, nil, nil
}

func (p *GithubProvider) ArtifactUpload(artifactCachePath string, sourceArtifactPurl *packageurl.PackageURL, destArtifactPurl *packageurl.PackageURL) error {
	if !p.IsSupportedArtifactType(ActionUpload, destArtifactPurl.Type) {
		return fErrors.ProviderUnsupportedActionError(fmt.Sprintf("%s %s", ActionSearch, destArtifactPurl.Type))
	}
	return nil
}

func (p *GithubProvider) ArtifactSetMetadata(metadata map[string]string, destArtifactPurl *packageurl.PackageURL) error {
	if !p.IsSupportedArtifactType(ActionSetMetadata, destArtifactPurl.Type) {
		return fErrors.ProviderUnsupportedActionError(fmt.Sprintf("%s %s", ActionSetMetadata, destArtifactPurl.Type))
	}
	return nil
}
