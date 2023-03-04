package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	argc := len(os.Args)
	if argc < 2 {
		fmt.Println("goterm <port> [baud]")
		os.Exit(1)
	}

	port := os.Args[1]
	if ! strings.HasPrefix(port, "/dev/") {
		fmt.Println("goterm <port> [baud]\n")
		fmt.Println("port must begin with /dev/")
		os.Exit(2)
	}

	baud := 115200
	if argc > 2 {
		f, err := strconv.ParseFloat(os.Args[2], 10) // allow scientific notation i.e. 1.5e6
		if err != nil {
			fmt.Println("goterm <port> [baud]\n")
			fmt.Println(err)
			os.Exit(3)
		}
		baud = int(f)
	}

	err := Monitor(port, baud)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
}
