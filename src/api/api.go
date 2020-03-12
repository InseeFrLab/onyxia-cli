package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"onyxia-cli/src/configuration"
)

func CreateNamespace(token string, namespaceName string, ownerType string, ownerId string) {
	request := &CreateNamespaceDTO{
		Namespace: Namespace{
			Id: namespaceName,
		},
		Owner: Owner{
			Type: ownerType,
			Id:   ownerId,
		},
	}
	fmt.Printf("Creating namespace %s as %s %s\n", namespaceName, request.Owner.Type, request.Owner.Id)
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(request)
	req, _ := http.NewRequest("POST", configuration.Configuration.API.URL+"/api/namespace", buf)

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		panic(e)
	}

	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	// Print the body to the stdout
	io.Copy(os.Stdout, res.Body)
}
