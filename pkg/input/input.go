// input is processing input files of the form one URL at a time
//
// It has an iterator interface that returns one URL at a time in a lazy fashion.
package input

import (
	"bufio"
	"errors"
	"io"
	"net/url"
)

// URLParseError indicates that the URL was incorrect
var URLParseError error = errors.New("invalid URL")

// Input represents out input file
type Input struct {
	scn *bufio.Scanner
	url *url.URL
	err error
}

// New creates an input file reader from an io.Reader. 
func New(rdr io.Reader) *Input {
	scn := bufio.NewScanner(rdr)
	scn.Split(bufio.ScanLines)
	return &Input{scn: scn}
}

// Next determines wether there is more input
func (i *Input) Next() bool {
	r := i.scn.Scan()
	if r {
		var err error
		i.url, err = url.Parse(i.scn.Text())
		if err != nil {
			i.err = URLParseError
		}
	}
	return r
}

// Url is the next url from the input.
//
// Url only becomes valid after the first call of Next.
func (i *Input) Url() url.URL {
	return *i.url
}

// Err is the next error if any
//
// Err only becomes valid after the first call of Next.
func (i *Input) Err() error {
	return i.err
}
