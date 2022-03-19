package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type task struct {
	Title        string
	Comment      string
	CreationDate timestamp
	UntilDate    timestamp
	Status       string
}

func (t task) String() string {
	var sb strings.Builder
	sb.WriteString("# ")
	sb.WriteString(t.Title)
	sb.WriteRune('\n')

	if t.Comment != "" {
		sb.WriteString("- ")
		sb.WriteString(t.Comment)
		sb.WriteRune('\n')
	}

	sb.WriteString("Creation date: ")
	sb.WriteString(t.CreationDate.String())
	sb.WriteRune('\n')

	sb.WriteString("To do until:   ")
	sb.WriteString(t.UntilDate.String())
	sb.WriteRune('\n')

	sb.WriteString("Status:        ")
	sb.WriteString(t.Status)
	sb.WriteRune('\n')

	return sb.String()
}

func (t *task) Save() {
	res, _ := json.MarshalIndent(t, "", "  ")
	fmt.Println(string(res))
}