package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "nana"}
	rootCmd.AddCommand(getPatchCmd())
	rootCmd.AddCommand(getDeleteFuncCmd())
	rootCmd.Execute()
}
