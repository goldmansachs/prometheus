//go:build !nobuiltinassets && !embedfs

package ui

import (
	"bytes"
	"net/http"
	"os"
	"path"
	"time"
)

// bindataFile implements http.File interface for go-bindata assets
type bindataFile struct {
	*bytes.Reader
	name    string
	size    int64
	modtime time.Time
}

func (f *bindataFile) Close() error               { return nil }
func (f *bindataFile) Stat() (os.FileInfo, error) { return f, nil }
func (f *bindataFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, os.ErrInvalid
}

// os.FileInfo interface
func (f *bindataFile) Name() string       { return path.Base(f.name) }
func (f *bindataFile) Size() int64        { return f.size }
func (f *bindataFile) Mode() os.FileMode  { return 0444 }
func (f *bindataFile) ModTime() time.Time { return f.modtime }
func (f *bindataFile) IsDir() bool        { return false }
func (f *bindataFile) Sys() interface{}   { return nil }

// bindataFileSystem wraps go-bindata's Asset function to implement http.FileSystem
type bindataFileSystem struct{}

func (fs bindataFileSystem) Open(name string) (http.File, error) {
	// Remove leading slash if present
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}

	// Try to get the asset
	data, err := Asset(name)
	if err != nil {
		// Try with .gz extension for compressed assets
		gzName := name + ".gz"
		data, err = Asset(gzName)
		if err != nil {
			return nil, os.ErrNotExist
		}
		name = gzName
	}

	// Get asset info
	info, err := AssetInfo(name)
	if err != nil {
		return nil, err
	}

	return &bindataFile{
		Reader:  bytes.NewReader(data),
		name:    name,
		size:    info.Size(),
		modtime: info.ModTime(),
	}, nil
}

var Assets = http.FileSystem(bindataFileSystem{})