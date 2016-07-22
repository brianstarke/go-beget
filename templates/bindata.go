// Code generated by go-bindata.
// sources:
// templates/ginCreator.tmpl
// templates/ginSearcher.tmpl
// templates/ginUpdater.tmpl
// templates/searchRequest.tmpl
// templates/sqlCreator.tmpl
// templates/sqlSearcher.tmpl
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

var _templatesGinsearcherTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x55\xc1\x52\xdb\x30\x10\x3d\x5b\x5f\xb1\xf5\xa1\x63\x67\x82\x45\x7b\xcc\x4c\x0e\x80\x0d\x4d\x87\x49\x68\x12\x0e\xbd\xa1\xd8\x1b\x47\xc5\x91\x8c\x24\x07\x98\x8c\xff\xbd\x92\xed\x29\x84\x42\x80\x03\x87\x36\x87\x44\xd9\x7d\xfb\xf6\xed\xfa\x8d\x5c\xb2\xf4\x9a\xe5\x08\x1a\x99\x4a\x57\x84\xf6\x88\x77\x96\x8c\x93\xe9\xd1\x3c\x89\xe1\x64\x12\x27\x7d\xf8\x39\xb9\x9c\x42\x12\x8f\xe6\x33\x38\x9a\x26\x70\x7a\x39\x1f\x9d\x27\x84\x78\x39\x0a\x54\xcc\x60\x06\x4b\x25\xd7\xb0\xdd\x46\xf3\xfb\x12\xc7\x6c\x8d\x75\x4d\xbc\x4a\x73\x91\xc3\xca\x98\x72\x40\x69\xce\xcd\xaa\x5a\x44\xa9\x5c\xd3\x85\xe2\x4c\x68\xc3\xd4\x35\xd2\x5c\x1e\x2c\x30\x47\x43\x3b\x2a\xa9\x68\xab\x03\x15\xe9\x51\x42\xf8\xba\x94\xca\x40\x40\x00\xfc\x47\x14\x39\x17\x07\xb9\x14\x3c\x75\x27\xff\x49\xf2\xd7\x5a\x72\x25\x05\xd5\x37\xc5\x9d\x4f\x42\x42\x28\x85\x31\xde\xee\xa8\x9b\x35\x4d\xbe\x31\x91\x15\xa8\xe0\x96\x17\x05\x28\x34\x95\x12\xc0\x60\xd5\x45\x97\x52\xc1\x33\x45\x53\xbc\xa9\x50\x1b\x1d\x59\x5e\x47\x3d\x5f\xe1\x9f\x12\xa6\x75\xb5\x46\x0d\x66\xc5\x8c\xfd\xc2\x3d\xf5\x6d\xd3\x05\xc2\xc5\x64\x36\xb7\x1b\x6c\xb6\xe5\xf8\xbe\xcf\x26\xe3\x88\x2c\x2b\x91\xee\x57\x1d\x64\x0b\xe8\xb9\x19\xa3\xf8\x38\x04\xbb\x87\xa8\x4b\x9c\xba\xd2\x2d\xf1\xba\x81\x1c\x53\x90\x42\xcf\x21\x4e\xa4\x30\x78\x67\x42\x97\xf6\x36\x4c\x81\xde\x37\xa2\xc5\xa0\x52\x30\x18\x42\x1a\x1d\x73\x91\x39\x65\xc1\x67\xad\xec\x4a\x3d\x8f\x2f\xc1\x25\x3f\x0d\x41\xf0\xa2\xe1\xf3\xd2\xa8\x41\x7c\x3d\x3c\xec\xb7\x7a\xb6\xbe\x85\x48\xe5\x0f\x1c\x34\x4a\xdc\x39\x08\xeb\xd0\x61\x5b\x71\xf6\x54\x3b\x32\x85\xba\x2a\x8c\xee\x43\xd7\xcf\x0e\x3e\xfb\x71\xfe\x8c\xb2\x66\xec\x30\x6a\xff\x04\x1f\x21\xe5\x71\x65\x27\x2b\x24\x0f\xa0\x9a\xd4\xcf\x1a\xea\x44\x56\xc2\xfc\x6b\x7e\x7a\x2c\xfa\xff\xb3\xd3\x9b\xdd\xd4\xac\xe1\xc3\xcd\xd4\x55\xa6\xae\x99\xad\x6c\x35\xd6\x6f\xf1\xd6\x19\x9a\xe3\xfb\x51\xfc\x4e\x77\xc1\x59\x32\xb7\xb0\x37\x9a\xeb\x5e\x56\x36\xb1\x41\x48\xa5\xbd\x9b\x55\x95\xba\x5b\x9d\xb5\x1c\xb2\x32\x68\xbb\x9a\x95\x0d\x94\xcc\xfe\x94\x4c\xb1\x35\xc8\xa5\x63\xbc\x1a\xf0\xec\x0a\x02\x1e\x61\x04\x57\x94\x95\x9c\x6e\xbe\x50\x17\x0b\x1b\x49\xce\xb4\xa3\x18\x6c\x4b\x60\x85\x96\x0f\x4e\xb6\xc1\x25\xc7\x22\xb3\x34\x4f\x5e\x1c\x96\x94\x6b\x48\x59\x51\x58\x09\xa3\xf8\x05\xef\xee\x2e\xe5\x35\xf7\x02\xec\xb7\x2f\xd8\x0f\xcf\x5a\x6f\x5e\xb8\xe9\x02\x9f\x67\xbe\x35\x44\x93\x58\x42\x81\x22\xe0\x59\x08\xc3\x21\x1c\x76\x70\x80\x3d\xbe\xf0\xc7\xd2\x4d\x58\x2a\xb9\xe1\x19\x66\xbe\x7d\xce\x6d\x4d\xf7\xac\xdd\xb1\x6e\xd9\xdf\x69\xd6\x66\xee\x53\xb7\xb9\x60\x07\x34\x8a\xfb\x76\x82\x07\xc5\x3b\x26\x7e\x55\xef\xae\x8f\x5f\x92\xfa\xf7\xdd\x18\x92\x1d\xa4\x73\xf0\xef\x00\x00\x00\xff\xff\xfe\xe3\x7c\x23\x50\x08\x00\x00")

func templatesGinsearcherTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesGinsearcherTmpl,
		"templates/ginSearcher.tmpl",
	)
}

func templatesGinsearcherTmpl() (*asset, error) {
	bytes, err := templatesGinsearcherTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/ginSearcher.tmpl", size: 2128, mode: os.FileMode(420), modTime: time.Unix(1468871773, 0)}
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

var _templatesSearchrequestTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x57\xdf\x6f\xdb\xb6\x13\x7f\x16\xff\x8a\x83\xbe\x2d\x6a\x07\xfe\x4a\x7b\x0e\x90\x87\xb6\x51\xba\x0c\x69\x83\xc5\xce\x86\xa2\x28\x16\x4a\xa2\x6c\xad\xb2\xe4\x92\xb4\xdb\x4c\xf0\xff\xbe\x3b\x92\xfa\x69\x27\x4d\xb6\x3d\x85\xe2\xf1\x7e\x7d\xee\xee\x73\xce\x86\x27\x5f\xf8\x52\x80\x12\x5c\x26\x2b\x16\x9e\x30\xef\x5d\xf4\x21\xba\x79\xbd\x88\xce\xe1\xed\xf5\x79\x34\x83\x8f\xd7\xb7\x37\x10\x9d\x5f\x2e\xe6\xf0\xfa\x26\x82\x8b\xdb\xc5\xe5\x55\xc4\x98\xb7\x14\xa5\x90\x5c\x8b\x14\x32\x59\xad\xa1\xae\x83\xc5\xfd\x46\x7c\xe0\x6b\xb1\xdf\x33\x6f\xab\xf2\x72\x09\x2b\xad\x37\xa7\x61\xb8\xcc\xf5\x6a\x1b\x07\x49\xb5\x0e\x63\x99\xf3\x52\x69\x2e\xbf\x88\x70\x59\xfd\x3f\x16\x4b\xa1\x43\x67\xaa\x92\xa1\x8d\x43\x48\x76\x12\x32\x96\xaf\x37\x95\xd4\x30\x61\x9e\x9f\xad\xb5\x8f\x7f\x94\x96\x68\x56\xf9\xe8\xde\xff\x91\xd1\xc6\x94\xcf\xa6\x8c\xb1\x30\x1c\x46\x78\x91\x8b\x22\x85\x5c\x01\x87\xcc\x1c\xbf\xa1\xb9\xbc\x04\xbd\x12\xc3\x87\x80\x3e\xb7\x89\x26\x03\x7a\xc5\xb5\x51\x89\x0b\x01\xba\x82\x58\xa0\x6e\xa1\x85\x44\x0c\xaa\x72\x06\x0a\xa3\x75\xc7\x4a\x82\x14\x7a\x2b\x4b\x91\x06\x4c\xa3\xb1\xa3\xde\x4b\x6d\x02\x8b\xca\xed\xfa\x15\xc2\x88\x4a\x2b\x51\x6c\xb2\x6d\x51\x0a\xa5\x58\x52\x61\x4e\x30\xa9\xeb\x17\xda\x29\xc2\xe9\x19\xf4\x51\xae\x6b\xc9\x4b\xac\xde\x8b\x7c\x06\x2f\xbe\x1b\xe9\xdc\x64\x4d\x11\x1a\x1f\x6a\xbf\xef\xe9\x9b\x8f\xef\x81\x4b\xac\xae\xf3\x0c\xc4\x57\xd4\x86\x9f\x46\xcf\x6c\x7c\x67\x90\x57\x9a\xd7\xb5\x28\x53\xe3\xcc\xfc\x9d\x9a\x90\xdf\x73\xa9\x56\xbc\x58\x88\xef\x88\xc8\x7a\x53\x88\xb5\x28\xb5\x32\xf5\x56\x54\xf0\xaa\xc0\xc0\x82\x4a\x2e\xc3\xcd\x97\x65\x28\xca\xa4\x4a\xb1\x70\xe1\xff\x48\xc1\xe9\x62\x91\xb3\x6d\x99\xc0\x44\x1d\xc1\x66\xda\xf7\x30\x99\xc2\xe4\xd3\xe7\xf8\x5e\x8b\x19\x08\x29\x2b\x39\x85\x9a\x79\x3b\x2e\x21\xe5\x9a\x83\x6d\x0a\x6c\x09\x85\x35\x4c\x56\xa0\x48\xfa\x24\x68\x12\xae\xa8\x30\xc7\xf1\x39\x65\x9e\x67\xec\x9f\x81\x6f\xae\x7f\x99\x5f\x7f\xb0\x22\x9f\xec\x5b\x54\xbc\x54\x64\x7c\x5b\x68\x7a\x6d\x2b\x0e\x65\x5e\xcc\x00\x1b\x36\x88\x28\xd6\x6c\xe2\xdf\x96\x4d\xc7\xac\x6d\x52\x70\xf7\x72\x77\x07\xd4\x6d\xd8\x44\x98\x96\xf2\xb1\x79\xa6\xcc\x43\x73\xce\x86\x4d\x77\x42\xfe\xa7\x33\xb2\xc8\xf6\x06\xf7\xdb\x72\xfd\x2f\x90\x6f\xb5\x7b\xd8\x9f\x1c\x03\x7f\xe0\x66\x12\xbb\x70\xa6\x16\x7d\x82\x17\x31\x27\x4c\xdd\x3c\x06\x0b\x99\xaf\x27\xf6\x63\x12\x63\xc0\x77\xfe\xdd\xb4\x57\x10\x2d\x9f\x57\x92\x03\xbc\x09\xdd\x13\x85\x95\x78\xa8\x58\x8f\x16\xe4\xf1\x5a\xbc\x7a\xa9\x5e\xd1\x28\x56\xf0\xc0\x9c\x52\x71\xb4\x1c\x94\xa7\x2b\xc8\x79\x6c\xde\x98\xe9\xb4\x42\x65\x38\xa4\xa4\x8b\x2a\x33\x67\x4b\x30\xe8\x60\x8b\xb9\x39\x8e\x99\xff\x7a\x05\x5f\xb7\x42\xde\x3f\x3a\x04\x3d\xeb\x38\x04\x16\x60\x03\xff\x7f\xde\xe9\x2e\x31\x8b\xfc\x79\x7c\xd8\xe7\x5d\xf2\xbe\x6f\x72\x3f\x61\x83\x88\xad\xcf\x1b\x81\x49\x21\x6b\x61\x11\x72\x24\x31\x64\x57\x25\x34\xe1\xb0\xe1\x12\x9f\x21\x57\x2a\xe2\x39\x66\xc9\x99\xb2\x21\xd6\x1b\x18\x0a\x00\x2e\x35\x24\xbc\x24\x7e\x55\x02\xa9\xbd\xc8\xff\x42\x5a\xe5\x65\x8a\x56\x94\x12\x29\x8b\x85\xfe\x26\x44\x49\xd2\x5d\x9e\x90\x1b\x05\xd4\x2d\x86\x77\x11\x64\x03\x76\xb3\x9f\x30\x06\x02\x1b\xd7\x83\x36\xb3\x12\xd0\x6e\x39\x2c\xf5\x30\x7e\x4b\xfa\x06\x6a\xb7\x46\x82\xc1\x03\x8b\x2b\xf3\x2e\x0c\xff\x2b\x1c\x90\x23\xc6\xac\x10\xee\xfe\x54\x55\x79\xea\xdb\x55\xa1\xfc\x3b\xe6\x5d\xcb\x54\xc8\x37\xf7\xc3\x00\x9a\x4b\xcf\x73\x0a\x95\xbd\x20\x05\xeb\x6e\xec\xc5\xb2\xb4\xd7\xda\xa7\x37\xf8\xfa\xc1\xe2\xb8\x78\xdc\xd2\x33\x67\xb5\x11\x49\x9e\xe5\x09\x01\x36\xdc\xe0\x8f\x80\xe4\xec\x74\x18\xd9\x40\x8e\xc4\xd6\x0f\x8d\xf2\xf8\x8d\x17\x5b\x1a\x01\x54\xcf\x78\x22\xea\x7d\xf3\x62\x47\x02\x03\xcd\xc6\xfe\x12\x80\x16\x77\xeb\xad\xbd\x6f\xc0\x71\xdf\xa4\xf3\xb6\x2a\xd3\x5c\xe7\x55\x39\x56\xea\x04\x4e\x2b\x69\x2e\x8e\xc3\xd4\x94\xc0\x20\x44\xdb\x1c\xd2\x5c\x8a\x44\xe7\x3b\xd1\x2e\xff\x67\x22\xd6\x98\xfc\x07\x60\x9d\x0b\x95\xe0\xf0\xd1\x90\xc4\x55\x55\x34\xf2\x14\xaf\x5d\xf8\x21\x5c\x36\x1b\x00\x8e\xf7\x69\x87\xb5\x79\xfe\x4e\xe8\x05\xd1\xc2\x01\x5d\xd1\xa2\x89\x89\x23\xb4\xa1\x47\x62\xaf\x86\x98\xe4\x68\x45\x0c\xec\x4f\x07\x26\x07\x1c\xd5\x31\x4a\xd0\x3e\x20\x4a\xd9\x37\x91\x34\xb3\xd3\x8f\x83\x26\xf5\xf7\x9f\x23\xfc\x9d\x99\x14\x1c\x27\x59\x3d\x3d\x0a\x67\x0e\x63\xf8\xf4\x79\xd4\x07\x14\x8e\x1b\x3f\xe2\xc7\x03\x79\x8d\x31\x79\x44\x44\x7f\xe0\xe2\xa6\x17\x96\x51\x55\x23\x37\x3c\xeb\x2c\x98\xb5\x37\x52\x47\xa1\x2d\xeb\x29\x64\x81\x39\xcc\xe8\xca\x74\x3b\x5d\x99\x83\xb9\x6a\xda\x98\x6e\x9b\xb3\x11\xb4\xad\x4a\x92\xf6\x83\x44\xfb\xd6\x35\x2d\x40\xbe\xd9\x60\x4f\x4c\xdc\xc5\xcc\xcd\xb1\x59\x4e\x2d\xe4\x4e\xd8\x21\xdd\x34\xe1\x18\xe9\xeb\x9b\xf3\xe8\x06\xde\x7c\x7c\x3e\xd8\xce\x22\x15\xbc\xc1\xa2\x25\xb5\x36\x8e\xb1\x88\x70\x72\x30\xa9\xf6\xb2\xc3\xab\xeb\xf7\x81\xbc\xbb\x9e\x51\x96\x6d\x52\x57\xf9\x3a\xd7\x07\x29\x5d\x5d\xbe\xbf\x5c\xb8\x7c\x9e\x9e\x8e\xb1\x85\xc9\xe0\xb4\xf4\xe3\x97\x81\x11\xf4\x80\xcc\x32\xda\x67\x07\x38\x5e\x5c\xcc\xa3\xe7\x7b\xb5\xd6\x8e\xb9\xb5\x12\xe7\xf7\x75\x9a\xba\x36\xe6\x69\x4a\xcc\xd4\x1f\x90\x27\x7a\x6b\x6d\x4c\xb2\x07\xe8\x67\x06\xbb\x31\x39\xe3\x3a\xfd\x01\x1b\xcf\x20\xf9\x21\xf7\x4e\x69\xf8\x68\x6a\x1e\xdc\x24\xbd\xbe\xc8\x9a\x66\x70\xb3\xb3\x6b\x26\xa7\x1b\x9c\xaa\x37\x36\xbd\xa9\x49\x7a\x33\x83\x23\xd3\x9b\xdd\x76\x68\xba\x3b\x9c\x9b\xa9\x43\x77\xde\x8d\x07\x42\x6e\x6b\x3a\x9a\x8b\x27\x62\xdc\x59\x7a\x04\xe4\x5c\x8d\x68\xdd\xfc\x2f\xd3\xb5\x3b\x9c\x1d\xdd\x1f\xc7\x20\xea\xcf\x4b\xdf\xec\x70\x4c\xdc\x4f\x87\x71\xcb\xce\xa3\xab\xe8\xed\xc2\x1a\x7b\x16\xc9\xd2\x7b\xcb\xb1\x2d\xd3\x5b\x23\x8e\x59\xcd\xed\x63\x84\x6a\xde\x5a\x3e\x35\xc7\x1e\xa7\xd1\x37\x2a\x04\x83\x9f\xbd\x63\x6e\x33\x01\xef\xd9\xdf\x01\x00\x00\xff\xff\xe3\x16\x77\x94\xba\x10\x00\x00")

func templatesSearchrequestTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSearchrequestTmpl,
		"templates/searchRequest.tmpl",
	)
}

