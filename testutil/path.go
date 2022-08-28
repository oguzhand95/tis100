//go:build test
// +build test

package testutil

import (
	"errors"
	"path/filepath"
	"runtime"
)

const testDataPath = "testdata"

func GlobalPathTo(paths ...string) string {
	_, currFile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("failed to detect testdata directory"))
	}

	return filepath.Join(filepath.Dir(currFile), testDataPath, filepath.Join(paths...))
}

func PathTo(paths ...string) string {
	return filepath.Join(testDataPath, filepath.Join(paths...))
}
