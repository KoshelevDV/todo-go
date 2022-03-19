package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type list []task

func (l list) String() string {
	var str strings.Builder
	for i := range l {

		str.WriteString(l[i].String())
		str.WriteRune('\n')
	}
	return str.String()
}

func (l list) ShowShort() {
	var str strings.Builder
	for i := range l {
		temp := fmt.Sprintf("%03d. %s\n", i, l[i].Title)
		str.WriteString(temp)
	}
	fmt.Print(string(str.String()))
}

func (l *list) Save() {
	res, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	os.WriteFile("data.json", res, 0644)
}

func Load(path string) *list {
	f, _ := os.ReadFile(path)
	var r list
	json.Unmarshal(f, &r)
	return &r
}
