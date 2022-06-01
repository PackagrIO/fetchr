package query

import (
	"github.com/package-url/packageurl-go"
	"github.com/packagrio/fetchr/pkg/models"
	"github.com/packagrio/fetchr/pkg/provider"
	"github.com/sirupsen/logrus"
)

type Pipeline struct {
	Config          *Configuration
	ArtifactPurlStr string
	Logger          *logrus.Logger
}

func (p *Pipeline) Start() ([]*models.QueryResult, error) {
	p.Logger.Infof("Starting Query")
	artifactPurl, err := packageurl.FromString(p.ArtifactPurlStr)
	if err != nil {
		return nil, err
	}

	//var sourceProvider provider.ProviderInterface
	sourceProvider, err := provider.DefaultGithubProvider()
	if err != nil {
		return nil, err
	}

	foundArtifacts, err := sourceProvider.ArtifactSearch(&artifactPurl)
	if err != nil {
		return nil, err
	}

	return foundArtifacts, err
}
