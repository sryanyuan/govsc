package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sryanyuan/govsc/command"
)

func main() {
	var cmdEntry = &cobra.Command{Use: "govsc"}
	cmdEntry.AddCommand(command.NewInit2Command())
	if err := cmdEntry.Execute(); nil != err {
		fmt.Println("Execute done with error:", err)
	}
}
