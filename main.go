package main

import (
	"os"

	"github.com/Dmytro-Hladkykh/ipfs-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
