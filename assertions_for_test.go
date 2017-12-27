package blawg

import (
	"testing"
	"strings"
	"os"
)

type Assertions struct {
	test *testing.T
}

func NewAssertions(t *testing.T) Assertions {
	return Assertions{t}
}

func (a Assertions) StringContains(str, substr string) {
	if !strings.Contains(str, substr) {
		a.test.Errorf(`Expected to find
"%s"
in
"%s"
`, substr, str)
	}
}

func (a Assertions) StringsEqual(str1, str2 string) {
	if str1 != str2 {
		a.test.Errorf("Expected '%s' to equal '%s'", str1, str2)
	}
}

func (a Assertions) NotError(err error) {
	if err != nil {
		a.test.Errorf("unexpected error: %s", err)
	}
}

func (a Assertions) FileExists(pathToFile string) {
	_, err := os.Stat(pathToFile)
	if err != nil {
		a.test.Errorf("Could not find file: %s", err)
	}
}

func (a Assertions) DirectoryExists(pathToDirectory string) {
	a.FileExists(pathToDirectory)
}
