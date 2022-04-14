package cmd

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/shvydky/txtrpl/text"
	"github.com/spf13/cobra"
)

// replaceCmd represents the replace command
var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Replace matched text with replacement text in-place.",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		inFile, err := text.OpenFile(*in, os.Stdin)
		if err != nil {
			return err
		}
		defer func() {
			if inFile != os.Stdin {
				inFile.Close()
			}
		}()

		outFile, err := text.CreateFile(*out, os.Stdout)
		if err != nil {
			return err
		}
		defer func() {
			if outFile != os.Stdout {
				outFile.Close()
			}
		}()

		var re *regexp.Regexp
		if *regex {
			re, err = regexp.Compile(*pattern)
		} else {
			re, err = regexp.Compile(regexp.QuoteMeta(*pattern))
			*replace = strings.ReplaceAll(*replace, "$", "$$")
		}
		if err != nil {
			return err
		}

		reader := bufio.NewReader(inFile)
		writer := bufio.NewWriter(outFile)

		char, _, err := reader.ReadRune()
		buffer := make([]rune, 0, 256)
		for err == nil {
			buffer = append(buffer, char)
			if char == '\n' {
				content := string(buffer)
				content = text.Replace(re, content, *replace)
				_, err = writer.Write([]byte(content))
				if err != nil {
					return err
				}
				buffer = buffer[:0]
			}
			char, _, err = reader.ReadRune()
		}
		writer.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(replaceCmd)
}
