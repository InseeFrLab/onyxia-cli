package auth

import (
	"encoding/json"
	"log"
	"onyxiactl/utils"

	oauth2ns "onyxiactl/utils/oauth2"

	"golang.org/x/oauth2"
)

//ExecuteGetToken ...
func ExecuteGetToken(clientID string, realm string, keycloakURL string) (*oauth2.Token, error) {
	conf := utils.GenerateOauthConfigFromParams(clientID, realm, keycloakURL)
	client, err := oauth2ns.AuthenticateUser(conf)
	if err != nil {
		log.Fatal(err)
	}
	return client.Token, err
}

//ExecuteGetUserInfo ...
func ExecuteGetUserInfo(clientID string, realm string, keycloakURL string) (*UserInfo, error) {
	conf := utils.GenerateOauthConfigFromParams(clientID, realm, keycloakURL)
	client, err := oauth2ns.AuthenticateUser(conf)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Get(keycloakURL + "/auth/realms/" + realm + "/protocol/openid-connect/userinfo")
	if err != nil {
		log.Fatal(err)
	}
	userInfo := new(UserInfo)
	err = json.NewDecoder(resp.Body).Decode(userInfo)
	return userInfo, err
}

//UserInfo ...
type UserInfo struct {
	Email             string
	EmailVerified     bool
	FamilyName        string
	GivenName         string
	Groups            []string
	Name              string
	PreferredUsername string
	Sub               string
}
