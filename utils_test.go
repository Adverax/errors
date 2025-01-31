package errors

import (
	"fmt"
	"testing"
)

type myPerson struct {
	Name string
	Age  int
}

var seq int

func next() int {
	seq++
	return seq
}

func newPerson(name string) (*myPerson, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	return &myPerson{
		Name: name,
		Age:  20 + next(),
	}, nil
}

func TestCheckList(t *testing.T) {
	tests := []struct {
		name    string
		items   []Option[*myPerson]
		wantErr bool
		wantLen int
	}{
		{
			name: "all valid",
			items: []Option[*myPerson]{
				WithItem(newPerson("John")),
				WithItem(newPerson("Alice")),
			},
			wantErr: false,
			wantLen: 2,
		},
		{
			name: "one invalid",
			items: []Option[*myPerson]{
				WithItem(newPerson("John")),
				WithItem(newPerson("")),
				WithItem(newPerson("Alice")),
			},
			wantErr: true,
			wantLen: 0,
		},
		{
			name: "all invalid",
			items: []Option[*myPerson]{
				WithItem(newPerson("")),
				WithItem(newPerson("")),
			},
			wantErr: true,
			wantLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list, err := CheckList(tt.items...)
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
			}
			if len(list) != tt.wantLen {
				t.Fatalf("expected %d items, got %d", tt.wantLen, len(list))
			}
		})
	}
}
