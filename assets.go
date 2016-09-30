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

var _templatesDefaultNix = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x52\x4d\x6f\xdb\x30\x0c\xbd\xfb\x57\x10\xee\x8e\xb5\xec\x05\x58\x07\xa4\xd8\x65\x08\xba\xed\xb2\x06\x45\xb7\x4b\x90\x83\x6c\x31\xb2\x10\x5b\x32\x44\xd9\xfb\x08\xf2\xdf\x47\xc9\x31\xe6\x1e\x0c\x3c\x92\x8f\xef\x91\xb4\xee\xe0\xb5\x35\x04\x27\xd3\x21\xfc\x92\x04\x1a\x2d\x7a\x19\x50\x41\xfd\x07\xb4\xdb\x58\xf3\x5b\x64\x17\xa0\xa0\xd0\x4e\xf7\x50\x8f\xa6\x53\x5f\xdc\x5e\x36\x67\xa9\xf1\x1e\x4e\x18\x9a\x56\x9b\x70\x43\xad\xbe\x81\xfa\xaf\xbf\x21\x9a\x2c\x5c\xb7\x59\xf6\xb6\x13\x3c\x36\x70\xc9\x00\xac\xec\x11\x3e\x41\x7e\x38\x80\xd8\x9f\xb5\xf8\x1e\xe3\xe3\xb1\x78\x77\x99\xd0\x93\x71\xf6\x9a\x3f\x32\xed\x16\xac\x99\x3f\x06\xc5\x73\xee\xf8\x13\x4f\xce\xf7\x32\x40\xbe\xa9\xaa\x87\xea\x7d\xb5\xc9\x67\x85\x79\x68\xd1\x99\x5a\x50\xf0\xc6\x6a\x12\x34\xd6\x33\x84\x0a\x3e\xf2\x10\xd3\x2c\xcf\x60\x2d\xfd\x82\x93\x49\x76\xc7\x23\x97\xb9\xce\x05\x73\x02\x8b\x20\x3e\xc7\x35\x5e\xa5\x26\xc8\xa3\x0b\xd7\xd2\x62\x4f\x5d\x4c\xb1\x44\x51\x84\x88\xa2\xd2\x7f\x6a\x92\x49\x2a\x68\x15\x14\xa9\x4d\x2f\xa7\xd8\xcb\xd0\xae\xcd\xbf\xf5\x83\xf3\x21\x65\x17\x7b\xf2\x0d\x33\x96\x5b\xa7\xbb\x01\x18\xdb\xa2\xe7\x88\x67\x7f\x4c\x89\xd1\x77\x6b\x9d\x9f\x0d\xbd\xe0\xe0\x16\x73\x56\x69\xe5\xe6\xc3\xc3\x9a\xf2\x55\x52\xbb\xd4\xaf\xc9\x49\xbb\x1d\x0e\x71\x11\x51\x2a\x06\x82\x7f\x7f\xca\xf3\x33\x79\xde\x3d\x6f\x41\x2a\x05\x3d\x06\xc9\xa7\x97\xd0\x86\x30\xd0\xb6\x2c\x99\xe4\x48\x38\xaf\x23\x1a\xce\x9a\xca\x5e\xda\x51\x76\xe5\x1d\x61\x53\x50\x90\x56\x49\xaf\x8a\xd8\x57\xc8\xc0\xe7\xaf\xc7\x80\xc4\xaa\x31\xc3\x5e\x97\xd9\xfe\x9a\xfd\x0b\x00\x00\xff\xff\x36\x91\xe6\xec\x8d\x02\x00\x00")

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

	info := bindataFileInfo{name: "templates/default.nix", size: 653, mode: os.FileMode(420), modTime: time.Unix(1475151286, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDepsNix = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8f\xbd\x0e\x82\x30\x14\x46\xf7\x3e\xc5\x17\x9c\xed\x40\xa2\x8b\x71\xd5\x95\x28\x4e\x4d\x87\x0a\x97\x42\x24\xd5\x94\xfa\x43\x0c\xef\x6e\x0b\x91\xd0\xa1\x69\x72\xce\xb9\xc3\xb7\x42\x5e\x37\x1d\xaa\xa6\x25\xbc\x55\x07\x4d\x86\xac\x72\x54\xe2\xda\x43\xdf\x53\xd3\x7c\x38\x13\x4c\x88\x35\xac\x32\x9a\xc0\x21\x25\x03\xbe\xfe\xc1\x07\x99\x2a\x6e\x4a\x53\xa6\x5c\x8d\x3d\x12\x21\xc0\x8f\x11\x94\x32\xd9\x8d\x6d\x45\xae\x08\xcd\x74\x09\xb8\xfe\x41\xff\x93\x43\x70\x3c\x0f\x64\xee\x81\xa7\x6d\xe3\xe0\xe2\xc1\xc2\x5b\x7a\xc5\xfe\xe4\xc1\xc2\x77\xb5\x4a\x37\xdb\x38\x39\x4f\x6c\xae\x86\xf0\x0d\xe3\x3e\x32\x65\xd8\x26\xd9\x2f\x00\x00\xff\xff\x48\x03\x79\x55\x14\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/deps.nix", size: 276, mode: os.FileMode(420), modTime: time.Unix(1475151286, 0)}
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
