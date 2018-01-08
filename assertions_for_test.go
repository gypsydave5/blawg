package blawg

import (
	"os"
	"strings"
	"testing"
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

func (a Assertions) True(b bool, s string) {
	if !b {
		a.test.Errorf(s)
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

func (a Assertions) ErrorMessage(err error, message string) {
	if err == nil {
		a.test.Errorf("expected an error, but received nil")
		return
	}

	actual := err.Error()
	if actual != message {
		a.test.Errorf("expected error message '%s' but got '%s'", message, actual)
	}
}

func (a Assertions) FileExists(pathToFile string) bool {
	_, err := os.Stat(pathToFile)

	return os.IsNotExist(err)
}

func (a Assertions) DirectoryExists(pathToDirectory string) {
	a.FileExists(pathToDirectory)
}

func (a Assertions) FileDoesNotExist(pathToFile string) {
	_, err := os.Stat(pathToFile)
	if err == nil {
		a.test.Errorf("Expected file '%s' not to exist", pathToFile)
	}
}
