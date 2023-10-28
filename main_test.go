package main_test

import (
	"demo"
	"testing"
)

func TestRun(t *testing.T) {
	if err := main.Run(); err != nil {
		t.Fatal(err)
	}
}
