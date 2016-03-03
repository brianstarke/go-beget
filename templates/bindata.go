// Code generated by go-bindata.
// sources:
// templates/search_request.tmpl
// DO NOT EDIT!

package templates

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

var _templatesSearch_requestTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xdf\x6f\xdb\xb6\x13\x7f\x96\xfe\x8a\x83\xbe\x2d\x2a\x07\xfe\x4a\x7b\x0e\x90\x87\xae\x76\x8b\x0e\x59\x82\x25\xce\x80\xa1\x28\x6a\x4a\xa2\x6c\x2e\xb6\xe8\x92\xb4\x9b\x4c\xf0\xff\xbe\x3b\x92\xfa\x61\xc7\x4e\xe2\x6d\x4f\xa6\xee\xc8\xfb\xf1\xb9\xbb\xcf\x79\xc5\xf2\x7b\x36\xe3\xa0\x39\x53\xf9\x3c\x4c\xcf\xc2\xe0\xd3\xf8\x6a\x7c\xf3\x7e\x32\x1e\xc1\x87\xeb\xd1\x78\x08\x7f\x5c\xdf\xdd\xc0\x78\xf4\x79\x72\x0b\xef\x6f\xc6\xf0\xf1\x6e\xf2\xf9\x72\x1c\x86\xc1\x8c\x57\x5c\x31\xc3\x0b\x28\x95\x5c\x42\x5d\x27\x93\xc7\x15\xbf\x62\x4b\xbe\xdd\x86\xc1\x5a\x8b\x6a\x06\x73\x63\x56\xe7\x69\x3a\x13\x66\xbe\xce\x92\x5c\x2e\xd3\x4c\x09\x56\x69\xc3\xd4\x3d\x4f\x67\xf2\xff\x19\x9f\x71\x93\x7a\x53\x52\xa5\x2e\x0e\xae\xc2\xb3\x34\x0c\xc5\x72\x25\x95\x81\x38\x0c\xa2\x72\x69\x22\xfc\xd1\x46\xa1\x59\x1d\xa1\xfb\xe8\x25\xa3\x8d\xa9\x28\x1c\x84\x94\xd8\x4e\x80\x1f\x05\x5f\x14\x20\x34\x30\x28\xed\xf1\x07\x5a\x13\x15\x98\x39\xdf\xcd\x04\xd0\xe5\x3a\x37\xa1\x99\x33\x63\xef\x67\x0b\x0e\x46\x42\xc6\xf1\xe1\xc2\x70\x85\xf9\xcb\x6a\x08\x1a\x23\xf5\x47\xa9\x40\x71\xb3\x56\x15\x2f\x12\xca\xc3\xa0\x31\x38\xe4\xbd\x32\x61\x98\x4b\x0c\x1c\xe2\xba\x7e\x63\xbc\x1a\xce\x2f\xa0\x0f\x65\x5d\x2b\x56\x61\x89\xde\x88\x21\xbc\x79\xb0\xda\x5b\x9b\x1a\x85\x62\x2d\xe9\xed\xb6\xf7\x7e\xbb\xfd\x86\x5f\x0f\x89\x8f\xbf\xae\x45\x09\xfc\x3b\x3e\x87\x9f\xf6\xee\xb9\x30\x2e\x40\x48\xc3\xea\x9a\x57\x85\xf5\x66\x7f\x09\xb2\x14\x7e\x65\x4a\xcf\xd9\x62\xc2\x1f\x30\xf7\xe5\x6a\xc1\x97\xbc\x32\xda\x56\x55\x53\x59\xe5\x02\x23\x4b\xa4\x9a\xa5\xab\xfb\x59\xca\xab\x5c\x16\x58\x9e\xf4\x7f\xf4\xc0\xbf\xc5\x52\x96\xeb\x2a\x87\x58\x1f\x80\x60\xd0\xf7\x10\x0f\x20\xfe\xf2\x35\x7b\x34\x7c\x08\x5c\x29\xa9\x06\x50\x87\xc1\x86\x29\x28\x98\x61\xe0\x4a\x8f\x85\xd7\x58\xaa\x7c\x0e\x9a\xb4\xaf\xc2\x26\x67\x9a\xf0\x3f\x02\xd0\x79\x18\x04\xd6\xc1\x05\x44\x56\xfc\xcb\xed\xf5\x95\x53\x45\xe4\xc0\xc1\x12\x14\xbc\x64\xeb\x85\xa1\xdb\xae\xb8\x50\x89\xc5\x10\xb0\x2f\x93\x31\x05\x5b\xc6\xd1\x5d\xd5\x34\xc7\xd2\x65\x05\xd3\xb7\x9b\x29\x50\x57\x61\xbf\x60\x5e\x3a\xc2\x3e\x19\x84\x01\x9a\xf3\x36\x5c\xbe\x31\xf9\x1f\x0c\xc9\x62\xb8\xb5\xc0\xdf\x55\xcb\x7f\x01\x7d\xfb\xba\x07\xfe\xd9\x21\xf4\x77\xdc\xc4\x99\x0f\x67\xe0\xe0\x27\x7c\x11\x74\x02\xd5\x8f\x5d\x32\x51\x62\x19\xbb\x8f\x38\xc3\x80\xa7\xd1\x74\xd0\xab\x88\x51\xa7\xd5\xe4\x09\xde\x84\xee\x99\xc6\x4a\x1c\xad\xd6\xb3\x15\x79\xbe\x18\xef\xde\xea\x77\x34\x73\x12\x8e\x0c\x24\x55\xc7\xa8\x9d\xfa\x74\x15\x19\x65\xf6\x8e\x1d\x50\xa7\xd4\x96\x2c\x2a\x12\xc8\xd2\x9e\x1d\x93\xa0\x83\x35\x26\xe7\xc9\xe4\xf6\xb7\x4b\xf8\xbe\xe6\xea\xf1\xd9\x31\xe8\x59\xc7\x31\x70\x08\x5b\xfc\xff\xfb\x5e\xf7\x99\x39\xec\x47\xd9\xd3\x4e\xef\xb2\x8f\x22\x9b\xfc\x1e\x75\x3a\xa7\x37\x1c\xb3\x42\xe6\xc2\x2a\x88\x8a\x13\x8f\x6a\x6e\x08\x88\x15\x53\x78\x0d\x89\x51\x43\x29\x55\xe8\x58\x98\xd2\x29\xa9\xa5\xfa\x86\x12\x80\xcf\x06\x72\x56\x11\x99\x6a\x8e\x1c\xbe\x10\x7f\x21\x87\xb2\xaa\x40\x2b\x5a\xf3\x22\xcc\xb8\xf9\xc1\x79\x45\xda\x8d\xc8\xc9\x8d\x06\xea\x17\x4b\xb2\x88\xb2\x45\xbb\x59\x44\x18\x03\xa1\x8d\x7b\xc0\xd8\x69\x39\x42\xbe\xbb\xf1\x3b\x7a\xb7\x58\xfb\x7d\x91\xec\x5c\x70\xc0\x86\xc1\x47\x4b\xf6\x1a\x47\xe4\x80\x31\xa7\x84\xe9\x9f\x5a\x56\xe7\x91\xdb\x0b\x3a\x9a\x86\xc1\xb5\x2a\xb8\xfa\xf9\x71\x37\x80\x46\x18\x04\xfe\x81\x74\x02\x7a\xe0\xdc\xed\x7b\x71\x44\x1d\xb4\xf6\xe9\x0e\xde\x3e\x5a\x1c\x1f\x8f\x5f\x6f\xf6\xac\x57\x3c\x17\xa5\xc8\x09\xb0\xdd\x55\xfd\x0c\x48\xde\x4e\x87\x91\x0b\xe4\x40\x6c\xfd\xd0\x28\x8f\xdf\xd9\x62\x4d\x33\x80\xcf\x4b\x96\xf3\x7a\xdb\xdc\xd8\x90\xc2\x42\xb3\x72\x2b\x1f\x5a\xdc\x9d\xb7\x56\xde\x80\xe3\xbf\xe9\xcd\x07\x59\x15\xc2\x08\x59\xed\x3f\xea\x14\xfe\x55\xde\x08\x0e\xc3\xd4\x94\xc0\x22\x44\xab\x1b\x0a\xa1\x78\x6e\xc4\x06\x29\xc3\x6f\xfa\x13\x11\x6b\x4c\xfe\x03\xb0\x46\x5c\xe7\x38\x7c\x34\x24\x99\x94\x8b\x46\x5f\xa0\xd8\x87\x9f\x76\x3b\x00\x0e\xf7\x69\x87\x75\x43\x34\x6a\x8f\xf3\x77\xae\x0f\xe0\x13\x37\x13\x62\x8e\x27\x9c\xd3\x11\x44\xd2\x5e\x20\x86\xc0\x38\x5e\x6d\xd9\x4f\x0b\xda\xfd\xf2\x75\xaf\x54\xe4\xc2\x4f\x08\x71\xd8\x13\x7d\x8d\x7e\x02\xe2\x8a\x6f\xb8\x5d\xe9\x86\x63\x3d\xdd\xe8\x2d\x17\x7a\x0b\x76\x37\xed\x3d\x47\xa5\x43\xfe\x1c\xca\xc4\x1e\x86\x24\xb2\x0d\x49\x22\x7b\xb0\xa2\xa6\xd3\x48\xda\x9c\xad\xa2\xed\x26\xd2\xb4\x1f\xa4\xda\xb6\xae\x69\x4b\xb1\xd5\x0a\xcb\x16\x7b\xc1\xd0\x8f\x9a\x5d\x20\x2d\x8c\x5e\x79\x12\x7a\xbe\x95\xa8\x2a\x4d\x72\x2d\x91\xb4\x86\xf7\x55\x94\xb8\xcf\x5b\xb7\xc2\x0e\x80\xae\xc7\x76\xf4\x9d\x78\x48\x61\x9f\x12\xe5\xa5\x58\x0a\xfa\xd3\x86\x8d\xd7\x0f\x4b\x25\x56\x71\x5a\xc2\x65\x89\xdb\xe3\x90\x2d\xa7\x79\xbd\xb1\xf7\x45\xe1\xfa\x20\x2e\x8f\x4c\xdf\x10\x36\xfb\xdc\x84\xdb\xe4\x05\x32\x1a\x42\xfe\x22\xf5\x0c\xa8\xb1\xa9\x23\x8f\x12\x69\xaf\x44\x65\x53\x17\xdf\x97\x9b\xa6\x2b\xbb\xa6\x94\xbd\x96\xec\x75\x64\xde\xeb\x47\x6c\xc7\xde\x5c\xb4\x0d\xd9\xc9\xb0\x27\x07\xaf\x07\xef\xb6\x6b\xbd\xe3\xe8\x09\xbd\x47\x57\xf6\x6f\x7a\xd7\x52\x70\x71\x90\x17\x0f\xe5\xde\xef\xc9\xbe\xd9\x93\x5b\xd1\x6d\x4d\xc7\x36\x2d\x8f\xb9\x35\xe9\x39\xc6\x4a\x9f\xa3\x16\x7b\xd7\x31\x8b\x3d\xf6\xa6\x9b\xbe\xf1\x41\xb2\xf3\x27\x6d\x7f\xca\xed\xbf\x84\x6d\xf8\x77\x00\x00\x00\xff\xff\xa8\x06\xd3\x6f\x50\x0f\x00\x00")

func templatesSearch_requestTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSearch_requestTmpl,
		"templates/search_request.tmpl",
	)
}

func templatesSearch_requestTmpl() (*asset, error) {
	bytes, err := templatesSearch_requestTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/search_request.tmpl", size: 3920, mode: os.FileMode(420), modTime: time.Unix(1457029286, 0)}
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
	"templates/search_request.tmpl": templatesSearch_requestTmpl,
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
		"search_request.tmpl": &bintree{templatesSearch_requestTmpl, map[string]*bintree{}},
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

