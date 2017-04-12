package services_test

import (
	"testing"

	"github.com/gophersgang/orchestra/services"
)

func TestGetProperGopath(t *testing.T) {
	envGopath := "/Users/someuser/Desktop/go-sandbox/vendor:/Users/someuser/Desktop/go-sandbox:/Users/someuser/go"
	projectPath := "/Users/someuser/Desktop/go-sandbox/src/subproj/"
	res, _ := services.GetProperGopath(projectPath, envGopath)
	expected := "/Users/someuser/Desktop/go-sandbox"
	if res != expected {
		t.Errorf("wrong result: expected: %s, was: %s", expected, res)
	}
}

func TestGetProperGopathError(t *testing.T) {
	envGopath := "/Users/someuser/Desktop/go-sandbox/vendor:/Users/someuser/Desktop/go-sandbox:/Users/someuser/go"
	projectPath := "/Users/someuser/Desktop/some-non-related-folder/src/subproj/"
	_, err := services.GetProperGopath(projectPath, envGopath)
	if err == nil {
		t.Errorf("error was expected!")
	}
}
