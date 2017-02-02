// Code generated by go-bindata.
// sources:
// cli/locales/en-US.json
// DO NOT EDIT!

package utils

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

var _localesEnUsJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x59\x4f\x6f\xe3\xb6\x13\xbd\xe7\x53\x0c\x72\xf1\x25\xbf\x20\x7b\x5b\xfc\x2e\x45\xb0\xfd\x83\x2d\xd2\x2e\x10\x67\x0b\x14\x45\x61\x8c\xa5\x91\xcc\x0d\x45\xaa\xe4\x30\x86\x51\xf4\xbb\x17\x43\xc9\xb6\x5c\x58\xa4\xbc\xb1\xb3\x97\x05\x56\xe6\xbc\xf7\x38\x6f\x38\xfc\x93\x3f\xae\x00\xfe\xbe\x02\x00\xb8\x56\xe5\xf5\xff\xe1\x1a\xdb\x76\x61\xb0\xa1\xeb\x9b\xee\x2b\x3b\x34\x5e\x23\x2b\x6b\xe2\xcf\xef\xaf\xaf\x00\xfe\xb9\x39\x16\x16\x3c\xd6\x63\x71\xf7\x0d\xea\x1a\x9b\xf7\xf0\xe1\xe1\xe3\x38\xc2\x0b\x39\x2f\x01\xc7\x31\xee\x6e\xef\x6e\xdf\x1d\x0f\x2e\x69\x19\xea\xa4\x80\xf9\xca\xae\xe1\xf1\x87\xf9\x13\x14\xb6\x69\x82\x51\x45\xfc\x09\xd6\x8a\x57\xb0\x57\x67\x0d\x3b\xab\x35\x39\x40\x0f\x8a\x41\x79\x68\xc9\x55\xd6\x35\x54\xde\x1e\xe7\xfe\xe2\xad\xc9\x53\xdb\xc0\x6d\x60\x50\x06\x7e\x9e\x7f\xfa\x15\x04\x11\xf9\x38\xe0\x06\x1b\x7d\x0a\xe0\xef\xf7\xbf\x3c\x24\x01\x95\x79\x41\xad\xca\x45\x70\x7a\x04\xf1\x63\x37\x02\x3e\x3f\x3e\x80\x75\x40\x4d\xcb\x9b\x91\xe9\x3a\xaa\x95\x67\xb7\x11\xb4\xa4\xca\xc7\x7e\xa0\x80\x66\xa0\xd8\x3e\x53\x3a\x87\x3b\xb0\x27\x19\x7a\x1c\xae\xd8\x99\x97\xd5\x36\xf0\x39\xa1\xee\xaf\x40\x9e\x17\x85\x6d\x5a\x4d\x4c\xe5\xa8\xb4\x38\x0e\x3e\xec\xc6\x8d\x98\x50\xd9\x45\x49\xbe\x70\xaa\xe5\xf1\x22\x17\x6f\x3d\xc8\x60\xf1\x53\x2a\x14\x97\x36\x30\xf0\x8a\xa0\xb0\xa6\x52\x75\x70\xdd\x77\x5b\xc5\x8f\x1f\x1e\x3e\x8e\x38\x15\x19\x73\x75\x74\x26\x2e\x17\x34\x2d\x4a\x4a\xe6\x29\x68\x82\xef\x29\x91\x22\x4f\xee\x45\x15\xb4\xd0\xca\xf3\x84\x54\x3d\x28\xcf\xd0\xa0\xc1\x9a\x4a\x68\x54\xe1\x6c\x0f\xe0\x47\x44\x1e\xe0\xa7\x12\xf3\x3a\xe4\x6e\x65\x26\x09\x3e\x75\x8b\xb7\x4b\xfc\x2d\xcc\x43\xdb\x5a\xc7\x54\xc2\x0b\xea\x40\x1e\xd0\x11\xcc\xa4\xb1\xcc\x6e\x60\x26\xfd\x60\x26\xcb\x72\xc6\xb8\xd4\x34\x1b\xb3\xc0\x06\x3e\x31\x77\xc3\x99\x41\xdf\x7c\x41\x70\x94\xa9\x41\x2c\x1d\x9b\xef\x80\x2b\x9f\xc7\x73\xb0\x6c\x13\x9c\x62\x7b\x5a\x11\xf4\xe3\x40\xb6\xb0\x3c\xea\x37\x72\x0a\x0b\x41\x3f\xc9\xaa\x2e\x24\x99\xad\x21\x6a\xd6\x94\x13\xf1\xce\x98\xfe\x21\xac\xec\x49\x19\x44\x51\x08\x6a\x6c\xef\x1d\x82\x31\xa6\x4f\x00\x3b\x34\xc6\x7a\x9a\x31\xac\x4c\x8c\x9e\x06\x3b\x08\x98\x00\xff\xad\x9a\x84\xf4\xe9\xc2\x11\x32\x4d\xd9\x8c\x88\xbb\x12\x11\x11\x80\x07\x4b\x79\x02\x41\x72\xef\x79\x1d\x74\xa5\x74\x1a\xff\x91\xb0\x04\xec\x0e\x5b\xd6\xf5\x67\x24\xa5\xe3\xc6\xc6\xa8\x8c\xf4\x1e\x84\x41\x0a\x64\x8f\xeb\x04\xb1\x85\x8e\xe5\x16\x04\xc5\x43\xe5\x6c\x03\x9e\x4b\x65\x60\xb9\x81\x92\x2a\x0c\x9a\x27\x88\x54\x46\x2c\x76\x54\x2a\x47\x45\xb6\x90\xa2\xe2\x3e\x25\x42\x88\x06\x22\x00\x0c\x00\x52\xa4\xa1\x2d\xa7\xb9\xfa\x39\x0e\xfc\xaa\xec\xf7\x1c\xa9\x69\xbc\x1e\xfd\xc2\xde\x76\x2c\x5f\xed\x6d\x2f\xf2\x6d\xbd\xad\x69\xca\x4e\xf1\xd3\x57\xae\x29\x41\x4f\xa9\x7f\x0d\xee\xab\x5b\xfc\x0e\xe9\x0c\xfd\x7d\x98\xcb\x73\x37\xf7\x1d\xf6\x59\x3b\x7b\xec\xe8\xb1\xb7\xe7\x68\x51\xa7\xaf\x3c\x3d\x27\x6a\x9d\x3e\x73\xed\x0f\xf2\x49\xb8\xee\x20\x7f\x50\xd3\x93\xeb\xa2\x87\xcf\x17\xf4\x01\x49\xbc\xa9\xcb\x2d\xc4\xb7\x54\xa8\x4a\x51\xb9\xad\x9c\x1b\xb1\xfc\x46\x52\x35\xd1\xa6\x5e\xc0\x79\x6a\xb3\x07\x3b\x57\x79\xee\x93\x73\x91\x0a\xed\xe1\x73\xd5\x72\x7f\x42\x99\x08\x56\x65\x5d\xe6\x78\xf8\xa3\x8c\x80\x18\xb2\x7d\x77\x91\xbb\x66\xbc\x67\xf6\xd7\xcf\x89\x5c\xc3\x90\xb1\xc2\xb1\xb0\xb1\x01\x1c\xa1\xd6\x1b\x58\xa3\x61\xe9\xfc\x1d\xc6\x7e\x09\x7c\x77\x9c\x8f\x1d\x56\x95\x2a\x16\x9e\xd1\x4d\x69\xbb\x73\x19\x07\xf1\x63\x45\xce\xc9\xee\xd3\x43\x08\x29\x82\xa1\xf5\xee\xda\x63\xab\x69\xeb\xe4\x50\x43\xf2\x20\x75\x71\xf6\x33\x1e\xfe\x0f\x81\x7b\x59\x59\xe0\xe1\x14\xd8\x6e\xa7\x37\x89\x62\x42\x33\xbe\x87\x96\x5c\x41\x86\xb1\xa6\xf8\xd0\xd1\x67\xef\xee\x7f\xef\xee\xee\x84\xcf\x93\x29\x23\xef\xa1\x94\x2c\x3f\xb5\x93\x8a\x87\x5a\x08\x6d\x04\xc7\xc6\x06\xc3\x43\x0d\x67\x70\x90\xd2\x0f\xb3\x6f\xc1\x7f\x91\x02\xa2\x76\xd1\xe9\xfd\x16\xe6\xe2\xd2\x4e\x6a\x0d\xf7\x32\xee\x42\x8b\xb3\xd3\x90\x9c\xfc\xc5\xd9\xcf\xf9\x2c\x43\x85\x6a\x69\xe1\x42\xe6\x31\x9d\x38\xb4\x80\x50\xa1\xd2\xc1\x11\xf8\x82\x0c\x3a\x65\x01\x4d\x09\xe8\x3d\x39\xde\xb6\x09\xf2\x0c\x58\xa3\x32\xbe\x7b\xd2\xc4\xb6\xd5\xfd\xb3\x7f\x5e\xc3\xa4\xcb\xf1\x9b\x28\x61\xdb\x5a\x6d\xeb\x4d\x36\xc9\x03\xd4\x99\x07\x6d\x6b\x55\xa0\x86\x6d\x78\x9e\x68\x3b\x01\x9f\x65\xfa\xef\x94\x47\x4f\x09\x7b\xf0\x62\x45\xc5\x73\x1a\xf9\x37\xd4\xaa\x8c\xff\xf3\x20\x89\x5b\x11\x2c\x69\x85\x2f\xca\xba\xed\x03\xf4\xe1\x5b\x62\x19\xba\x9a\xde\x2b\xca\xeb\xd0\x16\xcb\x45\x67\x6d\xba\xce\xe2\x10\x71\x4f\x99\x2f\x54\x70\x67\xa2\x44\x83\x32\x7d\xbf\x38\xc9\xc6\x15\x61\x49\x2e\xcd\x19\x0f\xb6\x9b\x88\xdd\xff\xf5\x41\x16\x6c\xf1\x2c\xb3\xec\xe2\x21\x78\x2a\xe5\xa6\x7a\x32\x7f\x8b\xcc\xe4\x72\x8b\x4b\xcb\x54\xad\xd1\x9b\xad\x00\x0f\xeb\x95\xf5\xe3\x7a\x1a\xe4\x62\x25\x17\xeb\x15\x41\x4f\x91\xd7\xf2\xc6\x57\xa4\x3d\xf1\x1a\x55\x9a\xf6\x49\x35\x24\xa6\xcb\x40\x58\x52\x65\x1d\x41\x2c\xdd\x6d\xa1\x39\xf2\x41\xb3\x1f\x2a\x62\x89\x09\x46\x71\xaf\xaa\xf1\xb3\x1b\x98\xc5\x7f\x9a\x09\x9a\xf2\x47\xe7\xf9\xb3\xea\x36\xe8\x67\xda\x2c\x2d\xba\xb2\x7b\x3c\xe8\xa0\xaf\xfe\xbc\xfa\x37\x00\x00\xff\xff\x23\xca\x1f\x25\xc5\x1d\x00\x00")

func localesEnUsJsonBytes() ([]byte, error) {
	return bindataRead(
		_localesEnUsJson,
		"locales/en-US.json",
	)
}

func localesEnUsJson() (*asset, error) {
	bytes, err := localesEnUsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/en-US.json", size: 7621, mode: os.FileMode(509), modTime: time.Unix(1486064538, 0)}
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
	"locales/en-US.json": localesEnUsJson,
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
	"locales": &bintree{nil, map[string]*bintree{
		"en-US.json": &bintree{localesEnUsJson, map[string]*bintree{}},
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
