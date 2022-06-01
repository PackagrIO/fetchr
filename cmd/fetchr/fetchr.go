package main

import (
	"fmt"
	"github.com/analogj/go-util/utils"
	"github.com/packagrio/fetchr/pkg/actions/query"
	"github.com/packagrio/fetchr/pkg/report"
	"github.com/packagrio/fetchr/pkg/version"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

var goos string
var goarch string

func main() {
	app := &cli.App{
		Name:     "fetchr",
		Usage:    "Universal tool to retrieve packages/artifacts using PackageUrls (PURL)",
		Version:  version.VERSION,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Jason Kulatunga",
				Email: "jason@thesparktree.com",
			},
		},
		Before: func(c *cli.Context) error {

			packagrUrl := "github.com/packagrio/fetchr"

			versionInfo := fmt.Sprintf("%s.%s-%s", goos, goarch, version.VERSION)

			subtitle := packagrUrl + utils.LeftPad2Len(versionInfo, " ", 53-len(packagrUrl))

			fmt.Fprintf(c.App.Writer, fmt.Sprintf(utils.StripIndent(
				`
			 ____   __    ___  __ _   __    ___  ____ 
			(  _ \ / _\  / __)(  / ) / _\  / __)(  _ \
			 ) __//    \( (__  )  ( /    \( (_ \ )   /
			(__)  \_/\_/ \___)(__\_)\_/\_/ \___/(__\_)
			%s

			`), subtitle))
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:  "query",
				Usage: "universal artifact search using Purl artifact ids",
				Action: func(c *cli.Context) error {

					pipeline := query.Pipeline{
						Config:          query.NewConfiguration(),
						ArtifactPurlStr: c.Args().Get(0),
						Logger:          logrus.New(),
					}
					queryResults, err := pipeline.Start()
					report.QueryReportMarkdown(queryResults)
					return err
				},

				Flags: []cli.Flag{},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}
