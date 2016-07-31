// Code generated by go-bindata.
// sources:
// templates/create.tmpl
// templates/search.tmpl
// templates/searchEnums.tmpl
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

var _templatesCreateTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\x3f\x6f\xdb\x30\x10\xc5\x67\xf3\x53\x1c\x84\xa0\xa0\x02\x97\x5a\x8a\x0e\x01\x32\x34\x96\x12\x18\x30\x9c\xd4\x76\x9b\xa1\xe9\x40\x49\x67\x99\x08\x45\x3a\x14\xd9\x3a\x10\xfc\xdd\x7b\x92\xec\xc0\xee\xbf\x85\x12\x4f\xef\x1e\xdf\xfd\xc4\xad\x2c\x9e\x65\x85\xd0\xb6\xe2\x61\x78\x9d\xcb\x1a\xf7\x7b\xc6\x92\x04\xee\xb2\x79\xb6\xf8\xb4\xca\x52\x98\xdc\xa7\xd9\x18\xb2\x74\xba\x5a\xc2\xe3\x74\x36\x83\x9b\x0c\x66\xf7\xcb\x15\xa9\x3a\x61\x85\x06\x9d\xf4\x58\xc2\xda\xd9\xfa\x0f\xb3\x84\x0a\xab\xd7\xed\xd1\x9a\x1a\x42\xa3\x4c\x05\x1b\xef\xb7\x57\x49\x52\x29\xbf\x09\xb9\x28\x6c\x9d\xe4\x4e\x49\xd3\x78\xe9\x9e\x31\xa9\xec\xfb\x1c\x2b\xf4\x49\xbf\x32\xa6\xea\xad\x75\x1e\x38\x1b\x45\xa5\xf4\x32\x97\x0d\x26\xcd\x8b\x8e\x18\x15\x4e\x2c\xb4\x34\x86\xea\x41\x39\x87\xf4\x31\x66\xac\x6d\x2f\x4c\xa8\x6f\x15\xea\xb2\x81\xab\x6b\xd0\x68\x40\x0c\xdb\x6e\xd0\x1f\xd2\x81\x32\x0d\x3a\x3f\xb1\x3a\xd4\xa6\x39\x4b\x0b\xd7\xf0\xed\xd4\x60\xbf\xff\xde\x78\x47\xe9\x5b\x36\x6a\x5b\x27\x0d\xc1\xbb\x50\x63\xb8\xd8\x75\xd6\x6f\xb6\x11\xf5\xec\x44\x9a\x0f\x26\xd1\xb8\x13\xa3\x29\xe9\xbc\x01\xed\xc4\x21\xf1\x3a\x3f\x69\x08\xd1\x80\x84\xdf\xeb\xe0\x2d\xf8\x0d\xc2\x71\x6e\xb6\x0e\xa6\x00\x5e\xc0\xe5\x99\x32\xfe\x9b\x2d\x2f\x73\xb8\x24\x4e\x22\xbd\x89\x81\x2b\xe3\x3f\x7e\x18\x03\x3a\x67\x5d\x0c\x34\xc3\x4b\x40\xf7\xda\x45\x3f\x22\x13\xd3\x3e\x06\x8f\x3a\x17\x99\xeb\x83\x4d\x14\x0b\x36\x1a\x1d\x08\xf1\x7f\xf3\x12\x42\xf4\xca\xaf\x52\x07\x6c\xf8\x7f\x10\x15\xa2\x67\x34\xb4\xb5\xad\x5a\x83\xf6\xa4\x83\x53\xd6\xe3\x03\xb5\xc3\xa3\x77\x5e\x86\xf5\x5a\xed\x78\xb4\xc8\x56\x5f\x16\xf3\xe9\xfc\x0e\x9e\x22\x55\x3e\x45\x43\xc2\x45\x30\x8f\x74\x19\x68\xea\x7e\xfb\xa0\x65\x81\x1b\xab\x4b\x74\xb7\xd6\xd5\xd2\xf3\xb7\x39\x53\xab\xb5\x74\x74\x3f\x46\xfd\x0d\x28\xa1\x67\x43\x5b\x82\xd3\x05\xed\xc9\x88\xcf\xdd\xba\xb0\x3f\x79\x2c\x96\x85\x34\xfc\x9d\x2a\xbb\x16\x87\x3e\x38\x43\x5d\x3d\x4b\xfa\xa9\xbf\x02\x00\x00\xff\xff\x9a\xf9\xcc\x9f\x49\x03\x00\x00")

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

	info := bindataFileInfo{name: "templates/create.tmpl", size: 841, mode: os.FileMode(420), modTime: time.Unix(1469977610, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSearchTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x1a\x6b\x73\xdb\xc6\xf1\xb3\xf8\x2b\x36\x68\x62\x03\x0a\x0b\x6a\x52\x37\x1f\x38\x61\x3a\xb6\x44\xd9\x6a\x14\x4a\x11\xe9\xfa\x83\xc7\x53\x81\xc0\x81\x44\x0c\xe2\xa8\x3b\x50\x8f\x30\xfa\xef\xdd\xdd\x3b\xbc\x28\x88\xa1\x14\x4f\x66\x9a\xd1\x84\xc0\xdd\xbe\x77\x6f\x1f\x07\x2f\x83\xf0\x73\x30\x13\xb0\x5e\xfb\xe7\xe6\x71\x14\x2c\xc4\xfd\x7d\xa7\xd3\xeb\xc1\xdb\xe1\x68\x78\xf1\x7a\x32\x3c\x82\xc3\xb3\xa3\x61\x17\x86\x47\x27\x93\x31\x7c\x38\x39\x3d\x85\x37\x43\x38\x3d\x1b\x4f\x10\x8a\x00\x67\x22\x13\x2a\xc8\x45\x04\xb1\x92\x8b\x07\xc4\x7a\xb8\x30\xb9\x5b\x16\xa4\x11\x61\xa5\x93\x6c\x06\xf3\x3c\x5f\xf6\x7b\xbd\x59\x92\xcf\x57\x53\x3f\x94\x8b\xde\x54\x25\x41\xa6\xf3\x40\x7d\x16\xbd\x99\xfc\xfb\x54\xcc\x44\xde\xe3\xff\x77\x3a\xc9\x62\x29\x55\x0e\x6e\x67\xcf\x89\x82\x3c\x98\x06\x5a\xf4\xf4\x55\xea\xe0\xbb\xc8\x42\x19\x21\xc5\xde\xaf\x5a\x66\xb4\x10\x2f\x72\xfa\x49\x64\x2f\x91\xab\x3c\x61\xa0\x0c\x49\x11\x47\x7a\xd6\xb9\x42\x70\xed\x74\xf0\xb9\xc6\xfe\xd7\x85\x4c\x94\xcc\x88\xec\x2d\xed\xe9\x2b\xa8\x6f\xa7\x41\x46\x7b\xab\x44\x29\x81\x24\x3d\x36\x52\x43\xb7\xb1\x08\x54\x38\xbf\x10\x57\x2b\xa1\x73\x48\x34\x04\xa0\x05\xea\x94\x26\xbf\x05\xd3\x54\x40\x73\x3f\x96\xaa\x89\xee\x77\x72\x7c\xde\x46\x12\x05\x5f\x85\x39\xac\x3b\x7b\xc7\x89\x48\x23\x0d\x1f\x3f\x35\xa0\x79\x15\x2e\xc9\x0e\x7d\x27\x66\x10\xe7\x92\x80\xd3\x5c\xa8\x16\x68\x5a\xae\xc0\x19\x88\xe0\xcf\x54\x24\xd4\x9b\x3b\xd8\x6f\x80\x17\xab\x16\x5e\x9a\x57\x82\x3f\x4d\x16\x09\xea\x9b\xe5\xc5\x5e\x4a\x0b\x4c\x29\x8e\xb5\x68\x6c\x49\x5e\xc1\xbd\xfb\x87\xf6\x33\xe2\xb3\xdd\x58\x78\xb8\x41\xeb\x27\x19\xe4\xf3\x0d\xab\x58\x43\x10\x81\x7c\x1e\x18\x53\x93\x81\x73\x09\x53\x01\x46\x13\x0c\x47\x99\x75\x41\x63\xd4\xd8\x47\xb4\xb7\x12\xf9\x4a\x65\x22\x6a\x33\xb5\xe5\x9e\xe5\x2c\xd8\x71\xb0\xba\x05\x91\xad\x16\x2f\x23\xf6\xd4\x5c\xa4\xcb\x78\x95\x66\x42\xeb\x4e\x28\x31\x4a\xc1\x5d\xaf\xbf\xce\x2d\x36\xf4\x07\x50\x0f\xf2\xf5\x5a\x05\x19\x1e\xac\xaf\x93\x2e\x7c\x7d\xcb\xbb\xc6\x63\xf7\xf7\x35\x2c\x7e\xb9\xf5\xad\x4e\xeb\x75\x12\x83\xb8\x42\x1c\x38\xd8\x00\x33\xa2\x0d\x20\x91\x79\xb0\x5e\x8b\x2c\x62\x16\xfc\x8b\x61\x48\xe2\xfe\x7b\x7c\x36\x82\x8c\x24\x61\xe1\x82\x2c\xff\x93\x62\x86\x75\xd1\xcc\x81\x41\x01\x1c\x5e\x25\x66\x66\xc7\xa9\x89\x81\x52\x1c\x4d\x19\x9d\x59\x19\x53\x6b\x76\x1e\x0b\x26\x63\x7e\x36\x9e\x45\x4f\xad\xb4\x00\xeb\xdc\xf1\x2f\xa7\x80\x11\xae\xee\x3a\xf1\x2a\x0b\xc1\xd5\x2d\x9e\xf1\xea\xd4\x5d\xaf\x90\x09\xcf\x82\xc6\x30\x09\xe7\xa0\xe9\x79\x9b\x46\x98\x32\xe0\x31\xeb\xf7\x3b\x7b\x7b\x46\x62\xab\xe3\xd1\xb4\xd0\x70\xaf\xb0\xf8\x1e\xfe\x15\x30\x8e\x8d\xdf\x9f\x03\xa5\xe7\x41\x3a\x11\xb7\x18\x85\x8b\x65\x2a\x16\x02\x4d\xcf\xb9\x4d\x53\x72\x93\x98\x35\x66\xbe\x54\xb3\xde\xf2\xf3\xac\x57\x66\xaa\xbf\x11\x82\xc5\x15\x6a\xab\xd6\x35\x0e\xa8\xb5\xfb\xf1\xd3\xf4\x2e\x17\x5d\x10\x4a\x49\xe5\x91\xca\xd7\x81\x02\x4a\x89\xd6\x22\x9d\x2f\x67\x10\xa6\x3a\x80\x7a\x28\xd4\xac\x11\x89\x38\x58\xa5\x79\xcd\x72\x59\x92\x76\x01\x93\xaf\x3f\x24\xe1\x62\xd7\x79\x9f\x15\xc7\x72\x61\xb4\x80\xcb\x6f\xae\x2f\xd9\xeb\x78\x52\x51\x0f\xed\xe0\x09\xf5\xea\x96\x35\xfa\xb9\xc4\xda\xeb\x12\x45\x6b\xe8\xf7\xd9\xe2\x4f\x98\xba\xc4\xae\x19\x7b\xbf\xcd\xda\x0d\x36\xee\xd4\x8a\xe3\x19\x73\x73\xb0\xe5\x8a\x8c\x68\x4b\x88\x3f\x51\xc9\xc2\x35\x2f\xee\x14\x05\xbe\x74\x2e\xbd\x9a\x07\x72\xb5\x8b\x0f\xc2\x0d\xbb\xef\x6b\xb4\xfa\x63\x8e\xd9\xea\x81\xed\xc6\x7f\xf9\x8d\x7e\x49\x09\x4e\xc2\x23\xd9\x8f\xbc\x91\xab\x86\x3f\x2a\x0f\xb4\xd5\x0e\x9b\xab\xf9\x59\x2f\x45\x98\xc4\x49\x48\x1c\x9b\x25\xbf\x8d\x9b\x41\x69\xd6\xb3\xb6\x7c\x5c\x2f\x66\x54\x51\xfe\x13\xa4\x2b\xca\x1b\x88\x1e\x07\xa1\x58\xdf\x17\x10\xd7\xb4\xc1\x35\x67\x49\xdd\x08\x7a\xab\x7c\x28\x4a\x8f\x7d\x27\xa0\x43\x99\x45\x49\x9e\xc8\x0c\xaa\x27\x0b\x16\x16\x0b\xed\x45\xaa\x28\x83\xa6\xbc\x53\x4f\x12\x25\x4a\x84\x79\x72\x2d\xca\x62\xf4\x14\x53\x14\xf4\x9e\x61\x8b\x13\x7d\x24\x74\x88\xc1\x40\x89\x70\x2a\x65\x5a\x40\x24\xb5\x8d\x52\x8b\xd7\x51\x54\xf7\x1a\x15\x34\x7c\x5e\x88\x7c\x2e\x39\x1f\x07\x51\x54\x39\x93\xde\x9b\xfd\x8a\x0f\x80\x52\xda\x90\xeb\x9a\xd2\x8b\x49\xbd\xb0\x31\xd2\x1c\xfe\x02\x41\x16\xf1\x72\x69\x43\x5a\x7f\x3d\x3a\x42\xe4\x93\x18\xee\xe4\x0a\x6e\xb0\x34\x11\x75\x79\x8d\x47\x2a\x89\x8c\xcd\xfa\xb6\x93\x04\xca\x65\x5a\x35\x19\xf3\x86\x56\x7e\x29\xbf\x8b\x2e\x4a\xd1\x8d\x5d\x70\xb0\xca\xe3\x0f\x9d\x68\x37\x7e\x70\xa0\x09\xd6\x5b\x33\x3a\x40\xec\x97\xe1\x30\x80\xb7\x93\x72\xb5\x72\xff\x00\xce\x2e\x78\x19\x8b\x99\x11\x87\xb4\x41\x99\x5f\x2a\x01\x41\x9a\x62\x2f\x97\xfb\x45\xf6\x50\x1b\xdc\x1a\x12\x7b\x95\xad\xdd\xf8\x11\x67\x76\xe1\x7a\x33\x92\xbb\x10\xc6\x33\xf0\x7d\x7f\x9b\x3e\xde\x36\xc6\x14\x3d\x31\xe5\x97\x16\x4c\xdc\x32\x91\xd5\x37\x25\xb8\x8b\xef\x85\x49\xfa\x00\xc3\x2b\x5a\x28\xad\xd1\x87\xd7\x19\x83\xf0\x79\xeb\x03\x1b\x0b\x93\xcb\x78\x89\xa9\x2e\x77\x59\x76\xaf\x4b\x89\x02\x39\xa2\x4d\xff\x8b\xc2\x13\x63\x93\xe8\x48\x0f\xe2\x17\xba\x2f\x62\xcf\x00\xa1\x03\x8b\x36\x74\x00\xc1\x72\x89\xc1\xe9\x56\x6b\xe8\x43\x4a\x9b\x36\xe7\x68\x55\x0f\x59\x6e\x74\x31\x38\xb5\x91\x5b\x53\xf4\x50\x8c\x69\x91\xe2\xb1\xc3\xc8\x7a\xd7\x88\x64\x92\x26\x9c\x07\x58\x63\xa6\xab\x24\xe5\xa3\x11\x6c\xef\xd4\x9f\xe4\x54\x92\xc0\xb5\x82\xa0\xab\xda\x6a\xc8\x1f\x38\x88\xb5\x66\xfc\xa6\x21\x68\xa9\x6b\x75\x44\xca\x6d\xf6\x18\x8b\xdc\xf4\xda\x18\x8a\xa6\xb3\xe2\xd7\x9d\x6c\xb0\xc5\x00\x3b\xea\x5f\x70\x77\xd3\xa2\xdf\xdf\x45\x57\x23\xf0\x00\x18\xa9\x5d\x27\x3b\x25\x94\x4a\x99\xf7\xbf\x4e\x2b\xc3\xcf\x95\xe5\xb0\xb2\x8b\x5e\x56\xe8\x01\x18\xb4\x76\xcd\xce\x69\xba\xde\xd0\x8b\xd3\x8a\xb1\x0a\x4d\xb1\x34\x9b\xf0\xee\x92\x60\x71\xe0\x98\xa2\xc6\x04\xc2\xef\x3a\xf9\x4d\xec\x1a\x9f\x96\x9f\x4b\x88\x23\xa6\xd3\x65\x22\x63\xa4\xb1\x93\x56\xa5\x02\x3e\x9e\xdd\xca\x2e\x15\x3d\xd8\x2f\x09\x7a\x16\xc6\x44\x44\xb9\x5a\x73\x6a\x51\xd7\x4a\xed\x69\x01\x6b\xc8\x6e\x7e\x7d\xae\x2b\x0d\xd7\x2d\x69\x37\xd9\x2c\x9a\x3b\x39\xdb\x2a\x33\x80\x17\x6d\xc5\xbb\x96\x5b\xed\x7f\x65\x8a\xad\xd7\xe8\x7e\x83\xb9\xcd\x9e\x9b\x51\xf3\xd6\xde\xa6\x8c\x39\xbd\xd1\x64\x74\x93\x60\xf1\x29\x2e\x59\x30\x34\x40\xdc\x8a\x70\x95\x9b\x2b\x05\xdc\xc7\x71\x2f\xe7\x66\x98\xc3\xc6\xd2\xb3\x63\x15\x51\xb4\x93\x92\x4b\xd3\x34\x2c\x53\xac\x35\x73\x99\xa2\xdc\xda\x63\x04\xec\x62\xd2\x24\x2c\xe7\x33\x4e\xed\x7a\xd7\x98\x7b\x20\x2d\xcd\x28\x86\x61\x17\x9b\xe7\x46\x75\x6b\x8e\x2c\xba\x9a\x57\x70\xee\x4d\x45\x56\x65\x41\x0f\x06\x03\x38\xe0\x0a\x42\x49\xd2\xd9\xc7\x21\xec\x1e\x44\x4a\xe3\x0a\xae\xd9\x7a\x13\x57\xf5\xa6\x4a\xa9\xb4\xcf\x48\x1a\xbe\xc5\xea\xde\x1c\x19\xbf\x05\xea\x19\x10\xe2\xde\x52\xae\x77\xf1\x17\xc9\x6c\x9e\xbb\x98\x81\x09\xc6\x16\xad\xa5\xbe\x4a\xb9\xdb\xbf\xf2\xc7\x85\x91\xdf\x50\x8c\x0a\x3e\x22\xe7\x95\x2d\x8f\xa5\x5a\x04\x88\x7e\xe5\x1f\xc9\x34\x0d\x94\x3d\x1e\x64\x14\x57\xf3\xcb\xb1\x92\x0b\xd7\x21\x5b\x92\xe3\xec\x70\xe9\x99\xca\x88\xd5\xed\x1c\x5b\x99\x24\x44\x1e\xda\x7d\x41\x6c\x3d\x63\x97\x32\x81\xfe\x68\x0d\xc2\x22\x0d\x80\x7e\xcc\x8e\xbb\x42\x23\x7f\xff\xca\x2d\x20\x3d\x2b\xbb\x41\xb6\x29\xa7\x0d\xdb\x9e\xee\x0a\xdd\x2c\x6c\xe0\xdb\xc0\xff\x6a\x40\x93\x00\xd3\x20\xe7\x49\xd3\xb0\x1a\xf7\x35\x41\xfd\x46\x57\xca\xfe\x90\x7c\x72\xaa\xf6\x01\x87\x93\x6f\x34\x1c\x0d\xc7\x87\x34\x70\x54\x98\xec\xab\xa6\xcf\x3c\xf2\x56\xe5\x7a\x4b\x6a\x3b\x0e\x3b\x78\x53\x59\x9b\x17\x10\xdf\x6b\x9c\x3b\xde\x9d\xc8\xf1\x55\xea\x7a\x1b\x27\xf0\x50\xae\xb2\xff\x9b\x03\x58\x08\xbb\xd3\xf9\xfb\x32\x71\xed\x1c\x9e\xbd\x1f\x4d\xdc\x7d\xec\x8c\xc6\x10\x66\xb9\xf3\xdc\x38\x7f\xc4\x15\xbb\x19\xa0\x49\x93\x35\xdb\x27\xbd\x58\x46\xab\x54\x99\x74\xd0\xdc\x9a\xb4\xc6\xe6\xd6\x2c\x48\xc5\xef\x67\xaa\xea\x63\x37\xf3\x8a\x69\x5a\x29\xfa\x84\xb5\xd9\xf0\x76\xa9\xe8\xc0\x70\xb9\xe2\x97\x18\x47\xff\xd8\xe7\x5e\xd9\xb3\x27\xa2\x3e\x70\x0c\xe0\x44\x8f\x56\x18\x44\xbf\xff\xbe\xb1\x3e\x92\x39\x6f\x70\x70\x0b\x78\x94\x7c\x19\xd4\x4c\xb9\x36\xb4\xe0\xd4\xa2\xec\xd1\xa8\xb7\xd6\x92\x7a\x6a\xb1\x79\x7a\x58\xfd\x12\x86\xde\x0a\xa0\xf2\xcc\x53\x22\x96\x14\x99\x45\xca\xd8\xaf\x1f\xa3\x0f\x73\xa1\x04\xef\x37\x10\x88\xd2\x56\x0c\x06\x20\x94\xba\x5b\xb7\x79\xb5\xa6\x7b\xdb\x30\x53\xbf\xfe\x63\x37\x16\x26\xdd\xbc\x03\xab\x59\x1b\x41\xf9\xb6\x65\x78\x45\x17\x26\xb2\x9a\x06\x9d\x81\x63\xb7\xd0\x1b\x0f\x77\xbf\x2a\xb7\x4f\x93\xcf\x62\xe3\xb2\xa5\x4c\x68\x08\x90\x83\x8b\x69\x2d\xd0\x34\xc5\x62\x3f\xa3\xdc\xef\x5f\xe1\xb8\x76\x72\x7a\xf2\xd3\x10\xfe\x45\x63\xaa\x49\x56\x5e\xc5\xec\x59\x04\x47\x67\x93\xc7\x89\xbe\x55\x02\xcf\x81\x9a\xcc\x83\x6c\x53\x8f\x1f\x9d\x87\x30\x67\xea\xa1\xbe\x3f\x56\xfa\x0a\xad\xdb\x69\xfd\xf0\x10\xa4\x8d\xd4\x0f\x25\x29\x13\xff\x8f\xe9\x8a\x5a\x9e\x8c\x61\xf4\xfe\xf4\xb4\xd5\x4c\x3b\xa0\xa2\x4d\x1e\xa0\xd7\xae\xc7\x36\xbd\x5d\x2b\x00\x9b\xd4\xf0\xaf\x66\xd7\x66\x65\xe9\x96\x81\x56\xd4\x8a\x21\x57\x02\x61\xa2\xd7\x14\x8a\x3c\xf8\x2c\x28\xa7\x63\xf0\x1f\xbd\xa1\x9b\x90\x8c\x6e\x87\xf0\xac\x52\xae\x37\x95\x83\x6e\x3d\xe8\x9a\xc8\x60\xa9\x27\x35\xb9\x0d\x8e\x6e\x34\xa5\x6c\x47\x9c\xba\x48\x47\xa3\xba\x1a\xf6\x37\xbe\xf3\xd4\x6e\x2f\x6b\xd5\xad\x2c\x50\x7b\x88\x3f\xce\x95\xbd\x90\xd0\x5c\x27\x38\xd1\x29\xbf\xa5\xbb\x33\x87\x9e\x40\x6a\x2d\x81\x35\x25\xae\x1a\xd3\xd2\x7d\xed\x72\xa6\x82\x88\xaf\x1f\xe9\x7b\x5a\xcd\x0e\xe8\x98\xe9\xad\xc9\xa4\xe9\xad\x3f\x12\x37\x47\x53\x54\x03\x9b\xae\xa5\xd4\xf9\x0c\x95\x70\x6a\x35\x01\x41\x6d\x2a\x77\xad\x7a\x5d\x68\xca\xcb\xa3\x72\xc3\x19\x5c\x0a\x9f\xe0\x0b\x1c\x36\x86\xa7\xc3\xc3\x09\x70\x3d\x7b\x9a\x1b\x98\x57\xcd\x0b\x58\x7d\x31\x92\xfe\xf1\x5d\xbd\xda\x3e\xd3\xe8\x55\x45\xdf\x66\xf3\x83\xee\x17\x35\x3b\x25\xd3\x22\x8e\x3e\x7e\xaa\xee\x24\xf7\x8c\x51\x59\x37\xb8\x8c\xa6\x7d\x87\x0a\xfe\xa5\xe1\x4b\x92\x0d\xea\xae\x7a\xb1\xcd\x57\x3b\x2a\x63\x97\x2c\xa5\x8f\x07\x9f\x7c\x16\xa1\xfe\x4d\x00\x95\x68\x71\xd1\x3b\xf4\x6d\x8a\x9d\x0b\xb9\xb1\xf8\x08\x85\x3d\xdb\xbb\xc9\xe4\x1c\xe6\x66\x8f\x6f\x0c\x69\xee\xdc\x7e\x31\xa4\x7d\x38\xc9\x4b\x1a\xdf\x1d\x1c\x94\xf7\x9b\x56\x2a\x42\xc7\x60\xd2\xab\x30\xc4\x3c\xd8\x85\x57\x07\xaf\x00\x95\xcb\x24\x76\x82\x70\x7e\x36\x9e\xd0\xd2\x01\x81\x4c\x83\x88\xbf\xd1\x75\xe1\x9f\x66\x21\xc8\xee\x18\x1b\xa9\x29\x13\x2b\xb6\xe3\xfb\x43\xa5\x1a\xd1\x46\x5f\x3f\xfc\xba\xc6\xd5\x18\xcf\xb7\x88\x37\x06\xe2\x42\xe8\xa5\xcc\xb4\xf8\xa0\x92\x9c\xae\x05\x30\xb8\xed\xba\x0d\xe7\xb5\x69\x2d\x94\xff\xb3\x99\xc9\xd1\x37\x0e\x29\xe0\x98\xc6\xe1\xc6\x67\xcc\x77\x02\x63\x4b\xb9\x8c\x4a\x8d\xe3\x4a\x63\x76\x3e\x46\xb7\x50\xba\x2d\xbc\x58\x74\x2a\xd3\x32\xa8\xcd\x57\x78\x64\x16\x44\xaf\xd3\xd4\x55\xfe\x1b\x19\xdd\x15\x7d\xd2\x46\x24\x3c\xce\xea\x4d\x10\x15\xe2\x12\x1c\xe5\xed\x63\x73\x1b\x79\x83\x51\x3c\x92\x1b\xe9\xb4\xf8\xb6\x58\xbc\x4e\x91\xa7\xd3\x22\xa6\xbd\x72\xde\x72\x8f\xb4\x67\x03\x9c\xee\xd6\xfd\xf2\x03\x91\x8b\xfa\xbd\xd0\xea\xcb\xa9\x11\xb3\x1e\xfc\xf5\x06\x22\x51\xfc\xc3\x01\x6a\x72\x9a\x8a\xf5\xb1\x44\x39\x6c\x5b\xf3\xa9\xc7\x0c\x49\x2d\x5a\x55\xe7\xb8\xf9\x0d\xa2\xd0\x07\x73\xcd\x66\x31\x41\x95\x2c\xd6\x93\xf5\x3a\xa1\x51\x23\x0b\xd2\xb1\x50\xd7\x42\xb1\x64\x5b\x14\x34\xb9\xf7\x99\xca\x91\x23\x2e\x8a\x0c\x63\x63\x8c\x9d\x63\xbf\x95\xba\x7f\x8d\x0e\xad\x2e\x62\xc6\x3b\x69\xf1\xa8\x14\x67\x3f\x11\x34\xf1\x3c\xe7\xf0\xb6\x9f\x17\x6b\x5a\x7b\xac\x59\x41\x8e\xba\xea\xff\x05\x00\x00\xff\xff\x6d\x6b\xf0\x37\xca\x23\x00\x00")

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

	info := bindataFileInfo{name: "templates/search.tmpl", size: 9162, mode: os.FileMode(420), modTime: time.Unix(1469813056, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSearchenumsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x55\x4d\x6f\xe3\x36\x10\x3d\x9b\xbf\x62\xa0\x62\xb1\xf2\xc2\xb5\xee\x01\x72\x48\x37\xc2\x22\x80\x6b\xb7\x89\x17\x3d\x14\x05\x42\x49\xb4\x44\x44\x22\x65\x92\xda\x34\x08\xf2\xdf\x3b\x43\xea\xd3\x49\x7b\x69\x0e\x31\x39\xf3\xe6\x69\x38\x7c\x33\x6c\x79\xfe\xc4\x4b\x01\xaf\xaf\xdb\xdf\xc2\x72\xcf\x1b\xf1\xf6\xc6\x58\x92\xc0\xb7\x74\x9f\xde\xdf\x1c\xd3\x5b\xf8\x7a\xb8\x4d\x37\x90\xde\xde\x1d\x1f\xe0\x8f\xbb\xdd\x0e\x7e\x49\x61\x77\x78\x38\x22\x8a\x80\xa5\x50\xc2\x70\x27\x0a\x38\x19\xdd\xbc\x23\x4b\xd0\x70\x7c\x69\x07\x6a\x0c\xe8\xac\x54\x25\x54\xce\xb5\x57\x49\x52\x4a\x57\x75\xd9\x36\xd7\x4d\x92\x19\xc9\x95\x75\xdc\x3c\x89\xa4\xd4\x3f\x67\xa2\x14\x2e\xf1\xff\x19\x93\x4d\xab\x8d\x83\x98\xad\xa2\x53\xe3\x22\xfc\xb1\xce\x20\x8d\x8d\xd8\xda\xa7\x7b\x68\x29\x09\x6d\x40\x5a\xe0\x0a\xf4\xb0\x75\x1a\x32\x81\x9f\xc4\xf4\xa4\x02\x0e\x0f\x82\x9b\xbc\xba\x17\xe7\x4e\x58\x07\x27\x59\x3b\x61\x98\xc3\xfc\x66\x0c\xca\x79\xca\xaf\x9a\xb2\x51\x8e\x08\x0b\x70\x95\x90\x06\xc2\x57\x21\xd7\x9d\xc2\xc0\x96\x1b\xf4\xc6\x03\xfb\xc3\xef\xbb\xa1\x1a\x52\xab\x35\xcb\x89\x80\x72\x4e\xcf\x13\xfb\x35\x48\xed\x38\x5b\xed\xb5\x4b\xcf\x6c\xb5\x93\x4f\xc2\x6f\xc2\xe2\x9b\x11\x58\x4a\x73\xac\xb8\x5a\x6c\x0e\xc6\x83\x85\xb5\x83\x73\x5a\x07\xdf\x9d\xdd\x77\x75\xed\xa9\xfc\x22\x94\xe5\x57\x6e\x6c\xc5\xeb\xa3\xf8\xdb\x01\xd6\xb0\x16\x8d\xa0\x03\x51\xf1\x2d\x55\x5f\xd7\x5c\x95\x5b\x6d\xca\xa4\x7d\x2a\x13\xa1\x72\x5d\xe0\xf9\x92\x9f\x28\xa0\x8f\xc5\xfa\x9c\x3a\x95\x43\xac\xc7\x43\xac\xe7\xbc\xf1\x1a\xe2\x3f\xff\xca\x5e\x9c\xd8\x80\x30\x86\xbc\xaf\x6c\xf5\x83\x1b\x28\xb8\xe3\x7d\xc5\x18\x5b\xd9\x67\xe9\xf2\x0a\x34\x79\x73\x6e\x05\xa4\xe7\x2b\xb6\x5a\x79\xd0\x35\x44\xe2\x1c\xf5\x76\x5f\x99\xb9\x4b\x91\x61\xf0\x52\x9d\xe6\xce\x1a\xf7\xb3\xc8\x4b\xb7\x0a\xa6\x01\x31\x2b\xe9\x1c\x55\x4e\xe6\x0f\x90\x54\xe0\x7f\x41\x93\x6b\xcc\x6c\xbc\x91\x45\x7e\xa3\xf5\x3d\xee\x92\xb8\x5e\x78\x06\x7c\xb8\xd9\x39\x4e\x7a\xcb\xec\xd4\x97\x00\x15\x4c\x11\x96\xbd\x10\x27\xde\xd5\x8e\xbc\x46\xb8\xce\x28\x50\xb2\xde\x00\x76\xd1\x36\xa5\xdb\x3a\xc5\xd1\x77\xc5\xb3\x5a\x50\xab\x34\xe1\x5a\xe1\xf1\xd3\x8f\x47\x52\x34\x75\x0f\x5e\xac\x8d\x36\xa0\xd7\x6c\xf5\xc6\x06\x8e\x70\xe1\x31\x7d\x6f\xbd\x21\x46\x16\xa6\xc6\x77\xd5\xfc\x0f\xc5\x8d\xd1\x33\xcd\x7d\x99\x44\xb7\x20\x8f\xb3\x3e\x89\x75\x50\x1d\xc9\x0a\xb5\x06\x57\xd7\xbd\xe4\xec\xf6\x68\x64\x13\x87\x4d\x9c\x61\x9a\x8f\xd1\xe3\x7a\x12\x22\x81\x07\x29\x92\xfa\xa8\x42\x5f\x34\x56\x8f\xba\x29\x58\x83\xf0\x46\x47\xdf\xb2\xc1\xe7\x75\x37\xba\x42\xf7\x8e\x51\xbb\x85\x73\xec\xee\xe0\x9f\x8b\x6d\xc4\x2c\x1a\xff\x1d\xce\xcb\xe1\x23\x6c\xe8\xfd\x3e\xa3\x49\x69\x53\x5e\xb3\x89\x71\x89\x5a\x92\x5e\xce\x93\x80\xee\x95\x36\xa2\x86\x29\x33\x9e\x74\xe9\x1e\x86\xcf\x47\xb2\xfb\x6f\xc5\x7d\xfe\x64\x3f\xd3\xe4\xd5\xb0\x18\xc6\x28\x3c\xbc\xa7\x85\xf4\x26\xb1\xe1\x80\x2e\x24\x8d\xda\x7e\xe8\x4b\xdc\xe6\x7e\xc8\xea\x13\x3c\x57\x02\x47\x36\x2e\x0d\x60\x9a\x38\xf9\xc3\xac\x07\x5b\xe9\xae\x2e\x28\x1c\xdf\x05\x5e\x14\x38\xba\xb9\x8f\xbe\xd9\xdf\x12\x18\x57\x87\x7b\xca\x0c\xa3\xc1\xfa\xd7\x02\x5a\x23\x3c\xb3\xb0\xe1\xa9\x98\x3e\x3c\x4c\xb7\xe1\xb9\xb0\x70\x42\x8e\x17\xdd\x19\xc8\x3a\xe7\xa6\x27\xe0\x06\x9f\x90\x29\x0c\x5b\x14\x3f\x87\x0d\x7c\x30\x30\xfb\x43\xf3\xe1\x9e\x5e\xb4\x7f\x02\x00\x00\xff\xff\xa2\x18\x65\x62\x9a\x07\x00\x00")

func templatesSearchenumsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSearchenumsTmpl,
		"templates/searchEnums.tmpl",
	)
}

func templatesSearchenumsTmpl() (*asset, error) {
	bytes, err := templatesSearchenumsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/searchEnums.tmpl", size: 1946, mode: os.FileMode(420), modTime: time.Unix(1469813078, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUpdateTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x57\x4f\x6f\xdb\xc6\x13\x3d\x93\x9f\x62\x7e\xfc\x39\x35\x65\xa8\x54\x0f\x45\x0f\x02\x7c\x48\x2a\x35\x70\xe1\xda\x6d\x2c\xa3\x07\xc3\x80\x96\xe4\x4a\xa2\x4d\x2e\x99\xdd\xa5\x63\x97\xd0\x77\xef\xcc\x2e\x49\x91\x96\x2c\x3b\x89\x03\xf4\x90\x1c\x22\x92\xbb\xf3\xef\xbd\x37\xb3\xeb\x82\x45\xb7\x6c\xc9\xa1\xaa\x82\x3f\xed\xe3\x19\xcb\xf8\x7a\xed\xba\xa3\x11\xbc\x9f\x9e\x4d\x3f\xbc\x9d\x4d\x27\xf0\xeb\xf9\x64\x3a\x84\xe9\xe4\x64\x76\x01\x7f\x9f\x9c\x9e\xc2\xbb\x29\x9c\x9e\x5f\xcc\x70\x17\x6d\x5c\x72\xc1\x25\xd3\x3c\x86\x85\xcc\xb3\x2d\x67\x23\xfc\x30\x7b\x28\x1a\xd7\x68\x50\xaa\x44\x2c\x61\xa5\x75\x31\x1e\x8d\x96\x89\x5e\x95\x61\x10\xe5\xd9\x28\x94\x09\x13\x4a\x33\x79\xcb\x47\xcb\xfc\xc7\x90\x2f\xb9\x1e\x99\xff\x5d\x37\xc9\x8a\x5c\x6a\xf0\x5d\xc7\x5b\x64\xda\xc3\x1f\xa5\x25\xba\x51\xf4\xc8\x45\x94\xc7\xf8\x32\xba\x51\xb9\xf0\xdc\x81\xc9\xbf\x17\xf6\xb7\x84\xa7\x31\x24\x0a\x18\x2c\xcc\xe3\x27\x0c\x9b\x08\xd0\x2b\xde\xdf\x08\xe8\xb6\x8c\x34\x39\xd0\x2b\xa6\x8d\x49\x98\x72\xd0\x39\x84\x1c\xca\x22\xa6\x42\x5d\x8d\xfb\x77\x06\x10\xda\xc4\x9e\x8a\x32\x3b\x44\x3c\x72\x09\x2b\x9e\x16\x8b\x32\x15\x5c\x29\x37\xca\xb1\x3c\xf0\xab\xea\x40\xd7\x86\x30\x3e\x86\x2e\x3a\x55\x25\x99\x40\x46\x0e\x92\x21\x1c\xdc\x9b\xd5\x4b\x13\x93\x92\x30\x31\xd4\x7a\xdd\xb1\x37\x2f\xf7\x41\x9d\x7b\x55\x25\x0b\xe0\x1f\xd1\x1a\x7e\x7a\xb4\xcd\xe6\x77\x0c\x49\xae\x59\x55\x71\x11\x9b\x60\xe6\xd7\xc2\xf5\x07\x93\x6a\xc5\xd2\x19\xbf\xc7\xa2\xb3\x22\xe5\x19\x17\x5a\x19\x96\x14\xd1\x94\xa7\x98\x58\x90\xcb\xe5\xa8\xb8\x5d\x8e\x5a\xc4\xff\x4f\x06\xb5\x2d\x97\xee\xa2\x14\x11\xf8\x6a\x07\x36\x83\x6e\x04\x7f\x00\xfe\xd5\x75\xf8\xa0\xf9\x10\xb8\x94\xb9\x1c\x40\xe5\x3a\x77\x4c\x02\xd6\xca\xc0\x52\xeb\xba\x8e\x42\x9a\xa2\x15\x28\x5a\x7d\x11\x34\x11\x53\x44\xcc\x6e\x7c\xc6\xae\xe3\x18\xff\xc7\xe0\x99\xcf\xbf\x5f\x9c\x9f\xd9\x25\x8f\xfc\x5b\x54\x9c\x98\x2f\x58\x99\x6a\xda\x2d\xb9\x2e\xa5\x00\x91\xa4\x43\x40\xd9\x05\x53\xca\x75\xe1\x7b\x97\xa2\x11\x45\x66\x8b\x82\xf9\x9b\xbb\x39\x90\xa0\x50\x27\x58\x96\xf2\x86\xa0\x06\xae\x83\xee\x6a\x1f\xb6\x5c\x9f\xe2\x0f\x86\xe4\xd1\xb5\x6d\x76\x29\xb2\xaf\x40\xbe\xb5\xee\x60\x7f\xb4\x0b\xfc\x5e\x18\x3f\xac\xd3\x19\x58\xf4\x09\x5e\xc4\x9c\x30\xad\xbb\x2a\x98\xc9\x24\xf3\xed\x8b\x1f\x62\xc2\x73\x6f\x3e\xe8\x10\xa2\xe5\xe7\x51\xb2\x85\x37\xa1\x7b\xa4\x90\x89\xa7\xc8\xda\x4b\xc8\x7e\x2e\x0e\xdf\xa8\x43\x6a\xc5\x1c\x9e\xe8\x53\x22\x47\xcb\x1e\x3d\x1b\x42\x26\xa1\xd9\x63\xba\xd3\x2e\x2a\x33\x26\x04\x7d\xc8\x17\xe6\xd9\xce\x10\x0c\x50\x62\x6d\xf5\x18\xb9\xf8\xeb\x14\x3e\x96\x5c\x3e\xec\x6d\x82\x8e\x77\x6c\x02\x0b\xb0\x81\xff\xd5\x95\x5e\x17\x66\x91\x9f\x84\xdb\x3a\xdf\x14\xef\x79\x75\xed\xbd\x8c\x6d\xcc\x0f\x1c\x8b\xc2\xa9\x85\x24\x24\x38\xc4\x70\x80\x2a\xae\x09\x87\x82\x49\xdc\xa6\xb9\x54\x34\xe7\xcc\x5c\x27\x03\x53\x4f\xd7\x4d\x00\x70\xa2\x21\x62\x82\x06\xa8\xe2\x38\xe3\xd3\xe4\x1f\x3c\x2c\x98\x88\xd1\x87\x52\x38\x4e\xd1\x36\xe4\xfa\x13\xe7\x82\x36\xdc\x25\x11\xc5\x51\x40\x72\x19\x02\xca\x13\x51\x36\x68\x37\x27\x0d\x26\x41\x68\xe3\x51\xa1\x4d\xb3\x04\x3b\x06\x72\x3f\x7b\x3b\xd5\x09\xdc\x93\x09\x49\xe3\x97\x9f\x61\x4e\x87\xc5\xd8\x4b\x62\x6f\xee\x3a\x76\xb7\x42\x09\x15\x57\xdb\xb4\x5d\xa3\x05\x97\x0b\x16\xf1\x6a\xdd\xd8\xd9\xa3\x40\xa1\xf1\xfa\xd5\xe7\x67\x29\xf7\xd4\xf2\x82\x41\x6a\x65\x95\xa0\x4a\x63\x52\x4f\xc6\x6e\xb9\x4f\x95\xd9\xef\xdd\x6a\xa8\xa7\xe9\x94\xba\xe5\x0f\x43\xb8\x63\x69\x69\x4e\x24\x2b\xbf\x52\x06\x0d\x2c\xe8\xd3\x51\x26\x02\x2d\xe3\xe6\xa0\x97\x03\x3a\x71\xe8\xe4\xc1\xe5\xff\x1d\x53\x2b\x19\x83\xde\xf4\xc4\x35\xfc\xb4\xa6\x9d\x9d\xec\xea\x8c\x7c\x35\xb8\xc6\x51\x60\xe2\xbb\x66\x93\xe4\x0a\x3b\xde\x26\xbf\x2b\x6f\xf2\xdf\x52\x30\x86\x8e\xcb\x21\xad\x20\xa9\x63\x30\xff\xb0\x88\x93\xc9\xb0\x71\x6a\xf2\x21\xfe\x9a\xfc\x7d\x1b\x68\xd0\x27\xf1\x5d\x22\x98\x7c\xf8\x2c\x1a\xad\xc9\x97\x11\x69\x6d\xbf\x53\xf9\xaa\x54\xbe\xf6\xb9\x8a\x54\x1e\xed\xe3\xf2\xd9\x03\x96\x6e\x37\xc9\x13\x18\x60\x45\x35\x21\xa6\xa2\xd6\x97\x1f\x0e\xe1\x87\x84\x48\xd9\xe6\xa4\x86\xc0\xb0\x41\x0e\x0c\x3c\x74\xc3\xbb\x22\xcc\xae\x03\xdf\x8c\x39\xb2\xa5\xc8\x37\xd0\x11\xcf\xfe\x09\xb7\x4f\x48\xe8\xbc\xa1\x0a\x23\x3c\x21\x44\x93\x1e\x05\x2d\x76\x1c\x81\xb8\x44\x85\x1c\x43\x11\xf4\x21\xab\x2f\x48\x18\x75\xf0\x8c\x0a\x1b\x01\x3a\xce\xcd\x55\xf1\x48\x6d\x1d\xa5\x23\x98\xee\x8e\xf3\xbd\x8d\xfa\xc5\x5d\xfe\x15\xe2\xa8\x5b\xfd\xbb\x3c\xfe\xb3\xf2\x78\xcf\x35\x5d\x11\x36\xa2\xb0\x05\x35\x76\xcd\x85\xa2\x2d\xe8\x85\x0a\x30\x6e\x71\xc4\xdb\xbb\x47\xd5\x06\x36\xb4\x6c\x42\xcf\xe8\x8a\x67\xae\x9e\xdf\x22\x83\xd6\x7b\xef\xea\xb9\xb9\x28\x06\xed\x06\xba\x29\xb6\x49\x35\x88\x7d\x8b\x94\x6a\xdf\x98\xd0\x6e\xb5\x50\x82\xb5\xa4\x5e\x78\xfe\x99\xbb\xf9\x73\x27\x60\xed\xf2\xca\x6c\x0e\x7a\x97\xf2\xad\xe3\xcb\xf2\x64\x0d\x6a\x50\xde\xc6\xb1\xf5\xf6\xfa\x98\xb4\xae\x7d\xfb\x47\xc6\x76\x8b\x34\xb5\x3d\xee\x29\xec\x88\xae\xbc\x37\x8d\xd1\x13\xfd\x8b\xbb\xfc\x51\xbf\x58\xa8\x36\xe0\xac\xdd\x7f\x03\x00\x00\xff\xff\xd5\x7b\x36\xc6\x3e\x12\x00\x00")

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

	info := bindataFileInfo{name: "templates/update.tmpl", size: 4670, mode: os.FileMode(420), modTime: time.Unix(1469813083, 0)}
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
	"templates/search.tmpl": templatesSearchTmpl,
	"templates/searchEnums.tmpl": templatesSearchenumsTmpl,
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
		"search.tmpl": &bintree{templatesSearchTmpl, map[string]*bintree{}},
		"searchEnums.tmpl": &bintree{templatesSearchenumsTmpl, map[string]*bintree{}},
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

