package utils

import (
	"context"
	"net/http"
	"time"

	oauth2ns "onyxiactl/utils/oauth2"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// GenerateOauthConfigFromParams ...
func GenerateOauthConfigFromParams(clientID string, realm string, keycloakURL string) *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: "",
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  keycloakURL + "/auth/realms/" + realm + "/protocol/openid-connect/auth",
			TokenURL: keycloakURL + "/auth/realms/" + realm + "/protocol/openid-connect/token",
		}}
	return conf
}

// GetLastToken ...
func GetLastToken() *oauth2.Token {
	token := new(oauth2.Token)
	token.AccessToken = viper.GetString("auth.accessToken")
	token.RefreshToken = viper.GetString("auth.refreshToken")
	token.Expiry = viper.GetTime("auth.expire")
	token.TokenType = viper.GetString("tokenType")
	return token
}

// GetAuthClient ...
func GetAuthClient() *http.Client {
	conf := GenerateOauthConfigFromParams(viper.GetString("auth.conf.ClientID"), viper.GetString("auth.conf.realm"), viper.GetString("auth.conf.keycloakURL"))
	ctx := context.Background()
	token := GetLastToken()
	if token.Expiry.Before(time.Now()) {
		client, _ := oauth2ns.AuthenticateUser(conf)
		SaveToken(client.Token)
		token = client.Token
	}
	return conf.Client(ctx, token)
}

//SaveToken ...
func SaveToken(token *oauth2.Token) {
	viper.Set("auth.accessToken", token.AccessToken)
	viper.Set("auth.expire", token.Expiry)
	viper.Set("auth.refreshToken", token.RefreshToken)
	viper.Set("auth.tokenType", token.TokenType)
	viper.WriteConfig()
}
