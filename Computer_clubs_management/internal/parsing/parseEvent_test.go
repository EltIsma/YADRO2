package parsing

import (
	"club_control/internal/domain"
	"reflect"
	"testing"
	"time"
)

func TestParseEvent(t *testing.T) {
	t1, _ := time.Parse("15:04", "09:00")
	t2, _ := time.Parse("15:04", "09:41")
	tests := []struct {
		input            string
		numTables        int
		seqCheck         time.Time
		wantEvent        *domain.Event
		wantError        string
		wantErrSubstring string
	}{
		{
			input:            "09:41 1 client1",
			numTables:        3,
			seqCheck:         t1,
			wantEvent:        domain.NewEvent(&domain.EventOptions{Time: t2, ID: 1, Client: "client1", Table: 0}),
			wantError:        "",
			wantErrSubstring: "",
		},
		{
			input:            "09:41 2 client2 1",
			numTables:        3,
			seqCheck:         t1,
			wantEvent:        domain.NewEvent(&domain.EventOptions{Time: t2, ID: 2, Client: "client2", Table: 1}),
			wantError:        "",
			wantErrSubstring: "",
		},
		{
			input:            "08:41 1 client1",
			numTables:        3,
			seqCheck:         t1,
			wantEvent:        nil,
			wantError:        "incorrect time: events must happen sequentially",
			wantErrSubstring: "08:41 1 client1",
		},
		{
			input:            "08:41 2 client1",
			numTables:        3,
			seqCheck:         t1,
			wantEvent:        nil,
			wantError:        "the line must contain 4 objects",
			wantErrSubstring: "08:41 2 client1",
		},
		{
			input:            "08:41 2 client1 4",
			numTables:        3,
			seqCheck:         t1,
			wantEvent:        nil,
			wantError:        "incorrect table number",
			wantErrSubstring: "08:41 2 client1 4",
		},
	}

	for _, test := range tests {

		gotEvent, line, gotError := ParseEvent(test.input, test.numTables, test.seqCheck)

		if !reflect.DeepEqual(gotEvent, test.wantEvent) {
			t.Errorf("parseConfig(%v) = %v, want %v", test.input, gotEvent, test.wantEvent)
		}

		if gotError == nil && test.wantError != "" {
			t.Errorf("parseConfig(%v) = nil, want error %v", test.input, test.wantError)
		} else if gotError != nil && test.wantError == "" {
			t.Errorf("parseConfig(%v) = %v, want no error", test.input, gotError)
		} else if gotError != nil && line != test.wantErrSubstring {
			t.Errorf("parseConfig(%v) = %v, want error containing substring %v", test.input, gotError, test.wantErrSubstring)
		}
	}
}
