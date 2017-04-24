package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/sryanyuan/govsc/command"
)

func main() {
	var cmdEntry = &cobra.Command{Use: "govsc"}
	cmdEntry.AddCommand(command.NewInitCommand())
	if err := cmdEntry.Execute(); nil != err {
		log.Println("Execute done with error:", err)
	}
}
