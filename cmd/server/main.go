package main

import (
	"fmt"

	"github.com/CastroEngineering/observability-tool/internal/version"
)

func main() {
	fmt.Println("server starting")
	fmt.Println(version.VersionNumber)
}
