package text

import (
	"regexp"
	"testing"
)

func TestReplace(t *testing.T) {
	re := regexp.MustCompile(`a(x*y)b`)
	text, err := Replace(re, "dwf wrtg ab e;rltg ayb || axyb wrrthy axxxxyb e", `$1`)
	if err != nil {
		t.Error(err)
	}
	_ = text
}
