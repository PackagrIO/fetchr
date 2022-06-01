package report

import (
	"code.cloudfoundry.org/bytefmt"
	"github.com/olekukonko/tablewriter"
	"github.com/packagrio/fetchr/pkg/models"
	"os"
	"strconv"
)

func QueryReportMarkdown(results []*models.QueryResult) {

	data := [][]string{}
	for _, result := range results {
		data = append(data, []string{result.ArtifactPurl, bytefmt.ByteSize(uint64(result.Size)), result.Checksum})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Artifact Id", "Size", "Checksum"})
	table.SetFooter([]string{"", "Total", strconv.Itoa(len(results))})
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}
