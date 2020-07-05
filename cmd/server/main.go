package main

import (
	"fmt"
	"os"

	"github.com/dhanusaputra/anywhat-server/pkg/cmd"
	"github.com/dhanusaputra/anywhat-server/pkg/service"
)

func main() {
	s := service.NewAnywhatService()
	if err := cmd.ListenGRPC(s, 8080); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
