package cmd

import (
	"log"
	"net/url"

	"github.com/ioannisGiak89/fsa-authorities/app/lib/factory"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all the authorities",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		baseUrl, err := url.Parse("http://api.ratings.food.gov.uk/")

		if err != nil {
			log.Fatal(err)
		}

		appFactory := factory.NewAppFactory()
		fsaClient := appFactory.BuildFsaClient(baseUrl)
		fsaService := appFactory.BuildFsaService(fsaClient)

		response, err := fsaService.GetAuthorities()

		if err != nil {
			log.Fatal(err)
		}

		var rows []table.Row
		for index, authority := range response.Authorities {
			rows = append(rows, table.Row{index, authority.Name, authority.SchemeType.String(), authority.LocalAuthorityId})
		}

		table := appFactory.BuildTable(rows, table.Row{"#", "Authority Name", "Scheme Type", "Local Authority ID"}, nil)
		table.Render()
	},
}
