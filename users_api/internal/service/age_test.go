package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name string
		dob  time.Time
	}{
		{
			name: "birthday passed",
			dob:  time.Now().AddDate(-25, -1, 0),
		},
		{
			name: "birthday upcoming",
			dob:  time.Now().AddDate(-25, 1, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			age := CalculateAge(tt.dob)
			if age <= 0 {
				t.Errorf("expected positive age, got %d", age)
			}
		})
	}
}
