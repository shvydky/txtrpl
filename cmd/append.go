/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/shvydky/txtrpl/text"
	"github.com/spf13/cobra"
)

// appendCmd represents the append command
var appendCmd = &cobra.Command{
	Use:   "append",
	Short: "Add replacement text as the last line of the output",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		outFile, err := text.CreateFile(*out, os.Stdout)
		if err != nil {
			return err
		}
		defer func() {
			if outFile != os.Stdout {
				outFile.Close()
			}
		}()

		buffer, found, err := text.ReadText(*in, *regex, *pattern)
		if err != nil {
			return err
		}
		if *appendIfNot {
			found = !found
		}
		outFile.Write(buffer)
		if found {
			line := *replace + "\r\n"
			outFile.WriteString(line)
		}

		return nil
	},
}

var appendIfNot *bool

func init() {
	rootCmd.AddCommand(appendCmd)
	appendIfNot = appendCmd.Flags().BoolP("not", "n", false, "Add text in case if condition is negative")
}
