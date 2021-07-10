package main

import "testing"

func TestRun(t *testing.T) {
	var tests = []struct {
		name    string
		wantErr bool
	}{
		{name: "Testing run function", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := run(); (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
