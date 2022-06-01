package query

import (
	"github.com/package-url/packageurl-go"
	"github.com/packagrio/fetchr/pkg/models"
	"github.com/packagrio/fetchr/pkg/provider"
	"github.com/sirupsen/logrus"
	"strings"
)

type Pipeline struct {
	Config          *Configuration
	ArtifactPurlStr string
	Logger          *logrus.Logger
}

func (p *Pipeline) Start() ([]*models.QueryResult, error) {
	p.Logger.Infof("Starting Query")

	var artifactPurl *packageurl.PackageURL
	var err error
	if strings.HasPrefix(p.ArtifactPurlStr, "pkg:") {
		parsedArtifactPurl, err := packageurl.FromString(p.ArtifactPurlStr)
		if err != nil {
			return nil, err
		}
		artifactPurl = &parsedArtifactPurl
	} else {
		//this is a local artifact
		artifactPurl = provider.NewLocalPackageUrl(p.ArtifactPurlStr)
	}

	//var sourceProvider provider.Interface
	sourceProvider, err := provider.CreateDefault(artifactPurl.Type)
	if err != nil {
		return nil, err
	}

	foundArtifacts, err := sourceProvider.ArtifactSearch(artifactPurl)
	if err != nil {
		return nil, err
	}

	return foundArtifacts, err
}
