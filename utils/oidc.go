package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type ID struct {
	Name               string
	Preferred_username string
	Email              string
	Groups       []string
}

func DisplayID(tokenString string) {
	var id = GetID(tokenString)
	fmt.Println(id.Name)
	fmt.Println(id.Preferred_username)
	fmt.Println(id.Email)
	fmt.Println(id.Groups)
}

func DisplayGroups(id ID, auto_complete bool) {
	for _, element := range id.Groups {
		if auto_complete == true {
			fmt.Printf("%s\n", strings.Replace(strings.Replace(element, " ", "\\\\ ", -1), "'", `\\\'`, -1))
		} else {
			fmt.Printf("%s\n", strings.Replace(strings.Replace(element, " ", "\\ ", -1), "'", `\'`, -1))
		}
	}
}

func GetGroups(id ID) []string {
	var list []string

	for _, element := range id.Groups {
		list = append(list, element)
	}
	return list
}

func stripCtlFromUTF8(str string) string {
	return strings.Map(func(r rune) rune {
		if r >= 32 && r != 127 {
			return r
		}
		return -1
	}, str)
}

func GetID(tokenString string) ID {
	dataString := strings.Split(tokenString, ".")[1]
	data, _ := base64.RawStdEncoding.DecodeString(dataString)
	var id ID
	if err := json.Unmarshal((data), &id); err != nil {
		panic(err)
	}
	return id
}
