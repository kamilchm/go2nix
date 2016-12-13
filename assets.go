// Code generated by go-bindata.
// sources:
// templates/default.nix
// templates/deps.nix
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesDefaultNix = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x52\x3d\x6f\xdb\x30\x10\xdd\xf5\x2b\x0e\x4a\xc7\x50\x72\x0d\x34\x83\x83\x2e\x85\x91\xb6\x4b\x63\x04\x69\x96\xc0\xc3\x49\x3c\x93\x84\x25\x52\x20\x29\xf5\xc3\xf0\x7f\xef\x91\x91\x5a\x75\x10\xf0\xf8\xee\xde\x7b\x77\x14\x6f\xe0\x59\x9b\x00\x27\xd3\x11\xfc\xc0\x00\x8a\x2c\x79\x8c\x24\xa1\xf9\x05\x3a\xc6\x21\xec\xea\x5a\x99\xa8\xc7\xa6\x6a\x5d\x5f\x9f\xb1\x37\x5d\xab\xfb\x5a\xb9\xad\x35\x3f\x61\x7a\x7d\x85\xea\x85\x7c\x30\xce\xc2\xf1\x58\x5c\x20\x44\x49\x76\xba\x85\x66\x34\x9d\xfc\xec\x0e\xd8\x9e\x51\xd1\x2d\x9c\x28\xb6\x9a\x9d\x66\xa4\xd5\x0c\x9a\xdf\x7e\x46\x61\xb2\x70\xdd\x15\xc5\xff\x4a\xf0\xd4\xc2\xa5\x00\xb0\xd8\x13\x7c\x84\x32\x25\x1e\xce\xaa\xfa\x96\xce\xc7\xa3\x18\x6d\x88\xd8\x74\x24\xde\x5d\xa6\xb7\x41\xae\xe5\x3d\xf7\xcf\x87\xb5\xe4\xfb\x20\x79\xb7\x3d\x7f\xd5\x83\xf3\x3d\x46\x28\xb7\x9b\xcd\x9d\xd8\xbc\x17\x9b\x6d\xc9\x66\x59\xe8\x69\x5a\x8b\x9e\x68\x32\xf3\x7a\x5c\xe6\x3a\x17\xcc\x09\x2c\x41\xf5\x29\x4d\xfa\x8c\x2a\x40\x99\xd4\x5c\xcb\xb3\x3f\x74\x89\x62\x0b\x21\x62\x42\xc9\xe9\x5f\xeb\x9c\xc2\x24\x59\x09\x22\xcb\xd4\xb2\xed\x01\xa3\x5e\x87\x7f\xed\x07\xe7\x63\x66\x97\xf8\xe0\x5b\xee\x58\xae\x33\x5f\x0d\x80\xb1\x9a\x3c\x9f\x78\xf6\xfb\x4c\x8c\xbe\x5b\xfb\xbc\xb4\xe1\x89\x06\xb7\x84\xb3\x8b\xc6\xed\x87\xbb\x75\xcb\x17\x0c\x7a\xa9\x5f\x73\x92\x72\x7b\x1a\xd2\x22\x55\x2d\x19\x54\xfc\xc3\x33\xcf\x8f\xe6\x71\xff\xb8\x03\x94\x12\x7a\x8a\xc8\x97\x8a\x7f\x1f\x0b\x37\xb9\x50\x39\xaf\x12\x1a\xce\x2a\xd4\x3d\xda\x11\xbb\xfa\x26\x50\x2b\xf8\x57\x59\x89\x5e\x8a\xa4\x13\x18\xa3\x37\xcd\x18\x29\xb0\x6b\x62\x38\xeb\xf2\x16\x7f\x2d\xfe\x04\x00\x00\xff\xff\x3f\xbc\x6a\x92\x9b\x02\x00\x00")

func templatesDefaultNixBytes() ([]byte, error) {
	return bindataRead(
		_templatesDefaultNix,
		"templates/default.nix",
	)
}

func templatesDefaultNix() (*asset, error) {
	bytes, err := templatesDefaultNixBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/default.nix", size: 667, mode: os.FileMode(420), modTime: time.Unix(1481621420, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDepsNix = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8f\x31\x4b\x43\x31\x14\x46\xf7\xf7\x2b\x3e\xea\x6c\x02\x05\x1d\x14\x37\xd1\xb5\x68\x75\x09\x19\x6e\xd3\xdb\x24\xf4\xbd\xe4\x91\xa4\xd5\x22\xfd\xef\xbe\xb4\xb4\x34\x43\x08\x9c\x73\xbe\xe1\xde\x61\xe9\x7c\xc6\xc6\xf7\x8c\x1f\xca\xb0\x1c\x38\x51\xe1\x35\x56\x07\xb8\x52\xc6\xfc\x24\xa5\xf5\xc5\xed\x56\xc2\xc4\x41\x6e\x69\xf0\xbd\x71\x83\xb4\x71\x1e\xfc\x2f\xf6\x4a\x41\x7c\x73\xca\x3e\x06\x68\xdd\xa9\x4e\xa9\x7b\x24\x0a\x96\x21\x5e\x79\xcc\x15\x02\x7f\xd3\x03\x6c\x5c\x90\xd9\x92\xe5\x05\x15\x87\x17\xcc\xea\xf8\xbd\x81\x5a\xcf\x9e\x4f\xed\x86\x8b\xa9\xcd\x79\x09\x94\xc3\xc8\x97\xc9\x5b\x75\x62\x59\xc9\xb5\x07\x76\xa9\x6f\x83\xaf\x09\xdc\xf8\xc4\xfb\xd6\x7f\x4c\xe0\xc6\x67\x47\xf3\x87\xc7\x36\xf9\x3c\xb3\x6b\x75\xac\xdf\xf1\x74\x23\x87\x75\xbd\x4d\x77\xff\x01\x00\x00\xff\xff\x39\x73\x16\x2a\x43\x01\x00\x00")

func templatesDepsNixBytes() ([]byte, error) {
	return bindataRead(
		_templatesDepsNix,
		"templates/deps.nix",
	)
}

func templatesDepsNix() (*asset, error) {
	bytes, err := templatesDepsNixBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/deps.nix", size: 323, mode: os.FileMode(420), modTime: time.Unix(1481621393, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/default.nix": templatesDefaultNix,
	"templates/deps.nix": templatesDepsNix,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"default.nix": &bintree{templatesDefaultNix, map[string]*bintree{}},
		"deps.nix": &bintree{templatesDepsNix, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

