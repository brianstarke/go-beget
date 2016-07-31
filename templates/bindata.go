// Code generated by go-bindata.
// sources:
// templates/create.tmpl
// templates/enums.tmpl
// templates/fields.tmpl
// templates/search.tmpl
// templates/update.tmpl
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

var _templatesCreateTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x92\xc1\x6e\xdb\x30\x0c\x86\xcf\xd1\x53\x10\x46\x30\xc8\x45\x26\x5f\x86\x1d\x0a\xf4\xb0\xc6\x6e\x11\x20\x48\xbb\x24\x5b\x0f\xeb\x0e\xb2\xcd\x38\x42\x65\x29\x95\xa5\x2d\x85\x91\x77\x1f\x6d\x27\xdb\xd2\x6d\xd8\x45\x89\x68\xf2\x27\xff\x8f\xda\xc9\xe2\x49\x56\x08\x6d\x2b\xee\x87\xbf\x0b\x59\xe3\xe1\xc0\x58\x92\xc0\x6d\xb6\xc8\x96\x1f\xd6\x59\x0a\xd3\xbb\x34\x9b\x40\x96\xce\xd6\x2b\x78\x98\xcd\xe7\x70\x9d\xc1\xfc\x6e\xb5\xa6\xac\x2e\xb1\x42\x83\x4e\x7a\x2c\x61\xe3\x6c\xfd\x87\x58\x42\x81\xf5\xcb\xee\x24\x4d\x05\xa1\x51\xa6\x82\xad\xf7\xbb\xcb\x24\xa9\x94\xdf\x86\x5c\x14\xb6\x4e\x72\xa7\xa4\x69\xbc\x74\x4f\x98\x54\xf6\x6d\x8e\x15\xfa\xa4\x3f\x19\x53\xf5\xce\x3a\x0f\x9c\x8d\xa2\x52\x7a\x99\xcb\x06\x93\xe6\x59\x47\x8c\x02\xbf\x49\x68\x69\x0c\xc5\x83\x72\x0e\xe9\x63\xcc\x58\xdb\x8e\x4d\xa8\x6f\x14\xea\xb2\x81\xcb\x2b\xd0\x68\x40\xcc\x4c\x83\xce\x0f\xc1\xce\xee\x37\xe9\x40\xf5\xb1\xa9\xd5\xa1\x36\xcd\xd9\xcc\x70\x05\x5f\xbe\x36\xde\xd1\xd4\x2d\x1b\xb5\xad\x93\x86\xa0\x8d\xd5\x04\xc6\xfb\x4e\xf2\x95\x5c\x44\x2d\xf7\x22\xcd\x87\xe2\x68\xd2\x95\xa0\x29\xa9\xcf\x00\x76\xea\x90\x68\x9d\x77\x18\x9a\x37\x20\xe1\x75\x1c\xbc\x05\xbf\x45\x38\xb9\x66\x9b\x60\x0a\xe0\x05\x5c\x9c\x65\xc6\x7f\x93\xe5\x65\x0e\x17\x44\x49\xa4\xd7\x31\x70\x65\xfc\xfb\x77\x13\x40\xe7\xac\x8b\x81\x9c\x3c\x07\x74\x2f\x9d\x81\x13\xb0\xa3\x11\x1e\x75\x2a\x32\xd7\x47\x99\x28\x16\x6c\x34\x3a\x92\xe1\xff\xe6\x24\x84\xe8\x33\x3f\x4b\x1d\xb0\xe1\xff\x05\x55\x88\x9e\xd4\x50\xdc\xb6\x6a\x03\xda\x53\x36\xfc\x5a\xd8\xe1\x30\x39\xb2\x3b\xfe\xf4\xfa\xab\xb0\xd9\xa8\x3d\x8f\x96\xd9\xfa\xd3\x72\x31\x5b\xdc\xc2\x63\xa4\xca\xc7\x68\x98\x73\x19\xcc\x03\x3d\x08\xf2\xde\x5f\xef\xb5\x2c\x70\x6b\x75\x89\xee\xc6\xba\x5a\x7a\xfe\xd3\x6d\x6a\xb5\x96\x8e\xde\xc8\xa8\xdf\x7f\x09\x3d\x21\xba\x12\xa2\x6e\xdc\x9e\x8f\xf8\xd8\x9d\x4b\xfb\x9d\xc7\x62\x55\x48\xc3\xdf\xa8\xb2\x2b\x71\xe8\x83\x33\x54\xd5\x13\xa5\xd5\xfe\x08\x00\x00\xff\xff\x8f\xc7\x71\x50\x4d\x03\x00\x00")

func templatesCreateTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesCreateTmpl,
		"templates/create.tmpl",
	)
}

