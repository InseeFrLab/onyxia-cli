package utils

import (
	"fmt"
	"io"
	"net/http"
	"errors"
)

func CallAPIGet(url, token string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		return []byte(""), e
	}
	if res.StatusCode != 200 {
		return []byte(""), errors.New(fmt.Sprintf("API returned code %d", res.StatusCode))
	}
	body, e := io.ReadAll(res.Body)

	return body, e
}
