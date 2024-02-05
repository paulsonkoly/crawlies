// ./crawlies
//
// A simple command line tool to download many files via http. Given an input
// file it reads URLs from the file and starts downloading on the specified
// number of threads.
//
// While dowloading it outputs using mpb a multi progress bar.
// Errors are aggregated and output at the end.
//
//	Usage of crawlies:
//
//	   -input string
//	       input file name
//	   -threadCnt int
//	       number of downloader threads (default 8)
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"sync"

	"github.com/paulsonkoly/crawlies/pkg/downloader"
	"github.com/paulsonkoly/crawlies/pkg/input"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

var inputFile = flag.String("input", "", "input file name")

var threadCnt = flag.Int("threadCnt", 8, "number of downloader threads")

type errorCollector struct {
	sync.Mutex
	errors []error
}

func newErrorCollector() *errorCollector {
	return &errorCollector{errors: make([]error, 0)}
}

func (e *errorCollector) addError(err error) {
	e.Lock()
	defer e.Unlock()
	e.errors = append(e.errors, err)
}

func main() {
	flag.Parse()
	if *inputFile == "" {
		log.Fatal("input file expected")
	}

	fInp, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fInp.Close()

	errors := newErrorCollector()
	urls := inputThread(input.New(fInp), errors)
	statuses := downloaderFanOut(urls)
	final := progressThread(statuses, errors)

	<-final

	for _, err2 := range errors.errors {
		fmt.Println(err2)
	}
}

func inputThread(i *input.Input, errors *errorCollector) <-chan url.URL {
	out := make(chan url.URL)
	go func() {
		defer close(out)
		for i.Next() {
			if i.Err() == io.EOF {
				return
			}
			if i.Err() != nil {
				errors.addError(i.Err())
				continue
			}
			url := i.Url()
			out <- url
		}
	}()
	return out
}

func downloaderFanOut(urls <-chan url.URL) <-chan downloader.Status {
	statuses := make(chan downloader.Status)
	wg := sync.WaitGroup{}
	wg.Add(*threadCnt)

	for i := 0; i < *threadCnt; i++ {
		go func() {
			for url := range urls {
				downloader.Download(&url, statuses)
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(statuses)
	}()
	return statuses
}

func progressThread(statuses <-chan downloader.Status, errors *errorCollector) <-chan struct{} {
	final := make(chan struct{})

	go func() {
		fNameToBar := map[string]*mpb.Bar{}
		progress := mpb.New()

		for status := range statuses {
			if status.Err != nil {
				errors.addError(status.Err)
				continue
			}
			if _, ok := fNameToBar[status.FileName]; !ok {
				bar := progress.AddBar(100, mpb.AppendDecorators(decor.Name(status.FileName, decor.WC{W: 70})))
				fNameToBar[status.FileName] = bar
			}
			bar := fNameToBar[status.FileName]
			bar.SetCurrent(int64(status.Percentage))
		}

		final <- struct{}{}
		close(final)
	}()

	return final
}