func templatesCreateTmpl() (*asset, error) {
	bytes, err := templatesCreateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/create.tmpl", size: 845, mode: os.FileMode(420), modTime: time.Unix(1469994694, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesEnumsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x55\x4d\x6f\xe3\x36\x10\x3d\x9b\xbf\x62\xa0\x62\xb1\xf2\xc2\xb5\xee\x01\x72\x48\x37\xc2\x22\x80\x6b\xb7\x89\x17\x3d\x14\x05\x42\x49\xb4\x44\x44\x22\x65\x92\xda\x34\x08\xf2\xdf\x3b\x43\xea\xd3\x49\x7b\x69\x0e\x31\x39\xf3\xe6\x69\x38\x7c\x33\x6c\x79\xfe\xc4\x4b\x01\xaf\xaf\xdb\xdf\xc2\x72\xcf\x1b\xf1\xf6\xc6\x58\x92\xc0\xb7\x74\x9f\xde\xdf\x1c\xd3\x5b\xf8\x7a\xb8\x4d\x37\x90\xde\xde\x1d\x1f\xe0\x8f\xbb\xdd\x0e\x7e\x49\x61\x77\x78\x38\x22\x8a\x80\xa5\x50\xc2\x70\x27\x0a\x38\x19\xdd\xbc\x23\x4b\xd0\x70\x7c\x69\x07\x6a\x0c\xe8\xac\x54\x25\x54\xce\xb5\x57\x49\x52\x4a\x57\x75\xd9\x36\xd7\x4d\x92\x19\xc9\x95\x75\xdc\x3c\x89\xa4\xd4\x3f\x67\xa2\x14\x2e\xf1\xff\x19\x93\x4d\xab\x8d\x83\x98\xad\xa2\x53\xe3\x22\xfc\xb1\xce\x20\x8d\x8d\xd8\xda\xa7\x7b\x68\x29\x09\x6d\x40\x5a\xe0\x0a\xf4\xb0\x75\x1a\x32\x81\x9f\xc4\xf4\xa4\x02\x0e\x0f\x82\x9b\xbc\xba\x17\xe7\x4e\x58\x07\x27\x59\x3b\x61\x98\xc3\xfc\x66\x0c\xca\x79\xca\xaf\x9a\xb2\x51\x8e\x08\x0b\x70\x95\x90\x06\xc2\x57\x21\xd7\x9d\xc2\xc0\x96\x1b\xf4\xc6\x03\xfb\xc3\xef\xbb\xa1\x1a\x52\xab\x35\xcb\x89\x80\x72\x4e\xcf\x13\xfb\x35\x48\xed\x38\x5b\xed\xb5\x4b\xcf\x6c\xb5\x93\x4f\xc2\x6f\xc2\xe2\x9b\x11\x58\x4a\x73\xac\xb8\x5a\x6c\x0e\xc6\x83\x85\xb5\x83\x73\x5a\x07\xdf\x9d\xdd\x77\x75\xed\xa9\xfc\x22\x94\xe5\x57\x6e\x6c\xc5\xeb\xa3\xf8\xdb\x01\xd6\xb0\x16\x8d\xa0\x03\x51\xf1\x2d\x55\x5f\xd7\x5c\x95\x5b\x6d\xca\xa4\x7d\x2a\x13\xa1\x72\x5d\xe0\xf9\x92\x9f\x28\xa0\x8f\xc5\xfa\x9c\x3a\x95\x43\xac\xc7\x43\xac\xe7\xbc\xf1\x1a\xe2\x3f\xff\xca\x5e\x9c\xd8\x80\x30\x86\xbc\xaf\x6c\xf5\x83\x1b\x28\xb8\xe3\x7d\xc5\x18\x5b\xd9\x67\xe9\xf2\x0a\x34\x79\x73\x6e\x05\xa4\xe7\x2b\xb6\x5a\x79\xd0\x35\x44\xe2\x1c\xf5\x76\x5f\x99\xb9\x4b\x91\x61\xf0\x52\x9d\xe6\xce\x1a\xf7\xb3\xc8\x4b\xb7\x0a\xa6\x01\x31\x2b\xe9\x1c\x55\x4e\xe6\x0f\x90\x54\xe0\x7f\x41\x93\x6b\xcc\x6c\xbc\x91\x45\x7e\xa3\xf5\x3d\xee\x92\xb8\x5e\x78\x06\x7c\xb8\xd9\x39\x4e\x7a\xcb\xec\xd4\x97\x00\x15\x4c\x11\x96\xbd\x10\x27\xde\xd5\x8e\xbc\x46\xb8\xce\x28\x50\xb2\xde\x00\x76\xd1\x36\xa5\xdb\x3a\xc5\xd1\x77\xc5\xb3\x5a\x50\xab\x34\xe1\x5a\xe1\xf1\xd3\x8f\x47\x52\x34\x75\x0f\x5e\xac\x8d\x36\xa0\xd7\x6c\xf5\xc6\x06\x8e\x70\xe1\x31\x7d\x6f\xbd\x21\x46\x16\xa6\xc6\x77\xd5\xfc\x0f\xc5\x8d\xd1\x33\xcd\x7d\x99\x44\xb7\x20\x8f\xb3\x3e\x89\x75\x50\x1d\xc9\x0a\xb5\x06\x57\xd7\xbd\xe4\xec\xf6\x68\x64\x13\x87\x4d\x9c\x61\x9a\x8f\xd1\xe3\x7a\x12\x22\x81\x07\x29\x92\xfa\xa8\x42\x5f\x34\x56\x8f\xba\x29\x58\x83\xf0\x46\x47\xdf\xb2\xc1\xe7\x75\x37\xba\x42\xf7\x8e\x51\xbb\x85\x73\xec\xee\xe0\x9f\x8b\x6d\xc4\x2c\x1a\xff\x1d\xce\xcb\xe1\x23\x6c\xe8\xfd\x3e\xa3\x49\x69\x53\x5e\xb3\x89\x71\x89\x5a\x92\x5e\xce\x93\x80\xee\x95\x36\xa2\x86\x29\x33\x9e\x74\xe9\x1e\x86\xcf\x47\xb2\xfb\x6f\xc5\x7d\xfe\x64\x3f\xd3\xe4\xd5\xb0\x18\xc6\x28\x3c\xbc\xa7\x85\xf4\x26\xb1\xe1\x80\x2e\x24\x8d\xda\x7e\xe8\x4b\xdc\xe6\x7e\xc8\xea\x13\x3c\x57\x02\x47\x36\x2e\x0d\x60\x9a\x38\xf9\xc3\xac\x07\x5b\xe9\xae\x2e\x28\x1c\xdf\x05\x5e\x14\x38\xba\xb9\x8f\xbe\xd9\xdf\x12\x18\x57\x87\x7b\xca\x0c\xa3\xc1\xfa\xd7\x02\x5a\x23\x3c\xb3\xb0\xe1\xa9\x98\x3e\x3c\x4c\xb7\xe1\xb9\xb0\x70\x42\x8e\x17\xdd\x19\xc8\x3a\xe7\xa6\x27\xe0\x06\x9f\x90\x29\x0c\x5b\x14\x3f\x87\x0d\x7c\x30\x30\xfb\x43\xf3\xe1\x9e\x5e\xb4\x7f\x02\x00\x00\xff\xff\xa2\x18\x65\x62\x9a\x07\x00\x00")

func templatesEnumsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesEnumsTmpl,
		"templates/enums.tmpl",
	)
}

