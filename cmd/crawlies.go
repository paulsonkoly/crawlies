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
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
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

type progressArbiter struct {
	sync.Mutex
	progress *mpb.Progress
}

func newProgressArbiter() *progressArbiter {
	return &progressArbiter{progress: mpb.New()}
}

func (p *progressArbiter) addBar(fileName string) *mpb.Bar {
	p.Lock()
	defer p.Unlock()
	return p.progress.AddBar(100, mpb.AppendDecorators(decor.Name(fileName)))
}

type inputArbiter struct {
	sync.Mutex
	input *input.Input
}

func newInputArbiter(i io.Reader) *inputArbiter {
	return &inputArbiter{input: input.New(i)}
}

func (i *inputArbiter) getUrl(errors *errorCollector) (*url.URL, bool) {
	i.Lock()
	defer i.Unlock()

	if !i.input.Next() {
		return nil, false
	}
	if i.input.Err() != nil && i.input.Err() != io.EOF {
		errors.addError(i.input.Err())
		return nil, false
	}
	url := i.input.Url()
	return &url, true
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

	inp := newInputArbiter(fInp)
	p := newProgressArbiter()
	errors := newErrorCollector()

	wg := sync.WaitGroup{}
	for i := 0; i < *threadCnt; i++ {
		wg.Add(1)
		go func() {
			downloaderThread(inp, p, errors)
			wg.Done()
		}()
	}
	wg.Wait()

	for _, err2 := range errors.errors {
		fmt.Println(err2)
	}
}

func downloaderThread(i *inputArbiter, p *progressArbiter, errors *errorCollector) {
	for {
		url, ok := i.getUrl(errors)
		if !ok {
			return
		}
		var bar *mpb.Bar

		status := make(chan downloader.Status)
		go func() { downloader.Download(url, status) }()

		for fin := false; !fin; {
			stat := <-status
			if stat.Err != nil {
				errors.addError(stat.Err)
				fin = true
				continue
			}
			if bar == nil {
				bar = p.addBar(stat.FileName)
			}
			bar.IncrBy(stat.Percentage - int(bar.Current()))
			if stat.Percentage == 100 {
				fin = true
			}
		}
	}
}
