package main

import (
	"strings"
)

const (
	COMPLETED   = "completed"
	ON_TRACK    = "on track"
	BEHIND      = "behind"
	NOT_STARTED = "not yet started"
)

type task struct {
	Title        string
	Comment      string
	CreationDate timestamp
	UntilDate    timestamp
	Status       string
}

func NewTask(t string, c string, cd timestamp, ud timestamp, s string) *task {
	task := task{
		Title:        t,
		Comment:      c,
		CreationDate: cd,
		UntilDate:    ud,
		Status:       s,
	}
	return &task
}

func (t task) String() string {
	var sb strings.Builder
	sb.WriteString("Title:         ")
	sb.WriteString(t.Title)
	sb.WriteRune('\n')

	if t.Comment != "" {
		sb.WriteString("Comment:       ")
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

// Not used
// func (t *task) Save() {
// 	res, _ := json.MarshalIndent(t, "", "  ")
// 	fmt.Println(string(res))
// }
