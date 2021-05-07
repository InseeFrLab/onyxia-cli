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
	"bytes"
	"encoding/json"
	"net/http"
	"onyxiactl/utils"

	"github.com/spf13/cobra"
)

var onboardCmd = &cobra.Command{
	Use:   "onboard",
	Short: "Welcome onboard :)",
	Long:  `Welcome onboard :)`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _  := cmd.Flags().GetString("token")
		onyxiaURL, _  := cmd.Flags().GetString("onyxiaURL")
		group, _ := cmd.Flags().GetString("group")

		id := utils.GetID(token)

		request := &OnboardingRequest{
			
		}

		if (group != "") {
			if (!contains(id.Groups,group)) {
				panic("User does not belong to group "+group)
			}
			println("Onboarding group ",group)
			request.Group = group
		}
	
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(request)
		req, _ := http.NewRequest("POST", onyxiaURL+"/onboarding", buf)
		req.Header.Add("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		res, e := client.Do(req)
		if e != nil {
			panic(e)
		}
		println("Welcome ! ",res.StatusCode)
	},
}

func init() {
	rootCmd.AddCommand(onboardCmd)
	onboardCmd.Flags().StringP("group", "g", "", "group to onboard")
}

type OnboardingRequest struct {
	Group     string     `json:"group"`
}


func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}