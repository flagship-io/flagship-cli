/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package flag

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [-i <flag-id> | --id=<flag-id>] [-d <data-raw> | --data-raw <data-raw>]",
	Short: "Edit a flag",
	Long:  `Edit a flag in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPEditFlag(FlagID, DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Printf("flag updated: %s", body)
	},
}

func init() {

	editCmd.Flags().StringVarP(&FlagID, "id", "i", "", "id of the flag you want to edit")
	editCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to edit your flag, check the doc for details")

	if err := editCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	if err := editCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	FlagCmd.AddCommand(editCmd)
}