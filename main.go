package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	list := list{}
	ok := CheckData("./data.json")
	if ok {
		list = *Load("./data.json")
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
