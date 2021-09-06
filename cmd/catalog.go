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

	"github.com/spf13/cobra"
	"onyxiactl/utils"
)

type Catalog struct {
	ID             string    `json:"id"`
	Description    string    `json:"description"`
	LastUpdateTime uint      `json:"lastUpdateTime"`
	Location       string    `json:"location"`
	Maintainer     string    `json:"maintainer"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	Type           string    `json:"type"`
}

type Catalogs struct {
	Catalogs []Catalog `json:"catalogs"`
}

func init() {
	rootCmd.AddCommand(catalogCmd)
	catalogCmd.AddCommand(catalogListCmd)
	catalogCmd.PersistentFlags().StringP("id", "i", "", "Catalog’s ID.")
}

var catalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "Manage the catalogs.",
	Long:  `Manage the catalogs. By default the command lists the available catalogs.`,
	Run:   func(cmd *cobra.Command, args []string) { catalogList(cmd) },
}

var catalogListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the available catalogs (default).",
	Long:  `List the available catalogs.`,
	Run:   func(cmd *cobra.Command, args []string) { catalogList(cmd) },
}

func catalogList(cmd *cobra.Command) error {
	token, _ := cmd.Flags().GetString("token")
	onyxiaURL, _ := cmd.Flags().GetString("onyxiaURL")
	id, _ := cmd.Flags().GetString("id")

	if len(id) == 0 {
		return catalogListAll(onyxiaURL+"/public/catalog", token)
	}

	return catalogListId(onyxiaURL+"/public/catalog", token, id)
}

func catalogListAll(url, token string) error {
	catalog := &Catalogs{}
	JsonCatalog, e := utils.CallAPIGet(url, token)
	if e != nil {
		panic(e)
	}

	e = json.Unmarshal(JsonCatalog, &catalog)
	if e != nil {
		panic(e)
	}

	for _, v := range catalog.Catalogs {
		utils.PrintStruct(v)
	}

	return nil
}

func catalogListId(url, token, id string) error {
	catalog := &Catalog{}
	JsonCatalog, e := utils.CallAPIGet(url+"/"+id, token)
	if e != nil {
		panic(e)
	}

	e = json.Unmarshal(JsonCatalog, &catalog)
	if e != nil {
		panic(e)
	}

	utils.PrintStruct(catalog)

	return nil
}
