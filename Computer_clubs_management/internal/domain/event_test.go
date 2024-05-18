package domain

import (
	"fmt"
	"testing"
	"time"
)

func TestValidateName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "valid name", input: "client1", want: true},
		{name: "name with special characters", input: "c@lient1", want: false},
		{name: "name with special characters in front", input: "**client1", want: false},
		{name: "name with special characters in back", input: "client1##", want: false},
		{name: "valid name", input: "client_-1", want: true},
		{name: "uppercase letters", input: "clientA", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ValidateName(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestValidateTime(t *testing.T) {
	tl, _ := time.Parse("15:04", "15:04")
	tests := []struct {
		tl   string
		want time.Time
		err  error
	}{
		{
			tl:   "15:04",
			want: tl,
			err:  nil,
		},
		{
			tl:   "invalid time",
			want: time.Time{},
			err:  fmt.Errorf("parsing time \"invalid time\": invalid time format"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.tl, func(t *testing.T) {
			ans, _ := ValidateTime(tt.tl)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
