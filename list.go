package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"
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

func (l list) ShowAll() {
	var str strings.Builder
	for i := range l {
		temp := fmt.Sprintf("%d\n%s\n", i+1, l[i].String())
		str.WriteString(temp)
	}
	fmt.Print(string(str.String()))
}

func (l list) ShowShort() {
	var str strings.Builder
	for i := range l {
		if l[i].Status != BEHIND {
			temp := fmt.Sprintf("%d. %s\n", i+1, l[i].Title)
			str.WriteString(temp)
		}
	}
	fmt.Print(string(str.String()))
}

func (l list) ShowOne(s string) {
	id, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Incorrect id %s\n", s)
	}
	if id > len(l) || id-1 < 0 {
		fmt.Printf("ToDo id is out of scope. Max id is %d\n", len(l))
		return
	}
	fmt.Println(l[id-1].String())
}

func (l list) SetStatus(_id, s string) {
	id, err := strconv.Atoi(_id)
	if err != nil {
		fmt.Printf("Incorrect id %s\n", _id)
		return
	}
	id -= 1
	switch strings.ToLower(strings.TrimSpace(s)) {
	case COMPLETED:
		l[id].Status = COMPLETED
	case ON_TRACK:
		l[id].Status = ON_TRACK
	case BEHIND:
		l[id].Status = BEHIND
	case NOT_STARTED:
		l[id].Status = NOT_STARTED
	default:
		fmt.Println(`Unknow status. You should use one of the next: "completed" "on track" "behind" "not yet started"`)
	}
}

func (l list) Save() {
	res, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path.Join(HOMEDIR, ".todoist", "data.json"), res, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func Load(path string) (r *list) {
	f, _ := os.ReadFile(path)
	json.Unmarshal(f, &r)
	return r
}
