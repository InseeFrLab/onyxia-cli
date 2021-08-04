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
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"onyxiactl/utils"
)

var catalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "List the available catalogs.",
	Long:  `List the available catalogs.`,
	Run: func(cmd *cobra.Command, args []string) {
		catalog := &CatalogResponse{}
		token, _ := cmd.Flags().GetString("token")
		onyxiaURL, _ := cmd.Flags().GetString("onyxiaURL")

		JsonCatalog, e := utils.CallAPIGet(onyxiaURL+"/public/catalog", token)
		if e != nil {
			panic(e)
		}

		e = json.Unmarshal(JsonCatalog, &catalog)
		if e != nil {
			panic(e)
		}

		for _, value := range catalog.Catalogs {
			j, _ := json.Marshal(value)
			fmt.Println(string(j))
		}
	},
}

func init() {
	rootCmd.AddCommand(catalogCmd)
}

type CatalogResponse struct {
	Catalogs []struct {
		/*		Catalog struct {
					Packages []struct {
						APIVersion  string   `json:"apiVersion"`
						AppVersion  string   `json:"appVersion"`
						Config      struct{} `json:"config"`
						Created     string   `json:"created"`
						Description string   `json:"description"`
						Digest      string   `json:"digest"`
						Home        string   `json:"home"`
						Icon        string   `json:"icon"`
						Name        string   `json:"name"`
						Sources     []string `json:"sources"`
						Urls        []string `json:"urls"`
						Version     string   `json:"version"`
					} `json:"packages"`
				} `json:"catalog"`
		*/
		ID          string `json:"id"`
		Description string `json:"description"`
		//		LastUpdateTime uint   `json:"lastUpdateTime"`
		Location   string `json:"location"`
		Maintainer string `json:"maintainer"`
		Name       string `json:"name"`
		Status     string `json:"status"`
		Type       string `json:"type"`
	} `json:"catalogs"`
}
