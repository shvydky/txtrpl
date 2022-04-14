package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "txtrpl",
	Short: "txtrpl is a simple text processing utility to users who are not familier with grep, sed, awk, etc",
	Long: `txtrpl implements several text manipulation methods that can be achieved by using
standard *nix utilities like grep, sed, awk. The motivation to have this utility mainly comes
from Windows where mentioned tools are not available.`,
}

var (
	pattern *string
	replace *string
	in      *string
	out     *string
	regex   *bool
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.txtrpl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	pattern = rootCmd.PersistentFlags().StringP("pattern", "p", "",
		"Search pattern. If combined with -r option should contain valid regex.")
	replace = rootCmd.PersistentFlags().StringP("text", "t", "",
		"Replace pattern. If combined with -r the group references can be used in form $<n> where $0 means 'all matched text' or in form $<name>.")
	regex = rootCmd.PersistentFlags().BoolP("regex", "r", false,
		"Expect pattern value to be a regular expression.")
	in = rootCmd.PersistentFlags().StringP("in", "i", "", "Name of the input file. STDIN is used by default")
	out = rootCmd.PersistentFlags().StringP("out", "o", "", "Name of the output file. STDOUT is used by default")
}
