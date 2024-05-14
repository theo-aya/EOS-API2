package config

import (
	"fmt"
	"github.com/spf13/viper"
)




type APIConfig struct {
	Eos string	
}

type Url struct {
	// PostUrl string
	// GetUrl string
	CreateFieldUrl string
	// Veg  string
	// Pro string
}



func LoadKey() (*APIConfig, error){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading configuration file: %w", err)
	}

	// Get values using Viper
	eos := viper.GetString("apiKeys.eos")
	return &APIConfig{
		Eos : eos,
	}, nil

}

func LoadUrl() (*Url, error){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading configuration file: %w", err)
	}

	// Get values using Viper
	// postUrl := viper.GetString("zoneVegUrls.postUrl")
	// getUrl := viper.GetString("zoneVegUrls.gettUrl")
	createFieldUrl := viper.GetString("fieldUrls.postcreateField")
	return &Url{
		CreateFieldUrl: createFieldUrl,
	}, nil

}