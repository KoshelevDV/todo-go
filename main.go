package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

var (
	HOMEDIR, _ = os.UserHomeDir()
)

func main() {
	list := list{}
	ok := CheckData(path.Join(HOMEDIR, ".todoist/data.json"))
	if ok {
		list = *Load(path.Join(HOMEDIR, ".todoist/data.json"))
	} else {
		err := os.MkdirAll(path.Join(HOMEDIR, ".todoist/"), 0744)
		if err != nil {
			fmt.Println(err)
		}
	}
	cmd := os.Args[1:]
	if len(cmd) < 1 {
		return
	}

	switch cmd[0] {
	case "show":
		if len(cmd) < 2 {
			fmt.Println("No id specified.\ntodoist show <N>")
			return
		}
		list.ShowOne(cmd[1])
	case "short":
		list.ShowShort()
	case "add":
		tt, err := CreateToDo()
		if err != nil {
			fmt.Println(">>> main.go >>>,", err)
		}
		list = append([]task(list), *tt)
		list.Save()
	case "set":
		if len(cmd) < 3 {
			fmt.Println("No field specified.\ntodoist show status [complete | \"on track\" | behind | \"not started\"]")
			return
		}
		list.SetStatus(cmd[2], cmd[3])
		list.Save()
	default:
		fmt.Printf("Unknow command\n")
	}

}

func CheckData(p string) bool {
	_, err := os.Stat(p)
	return !errors.Is(err, os.ErrNotExist)
}

func Interactive(cmd []string, list list) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cmd := strings.Fields(scanner.Text())

		switch cmd[0] {
		case "show":
			list.ShowOne(cmd[1])
		case "short":
			list.ShowShort()
		case "add":
			tt, err := CreateToDo()
			if err != nil {
				fmt.Println(">>> main.go >>>,", err)
				continue
			}
			list = append([]task(list), *tt)
		case "quit":
			list.Save()
			return
		default:
			fmt.Printf("Unknow command\n")
		}
	}
}
