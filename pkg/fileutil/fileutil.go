package fileutil

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

const (
	// PrivateFileMode grants owner to read/write a file.
	PrivateFileMode = 0600
	// PrivateDirMode grants owner to make/remove files inside the directory.
	PrivateDirMode = 0700
)

// IsDirWriteable checks if dir is writable by writing and removing a file
// to dir. It returns nil if dir is writable.
func IsDirWriteable(dir string) error {
	f := filepath.Join(dir, ".touch")
	if err := ioutil.WriteFile(f, []byte(""), PrivateFileMode); err != nil {
		return err
	}
	return os.Remove(f)
}

// ReadDir returns the filenames in the given directory in sorted order.
func ReadDir(dirpath string) ([]string, error) {
	dir, err := os.Open(dirpath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}

// TouchDirAll is similar to os.MkdirAll. It creates directories with 0700 permission if any directory
// does not exists. TouchDirAll also ensures the given directory is writable.
func TouchDirAll(dir string) error {
	// If path is already a directory, MkdirAll does nothing
	// and returns nil.
	err := os.MkdirAll(dir, PrivateDirMode)
	if err != nil {
		// if mkdirAll("a/text") and "text" is not
		// a directory, this will return syscall.ENOTDIR
		return err
	}
	return IsDirWriteable(dir)
}

// CreateDirAll is similar to TouchDirAll but returns error
// if the deepest directory was not empty.
func CreateDirAll(dir string) error {
	err := TouchDirAll(dir)
	if err == nil {
		var ns []string
		ns, err = ReadDir(dir)
		if err != nil {
			return err
		}
		if len(ns) != 0 {
			err = fmt.Errorf("expected %q to be empty, got %q", dir, ns)
		}
	}
	return err
}

func Exist(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

// ZeroToEnd zeros a file starting from SEEK_CUR to its SEEK_END. May temporarily
// shorten the length of the file.
func ZeroToEnd(f *os.File) error {
	// TODO: support FALLOC_FL_ZERO_RANGE
	off, err := f.Seek(0, io.SeekCurrent)
	if err != nil {
		return err
	}
	lenf, lerr := f.Seek(0, io.SeekEnd)
	if lerr != nil {
		return lerr
	}
	if err = f.Truncate(off); err != nil {
		return err
	}
	// make sure blocks remain allocated
	if err = Preallocate(f, lenf, true); err != nil {
		return err
	}
	_, err = f.Seek(off, io.SeekStart)
	return err
}
