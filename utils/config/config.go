package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

var v = viper.New()

func SetOptionalsDefault(grantType, scope string, expiration int) (*Config, error) {
	viper.Set("grant_type", grantType)
	viper.Set("scope", scope)
	viper.Set("expiration", expiration)

	return &Config{viper.GetViper()}, nil
}

func WriteCredentials(credendialsFile, clientId, clientSecret, accountId, accountEnvId string) (*Config, error) {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	if _, err := os.Stat(homeDir + "/.flagship"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(homeDir+"/.flagship", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	filepath, _ := filepath.Abs(homeDir + "/.flagship/" + credendialsFile)
	v.SetConfigFile(filepath)
	v.Set("client_id", clientId)
	v.Set("client_secret", clientSecret)
	v.Set("account_id", accountId)
	v.Set("account_environment_id", accountEnvId)
	err = v.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}

	return &Config{v}, nil

}

func WriteOptionals(credendialsFile, grantType, scope string, expiration int) (*Config, error) {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	if _, err := os.Stat(homeDir + "/.flagship"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(homeDir+"/.flagship", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	filepath, _ := filepath.Abs(homeDir + "/.flagship/" + credendialsFile)
	v.SetConfigFile(filepath)
	v.Set("grant_type", grantType)
	v.Set("scope", scope)
	v.Set("expiration", expiration)
	err = v.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}

	return &Config{v}, nil
}

func InitLocalConfigureConfig(credentialsFile string) *Config {

	if credentialsFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(credentialsFile)
	}

	if err := v.MergeInConfig(); err != nil {
		return &Config{v}
	}

	return &Config{v}
}

func WriteToken(credendialsFile, token string) (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}
	cobra.CheckErr(err)
	filepath, err := filepath.Abs(homeDir + "/.flagship/" + credendialsFile)
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}
	viper.SetConfigFile(filepath)
	viper.Set("token", token)
	err = viper.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}

	return &Config{viper.GetViper()}, err
}

func Binder(configureCmd *cobra.Command) {
	v.BindPFlag("client-id", configureCmd.PersistentFlags().Lookup("client_id"))
	v.BindPFlag("client-secret", configureCmd.PersistentFlags().Lookup("client_secret"))
	v.BindPFlag("account-id", configureCmd.PersistentFlags().Lookup("account_id"))
	v.BindPFlag("account-environment-id", configureCmd.PersistentFlags().Lookup("account_environment_id"))
}
