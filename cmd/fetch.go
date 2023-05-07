/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

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
	"github.com/Arjun259194/nwz/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.PersistentFlags().String("ctr", "", "Fetch top news of a country")
	fetchCmd.PersistentFlags().String("s", "", "Fetch top news with the key word passed")
}

func handler(cmd *cobra.Command, args []string) {
	apiClient := api.Api{
		ApiKey: "a481486cec3745ddae53821bb5132799",
	}

	country, _ := cmd.Flags().GetString("ctr")
	key, _ := cmd.Flags().GetString("s")

	var articles []api.Article

	if key != "" {
		articles = apiClient.GetNewsFromKeyWord(key)
	} else if country != "" {
		articles = apiClient.GetTopCountryNews(country)
	} else {
		articles = apiClient.GetTopCountryNews("us")
	}

	for i, article := range articles {
		if i > 10 {
			break
		}
		article.Display()
	}
}

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch news with different methods",
	Long:  `It let's you fetch news for a country or based on a keyword`,
	Run:   handler,
}
