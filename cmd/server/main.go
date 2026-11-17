package main

import (
	"fmt"
	"os"

	"github.com/CastroEngineering/observability-tool/internal/port"
	"github.com/CastroEngineering/observability-tool/internal/version"
)

func main() {
	fmt.Println("server starting")
	fmt.Println(version.Number)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: agent <port>")
		os.Exit(1)
	}
	portNumber, err := port.Parser(os.Args[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(portNumber)

}
