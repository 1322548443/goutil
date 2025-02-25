// Package fsutil Filesystem util functions, quick create, read and write file. eg: file and dir check, operate
package fsutil

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/goutil/internal/comfunc"
)

const (
	// MimeSniffLen sniff Length, use for detect file mime type
	MimeSniffLen = 512
)

// PathMatcher path matcher func
type PathMatcher func(path string) bool

// DetectMime detect file mime type. alias of MimeType()
func DetectMime(path string) string {
	return MimeType(path)
}

// MimeType get File Mime Type name. eg "image/png"
func MimeType(path string) (mime string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	return ReaderMimeType(file)
}

// ReaderMimeType get the io.Reader mimeType
//
// Usage:
//
//	file, err := os.Open(filepath)
//	if err != nil {
//		return
//	}
//	mime := ReaderMimeType(file)
func ReaderMimeType(r io.Reader) (mime string) {
	var buf [MimeSniffLen]byte
	n, _ := io.ReadFull(r, buf[:])
	if n == 0 {
		return ""
	}

	return http.DetectContentType(buf[:n])
}

// JoinPaths elements, alias of filepath.Join()
func JoinPaths(elem ...string) string {
	return filepath.Join(elem...)
}

// JoinSubPaths elements, like the filepath.Join()
func JoinSubPaths(basePath string, elem ...string) string {
	paths := make([]string, len(elem)+1)
	paths[0] = basePath
	copy(paths[1:], elem)
	return filepath.Join(paths...)
}

// SlashPath alias of filepath.ToSlash
func SlashPath(path string) string {
	return filepath.ToSlash(path)
}

// UnixPath like of filepath.ToSlash, but always replace
func UnixPath(path string) string {
	if !strings.ContainsRune(path, '\\') {
		return path
	}
	return strings.ReplaceAll(path, "\\", "/")
}

// ToAbsPath convert process. will expand home dir
//
// TIP: will don't check path
func ToAbsPath(p string) string {
	if len(p) == 0 || IsAbsPath(p) {
		return p
	}

	// expand home dir
	if p[0] == '~' {
		return comfunc.ExpandHome(p)
	}

	wd, err := os.Getwd()
	if err != nil {
		return p
	}
	return filepath.Join(wd, p)
}
