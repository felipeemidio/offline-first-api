package models

import (
	"errors"
	"strings"
	"time"
)

type Note struct {
	ID        uint64    `json:"id,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (note *Note) Check() error {
	if err := note.validate(); err != nil {
		return err
	}

	note.format()
	return nil
}

func (note *Note) validate() error {
	if note.Content == "" {
		return errors.New(`field "content" is required`)
	}

	return nil
}

func (note *Note) format() {
	note.Content = strings.TrimSpace(note.Content)
}
