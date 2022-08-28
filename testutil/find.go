//go:build test
// +build test

package testutil

import (
	"fmt"
	"os"
	"path/filepath"
)

// Walk walks the given path and returns the list of file paths seeking the files matching the given pattern.
// https://stackoverflow.com/a/55300382
func Walk(path, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(p)); err != nil {
			return err
		} else if matched {
			matches = append(matches, p)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find files with the given pattern: %w", err)
	}

	return matches, nil
}
