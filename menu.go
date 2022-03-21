package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func Menu(args []string, list list) {
	cmd := os.Args[1:]
	if len(cmd) < 1 {
		return
	}

	switch cmd[0] {
	case "show":
		Show(cmd, list)
	case "short":
		Short(list)
	case "add":
		Add(list)
	case "set":
		Set(cmd, list)
	default:
		fmt.Printf("Unknow command\n")
	}
}

func InteractiveMenu(cmd []string, list list) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cmd := strings.Fields(scanner.Text())

		switch cmd[0] {
		case "show":
			Show(cmd, list)
		case "short":
			Short(list)
		case "add":
			Add(list)
		case "set":

		default:
			fmt.Printf("Unknow command\n")
		}
	}
}

func Show(cmd []string, list list) {
	if len(cmd) < 2 {
		fmt.Println("No id specified.\ntodoist show <N>")
		return
	}
	list.ShowOne(cmd[1])
}

func Short(list list) {
	list.ShowShort()
}

func Add(list list) {
	tt, err := CreateToDo()
	if err != nil {
		fmt.Println(">>> main.go >>>,", err)
	}
	list = append([]task(list), *tt)
	list.Save()
}

func Set(cmd []string, list list) {
	if len(cmd) < 3 {
		fmt.Println("No field specified.\ntodoist show status [complete | \"on track\" | behind | \"not started\"]")
		return
	}
	list.SetStatus(cmd[2], cmd[3])
	list.Save()
}

func CheckData(p string) bool {
	_, err := os.Stat(p)
	return !errors.Is(err, os.ErrNotExist)
}

func CreateToDo() (*task, error) {
	scanner := bufio.NewScanner(os.Stdin)

	title := readString("Enter ToDo title: ", scanner)
	comment := readString("Enter comment for ToDo: ", scanner)
	untilDate, err := readTime("Write date until you have to complete your ToDo (format: yyyy/mm/dd hh:mm): ", scanner)
	if err != nil {
		return nil, err
	}
	status := NOT_STARTED

	return NewTask(title, comment, toTimestamp(int(time.Now().Unix())), untilDate, status), nil
}

func readString(t string, s *bufio.Scanner) string {
	fmt.Print(t)
	s.Scan()
	return strings.TrimSpace(s.Text())
}

func readTime(t string, s *bufio.Scanner) (ts timestamp, err error) {
	const layout = "2006/01/02 15:04"

	fmt.Print(t)
	s.Scan()

	value := strings.TrimSpace(s.Text())
	tmp, err := time.ParseInLocation(layout, value, time.Local)
	if err != nil {
		return ts, err
	}

	return toTimestamp(int(tmp.Unix())), nil
}
