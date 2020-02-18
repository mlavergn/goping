package main

import (
	"fmt"
	"os"

	ping "github.com/mlavergn/goping"
)

// Version export
const Version = "0.1.0"

func main() {
	host := "www.google.com"

	if len(os.Args) > 1 {
		host = os.Args[1]
	}

	alive := ping.Ping(host)
	message := " is unreachable"

	if alive {
		message = " is alive"
	}

	fmt.Println(host + message)
}
