package input_test

import (
	"strings"
	"testing"

	"github.com/paulsonkoly/crawlies/pkg/input"
)

func TestInput(t *testing.T) {
	inputData := "https://example.com\nhttps://example.org\nhttps://example.net\n"
	reader := strings.NewReader(inputData)
	inputReader :=  input.New(reader)
	expectedURLs := []string{"https://example.com", "https://example.org", "https://example.net"}

	for i, expectedURL := range expectedURLs {
		if !inputReader.Next() {
			t.Fatalf("expected more input at index %d", i)
		}

		if inputReader.Err() != nil {
			t.Fatalf("error parsing URL at index %d: %s", i, inputReader.Err().Error())
		}

		aURL := inputReader.Url()
    actualURL := (&aURL).String()
		if actualURL != expectedURL {
			t.Errorf("unexpected URL at index %d: got %s, want %s", i, actualURL, expectedURL)
		}
	}

	if inputReader.Next() {
		t.Fatal("unexpected more input available")
	}

	if err := inputReader.Err(); err != nil {
		t.Fatalf("unexpected error at the end: %s", err.Error())
	}
}
