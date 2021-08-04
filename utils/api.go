package utils

import (
	"io"
	"net/http"
)

func CallAPIGet(url, token string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		return []byte(""), e
	}
	body, e := io.ReadAll(res.Body)

	return body, e
}
