// Code generated by go-bindata.
// sources:
// templates/ginCreator.tmpl
// templates/ginUpdater.tmpl
// templates/searcher.tmpl
// templates/searcher_enums.tmpl
// templates/sqlCreator.tmpl
// templates/sqlUpdater.tmpl
// templates/updateRequest.tmpl
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

var _templatesGincreatorTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x91\x41\x6f\x9b\x40\x10\x85\xcf\xec\xaf\x98\x72\xa8\xc0\x72\x76\xa3\x4a\xbd\x58\xca\x21\xb1\x49\xeb\x2a\xb2\x5b\x1b\x1f\x7a\x5c\x60\x0c\xdb\xc0\xae\x3b\x2c\x4d\x2a\x8b\xff\xde\x5d\xa0\x55\x9c\x4a\xbd\xf5\x36\x62\xde\x3c\xbe\xf7\xf6\x24\xf3\x47\x59\x22\xe4\x84\xd2\x22\x13\x33\x16\x7c\x48\x36\xc9\xee\x36\x4d\x56\xb0\xdc\xae\x92\x39\x7c\xdd\x1e\x76\x90\xac\xd6\xe9\x1e\x6e\x77\x09\xdc\x1f\xd2\xf5\x43\xc2\x58\x50\xa2\x46\x72\x47\x05\x1c\xc9\x34\x70\x3e\xf3\xf4\xe7\x09\x37\xb2\xc1\xbe\x67\x41\xd7\x2a\x5d\x42\x65\xed\x69\x21\x44\xa9\x6c\xd5\x65\x3c\x37\x8d\xc8\x48\x49\xdd\x5a\x49\x8f\x28\x4a\x73\x95\x61\x89\x56\x4c\x56\x86\xc4\xc0\x61\x88\xcd\x04\x63\xaa\x39\x19\xb2\x10\x31\x80\xf0\x85\x43\xa9\xf4\x55\x69\xb4\xca\xfd\x14\xbe\x5a\x7e\x6b\x8c\x22\xa3\x45\xfb\xbd\x7e\x0e\x99\x5b\x5a\xc7\xd4\x42\x38\xd1\xad\x07\xcb\xbe\x0f\x59\xcc\x98\x10\xb0\xc1\xa7\x0b\xee\xe5\x50\xc3\x47\xa9\x8b\x1a\x09\x9e\x54\x5d\x03\xa1\xed\x48\x83\x84\x6a\xfa\x7a\x34\x74\x19\x76\x2a\xaf\xe5\xce\xd1\x9b\xa6\x15\xfe\x11\xcb\xb6\xed\x1a\x07\x60\x2b\x69\x9d\xc7\xe5\xdd\xe0\x9f\x21\x7c\x3e\xa4\x30\xf6\xf5\x69\xbf\xdd\x70\x76\xec\x74\xfe\x6f\xb4\xa8\xc8\x60\xe6\x33\xf2\xd5\x5d\x0c\xae\x07\x3e\x2d\xee\xfd\xe9\x99\x05\x13\xb5\x77\x8a\x72\x98\x79\xc5\xd2\x68\x8b\xcf\x36\xf6\xeb\xe0\x87\x24\xb0\x63\x39\xfc\xd5\xcb\x05\x48\x04\x8b\x1b\xc8\xf9\x9d\xd2\x85\x27\x8a\xde\x5a\x57\x57\x10\xa8\x23\xf8\xdd\x9b\x1b\xd0\xaa\x1e\x6c\x82\x9c\x0f\x82\xf7\xd7\xd7\xf3\x11\xe3\x1c\x3a\x89\xa1\x70\xe1\xa5\x3c\xf1\x73\x14\xf7\xb1\xd7\x8e\x4c\x6e\xea\x07\xb3\x62\x0e\xd3\x9f\x5c\xd4\xfd\x97\x87\xbf\xd3\x1a\x9f\x33\xe6\x63\xf2\xe8\x3f\x30\x4c\x97\xef\x5e\x5c\xaa\x22\x5c\xa8\x62\x10\xff\xd6\xf6\xac\x67\xbf\x02\x00\x00\xff\xff\xdc\x4b\xf0\x22\x28\x03\x00\x00")

func templatesGincreatorTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesGincreatorTmpl,
		"templates/ginCreator.tmpl",
	)
}

func templatesGincreatorTmpl() (*asset, error) {
	bytes, err := templatesGincreatorTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/ginCreator.tmpl", size: 808, mode: os.FileMode(420), modTime: time.Unix(1468871773, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGinupdaterTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x92\xc1\x6f\x9b\x30\x14\xc6\xcf\xf8\xaf\x78\xe3\x30\x41\xd4\xda\xd5\xa4\x5d\x22\xe5\x90\x26\x74\xed\x54\x25\x5b\x42\x0e\x3b\x1a\x78\x01\x2f\x60\x53\x63\xaf\x9d\x22\xfe\xf7\xd9\xc0\xa6\x6e\xaa\x72\xdb\x05\x3d\xf1\xbe\xdf\xf7\xbe\xf7\xe4\x96\xe7\x27\x5e\x22\xd8\xb6\xe0\x06\x09\x9b\x91\xe0\x53\xb2\x49\x76\xcb\x34\x59\xc3\x6a\xbb\x4e\xae\xe0\xdb\xf6\xb0\x83\x64\xfd\x90\xee\x61\xb9\x4b\xe0\xee\x90\x3e\x3c\x26\x84\x04\x25\x4a\xd4\x0e\x2a\xe0\xa8\x55\x03\xe7\x33\x4d\x7f\xb6\xb8\xe1\x0d\xf6\x3d\x09\x6c\x27\x64\x09\x95\x31\xed\x9c\xb1\x52\x98\xca\x66\x34\x57\x0d\xcb\xb4\xe0\xb2\x33\x5c\x9f\x90\x95\xea\x3a\xc3\x12\x0d\x9b\xac\x94\x66\x63\x0e\x4d\x66\x8c\x10\xd1\xb4\x4a\x1b\x88\x08\x40\xf8\xca\xa1\x14\xf2\xba\x54\x52\xe4\xbe\x0a\xff\x69\x7e\x6f\x94\xd0\x4a\xb2\xee\xa9\x7e\x09\x49\x4c\x08\x63\xb0\xc1\xe7\xbf\xc2\x1d\x86\x19\xf7\x5c\x16\x35\x6a\x78\x16\x75\x0d\x1a\x8d\xd5\x12\x38\x54\xd3\xdf\xa3\xd2\xf0\x06\xb4\xc3\x27\x8b\x9d\xe9\xa8\xf3\xf5\xd6\x69\x85\x7f\x10\xde\x75\xb6\xc1\x0e\x4c\xc5\x8d\xfb\xe0\x05\x7e\x1c\x9a\x21\x7c\x59\xa6\xab\x7b\x77\xc1\xe1\x5a\xde\xf0\xf3\x7e\xbb\xa1\xe4\x68\x65\x7e\x39\x76\x54\x64\x30\xf3\x4b\xd2\xf5\x6d\x0c\xee\x10\x74\x6a\xdc\x79\xf4\x4c\x82\x69\x23\xef\x14\xe5\x30\xf3\x8a\x95\x92\x06\x5f\x4c\xec\xdb\xc1\x0f\xae\xc1\x5e\xda\xd1\x69\x50\x6b\x98\x2f\x20\xa7\xb7\x42\x16\x3e\x59\xf4\xde\x6a\x77\xd3\x20\x10\x47\xf0\xcd\x77\x0b\x90\xa2\x1e\xfc\x82\x9c\x0e\x8a\x8f\x37\x37\x57\x63\x9e\x73\xe8\x24\x4a\x87\x73\x2f\xa5\x89\xaf\xa3\xb8\x8f\xbd\x76\x0c\xe7\xaa\x9e\x4c\x63\x16\x7e\xdd\xfd\xd7\xc7\x37\xf2\xf8\x5d\x63\x3a\xd6\xd1\xff\x98\x3f\x91\x1f\x5e\x91\xee\x81\x1a\xdb\x39\x34\x54\xa7\x70\x60\x7e\x23\x3d\xe9\xc9\xaf\x00\x00\x00\xff\xff\xd6\x42\x09\x81\x35\x03\x00\x00")

func templatesGinupdaterTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesGinupdaterTmpl,
		"templates/ginUpdater.tmpl",
	)
}

func templatesGinupdaterTmpl() (*asset, error) {
	bytes, err := templatesGinupdaterTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/ginUpdater.tmpl", size: 821, mode: os.FileMode(420), modTime: time.Unix(1468871773, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSearcherTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x1a\x6b\x73\xdb\xc6\xf1\xb3\xf8\x2b\x36\x68\x62\x03\x0a\x0b\x6a\x52\x37\x1f\x38\x61\x3a\xb6\x44\xd9\x6a\x14\x4a\x11\xe9\xfa\x83\xc7\x53\x81\xc0\x81\x44\x0c\xe2\xa8\x3b\x50\x8f\x30\xfa\xef\xdd\xdd\x3b\xbc\x28\x88\xa1\x54\x37\x33\xf5\x68\x86\xc0\xdd\xbe\x77\x6f\x1f\x07\x2f\x83\xf0\x73\x30\x13\xa0\x45\xa0\xc2\x79\xa7\xd3\xeb\xc1\xdb\xe1\x68\x78\xf1\x7a\x32\x3c\x82\xc3\xb3\xa3\x61\x17\x86\x47\x27\x93\x31\x7c\x38\x39\x3d\x85\x37\x43\x38\x3d\x1b\x4f\x10\x8a\x00\x67\x22\x13\x2a\xc8\x45\x04\xb1\x92\x0b\x58\xaf\xfd\xc9\xdd\x52\x8c\x82\x85\xb8\xbf\x87\x24\x2b\x16\x4e\x16\x4b\xa9\xf2\xfb\x7b\x42\x59\xe9\x24\x9b\xc1\x3c\xcf\x97\xfd\x5e\x6f\x96\xe4\xf3\xd5\xd4\x0f\xe5\xa2\x37\x55\x49\x90\xe9\x3c\x50\x9f\x45\x6f\x26\xff\x3a\x15\x33\x91\xf7\x2c\x7d\xa9\x7a\x46\x3a\xa1\x3a\x9d\x84\x89\x81\xdb\xd9\x73\xa2\x20\x0f\xa6\x81\x16\x3d\x7d\x95\x3a\xf8\x2e\xb2\x50\x46\x48\xbe\xf7\xab\x96\x19\x2d\xc4\x8b\x9c\x7e\x12\xd9\x4b\xe4\x2a\x4f\x18\x28\x43\xba\xc4\x9e\x9e\x75\xae\x10\x5c\x3b\x1d\x7c\xae\xc9\xf2\xeb\x42\x26\x4a\x66\x44\xf6\x96\xf6\x0c\x73\xd2\x04\x9c\x0d\x95\x90\x8a\xbe\x82\x3a\x72\x1a\x64\x84\xb9\x4a\x94\x12\xc8\xd0\x63\x8b\x36\x2c\x33\x66\x72\x17\xe2\x6a\x25\x74\x0e\x89\x86\x00\x8d\x8f\xea\xa7\xc9\x6f\xc1\x34\x15\xd0\xdc\x8f\xa5\x6a\xa2\xfb\x9d\x9c\x24\xd9\x42\x12\xd5\x5a\x85\x39\xac\x3b\x7b\xc7\x89\x48\x23\x0d\x1f\x3f\x35\xa0\x79\x15\x2e\xc9\x4a\x7d\x27\x66\x10\xe7\x92\x80\xd3\x5c\xa8\x16\x68\x5a\xae\xc0\x19\x88\xe0\xcf\x54\x24\xd4\x9b\x3b\xd8\x6f\x80\x17\xab\x16\x5e\x9a\x57\x82\x3f\x4d\x16\x09\xea\x9b\xe5\xc5\x5e\x4a\x0b\x4c\x29\x8e\xb5\x68\x6c\x49\x5e\xc1\xbd\xfb\x87\xf6\x33\xe2\xb3\xdd\x58\x78\xb8\x41\xeb\x63\xb4\xe5\xf3\x0d\xab\x58\x43\x10\x81\x7c\x1e\x18\x53\x93\x81\x73\x09\x53\x01\x46\x13\x8c\x5d\x99\x75\x41\xa3\x37\xed\x23\xda\x5b\x89\x7c\xa5\x32\x11\xb5\x99\xda\x72\xcf\x72\x16\xec\x38\x58\xdd\x82\xc8\x56\x8b\x97\x11\x7b\x6a\x2e\xd2\x65\xbc\x4a\x33\xa1\x75\x27\x94\x18\xd0\xe0\xae\xd7\x5f\xe7\x16\x1b\xfa\x03\xa8\x91\xea\xac\xd7\x2a\xc8\xf0\xec\x7d\x9d\x74\xe1\xeb\x5b\xde\x35\x9e\x24\x31\x8d\xef\xee\xef\x6b\xf8\xfc\x72\xeb\x5b\xed\xd6\xeb\x24\x06\x71\x85\xd8\x70\xb0\x01\x66\x84\x1c\x40\x22\xf3\x60\xbd\x16\x59\xc4\xcc\xf8\x17\x03\x92\x04\xff\xe7\xf8\x6c\x04\x19\xc9\xc4\x62\x06\x59\xfe\xc5\x04\x0e\xeb\x42\x9a\x23\x86\xa2\x38\xbc\x4a\x6c\xcd\x8e\x53\x13\x08\xe5\x39\x9a\x32\x3a\x33\x35\xe6\xd7\xec\x50\x16\x51\xc6\xfc\x6c\xbc\x8d\xde\x5b\x69\x01\xd6\xe1\xe3\x5f\x4e\x01\xa3\x5e\xdd\x75\xe2\x55\x16\x82\xab\x5b\xbc\xe5\xd5\xa9\xbb\x5e\x21\x13\x9e\x0f\x8d\xa1\x13\xce\x41\xd3\xf3\x6e\xba\x61\xba\x81\xc7\x3c\xd2\xef\xec\xed\x19\xd9\xad\xb6\x47\xd3\x42\xd7\xbd\xc2\x0b\x7b\xf8\x57\xc0\x38\x36\xba\x7f\x0e\x94\x9e\x07\xe9\x44\xdc\x62\x8c\x2e\x96\xa9\x58\x08\x74\x07\x27\x49\x4d\x59\x52\x62\x4e\x99\xf9\x52\xcd\x7a\xcb\xcf\xb3\x5e\x99\xe5\xfe\x42\x08\x16\x17\x33\xe3\x36\xfd\x6b\x1c\x50\x7f\xf7\xe3\xa7\xe9\x5d\x2e\xba\x20\x94\x92\xca\x23\xe5\xaf\x03\x05\x94\x4e\xad\x6d\x3a\xff\x0b\xd3\x30\xfd\x01\xd4\xc3\xa3\x66\x97\x48\xc4\xc1\x2a\xcd\x6b\x36\xcc\x92\xb4\x0b\x98\xc2\xfd\x21\x89\x19\xbb\xce\xfb\xac\x38\xbe\x0b\xa3\x0f\x5c\x7e\x73\x7d\xc9\x91\x80\x27\x1a\x35\xd2\x0e\x9e\x64\xaf\x6e\x63\xa3\xa9\x4b\xac\xbd\x2e\x51\xb4\x26\x7f\x9f\x2d\xfe\x0b\xa3\x97\xd8\x35\xb3\xef\xb7\xd9\xbd\xc1\xc6\x9d\x5a\x71\x3c\x63\x78\x0e\xc0\x5c\x91\x39\x6d\x21\xf2\x27\x2a\x59\xb8\xe6\xc5\x9d\xa2\xc0\x97\xce\xa5\x57\xf3\x45\xae\x9e\xe6\x8d\x70\xc3\x03\xfb\x1a\xed\xff\x98\x8b\xb6\xfa\x62\xbb\x1b\x5e\x7e\xa3\x5f\x52\x4a\x94\xf0\x48\xbe\x24\xbf\xe4\xaa\xe1\x99\xca\x17\x6d\xd5\xc6\x66\x77\x7e\xd6\x4b\x11\x26\x71\x12\x12\xc7\x06\x6c\x6b\x76\x36\x28\xcd\x0a\xd8\x96\xc1\xeb\xe5\x8f\x6a\xd0\xbf\x82\x74\x45\x59\x05\xd1\xe3\x20\x14\xeb\xfb\x02\xe2\x9a\x36\xb8\x4a\x2d\x4d\x33\x02\xe5\x43\x51\xac\xec\x3b\x01\x1d\xca\x2c\x4a\xf2\x44\x66\x50\x3d\x59\xb0\xb0\x58\x68\x2f\x6b\x45\xe1\x34\x0d\x01\xf5\x38\x51\xa2\x44\x98\x27\xd7\xa2\x2c\x5f\x4f\x31\x45\x41\xef\x19\xb6\x38\xd1\x47\x42\x87\x18\x0c\x94\x26\xa7\x52\xa6\x05\x44\x52\xdb\x28\xb5\x78\x1d\x45\x75\xaf\x51\x09\xc4\xe7\x85\xc8\xe7\x92\xb3\x75\x10\x45\x95\x33\xe9\xbd\xd9\xe1\xf8\x00\x28\xa5\x0d\xb9\xae\x29\xd6\x98\xf2\x0b\x1b\x23\xcd\xe1\x2f\x10\x64\x11\x2f\x97\x36\xa4\xf5\xd7\xa3\x23\x44\x3e\x89\xe1\x4e\xae\xe0\x06\x4b\x18\x51\x97\xd7\x78\xb8\x92\xc8\xd8\xac\x6f\x1b\x55\xa0\xfc\xa6\x55\x93\x31\x6f\x68\xe5\x97\xf2\xbb\xe8\xa2\x14\xdd\xd8\x05\x07\xfb\x02\xfc\xa1\xb3\xed\xc6\x0f\x8e\x36\xc1\x7a\x6b\x46\x07\x88\xfd\x32\x1c\x06\xf0\x76\x52\xae\x56\xee\x1f\xc0\xd9\x05\x2f\x63\xa9\x33\xe2\x90\x36\x28\xf3\x4b\x25\x20\x48\x53\xec\xfe\x72\xbf\xc8\x23\x6a\x83\x5b\x43\x62\xaf\xb2\xb5\x1b\x3f\xe2\xcc\x2e\x5c\x6f\x46\x72\x17\xc2\x78\x06\xbe\xef\x6f\xd3\xc7\xdb\xc6\x98\xa2\x27\xa6\x4c\xd3\x82\x89\x5b\x26\xb2\xfa\xa6\x40\x77\xf1\xbd\x30\x49\x1f\x60\x78\x45\x0b\xa5\x35\xfa\xf0\x3a\x63\x10\x3e\x6f\x7d\x60\x63\x61\x72\x19\x2f\x31\xe9\xe5\x2e\xcb\xee\x75\x29\x51\x20\x47\xb4\xe9\xbf\x51\x78\x62\x6c\x52\x1e\xe9\x41\xfc\x42\xf7\x45\xec\x19\x20\x74\x60\xd1\xb8\x0e\x20\x58\x2e\x31\x38\xdd\x6a\x0d\x7d\x48\x09\xd4\xe6\x1c\xad\xea\x21\xcb\xad\x31\x06\xa7\x36\x72\x6b\x8a\x1e\x8a\x31\x2d\x52\x3c\x76\x18\x59\xef\x1a\x91\x4c\xd2\x60\x8e\xc5\x6a\x33\x5d\x25\x29\x1f\x8d\x60\x7b\x6f\xff\x24\xa7\x92\x04\xae\x15\x04\x5d\xd5\x56\x4d\xfe\xc0\x41\xac\x35\xe3\x37\x0d\x41\x4b\x5d\xab\x23\x52\x6e\xb3\xc7\x58\xe4\xa6\x3b\xc7\x50\x34\x7d\x17\xbf\xee\x64\x83\x2d\x06\xd8\x51\xff\x82\xbb\x9b\x16\x13\xc2\x2e\xba\x1a\x81\x07\xc0\x48\xed\x3a\xd9\xb9\xa2\x54\xca\xbc\xff\x79\x5a\x19\x7e\xae\x2c\xc7\x9b\x5d\xf4\xb2\x42\x0f\xc0\xa0\xb5\x6b\x76\x6e\x46\xf6\x86\x5e\x9c\x56\x8c\x55\x68\x2a\xa6\x69\x86\x77\x97\x04\x8b\x23\xca\x14\x35\x26\x10\x7e\xd7\xc9\x6f\x62\xd7\xf8\xb4\xfc\x5c\x42\x1c\x31\x9d\x2e\x13\x19\x23\x8d\x9d\xb4\x2a\x15\xf0\xf1\xec\x56\x76\xa9\xe8\xc1\x7e\x49\xd0\xb3\x30\x26\x22\xca\xd5\x9a\x53\x8b\xba\x56\x6a\x4f\x0b\x58\x43\x76\xf3\xeb\x73\x5d\x69\xb8\x6e\x49\xbb\xc9\x66\xd1\xdc\xc9\xd9\x56\x99\x01\xbc\x68\x2b\xde\xb5\xdc\x6a\xff\x95\x29\xb6\x5e\xa3\xfb\x0d\xe6\x36\x7b\x6e\x46\xcd\x5b\x7b\x59\x33\xe6\xf4\x46\x73\xd3\x4d\x82\xc5\xa7\xb8\xc3\xc1\xd0\x00\x71\x2b\xc2\x55\x6e\x2e\x21\x70\x1f\xc7\xc2\x9c\xdb\x62\x0e\x1b\x4b\xcf\x0e\x5d\x44\xd1\xce\x51\x2e\xcd\xdf\xb0\x4c\xb1\xd6\xcc\x65\x8a\x72\x6b\x8f\x11\xb0\x8b\x49\x93\xb0\x9c\xde\x38\xb5\xeb\x5d\x63\xee\x81\xb4\x34\xb7\x18\x86\x5d\x6c\xa3\x1b\xd5\xad\x39\xc6\xe8\x6a\x86\xc1\xf9\x38\x15\x59\x95\x05\x3d\x18\x0c\xe0\x80\x2b\x08\x25\x49\x67\x1f\x07\xb3\x7b\x10\x29\x0d\x2e\xb8\x66\xeb\x4d\x5c\xd5\x9b\x2a\xa5\xd2\x3e\x23\x69\xf8\x16\xab\x7b\x73\xa0\xfc\x16\xa8\x67\x40\x88\x7b\x4b\xb9\xde\xcf\x5f\x24\xb3\x79\xee\x62\x06\x26\x18\x5b\xb4\x96\xfa\x2a\xe5\xbe\xff\xca\x1f\x17\x46\x7e\x43\x31\x2a\xf8\x88\x9c\x57\xb6\x3c\x96\x6a\x11\x20\xfa\x95\x7f\x24\xd3\x34\x50\xf6\x78\x90\x51\x5c\xcd\x2f\xc7\x4a\x2e\x5c\xbe\x97\x22\xc7\xd9\x81\xd3\x33\x95\x11\xab\xdb\x39\xb6\x32\x49\x88\x3c\xb4\xfb\x82\xd8\x7a\xc6\x2e\x65\x02\xfd\xd1\x1a\x84\x45\x1a\x00\xfd\x98\x1d\x77\x85\x46\xfe\xfe\x95\x5b\x40\x7a\x56\x76\x83\x6c\x53\x4e\x1b\xb6\x3d\xdd\x15\xba\x59\xd8\xc0\xb7\x81\xff\xd5\x80\x26\x01\xa6\x41\xce\x93\xa6\x61\x35\xee\x6b\x82\xfa\x8d\xae\x94\xfd\x21\xf9\xe4\x54\xed\x03\x0e\x27\xdf\x68\x38\x1a\x8e\x0f\x69\xe0\xa8\x30\xd9\x57\x4d\x9f\x79\xe4\xad\xca\xf5\x96\xd4\x76\x1c\x76\xf0\xa6\xb2\x36\x2f\x20\xbe\xd7\x38\x77\xbc\x3b\x91\xe3\xab\xd4\xf5\x36\x4e\xe0\xa1\x5c\x65\xff\x37\x07\xb0\x10\x76\xa7\xf3\xf7\x65\xe2\xda\x39\x3c\x7b\x3f\x9a\xb8\xfb\xd8\x19\x8d\x21\xcc\x72\xe7\xb9\x71\xfe\x88\x2b\x76\x33\x40\x93\x26\x6b\xb6\x4f\x7a\xb1\x8c\x56\xa9\x32\xe9\xa0\xb9\x35\x69\x8d\xcd\xad\x59\x90\x8a\xdf\xcf\x54\xd5\xc7\x6e\xe6\x15\xd3\xb4\x52\xf4\x09\x6b\xb3\xe1\xed\x52\xd1\x81\xe1\x72\xc5\x2f\xb1\x87\x78\x3e\xf7\xca\x9e\x3d\x11\xf5\x81\x63\x00\x27\x7a\xb4\xc2\x20\xfa\xfd\xf7\x8d\xf5\x91\xcc\x79\x83\x83\x5b\xc0\xa3\xe4\xcb\xa0\x66\xca\xb5\xa1\x05\xa7\x16\x65\x8f\x46\xbd\xb5\x96\xd4\x53\x8b\xcd\xd3\xc3\xea\x97\x30\xf4\x56\x00\x95\x67\x9e\x12\xb1\xa4\xc8\x2c\x52\xc6\x7e\xfd\x18\x7d\x98\x0b\x25\x78\xbf\x81\x40\x94\xb6\x62\x30\x00\xa1\xd4\xdd\xba\xcd\xab\x35\xdd\xdb\x86\x99\xfa\xe5\x20\xbb\xb1\x30\xe9\xe6\xbd\x58\xcd\xda\x08\xca\xb7\x2d\xc3\x2b\xba\x30\x91\xd5\x34\xe8\x0c\x1c\xbb\x85\xde\x78\xb8\xfb\x55\xb9\x7d\x9a\x7c\x16\x1b\x97\x2d\x65\x42\x43\x80\x1c\x5c\x4c\x6b\x81\xa6\x29\x16\xfb\x19\xe5\x7e\xff\x0a\xc7\xb5\x93\xd3\x93\x9f\x86\xf0\x0f\x1a\x53\x4d\xb2\xf2\x2a\x66\xcf\x22\x38\x3a\x9b\x3c\x4e\xf4\xad\x12\x78\x0e\xd4\x64\x1e\x64\x9b\x7a\xfc\xe8\x3c\x84\x39\x53\x0f\xf5\xfd\xb1\xd2\x57\x68\xdd\x4e\xeb\x87\x87\x20\x6d\xa4\x7e\x28\x49\x99\xf8\x7f\x4c\x57\xd4\xf2\x64\x0c\xa3\xf7\xa7\xa7\xad\x66\xda\x01\x15\x6d\xf2\x00\xbd\x76\x3d\xb6\xe9\xed\x5a\x01\xd8\xa4\x86\x7f\x35\xbb\x36\x2b\x4b\xb7\x0c\xb4\xa2\x56\x0c\xb9\x12\x08\x13\xbd\xa6\x50\xe4\xc1\x67\x41\x39\x1d\x83\xff\xe8\x0d\xdd\x84\x64\x74\x3b\x84\x67\x95\x72\xbd\xa9\x1c\x74\xeb\x41\xd7\x44\x06\x4b\x3d\xa9\xc9\x6d\x70\x74\xa3\x29\x65\x3b\xe2\xd4\x45\x3a\x1a\xd5\xd5\xb0\xff\xf1\x53\xf5\x29\xac\x39\xa1\xd6\xae\x34\x6b\x85\xae\xac\x55\x7b\x48\x6a\x9c\x2b\x7b\x37\xa1\xb9\x64\x70\xce\x53\x7e\x4b\xa3\x67\xce\x3f\x81\xd4\xba\x03\x6b\x55\x5c\x35\x56\xa6\x4b\xdc\xe5\x4c\x05\x11\xdf\x44\xd2\xa7\xba\x9a\x49\xd0\x47\xd3\x5b\x93\x54\xd3\x5b\x7f\x24\x6e\x8e\xa6\xa8\x11\xf6\x5f\x4b\xa9\xf3\x19\xea\xe3\xd4\xca\x03\x82\xda\xac\xee\x5a\x4d\xbb\xd0\x94\x97\xa7\xe6\x86\x5f\xb8\x2a\x3e\xc1\x2d\x38\x77\x0c\x4f\x87\x87\x13\xe0\xd2\xf6\x34\x8f\x30\xaf\x9a\x43\xb0\x10\x63\x50\xfd\xed\xbb\x7a\xe1\x7d\xa6\xd1\xab\xe2\xbe\xcd\xe6\x07\xdd\x2f\x6a\x76\xca\xab\x45\x48\x61\x44\x95\xd7\x93\x7b\xc6\xa8\xac\x1b\x5c\x46\xd3\xbe\x43\xb5\xff\xd2\xf0\x25\xc9\x06\x75\x57\xbd\xd8\xe6\xab\x1d\x95\xb1\x4b\x96\xd2\xc7\x83\x4f\x3e\x8b\x50\xff\x50\x80\x4a\xb4\xb8\xe8\x1d\xfa\x36\xc5\x26\x86\xdc\x58\x7c\xad\xc2\xf6\xed\xdd\x64\x72\x0e\x73\xb3\xc7\x97\x87\x34\x82\x6e\xbf\x23\xd2\x3e\x9c\xe4\x25\x8d\xef\x0e\x0e\xca\xab\x4e\x2b\x15\xa1\x63\x30\xe9\x55\x18\x62\x4a\xec\xc2\xab\x83\x57\x80\xca\x65\x12\x9b\x42\x38\x3f\x1b\x4f\x68\xe9\x80\x40\xa6\x41\xc4\x9f\xf5\xba\xf0\x77\xb3\x10\x64\x77\x8c\x8d\xd4\x94\x89\x15\xdb\xfc\xfd\xa1\x52\x8d\x68\xa3\x4f\x22\x7e\x5d\xe3\x6a\xa2\xe7\x0b\xc5\x1b\x03\x71\x21\xf4\x52\x66\x5a\x7c\x50\x49\x4e\x37\x04\x18\xdc\x76\xdd\x86\xf3\xda\x74\x19\xca\xff\xd9\x8c\xe7\xe8\x1b\x87\x14\x70\x4c\x0f\x71\xe3\x33\xe6\x3b\x81\xb1\xa5\x5c\x46\xa5\x1e\x72\xa5\x31\x51\x1f\xa3\x5b\x28\xf3\x16\x5e\x2c\x9a\x96\x69\x19\xd4\xe6\x03\x3f\x32\x0b\xa2\xd7\x69\xea\x2a\xff\x8d\x8c\xee\x8a\x96\x69\x23\x12\x1e\x67\xf5\x26\x88\x0a\x71\x09\x8e\x52\xf8\xb1\xb9\x98\xbc\xc1\x28\x1e\xc9\x8d\xcc\x5a\x7c\x84\x2c\x5e\xa7\xc8\xd3\x69\x11\xd3\xde\x3e\x6f\xb9\x52\xda\xb3\x01\x4e\xd7\xec\x7e\xf9\xd5\xc8\x45\xfd\x5e\x68\xf5\xe5\xd4\x88\x59\x0f\xfe\x90\x03\x91\x28\xfe\xd7\x01\xf5\x3b\x4d\xc5\xfa\x58\xad\x1c\xb6\xad\xf9\xea\x63\xe6\xa5\x16\xad\x6a\xe7\xf8\x91\xca\x50\xaa\x86\x69\x67\xb3\xc4\xa0\x76\x96\xc0\x93\x55\x3c\xa1\x01\x24\x0b\xd2\xb1\x50\xd7\x42\xb1\x90\x5b\x74\x35\x69\xf8\x99\x7a\x92\x4f\x2e\x8a\x64\x63\xc3\x8d\xfd\x64\xbf\xaa\xba\x7f\x8e\x0e\xad\xde\x62\xc6\x3b\x69\xf1\xa8\x14\x67\x3f\x11\x34\xf1\x3c\xe7\x48\xb7\x9f\x1f\x6b\x5a\x7b\xac\x59\x41\x8e\x7a\xed\xff\x04\x00\x00\xff\xff\x4e\x27\x65\x26\x35\x24\x00\x00")

func templatesSearcherTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSearcherTmpl,
		"templates/searcher.tmpl",
	)
}

func templatesSearcherTmpl() (*asset, error) {
	bytes, err := templatesSearcherTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/searcher.tmpl", size: 9269, mode: os.FileMode(420), modTime: time.Unix(1469589171, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSearcher_enumsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x55\x4d\x6f\xe3\x36\x10\x3d\x8b\xbf\x62\xa0\x62\xb1\xd2\xc2\x8d\xee\x01\x72\x48\x37\xc2\x22\x80\x1b\xa3\x89\x17\x3d\x14\x05\x4c\x49\xb4\x44\x58\x26\x65\x92\xda\x6d\x50\xec\x7f\xef\x0c\xa9\x4f\x27\xed\xa5\xbe\x58\xe4\xbc\x79\x1c\xce\xbc\x19\x76\xbc\x3c\xf1\x5a\x80\x15\xdc\x94\x0d\x63\x59\x06\x5f\xf2\xa7\xfc\xf9\x7e\x9f\x3f\xc0\xe7\xdd\x43\xbe\x81\xfc\xe1\x71\xff\x02\xbf\x3f\x6e\xb7\xf0\x4b\x0e\xdb\xdd\xcb\x1e\x51\x04\xec\xad\x54\x35\x34\xce\x75\xb7\x59\x56\x4b\xd7\xf4\xc5\x4d\xa9\xcf\x59\x61\x24\x57\xd6\x71\x73\x12\x59\xad\x7f\x2e\x44\x2d\x5c\x56\x0b\x25\x0c\x77\xda\x64\xe1\x28\x61\x18\x93\xe7\x4e\x1b\x07\x09\x8b\xe2\xe3\xd9\xc5\xf8\x67\x9d\x41\x4e\x1b\xb3\xd4\x87\xb2\xeb\x82\x0f\x48\x0b\x5c\x81\x1e\x97\x4e\x43\x21\xf0\x7c\x51\x81\x54\xc0\xe1\xc5\x53\x3e\x8b\x4b\x2f\xac\x83\xa3\x6c\x1d\xd2\xbb\xd7\x4e\x2c\x18\x94\xf3\x94\x9f\x35\x85\xa6\x1c\x11\x56\xe0\x1a\x21\x0d\x84\x53\xa1\xd4\xbd\x42\xc7\x8e\x1b\xb4\x26\x23\xfb\xcb\x6f\x5b\x18\x82\x97\x5a\xa5\xac\x24\x02\x8a\x39\xbf\xcc\xec\x77\x20\xb5\xe3\x2c\x7a\xd2\x2e\xbf\xb0\x68\x2b\x4f\xc2\x2f\xc2\xc7\x17\x23\x38\x12\xef\x1b\xae\x56\x8b\x9d\xf1\x60\x61\xed\x68\x9c\xbf\x83\xed\xd1\x3e\xf5\x6d\xeb\xa9\xfc\x47\x48\xcb\xaf\xdc\xd8\x86\xb7\x7b\xf1\x97\x03\xcc\x61\x2b\xce\x82\x2e\x44\x95\xb0\x54\x0a\xdd\x72\x55\xdf\x68\x53\x67\xdd\xa9\xce\x84\x2a\x75\x85\xf7\xcb\x7e\x22\x87\xc1\x17\xf3\x73\xec\x55\x09\x89\x9e\x2e\x91\x2e\x79\x93\x14\x92\x3f\xfe\x2c\x5e\x9d\xd8\x80\x30\x86\xac\x7f\xb3\xe8\x1b\x37\x50\x71\xc7\x87\x8c\x31\x16\xd9\xef\xd2\x95\x0d\x68\xb2\x96\xdc\x0a\xc8\x2f\xb7\x2c\x8a\x3c\xe8\x0e\x62\x71\x89\x87\x7d\x9f\x99\xa5\x49\xd1\xc6\x68\xa5\x3c\x2d\x8d\x2d\xae\x17\x9e\xd7\x66\x15\xb6\x46\xc4\x22\xa5\x4b\x54\x3d\x6f\xbf\x83\xa4\x04\xff\x0b\x9a\x4c\x53\x64\x53\x45\x56\xf1\x4d\xbb\x6f\x71\xd7\xc4\xed\xca\x32\xe2\x43\x65\x97\x38\xe9\x77\x16\xb7\xbe\x06\xa8\xb0\x15\x63\xda\x2b\x71\xe4\x7d\xeb\xc8\x6a\x84\xeb\x8d\x02\x25\xdb\x0d\x60\x17\xdd\xe4\x54\xad\x63\x12\x7f\x55\xbc\x68\x05\xb5\xca\x39\x94\x15\x0e\x1f\xbe\x1d\x48\xd1\xd4\x3d\x58\x58\x1b\x6f\x40\xa7\x2c\xfa\xc1\x46\x8e\x50\xf0\x84\xce\x4b\x37\xc4\xc8\x7e\x78\xbd\x7d\x55\xe7\xff\xa1\xb8\xc9\x7b\xa1\xb9\x4f\xb3\xe8\x56\xe4\x49\x31\x04\x91\x06\xd5\x91\xac\x50\x6b\x70\x7b\x37\x48\xce\xde\xec\x8d\x3c\x27\x61\x91\x14\x18\xe6\x21\x3e\xa4\xb3\x10\x09\x3c\x4a\x91\xd4\x47\x19\xfa\xa4\x31\x7b\xd4\x4d\x61\x37\x08\x6f\x32\x0c\x2d\x1b\x6c\x5e\x77\x93\x29\x74\xef\xe4\xb5\x5d\x19\xa7\xee\x0e\xf6\xa5\xd8\x26\xcc\xaa\xf1\xdf\xe0\xbc\x1c\xde\xc3\x86\xde\x1f\x22\x9a\x95\x36\xc7\xb5\x98\x18\xd7\xa8\x35\xe9\xf5\x3c\x09\xe8\x41\x69\x13\x6a\x9c\x32\xd3\x4d\xd7\xe6\x71\xf8\xbc\x27\xbb\xff\x56\xdc\xc7\x0f\xf6\x23\x4d\x5e\x0d\xab\x61\x8c\xc2\xc3\x3a\xad\xa4\x37\x8b\x0d\x07\x74\x25\x69\xd4\x0e\x43\x5f\xe2\xb2\xf4\x43\x56\x1f\xe1\x7b\x23\x70\x64\xe3\xa7\x01\x0c\x13\x27\x7f\x98\xf5\x60\x1b\xdd\xb7\x15\xb9\xe3\xbb\xc0\xab\x0a\x47\x37\xf7\xde\xf7\x4f\x0f\x04\xc6\xaf\xdd\x33\x45\x86\xde\xc3\x5b\x07\x9d\x11\x9e\x59\xd8\xf0\x54\xcc\x07\x8f\xd3\x6d\x7c\x2e\x2c\x1c\x91\xe3\x55\xf7\x06\x8a\xde\xb9\xf9\x09\xb8\xc7\x27\x64\x76\xc3\x16\xc5\xe3\xb0\x81\x77\x06\x16\x3f\xdc\xde\x3d\xd3\x8b\xf6\x4f\x00\x00\x00\xff\xff\x60\x32\x65\x14\x6c\x07\x00\x00")

