/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"

	"github.com/flagship-io/flagship/cmd/authorization"
	"github.com/flagship-io/flagship/cmd/campaign"
	"github.com/flagship-io/flagship/cmd/panic"
	"github.com/flagship-io/flagship/cmd/project"
	"github.com/flagship-io/flagship/cmd/user"
	"github.com/flagship-io/flagship/cmd/variation"
	"github.com/flagship-io/flagship/cmd/variation_group"
	"github.com/flagship-io/flagship/cmd/version"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	token        string
	outputFormat string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flagship",
	Short: "flagship manage your campaigns, project, users etc...",
	Long: `flagship is the main command, used to manage campaigns, projects, users, variation groups and variations
	
	Flagship is a feature flagging platform for modern developers. 
	Separate code deployments from feature releases to accelerate development cycles and mitigate risks.
	
	Complete documentation is available at http://flagship.io`,
	Run: func(cmd *cobra.Command, args []string) {
		getVersion, err := cmd.Flags().GetBool("version")
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}

		if getVersion {
			version.DisplayVersion()
			return
		}
		cmd.Root().Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandPalettes() {
	rootCmd.AddCommand(campaign.CampaignCmd)
	rootCmd.AddCommand(project.ProjectCmd)
	rootCmd.AddCommand(authorization.ConfigureCmd)
	rootCmd.AddCommand(authorization.AuthenticateCmd)
	rootCmd.AddCommand(panic.PanicCmd)
	rootCmd.AddCommand(user.UserCmd)
	rootCmd.AddCommand(variation_group.VariationGroupCmd)
	rootCmd.AddCommand(variation.VariationCmd)
	rootCmd.AddCommand(version.VersionCmd)
}
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP("version", "v", false, "CLI version")
	
  rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "access token")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "f", "table", "output format")
	
  viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("output_format", rootCmd.PersistentFlags().Lookup("output-format"))

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.flagship/credentials.yaml)")

	addSubCommandPalettes()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		homeDir, err := os.UserHomeDir()
		cobra.CheckErr(err)
		// Search config in home directory with name ".flagship" (without extension).
		viper.SetConfigFile(homeDir + "/.flagship/credentials.yaml")
	}

	// read in environment variables that match
	// If a config file is found, read it in.
	viper.MergeInConfig()
}
