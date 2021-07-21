package urlencode

import (
	"bufio"
	"fmt"
	"os"
)

// ReadInput takes input from the args or stdin.
//
// If args are not supplied it is assumed input is via stdin, which will block until input is
// received.
func ReadInput() string {
	data := ""

	if len(os.Args) > 1 {
		// get the supplied args
		for _, s := range os.Args[1:] {
			if len(s) < 2 {
				s = fmt.Sprintf("%02v", s)
			}
			data += s
		}

		return data
	}

	// no data so trying stdin
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		data = s.Text()
	}

	return data
}