func templatesSearcher_enumsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSearcher_enumsTmpl,
		"templates/searcher_enums.tmpl",
	)
}

func templatesSearcher_enumsTmpl() (*asset, error) {
	bytes, err := templatesSearcher_enumsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/searcher_enums.tmpl", size: 1900, mode: os.FileMode(420), modTime: time.Unix(1469557591, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSqlcreatorTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x93\x41\x8b\xdb\x30\x10\x85\xcf\xd6\xaf\x18\x4c\x28\x4e\x70\xad\x4b\xe9\x61\xa1\x87\xed\xc6\x5b\x02\x4b\x96\x66\xb3\x87\x52\x7a\x90\xed\xb1\xa3\xae\x2d\xa5\xb2\x4c\xb7\x18\xff\xf7\x8e\x64\x87\xdd\x50\xa7\xb4\x27\x83\xa5\xf7\xe6\x7b\x33\xa3\xa3\xc8\x9f\x44\x85\x90\x1b\x14\x16\x19\x5f\xb1\xe0\x53\xba\x4d\x77\xd7\xfb\x74\x0d\x37\xf7\xeb\x34\x86\x2f\xf7\x8f\x3b\x48\xd7\x9b\xfd\x03\x5c\xef\x52\xb8\x7d\xdc\x6f\xee\x52\xc6\x82\x0a\x15\x1a\x12\x15\x50\x1a\xdd\x40\xdf\x27\xfb\x5f\x47\xdc\x8a\x06\x87\x81\x05\x5d\x2b\x55\x05\x07\x6b\x8f\x57\x9c\x57\xd2\x1e\xba\x2c\xc9\x75\xc3\x33\x23\x85\x6a\xad\x30\x4f\xc8\x2b\xfd\x36\xc3\x0a\x2d\x9f\xac\xb4\xe1\x9e\x43\x1b\xb6\xe2\x8c\xc9\xe6\xa8\x8d\x85\x88\x05\xe1\x2b\x83\xef\x8d\x96\x46\x2b\xde\xfe\xa8\x9f\xc3\xf3\xa3\x59\xef\xc9\x31\x24\x62\x4b\x7c\x2d\x84\x13\xe9\xc6\xdb\x0f\x43\xc8\x96\x8c\x71\x7e\x1e\xe0\x66\x54\x41\x81\xa5\x54\x24\xb2\x07\x04\xa9\x2c\x9a\x52\xe4\x08\x25\x9d\xd8\x83\x6c\xe1\x84\xeb\x9c\xe7\x0d\x5e\x44\x3d\x0b\xfc\x4f\x8c\x72\xf0\x24\xc9\x99\x60\x09\x11\xdd\x7d\xff\x2e\x06\x34\x46\x9b\x25\x1b\x3c\xd5\xc3\xe7\xbb\x79\xdf\xe6\x58\x63\x83\xca\xb6\x20\xdc\xa5\x73\x94\x4b\xaa\xd6\x9a\x2e\xb7\x0e\xa5\xc8\x60\xe5\x5a\x98\xac\x3f\x4e\x95\xb6\xf8\xf3\x92\xcc\xa0\xed\x8c\x3a\x55\xca\x44\x4b\x43\x9f\xbb\xc8\xca\x4e\xe5\x7f\x31\x8a\x5e\x55\x5d\xce\xf7\x8b\xd0\xc6\x6a\xf0\xe6\x82\x49\x5f\x64\x57\x45\x36\x38\xea\xbe\x5f\xa8\xae\xb9\x95\x58\x17\x2d\x5c\x7d\x80\x1a\x15\x24\xfe\x9a\xc8\x6a\x1c\xff\x0f\x63\xba\xb1\xf5\x34\x8e\x16\x8d\xef\xd9\x99\x35\x25\xcc\xb5\x29\xe8\x18\xac\xf6\xc3\x2e\x9c\x07\x05\x1d\x23\x45\x06\x56\x17\x70\x96\xf0\x3f\x53\x75\xf9\x72\x5d\x77\x8d\xf2\xc0\x5f\xbf\xd1\x44\xe8\xa1\xd0\xdf\xa0\xef\x8d\x50\xf4\x14\x17\x32\x86\xc5\xb3\x3b\xfd\x33\x0a\xed\xee\x82\x9a\x97\x8d\xe6\x61\xec\x65\xa8\x0a\xf7\xe2\x28\xe7\xa9\x75\xd3\x2e\x24\x1b\x9f\x36\x32\x49\x91\x51\xc7\xe3\x71\xf5\x9d\xdf\x49\x0f\x13\x4b\x0c\xff\x52\x3d\x4f\x7c\xf9\x51\xdc\xf7\xb2\x84\xda\x92\x00\x5e\x86\x30\x0c\xf1\xc4\x33\x7d\xdc\x1a\xff\x0e\x00\x00\xff\xff\x0d\xa1\x76\x38\x64\x04\x00\x00")

func templatesSqlcreatorTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSqlcreatorTmpl,
		"templates/sqlCreator.tmpl",
	)
}

func templatesSqlcreatorTmpl() (*asset, error) {
	bytes, err := templatesSqlcreatorTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/sqlCreator.tmpl", size: 1124, mode: os.FileMode(420), modTime: time.Unix(1468871773, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSqlupdaterTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x53\xd1\x6a\xdb\x30\x14\x7d\xb6\xbe\xe2\xce\xb0\x62\x87\x4c\x7a\x19\x7b\x08\xf4\xa1\x9d\xbd\x62\x28\x2d\x4b\x93\x87\x31\xc6\x90\xed\x6b\x47\x6b\x2c\xbb\xb2\xdc\xb5\x84\xfc\xfb\x24\xcb\xde\x12\x88\x57\x28\xdb\x93\xb1\xae\xee\x39\xe7\x9e\xa3\xdb\xf0\xec\x9e\x97\x08\x5d\x93\x73\x8d\x84\xcd\x88\x77\x15\xdf\xc4\xcb\x8b\x55\x1c\xc1\xc7\xdb\x28\x9e\xc3\x97\xdb\xf5\x12\xe2\x28\x59\xdd\xc1\xc5\x32\x86\x4f\xeb\x55\x72\x1d\x13\xe2\x95\x28\x51\x99\xa6\x1c\x0a\x55\x57\xb0\xdb\xd1\xd5\x73\x83\x37\xbc\xc2\xfd\x9e\x78\x5d\x2b\x64\x09\x1b\xad\x9b\x05\x63\xa5\xd0\x9b\x2e\xa5\x59\x5d\xb1\x54\x09\x2e\x5b\xcd\xd5\x3d\xb2\xb2\x7e\x97\x62\x89\x9a\x0d\x50\xb5\x62\x4e\x87\x22\x33\x46\x88\xa8\x9a\x5a\x69\x08\x88\xe7\x17\x95\xf6\x0d\xa5\x7f\x00\xf4\xa3\xaa\x85\xaa\x25\x6b\x1f\xb6\x4f\x3e\x01\xf0\x5f\x22\x19\xa0\x7d\x12\x12\xc2\xd8\xb1\xde\xb5\xab\x81\x68\x41\x6f\x10\x84\x34\x3f\x05\xcf\x90\x68\x73\x65\xe2\xea\x78\x07\x76\xc4\x73\x87\x81\xa3\x58\xe2\x43\x87\xad\x3e\xd5\x36\x94\x42\x40\xa5\x6a\x45\xbc\x08\xb7\xa8\xf1\xf2\x39\x89\x82\x24\xb2\x90\x1f\xde\x8f\xb5\x7d\xaf\xf2\xee\xf3\xf5\x69\xf6\xaa\xd9\x62\x85\x52\xb7\xc0\xed\x25\x48\x79\x6b\x92\x18\xed\xeb\x65\x4f\xf5\xb6\x5a\x75\x99\xb6\xb2\xf3\x14\x66\xd6\x3f\x1a\x5d\x0e\x7c\x37\xf8\x73\xaa\x4d\xa1\xee\x94\xb4\x7c\x59\x2d\x0b\x51\x76\xea\x80\xb0\xe8\x64\xf6\x97\xe6\xe0\x80\x29\x3c\xed\xa7\x91\xe3\x18\xe0\x6c\x02\x64\x97\xa7\x8b\x3c\xdd\x0f\x4a\xdd\xa1\xd5\xf2\x88\xaa\xb7\x61\xda\x6f\xe3\x2c\xe8\xba\xf7\x89\xcb\x1c\xf0\x09\xb3\x4e\x63\x0b\x42\x3b\xe1\x81\x82\xd9\x04\x69\x08\xaf\xcb\xd6\xce\x63\x06\xbe\xd3\x6a\x0e\x8f\x7c\x6b\x2a\x73\x5b\x80\xc5\xf9\xe8\x19\xbd\x42\xed\x3a\x0d\x75\x70\x76\x84\x6f\x9e\xa8\x27\x8a\xbe\xe1\xcd\x39\x48\xb1\xb5\x70\xa3\x3f\x66\x19\x68\x6c\x39\x8a\xc0\xef\xbf\x30\xec\x8f\xdd\x38\x3b\x63\x61\xe9\xa7\xcd\x58\xc0\xdb\xd6\xef\xd5\x38\x98\x20\x0c\x89\x67\x4c\xf5\xbe\x3b\x89\xe7\xa0\x68\x9e\xd2\xd8\xb8\x14\x1c\x8f\x40\x83\xaf\xdf\x7e\xbf\xfb\xdd\x3e\xa4\x94\x5a\xa5\x83\x2e\xd3\x3b\x64\xf3\xe7\x59\x8f\xd2\xd0\x06\xe4\x8e\x87\xf9\x5f\x91\xc4\xf4\xb6\x58\x7b\xba\xde\xdc\xe9\xb9\xad\x83\x49\xb4\x80\x24\x9a\xbb\x71\x5f\x8e\xc7\x11\xba\x78\xd4\x3f\xce\xe4\xd0\xa4\xff\x9b\xc8\xaf\x00\x00\x00\xff\xff\x55\x14\x71\xcc\xe4\x05\x00\x00")

func templatesSqlupdaterTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSqlupdaterTmpl,
		"templates/sqlUpdater.tmpl",
	)
}

func templatesSqlupdaterTmpl() (*asset, error) {
	bytes, err := templatesSqlupdaterTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/sqlUpdater.tmpl", size: 1508, mode: os.FileMode(420), modTime: time.Unix(1468871773, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesUpdaterequestTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x57\x4d\x6f\xdb\x46\x10\x3d\x73\x7f\xc5\x94\x4d\x6a\xca\x50\xc9\x1e\x8a\x1e\x04\xf8\xe0\x56\x4a\xe0\x22\xb5\x51\x5b\x3e\x14\x86\x01\x2d\xc5\x95\x44\x5b\x5c\x32\xbb\x4b\xc7\x2e\xa1\xff\xde\x99\x5d\x8a\x22\x2d\x59\x76\x12\x07\xe8\x21\xb9\x44\xd2\xee\x7c\xbd\xf7\x66\x66\x5d\xf0\xe9\x2d\x9f\x0b\x28\x8b\x84\x1b\xc1\xa2\x43\xe6\xbd\x1f\x9d\x8e\xce\x8f\xc7\xa3\x21\xfc\x71\x36\x1c\xf5\xe1\x9f\xb3\xcb\x73\x18\x0d\x4f\xc6\x17\x70\x7c\x3e\x82\x77\x97\xe3\x93\x0f\x23\xc6\xbc\xb9\x90\x42\xa1\x51\x02\x33\x95\x67\x50\x55\xe1\xf8\xa1\x10\xa7\x3c\x13\xab\x15\xf3\x4a\x9d\xca\x39\x2c\x8c\x29\x06\x51\x34\x4f\xcd\xa2\x8c\xc3\x69\x9e\x45\xb1\x4a\xb9\xd4\x86\xab\x5b\x11\xcd\xf3\x9f\x63\x31\x17\x26\xaa\x5d\xe5\x2a\x72\x79\x28\x76\x18\x31\x96\x66\x45\xae\x0c\x04\xcc\xf3\x67\x99\xf1\xf1\x3f\x6d\x14\x7a\xd5\xf4\x51\xc8\x69\x9e\xe0\x97\xe8\x46\xe7\xd2\x67\x3d\xc6\xa2\xa8\x9b\xc3\xbb\x54\x2c\x13\x48\x35\x70\x98\xd9\x8f\x9f\x30\x8b\x54\x82\x59\x88\xee\x45\x40\xb7\xe5\xd4\x90\x03\xb3\xe0\xc6\x9a\xc4\x4b\x01\x26\x87\x78\x8d\x4c\xc2\x0c\xde\xdf\x19\x40\x1a\x1b\x7b\x24\xcb\xec\x00\xb1\xc8\x15\x2c\xc4\xb2\x98\x95\x4b\x29\xb4\x66\xd3\x1c\xab\x85\xa0\xaa\xde\x98\xda\x10\x06\x47\xd0\x86\xaa\xaa\x14\x97\x48\xc1\x9b\xb4\x0f\x6f\xee\xed\xe9\xa5\x8d\x49\x49\xd8\x18\x7a\xb5\x6a\xd9\xdb\x2f\xf7\x61\x9d\x7b\x55\xa5\x33\x10\x1f\xd1\x1a\x7e\x79\x74\xcd\xe5\x77\x04\x69\x6e\x78\x55\x09\x99\xd8\x60\xf6\x7f\x07\xd7\x5f\x5c\xe9\x05\x5f\x8e\xc5\x3d\x16\x9d\x15\x4b\x91\x09\x69\xb4\x25\x4d\x13\x6b\xf9\x12\x13\x0b\x73\x35\x8f\x8a\xdb\x79\xd4\x20\xfe\x23\x19\xd4\xb6\x48\xd5\xac\x94\x53\x08\xf4\x0e\x6c\x7a\xed\x08\x41\x0f\x82\xab\xeb\xf8\xc1\x88\x3e\x08\xa5\x72\xd5\x83\x8a\x79\x77\x5c\x01\xd6\xca\xc1\x51\x8b\xb2\xd2\x48\xd3\x74\x01\x9a\x4e\x5f\x04\xcd\x94\x6b\x22\x66\x37\x3e\x03\xe6\x79\xd6\xff\x11\xf8\xf6\xe7\x3f\x2f\xce\x4e\xdd\x91\x4f\xfe\x1d\x2a\x5e\x22\x66\xbc\x5c\x1a\xba\xad\x84\x29\x95\x04\x99\x2e\xfb\x80\xb2\x0b\x47\x94\xeb\x2c\xf0\x2f\xe5\x5a\x14\x99\x2b\x0a\x26\x6f\xef\x26\x40\x82\x42\x9d\x60\x59\xda\xef\x83\xee\x31\x0f\xdd\xd5\x3e\x5c\xb9\x01\xc5\xef\xf5\xc9\x23\x5b\x59\xdc\x2f\x65\xf6\x15\xc8\x37\xd6\x2d\xec\x0f\x77\x81\xdf\x09\x13\xc4\x75\x3a\x3d\x87\x3e\xc1\x8b\x98\x13\xa6\x75\x57\x85\x63\x95\x66\x81\xfb\x12\xc4\x98\xf0\xc4\x9f\xf4\x5a\x84\x18\xf5\x79\x94\x6c\xe1\x4d\xe8\x1e\x6a\x64\xe2\x29\xb2\xf6\x12\xb2\x9f\x8b\x83\xb7\xfa\x80\x5a\x31\x87\x27\xfa\x94\xc8\x31\xaa\x43\xcf\x86\x90\x61\x6c\xef\xd8\xee\x74\x87\xda\x8e\x09\x49\x3f\xe4\x33\xfb\xd9\xcd\x10\x0c\x50\x62\x6d\xf5\x18\xb9\xf8\xfb\x03\x7c\x2c\x85\x7a\xd8\xdb\x04\x2d\xef\xd8\x04\x0e\x60\x0b\xff\xab\x2b\xbd\x2e\xcc\x21\x3f\x8c\xb7\x75\xbe\x29\xde\xf7\xeb\xda\x3b\x19\xbb\x98\xe7\x02\x8b\xc2\xa9\x85\x24\xa4\x38\xc4\x70\x80\x6a\x61\x08\x87\x82\x2b\xbc\x86\xf3\x59\xd3\x9c\x23\x6b\x3b\x1d\x6d\x3d\x6d\x37\x21\xc0\x89\x81\x29\x97\x34\x40\xb5\xc0\x91\xbf\x4c\xff\xc5\x45\xc1\x65\x82\x3e\xb4\xc6\x71\x8a\xb6\xb1\x30\x9f\x84\x90\x74\xe1\x2e\x9d\x52\x1c\x0d\x24\x97\x3e\xa0\x3c\x11\x65\x8b\xf6\x7a\xcb\x60\x12\x84\x36\x6e\x0e\x63\x9b\x25\xdc\x31\x90\xbb\xd9\xbb\xa9\x4e\xe0\x9e\x0c\x49\x1a\xbf\xfd\x0a\x13\x5a\x16\x03\x3f\x4d\xfc\x09\xf3\xdc\x6d\x8d\x12\x2a\xae\xb6\x69\xbb\x46\x0b\xa1\x66\x7c\x2a\xaa\xd5\xda\xce\xad\x02\x8d\xc6\xab\x57\x9f\x9f\xa5\xda\x53\xcb\x0b\x06\xa9\x93\x55\x8a\x2a\x4d\x48\x3d\x19\xbf\x15\x01\x55\xe6\x7e\x6f\x57\x43\x3d\x4d\x5b\xea\x56\x3c\xf4\xe1\x8e\x2f\x4b\xbb\x91\x9c\xfc\x4a\x15\xae\x61\x41\x9f\x9e\xb6\x11\xe8\x18\x2f\x87\x9d\x1c\xd0\x89\x47\x9b\x07\x8f\x7f\x38\xa2\x56\xb2\x06\x9d\xe9\x89\x67\xf8\xd3\x8a\x6e\xb6\xb2\xab\x33\x0a\x74\xef\x1a\x47\x81\x8d\xcf\xec\x25\x25\x34\x76\xbc\x4b\x7e\x57\xde\xe4\xbf\xa1\x60\x00\x2d\x97\x7d\x3a\x41\x52\x07\x60\xff\x61\x11\x27\xc3\xfe\xda\xa9\xcd\x87\xf8\x5b\xe7\x1f\xb8\x40\xbd\x2e\x89\xbf\xa7\x92\xab\x87\xcf\xa2\xd1\x99\x7c\x19\x91\xce\xf6\x3b\x95\xaf\x4a\xe5\x6b\xef\x55\xa4\xf2\x70\x1f\x97\xcf\x2e\x58\x7a\xdd\xa4\x4f\x60\x80\x15\xd5\x84\xd8\x8a\x1a\x5f\x41\xdc\x87\x9f\x52\x22\x65\x9b\x93\x1a\x02\xcb\x06\x39\xb0\xf0\xd0\x0b\xef\x8a\x30\xbb\x0e\x03\x3b\xe6\xc8\x96\x22\xdf\x40\x4b\x3c\xfb\x27\xdc\x3e\x21\xa1\xf3\x35\x55\x18\xe1\x09\x21\xda\xf4\x28\x68\xb1\x63\x05\xe2\x11\x15\x72\x04\x45\xd8\x85\xac\x7e\x20\x61\xd4\xde\x33\x2a\x5c\x0b\xd0\xf3\x6e\xae\x8a\x47\x6a\x6b\x29\x1d\xc1\x64\x3b\xf6\x7b\x13\xf5\x8b\xbb\xfc\x2b\xc4\x51\xb7\xfa\x77\x79\xfc\x6f\xe5\xf1\x5e\x18\x7a\x22\x6c\x44\x51\xff\x15\x1a\x76\x1f\x14\x4d\x41\x2f\x54\x80\x75\x8b\x23\xde\xbd\x3d\xaa\x26\xb0\xa5\x65\x13\x7a\x4c\x4f\x3c\xfb\xf4\xfc\x16\x19\x34\xde\x3b\x4f\xcf\xcd\x43\x31\x6c\x2e\xd0\x4b\xb1\x49\x6a\x8d\xd8\xb7\x48\xa9\xf6\x8d\x09\xed\x56\x0b\x25\x58\x4b\xea\x85\xfb\xcf\xbe\xcd\x9f\xdb\x80\xb5\xcb\x2b\x7b\x39\xec\x3c\xca\xb7\xd6\x97\xe3\xc9\x19\xd4\xa0\x1c\x27\x89\xf3\xf6\xfa\x98\x34\xae\x03\xf7\x47\xc6\x76\x8b\xac\x6b\x7b\xdc\x53\xd8\x11\x6d\x79\x6f\x1a\xa3\x23\xfa\x17\x77\xf9\xa3\x7e\x71\x50\x6d\xc0\x59\xb1\xff\x02\x00\x00\xff\xff\xf4\x5c\x8f\xfc\x2f\x12\x00\x00")

func templatesUpdaterequestTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesUpdaterequestTmpl,
		"templates/updateRequest.tmpl",
	)
}

func templatesUpdaterequestTmpl() (*asset, error) {
	bytes, err := templatesUpdaterequestTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/updateRequest.tmpl", size: 4655, mode: os.FileMode(420), modTime: time.Unix(1468871773, 0)}
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
	"templates/ginCreator.tmpl": templatesGincreatorTmpl,
	"templates/ginUpdater.tmpl": templatesGinupdaterTmpl,
	"templates/searcher.tmpl": templatesSearcherTmpl,
	"templates/searcher_enums.tmpl": templatesSearcher_enumsTmpl,
	"templates/sqlCreator.tmpl": templatesSqlcreatorTmpl,
	"templates/sqlUpdater.tmpl": templatesSqlupdaterTmpl,
	"templates/updateRequest.tmpl": templatesUpdaterequestTmpl,
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
		"ginCreator.tmpl": &bintree{templatesGincreatorTmpl, map[string]*bintree{}},
		"ginUpdater.tmpl": &bintree{templatesGinupdaterTmpl, map[string]*bintree{}},
		"searcher.tmpl": &bintree{templatesSearcherTmpl, map[string]*bintree{}},
		"searcher_enums.tmpl": &bintree{templatesSearcher_enumsTmpl, map[string]*bintree{}},
		"sqlCreator.tmpl": &bintree{templatesSqlcreatorTmpl, map[string]*bintree{}},
		"sqlUpdater.tmpl": &bintree{templatesSqlupdaterTmpl, map[string]*bintree{}},
		"updateRequest.tmpl": &bintree{templatesUpdaterequestTmpl, map[string]*bintree{}},
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

