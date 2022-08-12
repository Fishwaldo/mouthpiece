package mouthpiece

import (
	"github.com/spf13/viper"
)

type OAuthConfig struct {
	ClientID     string `json:"clientid" doc:"OAuth Client ID"`
}

type FEConfig struct {
	OAuthProviders map[string]OAuthConfig `json:"oauthproviders" doc:"Provider OAuth Config for Frontend"`	
}

func GetFEConfig() (config *FEConfig) {
	config = &FEConfig{}
	config.OAuthProviders = make(map[string]OAuthConfig)
	if (viper.GetBool("auth.github.enabled")) {
		config.OAuthProviders["github"] = OAuthConfig{
			ClientID: viper.GetString("auth.github.client_id"),
		}
	}
	if (viper.GetBool("auth.google.enabled")) {
		config.OAuthProviders["google"] = OAuthConfig{
			ClientID: viper.GetString("auth.google.client_id"),
		}
	}
	if (viper.GetBool("auth.dev.enabled")) {
		config.OAuthProviders["dev"] = OAuthConfig{
			ClientID: "123456",
		}
	}
	return config
}