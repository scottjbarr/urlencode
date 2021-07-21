package main

import (
	"fmt"
	"net/url"

	"github.com/scottjbarr/urlencode"
)

func main() {
	data := urlencode.ReadInput()

	s, err := url.QueryUnescape(data)
	if err != nil {
		panic(err)
	}

	// write the bytes out as a string
	fmt.Printf("%s\n", s)
}
