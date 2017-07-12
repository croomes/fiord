package fio

import (
	"encoding/json"
	"fmt"
	"io"
)

func Decode(input io.Reader) (Report, error) {

	var report Report

	decoder := json.NewDecoder(input)

	err := decoder.Decode(&report)
	if err != nil {
		fmt.Printf("caught error: %v\n", err)
	}
	return report, nil
}
