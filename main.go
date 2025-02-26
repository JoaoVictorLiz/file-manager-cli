package main

import (
	"fmt"
	"os"
	"joaovictorliz.com/file-manager-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}