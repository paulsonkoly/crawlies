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

var URLParseError error = errors.New("invalid URL")

type Input struct {
	scn *bufio.Scanner
	url *url.URL
	err error
}

func New(rdr io.Reader) *Input {
	scn := bufio.NewScanner(rdr)
	scn.Split(bufio.ScanLines)
	return &Input{scn: scn}
}

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

func (i *Input) Url() url.URL {
	return *i.url
}

func (i *Input) Err() error {
	return i.err
}
