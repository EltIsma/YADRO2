package parsing

import (
	"bufio"
	"club_control/internal/domain"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestParseConfig(t *testing.T) {
	t1, _ := time.Parse("15:04", "09:00")
	t2, _ := time.Parse("15:04", "17:00")
	tests := []struct {
		input            string
		wantConfig       *domain.Config
		wantError        string
		wantErrSubstring string
	}{
		{
			input:            "3\n09:00 17:00\n10",
			wantConfig:       domain.NewClub(3, t1, t2, 10),
			wantError:        "",
			wantErrSubstring: "",
		},
		{
			input:            "invalid\n09:00 17:00\n10",
			wantConfig:       nil,
			wantError:        "invalid input: invalid number of tables",
			wantErrSubstring: "invalid",
		},
		{
			input:            "3\ninvalid 17:00\n10",
			wantConfig:       nil,
			wantError:        "parsing time \"invalid\": invalid time format",
			wantErrSubstring: "invalid 17:00",
		},
		{
			input:            "3\n09:00 10\n10",
			wantConfig:       nil,
			wantError:        "parsing time \"invalid\": invalid time format",
			wantErrSubstring: "09:00 10",
		},
		{
			input:            "3\n09:00 17:00\n-10",
			wantConfig:       nil,
			wantError:        "invalid input: invalid price",
			wantErrSubstring: "-10",
		},
		{
			input: "3\n09:00 17:00\n10\n\n",
			wantConfig:  domain.NewClub(3, t1, t2, 10),
			wantError:        "",
			wantErrSubstring: "",
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		scanner := bufio.NewScanner(r)
		gotConfig,line, gotError := ParseConfig(scanner)

		if !reflect.DeepEqual(gotConfig, test.wantConfig) {
			t.Errorf("parseConfig(%v) = %v, want %v", test.input, gotConfig, test.wantConfig)
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
