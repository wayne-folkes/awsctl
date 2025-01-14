package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/omerh/awsctl/aws/helper"
	"gitlab.com/omerh/awsctl/aws/outputs"
)

var listRegionsCmd = &cobra.Command{
	Use:   "regions",
	Short: "List all regions",
	Run: func(cmd *cobra.Command, Args []string) {
		regions, _ := helper.GetAllAwsRegions()
		out, _ := cmd.Flags().GetString("out")

		switch out {
		case "json":
			outputs.PrintGenericJSONOutput(regions, "")
		default:
			log.Println("Listing all regions...")
			for _, region := range regions {
				fmt.Println(region)
			}
		}
	},
}
