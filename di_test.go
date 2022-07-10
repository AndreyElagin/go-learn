package main

import (
	"fmt"
	"testing"
)

type HardcodedDataStore struct {
	id   string
	name string
}

func (ds HardcodedDataStore) UserNameForID(userID string) (string, bool) {
	if userID == ds.id {
		return ds.name, true
	} else {
		return "default vasian", true
	}
}

func TestSayHello(t *testing.T) {
	l := LoggerAdapter(LogOutput)
	ds := HardcodedDataStore{
		id:   "666",
		name: "satan",
	}
	logic := NewSimpleLogic(l, ds)

	var tests = []struct {
		id   string
		want string
	}{
		{"666", "Hello, satan"},
		{"777", "Hello, default vasian"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%s", tt.id, tt.want)
		t.Run(testName, func(t *testing.T) {
			resp, _ := logic.SayHello(tt.id)
			if resp != tt.want {
				t.Errorf("got %s, want %s", resp, tt.want)
			}
		})
	}
}
