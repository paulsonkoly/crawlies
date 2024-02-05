package downloader_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/paulsonkoly/crawlies/pkg/downloader"
)

func TestDownload(t *testing.T) {
	// Mock server for testing purposes
	server := startMockServer()
	defer server.Close()

	// Mock URL for testing
	mockURL, _ := url.Parse(server.URL + "/x.txt")

	// Create a channel to receive download status updates
	statusChan := make(chan downloader.Status)
	defer close(statusChan)

	// Perform the download operation
	go downloader.Download(mockURL, statusChan)

	// Check if the download progresses and completes
	var lastStatus downloader.Status
	timeout := time.After(5 * time.Second) // Timeout for download test
	for {
		select {
		case status := <-statusChan:
			if status.Err != nil {
				t.Fatalf("download error: %s", status.Err.Error())
			}
			lastStatus = status
		case <-timeout:
			t.Fatal("download took too long to complete")
		}

		// Check if download reached 100%
		if lastStatus.Percentage == 100 {
			break
		}
	}

	// Check if the downloaded file exists
	downloadedFilePath := "./" + lastStatus.FileName
	if _, err := os.Stat(downloadedFilePath); os.IsNotExist(err) {
		t.Fatalf("downloaded file %s does not exist", downloadedFilePath)
	}

	// Clean up downloaded file
	os.Remove(downloadedFilePath)
}

// startMockServer starts a mock HTTP server for testing purposes
func startMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    mockData := "Some mock data"
		// Send mock response with content-length header
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(mockData)))
		w.WriteHeader(http.StatusOK)
		// Write some data to the response body
		w.Write([]byte(mockData))
	}))
}

