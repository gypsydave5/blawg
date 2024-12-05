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
		a.test.Helper()
		a.test.Errorf(`Expected to find
"%s"
in
"%s"
`, substr, str)
	}
}

func (a Assertions) StringDoesNotContain(str, substr string) {
	if strings.Contains(str, substr) {
		a.test.Helper()
		a.test.Errorf(`Expected not to find
"%s"
in
"%s"
`, substr, str)
	}
}

func (a Assertions) True(b bool, s string) {
	if !b {
		a.test.Helper()
		a.test.Errorf(s)
	}
}

func (a Assertions) StringsEqual(str1, str2 string) {
	if str1 != str2 {
		a.test.Helper()
		a.test.Errorf("Expected '%s' to equal '%s'", str1, str2)
	}
}

func (a Assertions) NotError(err error) {
	if err != nil {
		a.test.Helper()
		a.test.Errorf("unexpected error: %s", err)
	}
}

func (a Assertions) ErrorMessage(err error, message string) {
	if err == nil {
		a.test.Helper()
		a.test.Errorf("expected an error, but received nil")
		return
	}

	actual := err.Error()
	if actual != message {
		a.test.Helper()
		a.test.Errorf("expected error message '%s' but got '%s'", message, actual)
	}
}

func (a Assertions) StringContainsInOrder(s string, first string, second string) {
	a.StringContains(s, first)
	a.StringContains(s, second)
	split := strings.Split(s, first)
	if strings.Contains(split[0], second) {
		a.test.Helper()
		a.test.Errorf("expected '%s' to appear after '%s', but it appears before in string '%s'", second, first, s)
	}
}

func (a Assertions) FileExists(pathToFile string) {
	_, err := os.Stat(pathToFile)
	if os.IsNotExist(err) {
		a.test.Helper()
		a.test.Errorf("expected file '%s' to exist", pathToFile)
	}
}

func (a Assertions) DirectoryExists(pathToFile string) {
	d, err := os.Stat(pathToFile)
	if os.IsNotExist(err) {
		a.test.Helper()
		a.test.Errorf("expected directory '%s' to exist", pathToFile)
		return
	}
	if !d.IsDir() {
		a.test.Helper()
		a.test.Errorf("expected '%s' to be a directory, but it isn't", pathToFile)
	}
}

func (a Assertions) FileDoesntExist(pathToFile string) {
	_, err := os.Stat(pathToFile)
	if !os.IsNotExist(err) {
		a.test.Helper()
		a.test.Errorf("expected file '%s' not to exist", pathToFile)
	}
}
