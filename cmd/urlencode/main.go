package main

import (
	"fmt"
	"net/url"

	"github.com/scottjbarr/urlencode"
)

func main() {
	data := urlencode.ReadInput()

	// write the bytes out as a string
	fmt.Printf("%s\n", url.QueryEscape(data))
}
