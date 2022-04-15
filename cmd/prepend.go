/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/shvydky/txtrpl/text"
	"github.com/spf13/cobra"
)

// prependCmd represents the prepend command
var prependCmd = &cobra.Command{
	Use:   "prepend",
	Short: "Add replacement text as the first line of the output",
	RunE: func(cmd *cobra.Command, args []string) error {
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
		if *prependIfNot {
			found = !found
		}
		if found {
			line := *replace + "\r\n"
			outFile.WriteString(line)
		}
		outFile.Write(buffer)

		return nil
	},
}

var prependIfNot *bool

func init() {
	rootCmd.AddCommand(prependCmd)
	prependIfNot = prependCmd.Flags().BoolP("not", "n", false, "Add text in case if condition is negative")
}
