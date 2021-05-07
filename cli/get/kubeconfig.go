package get

import (
	"encoding/json"
	"fmt"
	"net/http"
	"onyxiactl/utils"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

//OnboardingResponse ...
type OnboardingResponse struct {
	ApiserverURL string
	Token        string
	Namespace    string
	User         string
	ClusterName  string
	Onboarded    string
}

//ExecuteGetKubeConfig ...
func ExecuteGetKubeConfig(group string, write bool, destination string) {
	var resp *http.Response
	var err error
	client := utils.GetAuthClient()
	if group == "" {
		resp, err = client.Get(viper.GetString("onboardingURL") + "/api/cluster")
	} else {
		resp, err = client.Get(viper.GetString("onboardingURL") + "/api/cluster/credentials/" + group)
	}

	if err != nil {
		panic(err)
	}
	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	onboardingResponse := new(OnboardingResponse)
	err = json.NewDecoder(resp.Body).Decode(onboardingResponse)

	config := clientcmdapi.NewConfig()
	config.Clusters[onboardingResponse.ClusterName] = &clientcmdapi.Cluster{
		Server:                onboardingResponse.ApiserverURL,
		InsecureSkipTLSVerify: true,
	}
	config.AuthInfos[onboardingResponse.User] = &clientcmdapi.AuthInfo{
		Token: onboardingResponse.Token,
	}
	config.AuthInfos[onboardingResponse.ClusterName] = &clientcmdapi.AuthInfo{
		Token: onboardingResponse.Token,
	}
	config.Contexts[onboardingResponse.ClusterName] = &clientcmdapi.Context{
		Cluster:   onboardingResponse.ClusterName,
		AuthInfo:  onboardingResponse.User,
		Namespace: onboardingResponse.Namespace,
	}
	config.CurrentContext = onboardingResponse.ClusterName
	pathOptions := clientcmd.NewDefaultPathOptions()
	home, _ := homedir.Dir()
	filePath := filepath.Join(home, ".kube", "config")
	pathOptions.GlobalFile = filePath
	if write {
		if destination != "" {
			pathOptions.GlobalFile = filepath.Join(filepath.Clean(destination), "kubeconfig")
		} else {
			fmt.Println(color.GreenString("kubeconfig set to follow context: "), onboardingResponse.ClusterName)
		}
		if err := clientcmd.ModifyConfig(pathOptions, *config, true); err != nil {
			fmt.Println("Unexpected error:" + err.Error())
		}

	}
	f, _ := clientcmd.Write(*config)
	fmt.Printf("%s", f)
}
