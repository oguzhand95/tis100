package parser_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/oguzhand95/tis100/parser"
	"github.com/oguzhand95/tis100/testutil"
)

type testCase struct {
	Input    string `json:"input" yaml:"input"`
	Expected string `json:"expected" yaml:"expected"`
}

func TestParse(t *testing.T) {
	paths, err := testutil.Walk(testutil.PathTo("syntax"), "*.yaml")
	require.NoError(t, err)

	for _, path := range paths {
		tc := readTestCaseFromFile(t, path)
		t.Run(filepath.Base(path), func(t *testing.T) {
			p, err := parser.Parse(strings.NewReader(tc.Input), "")
			require.NoError(t, err)

			y, err := yaml.Marshal(p)
			require.NoError(t, err)

			t.Log(string(y))

			require.YAMLEq(t, tc.Expected, string(y))
		})
	}
}

func readTestCaseFromFile(t *testing.T, path string) *testCase {
	t.Helper()

	b, err := os.ReadFile(path)
	require.NoError(t, err)

	var testCase *testCase
	err = yaml.Unmarshal(b, &testCase)
	require.NoError(t, err)

	return testCase
}
