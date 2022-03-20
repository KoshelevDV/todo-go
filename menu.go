package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func CreateToDo() (t *task) {
	scanner := bufio.NewScanner(os.Stdin)

	title := readString("Enter ToDo title: ", scanner)
	comment := readString("Enter comment for ToDo: ", scanner)
	untilDate := readTime("Write date until you have to complete your ToDo (format: dd/mm/yyyy hh:mm): ", scanner)
	status := NOT_STARTED

	fmt.Println(NewTask(title, comment, toTimestamp(int(time.Now().Unix())), untilDate, status).String())
	return NewTask(title, comment, toTimestamp(int(time.Now().Unix())), untilDate, status)
}

func readString(t string, s *bufio.Scanner) string {
	fmt.Print(t)
	s.Scan()
	return s.Text()
}

func readTime(t string, s *bufio.Scanner) (ts timestamp) {
	const layout = "2006-01-02 15:04"

	fmt.Print(t)
	s.Scan()

	value := strings.TrimSpace(s.Text())
	tmp, err := time.ParseInLocation(layout, value, time.Local)
	if err != nil {
		panic(err)
	}

	return toTimestamp(int(tmp.Unix()))
}