func templatesSearchrequestTmpl() (*asset, error) {
	bytes, err := templatesSearchrequestTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/searchRequest.tmpl", size: 4282, mode: os.FileMode(420), modTime: time.Unix(1468871773, 0)}
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

var _templatesSqlsearcherTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x56\xdf\x6f\xdb\x36\x10\x7e\x96\xfe\x8a\x9b\x80\x06\xb2\xe1\x49\xdd\xf6\x66\x20\x0f\x69\xac\x04\x01\x02\x17\x8d\xd3\x87\x21\x08\x50\x4a\x3a\xd9\x6c\x25\xd2\x25\xa9\xb4\x81\xe1\xff\x7d\x47\x4a\x72\xa4\xcc\x0a\xba\x75\x1b\x06\xf4\xc5\x36\xa5\xfb\xf1\xdd\x77\xdf\x1d\xbd\x65\xd9\x27\xb6\x46\xd0\xc8\x54\xb6\xf1\xe3\xa9\xef\x5d\x26\xcb\xe4\xe6\xec\x36\x59\xc0\xf9\xdb\x45\x32\x83\xdf\xdf\xbe\xbf\x81\x64\x71\x75\xbb\x82\xb3\x9b\x04\x2e\xde\xdf\x5e\x5d\x27\xbe\xef\xad\x51\xa0\x62\x06\x73\x28\x94\xac\x60\xb7\x8b\x6e\x1f\xb7\xb8\x64\x15\xee\xf7\xbe\x57\x6b\x2e\xd6\xb0\x31\x66\x3b\x8f\xe3\x35\x37\x9b\x3a\x8d\x32\x59\xc5\xa9\xe2\x4c\x68\xc3\xd4\x27\x8c\xd7\xf2\xe7\x14\xd7\x68\xe2\x36\x94\x54\x71\x83\x03\x95\x3f\x8d\x7d\x9f\x57\x5b\xa9\x0c\x84\xbe\x17\x14\x95\x09\xe8\x4b\x1b\x45\x61\x75\x40\xe9\x83\x5e\xd0\x8f\x95\xe4\x4a\x8a\x58\x7f\x2e\xbf\x06\xc3\x57\x47\xf3\x75\x59\x28\x0e\x80\x21\xd4\x1a\x82\x16\xff\x95\xcb\xb9\xdf\x07\xfe\xc4\xf7\xe3\x78\x58\xd6\xaa\xf5\x03\xae\xc1\x6c\x10\xb8\x30\xa8\x0a\x96\xa1\x6f\x83\x8c\xd9\x76\x46\xb0\xf3\xbd\xe6\x69\xd8\x00\xb8\xc1\xcf\x35\x6a\x73\xcc\xaf\x7d\x35\x81\xf0\xee\xde\x01\x8c\x06\x46\x33\x40\xa5\xa4\x9a\xf8\xde\xb9\xac\x85\xf9\x2b\xf1\x08\xce\x6f\xbf\x3e\xf9\x5f\xa2\x79\xf3\x78\xc1\xb1\xcc\xc3\xc2\x7e\x0e\x9d\xdd\x8b\x19\x3c\xb0\xb2\xee\x55\xbb\xdb\x53\x9c\xe9\x4b\xb0\xf6\x8e\xbb\xd5\xbb\xeb\x11\x4a\xaa\x6d\x89\x15\x0a\xa3\x81\x59\x2b\x48\x99\x26\x1d\x1d\x9a\xef\xd8\x1c\xf5\x26\x11\xd4\x99\xb1\x6c\xe6\x29\x4c\x6d\xcf\xa3\xc5\x9b\x36\xe5\x12\xbf\x8c\xfa\x29\x34\xb5\x12\x36\x65\x26\x45\xc1\xd7\xb5\xa2\x9c\x0a\xb7\x12\x0a\xa9\xe0\x51\xd6\x7e\x51\x8b\xec\xa5\x10\x61\x2f\xe1\x64\xa4\xdd\x04\xab\x49\x04\x27\x63\x61\x76\x79\x3a\xcf\xd3\x7d\xc7\x92\x7b\x6a\x31\x3d\xa0\x72\x8c\x8c\xf7\x8f\x7a\x00\x46\x3a\xca\x98\xc8\x01\xbf\x62\x56\x1b\x52\x2f\x37\x0d\xf6\x50\xc1\x74\x2c\xeb\x04\xfe\x61\xf1\x35\xa5\xea\xba\x24\xd0\xf3\x53\x38\x6a\xbb\xa3\x1a\x3d\x62\x6c\x65\x54\x2b\x23\xed\xdc\xad\x43\xd7\xee\xe8\xb2\xdd\x24\x2b\x2c\x31\x33\x84\x3f\x3c\x19\x60\x9c\x41\xc1\x4a\x8d\x34\x90\x1e\x2f\x9c\xf7\x4f\xa7\x20\x78\x69\x01\x74\x64\xd3\x91\xcc\x2a\x13\x25\x16\x5b\x11\x06\xee\x1b\xda\xcd\x62\x77\x91\x25\xcd\x36\xfa\x05\x76\xe7\xf0\x4a\x07\x0e\x5f\x13\x26\x9c\xd0\x84\xd8\x0a\x6c\xce\x53\x50\x51\x9e\x46\x0d\xc8\xf0\xa4\xad\x7c\x06\xc3\xea\x22\xe2\xac\x3f\x28\x51\x14\x1d\xc7\x4d\x4f\xda\x6d\x16\x9d\x4b\x61\x18\x17\x3a\xec\x25\x9e\x41\x20\x24\x28\xf9\x45\xdb\x9e\x37\xc9\x88\x31\x13\x38\xda\x0f\x65\x1f\x50\x50\x58\x7a\xbc\x6f\xe0\x3e\x7f\x49\x71\x5b\xb1\xb9\x75\xf1\x5f\x68\xed\xfb\xf6\x92\xad\xf1\x81\xa9\xae\x02\xd2\x56\x33\xf4\xb6\xf4\xa6\x04\x67\x0e\x1f\x68\x92\x82\x4c\x98\xe0\x43\x53\xf8\xf7\x2b\x8d\xd2\xbc\x2c\xb4\xd7\x3f\xa2\xcc\x9e\x15\xbd\x94\x87\xce\x34\x06\x98\x07\x93\x31\xf9\xdd\xbd\xbe\x8f\x5c\xcf\xfa\x3a\x7c\xba\x76\xec\x5d\xca\x1a\x45\x0a\x8e\x82\xae\xc9\x0a\xcd\x46\xe6\x56\x7b\x6d\xa0\x8f\x35\x91\x26\x05\x76\xf8\xd2\x47\xa8\x98\xc9\x36\x96\x6e\x29\xc8\xdb\xfe\xd7\x28\x11\xdc\x05\xf6\x2d\xe2\xfc\x57\x2e\xbd\x4e\xb3\xdf\xaa\x79\x2b\xd7\xfe\x39\x3a\xcb\xf3\x0b\x5e\x52\x3a\xfa\xc3\xe3\x39\x5c\x33\xfa\xe1\x60\xd8\x1f\x07\x11\x27\xef\x06\xc7\xb3\xe5\x82\xce\x93\xe7\xd1\xae\x79\xc5\x0d\xa9\xe9\x97\xff\xcd\x02\xee\xf5\xfc\xe8\x90\xc0\x1d\x5d\xaf\xae\xee\xf9\x2b\x7d\xdf\x0e\x89\x3b\x47\x8b\xd4\xf9\x59\x3b\xab\xd9\x3f\x4f\xce\xd3\xb2\x80\x23\x2d\x1a\x4c\x16\xc1\xe8\xc6\xea\x6f\x4f\x55\xbf\x66\x2b\xea\xbe\xec\x0f\xb1\x1b\xb5\xff\x11\x00\x00\xff\xff\x18\x78\x4f\x87\x62\x0b\x00\x00")

func templatesSqlsearcherTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSqlsearcherTmpl,
		"templates/sqlSearcher.tmpl",
	)
}

func templatesSqlsearcherTmpl() (*asset, error) {
	bytes, err := templatesSqlsearcherTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/sqlSearcher.tmpl", size: 2914, mode: os.FileMode(420), modTime: time.Unix(1468871773, 0)}
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
	"templates/ginSearcher.tmpl": templatesGinsearcherTmpl,
	"templates/ginUpdater.tmpl": templatesGinupdaterTmpl,
	"templates/searchRequest.tmpl": templatesSearchrequestTmpl,
	"templates/sqlCreator.tmpl": templatesSqlcreatorTmpl,
	"templates/sqlSearcher.tmpl": templatesSqlsearcherTmpl,
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
		"ginSearcher.tmpl": &bintree{templatesGinsearcherTmpl, map[string]*bintree{}},
		"ginUpdater.tmpl": &bintree{templatesGinupdaterTmpl, map[string]*bintree{}},
		"searchRequest.tmpl": &bintree{templatesSearchrequestTmpl, map[string]*bintree{}},
		"sqlCreator.tmpl": &bintree{templatesSqlcreatorTmpl, map[string]*bintree{}},
		"sqlSearcher.tmpl": &bintree{templatesSqlsearcherTmpl, map[string]*bintree{}},
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

