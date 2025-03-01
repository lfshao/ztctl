package main

import (
	"os"

	"github.com/lfshao/ztctl/cmd"
	"github.com/lfshao/ztctl/pkg/output"
)

func main() {
	if err := cmd.Execute(); err != nil {
		output.DefaultPrinter.PrintMsg(output.LevelError, "%v", err)
		os.Exit(1)
	}
}
