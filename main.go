package tempfile

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	index   = 0
	tempDir = os.TempDir()
)

func nextIndex() (i int) {
	i = index
	index++
	return i
}

// Path returns the absolute path to a temporary file that
// doesn't exist yet.
func Path() string {
	name := fmt.Sprintf("%d_%d", os.Getpid(), nextIndex())
	path := filepath.Join(tempDir, name)
	os.Remove(path)
	return path
}
