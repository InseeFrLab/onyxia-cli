package utils

import (
	"fmt"
	"encoding/json"
	"os"
)

func PrintStruct(v interface{}) {
	j, _ := json.Marshal(v)
	fmt.Println(string(j))
}

func PrintStructs(v []interface{}) {
	for i := range(v) {
		PrintStruct(i)
	}
}

func PrintMessage(m string) {
	fmt.Print(m)
}

func PrintErrorMessageAndExit(m string) {
	fmt.Print(m)
	os.Exit(1)
}