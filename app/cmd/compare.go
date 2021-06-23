package cmd

import (
	"errors"
	"log"
	"net/url"
	"strings"

	"github.com/ioannisGiak89/compare-fsa-ratings/app/lib/factory"
	"github.com/ioannisGiak89/compare-fsa-ratings/app/model"
	"github.com/spf13/cobra"
)

var schemeType string
var authorityIds []string

func init() {
	compareCmd.Flags().StringVarP(
		&schemeType,
		"schemeType",
		"s",
		"",
		"The scheme type the authorities belong to",
	)
	err := compareCmd.MarkFlagRequired("schemeType")

	if err != nil {
		log.Fatal(err)
	}

	compareCmd.Flags().StringSliceVarP(
		&authorityIds,
		"authorityIds",
		"a",
		[]string{},
		"A comma separated list with the local authority IDs to compare",
	)
	newErr := compareCmd.MarkFlagRequired("authorityIds")

	if newErr != nil {
		log.Fatal(err)
	}
}

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare the food hygiene rating distribution for two or more authorities",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		baseUrl, err := url.Parse("http://api.ratings.food.gov.uk/")

		if err != nil {
			log.Fatal(err)
		}

		appFactory := factory.NewAppFactory()
		fsaClient := appFactory.BuildFsaClient(baseUrl)
		fsaService := appFactory.BuildFsaService(&fsaClient)

		var fsaSchemeRatingDistributions []model.FsaSchemeRatingDistribution
		for _, authorityID := range authorityIds {
			authority, err := fsaService.GetAuthorityByID(authorityID)
			if err != nil {
				log.Fatal(err)
			}

			if authority.SchemeType.String() != strings.ToUpper(schemeType) {
				log.Fatal(errors.New("authority with ID does not belong to scheme: " + schemeType))
			}

			e, err := fsaService.GetEstablishments(authorityID)
			if err != nil {
				log.Fatal(err)
			}

			if authority.SchemeType.String() == "FHRS" {
				fsaSchemeRatingDistributions = append(fsaSchemeRatingDistributions, model.NewFhrsSchemeRatingDistribution(authority, e.Establishments))
			} else {
				fsaSchemeRatingDistributions = append(fsaSchemeRatingDistributions, model.NewFhisSchemeRatingDistribution(authority, e.Establishments))
			}
		}

		table := appFactory.BuildCompareTable(fsaSchemeRatingDistributions)
		table.CreateTableAndRender()
	},
}
