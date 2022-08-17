/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version will match the tag
var Version = "v0.2.0"

func DisplayVersion() {
	fmt.Printf("Flagship CLI verison : %s\n", Version)
}

// VersionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get Flagship CLI version",
	Long:  `Get Flagship CLI version`,
	Run: func(cmd *cobra.Command, args []string) {
		DisplayVersion()
	},
}
