package tempfile

import (
	"regexp"
	"testing"
)

func Test_Path(t *testing.T) {
	path := Path()
	matches, _ := regexp.MatchString("/\\d+_\\d+$", path)
	if !matches {
		t.Errorf("%s isn't a valid tempfile path", path)
	}
	if Path() == Path() {
		t.Error("Path() should return a unique path each time")
	}
}
