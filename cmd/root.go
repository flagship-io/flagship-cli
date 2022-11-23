/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package cmd

import (
	"os"

	"github.com/flagship-io/flagship/cmd/analyse"
	"github.com/flagship-io/flagship/cmd/authorization"
	"github.com/flagship-io/flagship/cmd/campaign"
	"github.com/flagship-io/flagship/cmd/flag"
	"github.com/flagship-io/flagship/cmd/panic"
	"github.com/flagship-io/flagship/cmd/project"
	targetingkey "github.com/flagship-io/flagship/cmd/targeting_key"
	"github.com/flagship-io/flagship/cmd/token"
	"github.com/flagship-io/flagship/cmd/user"
	"github.com/flagship-io/flagship/cmd/variation"
	"github.com/flagship-io/flagship/cmd/variation_group"
	"github.com/flagship-io/flagship/cmd/version"
	"github.com/flagship-io/flagship/utils/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	cmdToken     string
	outputFormat string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flagship",
	Short: "flagship manage your campaigns, project, users etc...",
	Long: `flagship is the main command, used to manage campaigns, projects, users, variation groups and variations
	
	Flagship is a feature flagging platform for modern developers. 
	Separate code deployments from feature releases to accelerate development cycles and mitigate risks.
	
	Complete documentation is available at https://docs.developers.flagship.io/docs/flagship-command-line-interface`,
	Run: func(cmd *cobra.Command, args []string) {
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
	rootCmd.AddCommand(flag.FlagCmd)
	rootCmd.AddCommand(targetingkey.TargetingKeyCmd)
	rootCmd.AddCommand(version.VersionCmd)
	rootCmd.AddCommand(token.TokenCmd)
	rootCmd.AddCommand(analyse.AnalyseCmd)
}
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cmdToken, "token", "t", "", "access token to manage flagship resources")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output-format", "f", config.OutputFormat, "output format for the get and list subcommands for flagship resources. Only 3 format are possible: table, json, json-pretty")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("output_format", rootCmd.PersistentFlags().Lookup("output-format"))

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file that contains your credentials (default is $HOME/.flagship/credentials.yaml)")

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
