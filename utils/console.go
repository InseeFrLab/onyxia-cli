package utils

import (
	"fmt"
	"encoding/json"
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
