package main

import (
	"fmt"
	"os"
	"path"
)

var (
	HOMEDIR, _ = os.UserHomeDir()
)

func main() {
	list := *InitData()

	Menu(os.Args, list)

}

func InitData() (list *list) {
	ok := CheckData(path.Join(HOMEDIR, ".todoist/data.json"))
	if ok {
		return Load(path.Join(HOMEDIR, ".todoist/data.json"))
	} else {
		err := os.MkdirAll(path.Join(HOMEDIR, ".todoist/"), 0744)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
}
