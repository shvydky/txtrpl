package text

import (
	"bufio"
	"os"
	"regexp"
)

var (
	findRef   regexp.Regexp = *regexp.MustCompile(`\\((?P<num>\d+)|(?P<name>[^\s]+))`)
	numIndex  int           = findRef.SubexpIndex("num")
	nameIndex int           = findRef.SubexpIndex("name")
)

func OpenFile(name string, def *os.File) (*os.File, error) {
	if name != "" {
		return os.Open(name)
	}
	return def, nil
}

func CreateFile(name string, def *os.File) (*os.File, error) {
	if name != "" {
		return os.Create(name)
	}
	return def, nil
}

func Replace(re *regexp.Regexp, text, replace string) string {
	content := []byte(text)
	template := []byte(replace)
	result := []byte{}

	// For each match of the regex in the content.
	passed := 0
	for _, submatches := range re.FindAllSubmatchIndex(content, -1) {
		// Apply the captured submatches to the template and append the output
		// to the result.
		if passed < submatches[0] {
			result = append(result, content[passed:submatches[0]]...)
		}
		result = re.Expand(result, template, content, submatches)
		passed = submatches[1]
	}

	if passed <= len(content) {
		result = append(result, content[passed:]...)
	}

	return string(result)
}

func ReadText(name string, regex bool, pattern string) ([]byte, bool, error) {
	inFile, err := OpenFile(name, os.Stdin)
	if err != nil {
		return nil, false, err
	}
	defer func() {
		if inFile != os.Stdin {
			inFile.Close()
		}
	}()
	var re *regexp.Regexp
	if regex {
		re, err = regexp.Compile(pattern)
	} else {
		re, err = regexp.Compile(regexp.QuoteMeta(pattern))
	}
	if err != nil {
		return nil, false, err
	}

	reader := bufio.NewReader(inFile)

	char, _, err := reader.ReadRune()
	buffer := make([]rune, 0, 256)
	passed := 0
	found := false
	for err == nil {
		buffer = append(buffer, char)
		if char == '\n' {
			if re.MatchString(string(buffer[passed:])) {
				found = true
			}
			passed = len(buffer)
		}
		char, _, err = reader.ReadRune()
	}

	return []byte(string(buffer)), found, nil
}