func templatesEnumsTmpl() (*asset, error) {
	bytes, err := templatesEnumsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/enums.tmpl", size: 1946, mode: os.FileMode(420), modTime: time.Unix(1469813078, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFieldsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x54\x4d\x6f\x9b\x4c\x10\x3e\x9b\x5f\x31\xe2\x4d\x14\x88\xfc\x9a\x9e\x23\xe5\xd0\xd6\x4e\x95\xca\x4d\xda\xda\x51\x0f\x55\x25\x2f\x78\xf8\x50\x60\x97\xec\x0e\xa9\x23\xe4\xff\xde\xd9\x05\x57\x38\x75\xac\x4a\xa9\x64\x99\x05\xf6\x99\x79\x3e\x76\xa8\x45\x72\x2f\x32\x84\xb6\x9d\x7c\xee\x96\x37\xa2\xc2\xed\xd6\xf3\xa2\x08\x3e\xcc\x6e\x66\x5f\xdf\x2e\x67\x53\x78\x7f\x3b\x9d\x8d\x61\x36\xbd\x5e\x2e\xe0\xdb\xf5\x7c\x0e\xef\x66\x30\xbf\x5d\x2c\x79\x97\xdd\x98\xa1\x44\x2d\x08\xd7\x90\x6a\x55\xfd\x51\x2c\xe2\x07\xcb\xa7\x7a\x57\x9a\x01\x8d\x29\x64\x06\x39\x51\x7d\x11\x45\x59\x41\x79\x13\x4f\x12\x55\x45\xb1\x2e\x84\x34\x24\xf4\x3d\x46\x99\xfa\x3f\xc6\x0c\x29\x72\xff\x9e\x57\x54\xb5\xd2\x04\x81\x37\xf2\xd3\x8a\x7c\xbe\x18\xd2\x5c\xc6\xf8\x5e\xe8\xe8\xee\x75\xb9\x2a\xb0\x5c\x43\x61\x40\x40\xea\x96\x3f\xb9\x4b\x21\x81\x72\xdc\xdf\x08\x5c\xa5\x49\xc8\x16\xa0\x5c\x90\x83\xc4\x25\x02\x29\x88\x91\xb1\x25\xa1\x66\x61\x4a\x8e\xc1\x70\xff\x7e\xa9\x34\x68\xa4\x46\x4b\x5c\x4f\x3c\xe2\x62\x07\xbb\x4b\x72\xc4\xae\x44\xb3\x01\x94\x4d\x75\xc6\x06\x31\x32\xc7\xb2\x4e\x9b\x52\xa2\x31\x5e\xa2\x58\x2f\x04\x6d\x7b\x42\x3d\x1a\x2e\x2e\x61\x68\x57\xdb\x6a\x21\x39\xa2\x93\x62\x0c\x27\x1b\xf7\xd6\x95\x37\xdb\xed\x00\xe5\x6e\x36\x93\x5e\x53\xdb\x16\x29\xe0\x03\x63\xe0\xcd\xb3\x6d\x1d\xb5\x4b\x28\x14\x89\xb6\x45\xb9\x76\x2d\xdc\x95\x6d\xb4\x74\x3f\x2e\x6e\x6f\x40\x5a\x26\x8e\x9c\x90\xf4\x4a\x9a\xc9\x90\x5a\x17\x1a\x13\xf0\xdd\x53\xdb\xac\x7b\xe3\x0f\x68\x30\x8b\x69\xec\xe0\xae\x55\x67\xb5\x71\xe1\x39\x62\x2a\x75\xeb\x2e\x59\x4e\xaa\x31\x08\x7d\xb8\x8b\x2f\x73\x78\x68\x50\x3f\x79\x69\x23\x13\x08\xcc\x81\x64\xc2\x61\xf5\x20\xdc\x71\x6a\xbd\x91\xe1\x63\x92\xe4\x60\xec\xfa\x98\x22\x61\x6c\xe2\x87\xdd\xbf\xf0\x46\xa3\x8e\x71\xaf\x71\x1a\xef\x14\x8e\x76\x8e\x8f\xf8\xb7\xdb\xe3\x7b\xdd\xb8\x7d\x12\xda\xe4\xa2\x5c\xe2\x86\x4f\x61\x55\x97\x58\x21\x5b\xef\xa6\xc4\xd8\x31\x51\x25\xd3\x99\x28\x9d\x45\xf5\x7d\x16\xa1\x4c\xd4\x9a\x59\x47\xff\x59\x40\x8f\x45\x7d\x54\xf5\xa0\x03\xab\x0e\xbe\xff\x88\x9f\x08\xc7\x80\x5a\x2b\x1d\x5a\xc9\x8f\x42\xc3\x5a\x90\xe8\x1d\xf1\xfe\x9d\x21\xae\xea\x25\x0c\x8f\xc2\xc0\x8d\x35\xa6\xa2\x29\x69\xe0\x9c\x2c\xca\x31\xf0\x9c\x4f\x66\x96\x5c\x1a\xf8\x77\x72\x37\x96\x55\xa7\x02\x56\xa7\x8f\x2b\x97\x3a\x4f\x2a\xeb\x30\x3e\x4f\x68\x38\x74\xb6\xd3\x17\xd8\xd6\xe1\xd8\x56\xec\x8d\xbe\x93\xd5\x2b\xac\xfe\x8d\x1e\x98\x7d\x7e\xc8\xed\xbd\x36\x41\xdc\xd3\x09\x3b\xbb\xdd\x61\x23\x6d\x4d\xec\x3f\x63\x93\xa5\x2e\xaa\xa0\xbb\x09\x62\x26\xbc\xf2\x57\xe1\x20\x01\xd2\x7f\x93\x41\xf2\xcc\xf7\x73\xc3\xae\xbf\x14\xcc\xd1\x04\x8e\x9b\x7f\x76\x6a\xce\xec\x07\x4e\xc1\x0b\x5f\x3f\x9b\x06\xe9\xbd\x3c\xba\x04\x7e\x05\x00\x00\xff\xff\xe9\x72\xf3\x6a\x73\x06\x00\x00")

func templatesFieldsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesFieldsTmpl,
		"templates/fields.tmpl",
	)
}

