package main

import (
	"bytes"
	"os"
	"sync"
	"testing"
)

func TestPrintMessage(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	go printMessage(&wg)
	wg.Wait()

	w.Close()
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("failed to read from pipe: %v", err)
	}
	os.Stdout = old

	got := buf.String()
	want := "Hello from goroutine!\n"
	if got != want {
		t.Errorf("printMessage() = %q, want %q", got, want)
	}
}
