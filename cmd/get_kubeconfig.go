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
	"onyxiactl/cli/get"

	"github.com/spf13/cobra"
)

// kubeconfigCmd represents the kubeconfig command
var kubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "get kubeconfig from kube onboarding",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		generate, _ := cmd.Flags().GetBool("generate")
		destination, _ := cmd.Flags().GetString("destination")
		group, _ := cmd.Flags().GetString("group")
		get.ExecuteGetKubeConfig(group, generate, destination)
	},
}

func init() {
	getCmd.AddCommand(kubeconfigCmd)
	kubeconfigCmd.Flags().BoolP("generate", "g", false, "generate kubeconfig")
	kubeconfigCmd.Flags().StringP("destination", "d", "", "by default is in $HOME/.kube/config")
	kubeconfigCmd.Flags().StringP("group", "", "", "")
}