func templatesFieldsTmpl() (*asset, error) {
	bytes, err := templatesFieldsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/fields.tmpl", size: 1651, mode: os.FileMode(420), modTime: time.Unix(1469983969, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSearchTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x59\x6d\x6f\xdb\x46\xf2\x7f\x6d\x7f\x8a\x29\x81\x7f\xfe\xa4\xab\xa3\x8c\x5e\xae\x2f\x84\x2a\x87\xd8\x52\x12\x5d\x55\xc9\xb5\x94\xcb\x8b\x20\xb8\x50\xe4\x52\x62\x43\x71\xe5\x5d\xca\x76\xea\xfa\xbb\xdf\xcc\xec\xf2\xd1\xb2\x2a\x1b\x41\x81\x3b\x1c\x1a\x71\x77\x66\x76\x7e\xf3\xbc\xeb\x4d\x10\x7e\x09\x96\x02\xee\xee\xfc\x0b\xf3\x73\x12\xac\xc5\xfd\xfd\xf1\x71\xb7\x0b\x6f\x87\x93\xe1\xe5\xeb\xf9\x70\x00\xe7\xd3\xc1\xb0\x03\xc3\xc1\x68\x3e\x83\x0f\xa3\xf1\x18\xce\x86\x30\x9e\xce\xe6\x48\x45\x84\x4b\x91\x09\x15\xe4\x22\x82\x58\xc9\xf5\x03\x61\x5d\x5c\x98\x7f\xdd\x14\xa2\x91\x61\xab\x93\x6c\x09\xab\x3c\xdf\xf4\xba\xdd\x65\x92\xaf\xb6\x0b\x3f\x94\xeb\xee\x42\x25\x41\xa6\xf3\x40\x7d\x11\xdd\xa5\xfc\xdb\x42\x2c\x45\xde\xe5\xff\x1e\x1f\x27\xeb\x8d\x54\x39\xb8\xc7\x47\x4e\x14\xe4\xc1\x22\xd0\xa2\xab\xaf\x52\x07\xbf\x45\x16\xca\x08\x25\x76\x7f\xd3\x32\xa3\x85\x78\x9d\xd3\x3f\x89\xec\x26\x72\x9b\x27\x4c\x94\xa1\x28\x3a\x91\x7e\xeb\x5c\x21\xb9\x76\x8e\xf1\x77\xed\xf8\xdf\xd6\x32\x51\x32\x23\xb1\xb7\xb4\xa7\xaf\xa0\xbe\x9d\x06\x19\xed\x6d\x13\xa5\x04\x8a\xf4\xd8\x48\x0d\x6c\x33\x11\xa8\x70\x75\x29\xae\xb6\x42\xe7\x90\x68\x08\x40\x0b\xc4\x94\x26\xbf\x07\x8b\x54\x40\x73\x3f\x96\xaa\xc9\xee\x1f\xe7\xf8\x7b\x9f\x48\x54\x7c\x1b\xe6\x70\x77\x7c\xf4\x26\x11\x69\xa4\xe1\xe3\xa7\x06\x35\xaf\xc2\x67\xb2\x43\xcf\x89\x99\xc4\xf9\x4c\xc4\x69\x2e\xd4\x0e\x6a\x5a\xae\xc8\x99\x88\xe8\xa7\x2a\x12\xea\xec\x2b\x9c\x34\xc8\x8b\x55\x4b\x2f\xcd\x27\xd1\x8f\x93\x75\x82\x78\xb3\xbc\xd8\x4b\x69\x81\x25\xc5\xb1\x16\x8d\x2d\xc9\x2b\xb8\x77\xff\xd0\x7e\x56\x21\x36\x9c\x51\x07\xf4\x46\x84\x49\x9c\x84\x90\xcb\x26\xed\x0e\x5b\x59\xf6\x96\x91\xda\x34\x6d\x0b\x91\x9a\xff\x0e\xd2\xad\x20\x2d\x85\x8a\x83\x50\xdc\xdd\x17\x14\xd7\xb4\xc1\x40\x36\x14\xe2\xe8\xb1\xf2\x47\x81\xc7\x7e\x13\xd1\xb9\xcc\xa2\x24\x4f\x64\x06\xd5\x2f\x4b\x16\x16\x0b\xbb\x91\x17\xb6\x35\x31\x43\x81\x1e\x25\x4a\x84\x79\x72\x2d\x20\x5f\x05\x1c\x4c\x4f\x31\x45\x21\xef\x19\xb6\x18\xe9\x81\xd0\xa1\xc8\x28\xa3\x60\x21\x65\x5a\x50\x24\xb5\x8d\x12\xc5\xeb\x28\xaa\x7b\x6d\x25\x52\xb4\x07\xac\x45\xbe\x92\x11\x29\x1a\x44\x51\xe5\x4c\xfa\x6e\x26\x81\x0f\x80\x5a\x46\x22\x0e\xb6\x69\xde\x21\x81\xf9\x4a\x40\x61\x53\x92\x39\xfc\x15\x82\x2c\xe2\xe5\xd2\x86\xb4\xfe\x7a\x32\x40\xe6\x51\x0c\x5f\xe5\x16\x6e\x02\x0c\x30\x94\x2e\xaf\x85\x52\x49\x64\x6c\xd6\xb3\xe5\x09\xae\x03\x0c\x0a\xd5\x3c\x98\x37\xb4\xf2\x4b\xfd\x5d\x74\x51\x8a\x6e\xec\x80\xa3\x44\x84\xff\xc4\xdb\x2c\x74\xe3\x56\x0e\x18\x5a\xef\x8e\xd9\x01\x62\xbf\x0c\x87\x3e\xbc\x9d\x97\xab\x95\xfb\xfb\x30\xbd\xe4\xe5\x7b\xcf\xaa\x43\x68\x50\xe7\xff\x57\x02\x82\x34\xc5\x02\x91\xfb\xc7\x74\x16\xb8\xa8\xe3\xc9\xe3\xc9\xef\x55\xb6\x76\xe3\x47\x9c\xd9\x81\xeb\x76\x24\x77\x20\x8c\x97\xe0\xfb\xfe\x3e\x3c\xde\xbe\x83\x29\x7a\x62\xe8\xf5\x77\x25\x1b\x6e\x99\xc8\xea\x01\xeb\xd4\xc1\xef\xc2\x24\x3d\x80\xe1\x15\x2d\x94\xd6\xe8\xc1\xeb\x8c\x49\x38\xdf\x7a\xc0\xc6\x5a\xe7\xfe\x6c\x83\xf5\x38\x77\x59\x77\x0f\xf7\x31\xb4\x8e\xa8\x3a\xfe\x07\x95\xa7\x83\x55\x90\x61\x93\x22\x1c\x74\x5e\xe8\xbe\x88\x3d\x43\x84\x0e\x2c\x6a\x5b\x1f\x82\xcd\x06\x83\xd3\xad\xd6\xd0\x87\x58\xa4\x8f\x94\xc8\xb7\x2a\x43\x67\xd7\x43\x96\xab\x27\x06\xa7\x36\x7a\x6b\x8a\x1e\x8a\x31\x2d\x52\x4c\x3b\x8c\xac\x77\x8d\x48\x26\x6d\xc2\x55\x90\x64\xb0\xd8\x26\x29\xa7\x46\xb0\xbf\xfc\x3f\xc9\xa9\xa4\x81\x6b\x15\x41\x57\x3d\x74\xec\x9f\x3a\x88\x51\x33\x7f\xd3\x10\xb4\xd4\xb1\x18\x51\xf2\x2e\x7b\xcc\x44\x6e\x0a\x38\x86\xa2\x66\x23\xf0\xe7\x41\x36\xd8\x63\x80\x03\xf1\x17\xa7\xbb\x69\xd1\x44\x0e\xc1\x6a\x14\xee\x03\x33\xed\xc6\x64\x5b\x4f\x09\xca\x7c\xff\x75\xa8\xcc\x79\xae\x2c\x3b\xe0\x21\xb8\xac\xd2\x7d\x30\x6c\xbb\x91\x5d\xd0\xc8\xd6\xc2\xc5\x65\xc5\x58\x85\x46\xa3\x08\xb0\xf6\xd0\xee\x86\x68\xb3\xed\x7a\x81\x88\x89\x84\xbf\x75\xf2\xbb\x38\x34\x3e\xed\x79\x2e\x31\x4e\x58\x4e\x87\x85\xcc\x50\xc6\x41\xa8\x4a\x00\x3e\xe6\x6e\x65\x97\x4a\x1e\x9c\x94\x02\x3d\x4b\x63\x22\xa2\x5c\xad\x39\xb5\xe8\x6b\x25\x7a\x5a\xc0\x1e\x72\x98\x5f\x9f\xeb\x4a\x73\xea\x9e\xb2\x9b\xb4\x9b\xe6\x41\xce\xb6\x60\xfa\xf0\x62\x57\xf3\xae\xd5\x56\xfb\xbf\xb2\xc4\xd6\x7b\x74\xaf\x71\xb8\xad\x9e\xed\xa8\x79\x6b\x47\xf4\x19\x97\xb7\xd9\xaf\x63\xb8\x49\xb0\xf9\x14\x93\x3b\x86\x06\x88\x5b\x11\x6e\x73\x33\xa7\xe2\x3e\x8e\xe1\xb9\x58\x8b\xcc\x44\x96\x95\x47\x16\xc7\x4d\x92\x68\x66\x68\x70\x6f\x70\x40\x86\x4d\x8a\xbd\x66\x25\x53\xd4\x5b\x7b\xcc\x80\x53\x4c\x9a\x84\xd8\xca\x63\x66\xe2\xd2\xae\x0f\x8d\xb9\x07\xda\xba\x1e\x32\xf1\x81\x1d\x1c\x62\x1b\xdd\x0d\x1b\xbe\x54\x1e\x59\x94\xdb\xbc\xd5\x0b\x4d\x90\xc4\x90\x8a\xac\xaa\x82\x1e\xf4\xfb\x70\xca\x1d\x84\x8a\xa4\x73\x82\x37\x81\x7b\x10\xa9\x16\xbc\x66\xfb\x4d\x5c\xf5\x9b\xaa\xa4\xd2\x3e\x33\x69\xf8\x1e\xbb\xfb\x60\xc1\xeb\xa4\x3a\x2a\xf6\x3d\xd0\xcc\x80\x14\xf7\x56\xb2\xbd\x5d\xf8\x73\x95\xac\x2f\x93\xe5\x2a\x77\xb1\x02\x13\x8d\x6d\x5a\x1b\xbc\x5f\xd0\x29\xfa\xca\x9f\x15\x46\x3e\xa3\x18\x15\x9c\x22\x17\x95\x2d\xdf\x48\xb5\x0e\x90\xfd\xca\x1f\xc8\x34\x0d\x94\x4d\x0f\x32\x8a\xab\xf9\xe3\x0d\x5e\xb8\x5c\x87\x6c\x49\x8e\x33\xc6\x74\x3c\xd3\x19\xb1\xbb\x5d\xe0\x28\x93\x84\x78\x86\x76\x5f\xd0\xb1\x9e\xb1\x4b\x59\x40\x5f\x59\x83\xb0\x4a\x7d\xa0\x7f\xcc\x8e\xbb\x45\x23\xff\xf8\xd2\x2d\x28\x3d\xab\xbb\x61\xb6\x25\x67\x17\xb7\xcd\xee\x8a\xdd\x2c\xb4\xf8\x6d\xe0\x7f\xd7\x87\x2c\x49\x59\x06\x39\x4f\x9a\x81\xd5\xb8\xaf\x49\xea\x37\xa6\x52\xf6\x87\xe4\xcc\xa9\xc6\x87\xd8\x75\xfe\x4f\xc3\x60\x38\x3b\x47\x5b\xd7\x38\xd9\x57\x4d\x9f\x79\xe4\xad\xca\xf5\x56\xd4\x7e\x1e\x76\x70\x1b\xac\xad\x0b\xc8\xef\x35\xf2\x8e\x77\xe7\x72\x76\x95\xba\x5e\x2b\x03\xcf\xe5\x36\xfb\x9f\x49\xc0\x42\xd9\x83\xf2\xef\xdb\xc4\xb5\x73\x3e\x7d\x3f\x99\xbb\x27\x38\x19\xcd\x20\xcc\x72\xe7\xb9\x71\xfe\x88\x2b\x0e\x33\x40\x53\x26\x23\x3b\x21\x5c\xac\xa3\x05\x55\x16\x1d\x34\xb7\x26\xd4\x38\xdc\x9a\x05\xa9\xf8\x7b\xaa\xaa\x39\xb6\x5d\x57\xcc\xd0\x4a\xd1\x27\xac\xcd\x86\xb7\x1b\x45\x09\xc3\xed\x8a\x3f\x62\x0f\xf9\x7c\x9e\x95\x3d\x9b\x11\xf5\x0b\x47\x1f\x46\x7a\xb2\xc5\x20\xfa\xe3\x8f\xd6\xfa\x44\xe6\xbc\xc1\xc1\x2d\xe0\x51\xf1\x65\x50\xb3\xe4\xda\xa5\x05\x6f\x2d\xca\xa6\x46\x7d\xb4\x96\x34\x53\x8b\x76\xf6\x30\xfc\x92\x86\xbe\x0a\xa2\x32\xe7\xa9\x10\x4b\x8a\xcc\xa2\x64\x9c\xd4\xd3\xe8\xc3\x4a\x28\xc1\xfb\x0d\x06\x92\xb4\x97\x83\x09\x88\xa5\xee\xd6\x7d\x5e\xad\x61\xdf\x75\x99\xf1\x8a\x9c\xb2\x7e\x2d\xaf\xa1\x65\x51\xd2\x98\x6c\xe1\xaa\x6e\x6d\x24\x0d\x71\xde\xc2\x9b\x4e\x0f\x95\x94\xd5\x6d\xd0\xe9\x3b\x76\x0b\xbd\xf1\x70\xf7\xbb\x72\x7b\x9c\x7c\x11\xb4\x6b\xe3\xb5\x51\xd0\x90\x20\x07\x17\xcb\x5a\xa0\xe9\x16\x8b\xf3\x8c\x72\x7f\x7c\x89\xd7\xb5\xd1\x78\xf4\xf3\x10\xfe\x49\xd7\x54\x53\xac\xbc\xea\xb0\x67\x09\x9c\x4c\xe7\x8f\x0b\x7d\xab\x04\xe6\x81\x9a\xaf\x82\xac\x8d\xe3\x95\xf3\x90\x66\xaa\x1e\xe2\x7d\x55\xe1\x15\x5a\xef\x96\xf5\xd3\x43\x92\x5d\xa2\x7e\x2a\x45\x99\xf8\x7f\x0c\x2b\xa2\x1c\xcd\x60\xf2\x7e\x3c\xde\x69\xa6\x03\x58\xd1\x26\x0f\xd8\xed\x5b\xc5\x2e\x6f\xd7\x1a\x40\x5b\x1a\xfe\xbf\x66\xd7\x66\x67\xe9\x94\x81\x56\xf4\x8a\x21\x77\x02\x61\xa2\xd7\x34\x8a\x3c\xf8\x22\xa8\xa6\x63\xf0\x0f\xce\xe8\x25\x24\xa3\xd7\x21\xcc\x55\xaa\xf5\xa6\x73\xd0\xab\x07\x3d\x13\x19\x2e\xf5\xa4\x21\xb7\x71\xa2\x1b\x2d\xa8\xda\xd1\x49\x1d\x94\xa3\x11\xae\x86\x93\xd6\xe3\xa1\x67\x2a\x3f\x25\x40\xad\xbb\x95\x0d\xea\x08\xf9\x67\xb9\xb2\x0f\x12\x9a\xfb\x04\x17\x3a\xe5\xef\x98\xee\x4c\xd2\x13\x49\x6d\x24\xb0\xa6\xc4\x55\x63\x5a\x3c\xe6\xfd\x66\xa9\x02\x7a\xde\x91\x64\x89\xdb\x9a\x1d\xd0\x31\x8b\x5b\x53\x49\xd3\x5b\x7f\x22\x6e\x06\x0b\x84\x81\x43\xd7\x46\xea\x7c\x89\x20\x9c\x5a\x4f\x40\x52\x5b\xca\x5d\x0b\xaf\x03\x4d\x7d\xf9\xaa\xdc\x70\x06\xb7\xc2\x27\xf8\x02\x2f\x1b\xc3\xf1\xf0\x7c\x0e\xdc\xcf\x9e\xe6\x06\x3e\xab\xe6\x05\xec\xbe\x18\x49\x7f\xff\xa1\xde\x6d\x9f\x69\xf4\xaa\xa3\xef\xb3\xf9\x69\xe7\x9b\x9a\x9d\x8a\x69\x11\x47\x1f\x3f\x55\x6f\x92\x47\xc6\xa8\x8c\x0d\x3e\x47\x8b\x9e\x43\x0d\xff\xb3\x39\x97\x34\xeb\xd7\x5d\xf5\x62\x9f\xaf\x0e\x04\x63\x97\xac\xa4\x8f\xa7\x9f\x7c\x56\xa1\x43\x3c\xd6\xdf\x08\x62\x87\x8b\xde\xa1\x6f\x53\x9c\x5c\xc8\x8d\x46\x88\xa6\x99\xed\xdd\x7c\x7e\x01\x2b\xb3\xc7\x2f\x86\x74\xef\xdc\xff\x30\xa4\x7d\x18\xe5\xa5\x8c\x1f\x4e\x4f\xcb\xf7\x4d\xab\x15\xb1\x63\x30\xe9\x6d\x18\x62\x1d\xec\xc0\xcb\xd3\x97\x80\xe0\x32\x89\x93\x20\x5c\x4c\x67\x73\x5a\x3a\x25\x92\x45\x10\xc1\xbf\x66\xd3\x49\x07\xfe\x61\x16\x82\xec\x2b\x73\xa3\x34\x65\x62\xc5\x4e\x7c\x7f\x0a\xaa\x11\x6d\xf4\x77\x12\xbf\x8e\xb8\xba\xc6\xf3\x2b\xe2\x8d\xa1\xb8\x14\x7a\x23\x33\x2d\x3e\xa8\x24\xa7\x67\x01\x0c\x6e\xbb\x6e\xc3\xf9\xce\x8c\x16\xca\xff\xc5\xdc\xc9\xd1\x37\x0e\x01\x70\xcc\xe0\x70\xe3\x33\xe7\x3b\x81\xb1\xa5\x5c\x66\xa5\xc1\x71\xab\xb1\x3a\xbf\x41\xb7\x50\xb9\x2d\xbc\x58\x4c\x2a\x8b\x32\xa8\xcd\x9f\x76\xf0\xb0\x20\x7a\x9d\xa6\xae\xf2\xcf\x64\xf4\xb5\x98\x93\x5a\x91\xf0\xf8\x51\x67\x41\x54\xa8\x4b\x74\x54\xb7\xdf\x98\xd7\xc8\x1b\x8c\xe2\x89\x6c\x95\x53\x8c\x55\xeb\x2a\xf3\xb9\xc0\x33\x9d\x1d\x6a\xda\x27\xe7\x3d\xef\x48\x47\x36\xc0\xe9\x6d\xdd\x7f\x9f\xad\x03\xa5\x57\x41\xea\x22\xbe\x17\x5a\x7d\x3b\x18\x31\xe3\x18\x72\xa9\x8e\x44\xf1\xd7\x28\x1a\x72\x9a\xc0\x7a\xd8\xa2\x1c\xb6\xad\xcf\xc4\xe6\x92\xb4\x03\x55\x95\xc7\xcd\xbf\x41\x14\x78\xb0\xd6\xb4\x9b\x09\x42\xb2\x5c\x4f\xc6\x35\xa2\xab\x46\x16\xa4\x33\xa1\xae\x85\x62\xcd\xf6\x00\x34\xb5\xf7\x99\xe0\xc8\x11\x97\x45\x85\xb1\x31\xc6\xce\xf9\xc5\xba\xe6\xaf\xc1\xb0\xd3\x45\x7c\xf0\x41\x28\x1e\xd5\x62\xfa\x33\x51\xb7\xce\x34\x83\xad\x5b\x83\xee\x31\xbc\x42\x26\x8d\xd6\xff\x0d\x00\x00\xff\xff\x5d\x44\x41\x4e\x24\x1e\x00\x00")

func templatesSearchTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSearchTmpl,
		"templates/search.tmpl",
	)
}

func templatesSearchTmpl() (*asset, error) {
	bytes, err := templatesSearchTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/search.tmpl", size: 7716, mode: os.FileMode(420), modTime: time.Unix(1469993477, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUpdateTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xdf\x6f\xdb\x38\x12\x7e\xb6\xfe\x8a\x39\x1f\xae\x90\x0a\x9f\x1c\x1c\x7a\xfb\x60\x20\x0f\x4d\xed\xb4\xd9\x4d\x93\x6c\xed\xa2\x0f\x41\xb0\xa5\xac\xb1\xcd\x46\x26\x15\x92\x8a\x9b\x0d\xf2\xbf\xef\x0c\x29\xd9\x72\x1a\x3b\xd9\x62\xf7\x61\x1b\x53\xf3\xf3\xfb\x3e\xce\xb0\x14\xd3\x6b\x31\x47\xb8\xbf\x4f\x2f\xc2\x9f\x67\x62\x89\x0f\x0f\x51\xd4\xef\xc3\xfb\xd1\xd9\xe8\xd3\xdb\xc9\x68\x08\xef\xce\x87\xa3\x1e\x8c\x86\x27\x93\x31\x7c\x39\x39\x3d\x85\xa3\x11\x9c\x9e\x8f\x27\x64\xc5\x86\x73\x54\x68\x84\xc3\x1c\x66\x46\x2f\x7f\x08\xd6\xa7\x83\xc9\x5d\xd9\x84\x26\x87\xca\x4a\x35\x87\x85\x73\xe5\xa0\xdf\x9f\x4b\xb7\xa8\xb2\x74\xaa\x97\xfd\xcc\x48\xa1\xac\x13\xe6\x1a\xfb\x73\xfd\xdf\x0c\xe7\xe8\xfa\xfe\xff\x51\x24\x97\xa5\x36\x0e\xe2\xa8\xd3\xcd\x85\x13\x99\xb0\xd8\xb7\x37\x45\x97\x7e\xa3\x9a\xea\x9c\x22\xf6\xbf\x59\xad\xf8\x60\xb6\x74\xfc\x8f\xd4\x7d\xa9\x2b\x27\xbd\x91\xa2\x50\x9c\xb1\x1b\x45\x1d\x7b\x03\xdd\x56\xda\x42\x28\x45\xb1\x2a\x69\x0c\x92\x6d\xe2\xbb\xdf\x2a\xfa\x73\x49\x39\xf1\x13\xde\x54\x68\x1d\xe4\x38\x93\x0a\x2d\x08\xb0\xe8\x40\xcf\xa0\x14\x86\xcc\x1c\x1a\x0b\x33\x6d\x7c\x87\xec\xc0\x4d\x6e\x85\x49\x01\x4e\x1c\x4c\x85\x82\x0c\xc9\x97\xba\x2d\xe4\x9f\x04\x9b\x50\x39\xc5\xb0\x16\x73\xf6\xcd\xd0\xad\x10\x15\x1b\xdc\xca\x29\xe7\xb1\xf0\xeb\xf8\xfc\xac\x07\xda\x10\x74\x64\xef\xf4\x1a\x73\x2a\x62\xfc\xfb\x29\x10\x68\x0e\x97\xa8\x5c\x1a\x39\x4a\xb7\xaf\x7a\xeb\x4c\x35\x75\x70\x1f\x75\x4e\x86\x20\x95\xfb\xe5\x0d\x7c\x65\xe0\x06\x5d\x99\x77\xbf\x46\x9d\x60\x6d\x61\x29\xca\xcb\xad\x30\xc7\x12\x8b\xfc\x8a\x3c\xd0\xcc\xc4\x14\xef\x1f\x1a\xbf\x2a\x78\x90\x73\x10\xce\x19\xae\xf6\xe4\x9f\x1a\xf4\xf1\x05\x28\x5c\x05\xa0\x10\x4c\xfd\x91\xe0\x83\xc9\x82\x80\x8b\x66\x95\x9a\x42\x5c\xc1\xeb\xdd\xa1\x92\xfd\x99\x62\x99\x87\xfe\x92\x7d\x41\x18\x08\x83\xae\x32\x0a\x5e\xed\xb6\x22\x23\x82\x6b\x00\x32\xef\x45\x9d\x87\xba\xcf\xb7\x79\x1e\x8c\x40\x72\x3b\x53\xad\x6e\x51\x49\x52\x23\x02\xa9\x61\xa1\x73\xdf\x8e\xc8\x59\x9b\x75\xa3\x96\xb9\x73\x0b\xb2\xaf\x3b\x26\x49\xbc\x5b\x08\xa9\x44\x56\x60\xfa\xb2\xa6\xd7\x69\xe3\x19\x33\x12\xf0\xf2\xe4\xf4\xe0\x56\x14\x15\x42\x8b\xa3\x67\x7b\xaf\xd2\x9a\xf0\x4b\x1f\xed\x0a\x0e\x43\x90\x68\x0d\x4b\x55\xb7\xfb\xbe\x96\x5c\xb0\x67\xd5\x05\x03\xea\x69\x81\xb5\x0a\x0d\xb7\xea\xe5\x5c\x50\xfa\x85\x2e\x72\x34\x21\x9e\xf5\xf2\x25\xe9\xa3\x31\x74\x47\x5e\xd4\xe9\x0f\x19\xe3\x04\xe2\x90\xa4\x07\x97\x6d\x29\xf6\x42\xd8\x84\x3b\xa2\xa9\x00\x83\x43\xb0\x37\x29\x91\x36\x6e\x2e\xc6\x51\x25\xb9\x18\x3e\xbb\xd8\xd4\x76\xac\xcd\x52\xb8\x98\x6c\x87\xba\x28\x84\x49\xf8\x7b\x8d\x6e\x97\x4b\x63\x5e\x42\x6d\x5d\x9a\x0b\x1d\x26\x74\xd6\x86\x9a\x12\x19\xa1\x68\x84\xae\x71\xe4\x12\x7c\x0d\x5c\x42\x91\x8e\xd1\x05\x9e\xd2\x61\xe6\x49\xe2\x68\x71\x52\xfb\x27\xac\xa6\xa8\x65\xfe\x65\x81\x06\xb9\x9e\xd1\xcd\x3d\x5f\xc8\x01\x05\x3e\x19\x3e\x24\x1b\x3a\xd8\x6a\xa2\xc7\x37\x45\x9c\xd4\xc4\x8c\xbe\xe3\xb4\x6a\x50\x82\x95\x2c\x0a\x70\xe2\x9a\x47\x03\xdb\x0e\x8f\x58\x99\x0a\xa7\x4e\x6a\xe5\xb9\xc1\x60\x1f\x74\x18\x74\xf9\x32\x3e\xb6\x12\xc5\x79\x06\xaf\x43\x82\x24\xa0\x5f\x83\xdf\x5b\x13\x4e\xa7\x0c\x50\x95\x3e\x41\x24\x35\x24\x67\xde\xe2\x5f\x87\xa0\x64\xe1\x61\xab\x5b\xa4\xd3\x80\xcb\x1f\x21\xc6\x21\xe4\x59\xca\xc9\xe3\x56\xf8\x16\x24\x6c\xbf\x6f\xf4\x7c\xa0\xae\x0b\x22\x9b\x7b\x6c\x34\x4b\x4a\xfc\x30\x99\x5c\xc0\x22\x7c\x03\x0f\x40\x3d\xbd\x77\x63\x60\x53\x9e\xe0\x4d\x8c\xff\x1d\x1c\x78\x40\x59\xff\x06\x6d\x55\x38\xcb\xee\x04\xb3\xad\xa6\x34\xb9\x09\x82\x37\x07\x6f\x80\xfa\x54\xda\x11\x1d\x17\xb4\x33\xf9\xe8\x80\x4d\x32\x91\xd7\x63\xfd\xff\xe1\x40\xa8\x3b\xef\x4d\xd1\x4c\x00\xb4\x1e\x08\xcf\x36\xb5\x45\x05\x2f\xb9\xb4\xdd\xf1\x66\xc0\x71\xb4\x78\x15\x2c\x3e\xa1\x2d\xb5\xb2\xf8\xc5\x48\xba\x45\x3d\x30\xf0\xba\x3e\xaf\xc9\x66\x3e\xa8\x72\x93\x7e\x0c\xb3\x8c\x68\xea\x72\x03\x5d\xff\xa5\xb3\x4a\xbd\xe7\x07\x14\x74\x8d\x62\xef\xca\x77\xad\xb2\x67\xda\x1d\xeb\x4a\xe5\x09\x5b\x85\xc4\xf4\x17\xb3\xd9\xc9\xd6\x92\x08\x7b\x99\x92\x89\xfc\x6d\x51\xc4\x26\x3d\xd2\xf9\x1d\x33\xfa\x84\x28\x76\xa7\x3a\x12\x79\x53\x2e\xdb\xd1\xda\x4f\x8f\x4b\x9a\x0e\x2e\x5e\xf5\xa0\x7b\xa6\x1f\x6f\x17\xa9\x6a\xaa\xc2\xcf\x8c\x72\x76\x9f\x28\xf3\x56\xd0\x9e\x35\x7b\x54\x40\x36\x41\x96\xbc\xff\xd2\xcf\x6a\x29\x8c\x5d\x88\x22\xa6\xfe\x5e\x55\xe6\x9f\x6b\x63\xe6\xfb\x18\xf9\xbb\x95\x63\xf3\x60\xd8\x6c\x93\xa6\x93\x01\xfc\xc7\x76\x3d\xb6\xa9\x37\x8e\x93\x27\xba\x0a\x15\x57\x26\x7d\x7c\x8b\xff\x76\xbd\x27\x3c\x77\x95\x28\xc6\xf4\x40\x41\xe3\x33\xee\x29\x3c\x4c\x9b\x9f\x2c\x7a\x67\x0d\xe7\xbf\xb1\xf5\xe3\x8c\xfa\xda\x4f\xe9\x75\x90\x66\x53\x7f\x0c\x04\x4d\xf0\x3b\x89\x60\x59\x16\x7e\x23\x58\x7f\x11\x2c\x3f\x41\x35\x3d\x01\xe7\xa9\x36\xf3\x7e\x79\x3d\xef\xaf\xdf\x93\xff\x66\x87\xda\x17\x37\x3b\x6b\xdf\x88\x6c\x65\xe2\x65\x75\x79\x95\xdd\x39\xdc\xda\x4e\x7e\x7d\x49\x5a\x09\x39\x5f\x84\x25\x0d\xea\x98\x9f\x5a\xe1\xbc\xbd\xd3\x9a\x7d\x73\x8d\x77\xcf\x6d\x9b\xf5\xbd\x22\xdb\x74\xab\x84\x1d\xdc\xd6\x03\x81\x7e\xf7\xc2\xb4\x0d\x70\xb7\x8a\xab\x0b\x8a\x6d\xb2\x79\x12\x78\xa3\x30\xe8\x42\xed\x4f\x95\xcd\xf1\xd7\x4f\xc2\x01\xb4\x42\xf6\xf8\x8b\xdf\x69\xfe\x3f\x5e\x6c\xbd\x26\xa6\x2f\xc7\x5f\xa7\xba\xfc\x38\xe4\x69\x76\xdc\xfa\x92\xfd\x14\x87\x6b\x6f\x7c\xe1\xcb\x63\x2b\x5d\x9c\x41\xe0\xb1\xb5\xe7\x78\x46\xc8\x1d\x08\x50\x43\x35\x1d\x4f\xcc\x07\xf9\xb2\xcd\xc7\xe0\x10\xee\xf2\x92\x01\xbb\x4a\xe3\xf0\x90\x8d\x42\xe2\x6f\xd0\x12\xce\xfe\x37\xfa\x3e\x11\x51\xf0\x86\x27\xca\xb0\x43\x84\xbe\x3a\x4e\x5a\xc2\x8f\x99\xd6\x63\xa5\x4c\xb7\x11\x0b\x78\xc5\x94\x35\x79\x46\x82\x8d\xfa\x3a\x9d\x6f\x97\xe5\x23\xa9\x6d\x54\x4e\x50\x6e\x64\x42\x21\x48\x15\x7f\x05\x00\x00\xff\xff\xfd\x2a\x8e\x8b\xbc\x0e\x00\x00")

func templatesUpdateTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesUpdateTmpl,
		"templates/update.tmpl",
	)
}

func templatesUpdateTmpl() (*asset, error) {
	bytes, err := templatesUpdateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/update.tmpl", size: 3772, mode: os.FileMode(420), modTime: time.Unix(1469993571, 0)}
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
	"templates/create.tmpl": templatesCreateTmpl,
	"templates/enums.tmpl": templatesEnumsTmpl,
	"templates/fields.tmpl": templatesFieldsTmpl,
	"templates/search.tmpl": templatesSearchTmpl,
	"templates/update.tmpl": templatesUpdateTmpl,
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
		"create.tmpl": &bintree{templatesCreateTmpl, map[string]*bintree{}},
		"enums.tmpl": &bintree{templatesEnumsTmpl, map[string]*bintree{}},
		"fields.tmpl": &bintree{templatesFieldsTmpl, map[string]*bintree{}},
		"search.tmpl": &bintree{templatesSearchTmpl, map[string]*bintree{}},
		"update.tmpl": &bintree{templatesUpdateTmpl, map[string]*bintree{}},
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

