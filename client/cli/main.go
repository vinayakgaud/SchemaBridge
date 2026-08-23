package main

import (
	"fmt"
	"log"

	"github.com/vinayakgaud/schemabridge/client/cli/commands"
)

func main() {
	fmt.Println("CLI Tool is executing")

	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
