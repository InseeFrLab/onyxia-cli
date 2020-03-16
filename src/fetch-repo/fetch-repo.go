package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

func getURL(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	return resp, err
}

func fetchRepo(url string) (repo.IndexFile, error) {
	resp, err := getURL(url + "/index.yaml")
	file, _ := ioutil.ReadAll(resp.Body)
	var indexFile repo.IndexFile
	yaml.Unmarshal(file, &indexFile)
	return indexFile, err
}

// Name     string `json:"name"`
// URL      string `json:"url"`
// Username string `json:"username"`
// Password string `json:"password"`
// CertFile string `json:"certFile"`
// KeyFile  string `json:"keyFile"`
// CAFile   string `json:"caFile"`
func fetchChart(chartURL string) {
	entry := &repo.Entry{Username: "", CAFile: "", CertFile: "", KeyFile: "", Name: "", Password: "", URL: ""}
	mygetter := getter.All(&cli.EnvSettings{})
	repo.NewChartRepository(entry, mygetter)
}

func main() {
	index, _ := fetchRepo("https://charts.bitnami.com/bitnami")
	fmt.Println(index.APIVersion)
	entries := index.Entries
	for key, values := range entries {
		for _, value := range values {
			fmt.Println("Key:", key, "value:", *&value)
		}
	}
}
