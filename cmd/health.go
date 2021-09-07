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
	"github.com/spf13/cobra"
	"onyxiactl/utils"
)

func init() {
	rootCmd.AddCommand(healthCmd)
}

var healthCmd = &cobra.Command{
	Use:     "health",
	Aliases: []string{"ping"},
	Short:   "Ping the provider.",
	Long:    `Call the healthcheck URL.`,
	Run:     func(cmd *cobra.Command, args []string) { health(cmd) },
}

func health(cmd *cobra.Command) error {
	onyxiaURL, _ := cmd.Flags().GetString("onyxiaURL")

	_, e := utils.CallAPIGet(onyxiaURL+"/public/healthcheck", "")
	if e != nil {
		utils.PrintErrorAndExit("CRIT: provider is NOT healthy, " + e.Error() + ".")
	}

	utils.PrintMessage("OK: provider is healthy.")
	return nil
}
