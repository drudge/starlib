package dataframe

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qri-io/starlib/testdata"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarktest"
)

func runScript(t *testing.T, scriptFilename string) (string, error) {
	t.Helper()
	output := "\n"
	printCollect := func(thread *starlark.Thread, msg string) {
		output = fmt.Sprintf("%s%s\n", output, msg)
	}

	thread := &starlark.Thread{Load: testdata.NewModuleLoader(Module)}
	thread.Print = printCollect
	starlarktest.SetReporter(thread, t)
	thread.SetLocal(keyOutputConfig, &OutputConfig{})

	_, err := starlark.ExecFile(thread, scriptFilename, nil, nil)
	return strings.Trim(output, "\n"), err
}

func expectScriptOutput(t *testing.T, scriptFilename, expectFilename string) {
	t.Helper()
	output, err := runScript(t, scriptFilename)
	if err != nil {
		t.Fatal(err)
	}
	expect := mustReadFile(t, expectFilename)

	expect = strings.Trim(expect, "\n")
	output = strings.Trim(output, "\n")
	if diff := cmp.Diff(expect, output); diff != "" {
		t.Errorf("mismatch. (-want +got):\n%s", diff)
	}
}

func mustReadFile(t *testing.T, filename string) string {
	t.Helper()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}
