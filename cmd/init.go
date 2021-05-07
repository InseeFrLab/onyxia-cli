/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	_init "onyxiactl/cli/init"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config for onyxiactl",
	Long:  `Initialize config for onyxiactl`,
	Run: func(cmd *cobra.Command, args []string) {
		clientID, _ := cmd.Flags().GetString("clientID")
		realm, _ := cmd.Flags().GetString("realm")
		keycloakURL, _ := cmd.Flags().GetString("keycloakURL")
		updateFile, _ := cmd.Flags().GetBool("updateFile")
		onboardingURL, _ := cmd.Flags().GetString("onboardingURL")
		_init.Execute(clientID, realm, keycloakURL, onboardingURL, updateFile)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("keycloakURL", "k", "https://keycloak.dev.insee.io", "eg https://keycloak.example.com")
	initCmd.Flags().StringP("realm", "r", "dev", "realm to use")
	initCmd.Flags().StringP("clientID", "c", "onboarding", "clientID")
	initCmd.Flags().StringP("onboardingURL", "", "https://dev.insee.io", "eg https://onboarding.example.com")
	initCmd.Flags().BoolP("updateFile", "", false, "force the update")
}
