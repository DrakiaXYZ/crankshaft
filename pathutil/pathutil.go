// Package pathutil implements utilities for handing file paths.
package pathutil

import (
	"bufio"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var getCurrentUser = user.Current

// SubstituteHomeDir takes a path that might be prefixed with `~`, and returns
// the path with the `~` replaced by the user's home directory.
func SubstituteHomeDir(path string) string {
	usr, _ := getCurrentUser()
	homeDir := usr.HomeDir
	if strings.HasPrefix(path, "~") {
		return filepath.Join(homeDir, path[2:])
	}
	return path
}

/*
AddExtPrefix adds a prefix to the path's file extension and returns a new path.

For example: AddExtPrefix("foo/bar/baz.js", ".bak") = "foo/bar/baz.bak.js"
*/
func AddExtPrefix(path string, extPrefix string) string {
	dir := filepath.Dir(path)
	filename := filepath.Base(path)
	ext := filepath.Ext(filename)

	filenameWithExtPrefix := strings.TrimSuffix(filename, ext) + extPrefix + ext

	return filepath.Join(dir, filenameWithExtPrefix)
}

// Copy copies the file from a source path to a destination path, overwriting
// the destination if it already exists.
func Copy(from, to string) error {
	fIn, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fIn.Close()

	fOut, err := os.Create(to)
	if err != nil {
		return err
	}
	defer fOut.Close()

	if _, err := io.Copy(fOut, fIn); err != nil {
		return err
	}

	return nil
}

// FileLines reads in a file and returns it as an array of lines.
func FileLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	fileLines := []string{}
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	return fileLines, nil
}