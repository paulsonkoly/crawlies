package downloader

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

const bufsize = 4096

// IOError happens when there is a problem in writing the downloaded file
var IOError = errors.New("io error")

type downInfo struct {
	rdr   io.ReadCloser
	iSize int
}

// Status is the progress status of a download
type Status struct {
	FileName   string // FileNAme refers to the file being downloaded
	Percentage int    // Percentage refers to the dowload progression
	Err        error  // Err indicates that there was an issue during download
}

// Download executes a download, reporting status on the status channel
func Download(url *url.URL, status chan<- Status) {
	rInp, err := openIStream(url)
	if err != nil {
		status <- Status{Err: err}
		return
	}
	defer rInp.rdr.Close()

	wOut, err := openWFile(url.Path)
	if err != nil {
		status <- Status{Err: err}
		return
	}
	defer wOut.Close()

	var cpd, soFar int64
	for err != io.EOF || cpd > 0 {
		cpd, err = io.CopyN(wOut, rInp.rdr, bufsize)
		if err != nil && err != io.EOF {
			status <- Status{Err: err}
			return
		}
		soFar += cpd
		if rInp.iSize > 0 {
			status <- Status{FileName: path.Base(url.Path), Percentage: int((soFar * 100) / int64(rInp.iSize))}
		} else {
			status <- Status{FileName: path.Base(url.Path)}
		}
	}
	status <- Status{FileName: path.Base(url.Path), Percentage: 100}
}

func openIStream(url *url.URL) (*downInfo, error) {
	client := http.DefaultClient
	resp, err := client.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("Network error %w", err)
	}

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("Non 2xx response %d for %v", resp.StatusCode, url)
	}

	cs := resp.Header.Get("Content-Length")
	var iSize int
	if cs != "" {
		iSize, _ = strconv.Atoi(cs)
	}

	return &downInfo{rdr: resp.Body, iSize: iSize}, nil
}

func openWFile(fn string) (io.WriteCloser, error) {
	uPath := "./" + fn
	dirName := path.Dir(uPath)
	fName := path.Base(uPath)
	if err := os.MkdirAll(dirName, 0o700); err != nil {
		return nil, fmt.Errorf("mkdir error %w", err)
	}

	f, err := os.OpenFile(dirName+"/"+fName, os.O_CREATE|os.O_WRONLY, 0o640)
	if err != nil {
		return nil, fmt.Errorf("file io error %w", err)
	}
	return f, nil
}
