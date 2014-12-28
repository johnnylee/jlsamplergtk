package jlsamplergtk

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _data_gui_glade = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5d\xcd\x72\xdb\x36\x10\xbe\xe7\x29\x58\x1e\x32\xed\x74\x14\xdb\xf9\x4f\x63\x2b\x13\x67\xea\x34\xa9\x92\xe9\x34\x6e\x7b\xe4\x40\x24\x24\x22\x06\x01\x16\x04\x6d\xe9\x29\x7a\xec\xfb\xf5\x49\x0a\x52\xb2\x23\x5a\x00\x08\xfe\x58\x26\x29\xdc\x44\x6a\xb1\xc0\x2e\xf6\xc3\xee\x02\x4b\xf2\xf8\xcd\x22\xc2\xce\x25\x64\x09\xa2\xe4\xc4\x3d\x7a\x74\xe8\x3a\x90\xf8\x34\x40\x64\x7e\xe2\xfe\x71\x7e\x36\x7a\xe9\xbe\x19\x3f\x38\x46\x84\x43\x36\x03\x3e\x1c\x3f\x70\x9c\xe3\xef\x46\x23\xe7\xe6\xce\x88\xc1\xbf\x53\xc4\x60\xe2\xcc\xf9\xc5\x8f\xce\x93\x47\x87\xce\x68\x94\x93\xd1\xe9\x57\xe8\x73\xc7\xc7\x20\x49\x4e\xdc\xf7\xfc\xe2\x2f\x44\x02\x7a\xe5\x3a\x28\x38\x71\x23\x80\xc8\xfa\x3a\x23\x16\xe4\x31\xa3\x31\x64\x7c\xe9\x10\x10\xc1\x13\xd7\x07\xc4\x9b\x51\x3f\x4d\xdc\xf1\x19\xc0\x09\x3c\x3e\xb8\x26\x90\xd3\x73\xc4\x31\x74\x1d\xce\x00\x49\x30\xe0\x60\x8a\xc5\xcd\x25\x14\xcd\x3f\x4e\x9c\x2f\x20\x8a\x31\x64\x5b\x3c\xfc\x10\xe1\x60\xf5\x5b\x36\xe0\x53\xba\x58\x8d\x96\xd3\x38\xfb\x7d\x4d\xb9\xdd\xfb\x25\x4a\x90\xe8\xd1\x1d\x9f\xb3\x74\x6b\xa8\x75\xc4\x93\xb5\x89\x00\x9b\x23\xe2\x61\x38\xe3\xee\xf8\x79\x85\x16\x0c\xcd\xc3\x8a\x4d\x84\xc4\xd5\x1a\x4c\x29\xe7\x34\x32\x6c\x43\x19\x82\x84\x03\x2e\xac\xce\x1d\x0b\xf3\xe3\xc8\x07\xd8\xa4\x61\x12\x03\x5f\xd8\xa6\x3b\x7e\x29\xa5\x2e\xcc\x67\xc9\x9c\x26\x2b\x9b\x98\x50\x10\x14\xe7\xb6\xd6\xfc\xd6\x9d\x63\xad\x90\x52\x5d\xca\x05\x95\x0b\x7b\x4e\xe7\x73\x0c\x4f\x53\x31\x35\x64\x25\x35\x16\xe2\xae\xd1\x70\xca\x89\x7b\x9b\xc5\xf6\x60\x30\x98\x42\x2c\xc5\x95\x97\xa9\x4e\x05\x2d\x1d\xcb\x34\x81\x1e\xf0\xb3\xc9\xf7\x40\x1c\x43\x20\x58\xfb\xb0\x44\x47\xb5\x27\xa5\x64\x62\xaa\x36\x65\xd0\x87\xe8\x12\x26\x5e\x00\x67\x20\xc5\xbc\x3a\x07\x03\x4c\xde\xb5\xe2\x32\x3e\x29\x09\x20\xc3\x88\x94\xaa\xef\xf8\x60\x65\x56\x5b\xf7\x85\x95\x5e\x08\x33\x2d\xef\x0e\x2e\x62\x40\x82\x1a\xe3\x9c\x21\x8c\xab\x2b\x38\xa6\x09\x5a\xad\x2c\x87\x3a\xa9\xa4\xc3\x3f\x3e\x90\x20\xcb\x14\x6d\x6f\x7d\x1f\xe2\xc9\x0a\x2f\x19\xd6\x40\x76\x9d\xe3\xe7\xc8\x00\x68\x6d\x58\x73\x65\x0d\xab\xd1\xfd\x51\xa8\x47\xc8\x96\xad\xd3\x39\xed\x4f\x83\xb4\x90\xa3\xdd\x5a\xc8\xcf\x84\xb3\xe5\xca\x38\xbe\x0a\xe6\x9f\xc5\x68\x76\x64\x19\x55\x9b\x22\xb2\xee\xd5\xf3\x43\xc0\xdc\xf1\x7f\xff\xfe\x53\xa5\x39\x87\x0b\xae\x08\xc6\x4a\x1d\x46\x7b\xe6\x54\x55\xe8\xa6\xd6\xf4\x78\xb7\xd6\xf4\x25\x46\x84\x40\xb6\xe1\xd8\xd7\x37\x3a\xb1\xd8\xf4\x78\x55\x78\xd2\xc6\x3c\xca\xe4\x97\xcb\x5e\x4b\xee\xea\x32\x57\xf6\x93\x12\x59\xb7\xe4\x34\x09\xb9\xcf\x58\xb6\xce\xe5\x56\x3a\xcb\x7e\xde\xf6\x86\xf7\x1a\x6c\x87\x26\x8b\xc5\xd6\x08\xeb\x34\xca\x7d\xad\xb7\x00\x18\xcd\x75\xb1\xc9\x56\x36\x10\x02\x91\x25\x7b\x7c\x19\x0b\x8d\x10\x4a\xd4\xfd\x99\x86\x29\xd9\x00\x22\xe1\xd7\xd7\x51\xca\xf5\x65\x67\x83\x94\xd2\x94\x54\xa9\x71\x91\x29\x7b\x31\x08\x82\x3c\x9b\x3a\xd2\x2c\xcf\x2a\xed\xc9\x35\x78\x93\x43\x4e\xe9\xe2\x89\x44\x6b\x8d\x34\xd7\x54\x7b\xb2\xf6\xc6\xf9\xb6\x8e\x89\x3e\xf7\x36\xd3\xa5\x5c\x9f\xe7\x94\xe2\x29\x60\xd7\x7b\x2d\xf9\x85\xcc\x1a\x5b\xd1\x6d\x1b\xfa\x2d\x17\x52\x2d\xe8\x56\x36\x0e\x39\x17\x6a\x4d\xe4\xe9\xb8\x6e\xd4\x75\xd3\xc1\x56\x55\xd9\x96\x3a\x65\x7c\x32\x53\xe0\x28\xf6\x94\xd1\x64\xbe\x03\x91\xac\x15\x58\xb7\x97\xbb\x50\xa3\x3a\xaf\x9a\x78\x62\xcc\x4d\x86\x6a\x9c\xb8\x17\x18\x29\x82\xb1\x62\x67\xaa\xc0\x4c\x37\x24\xe3\x20\x4d\xc7\x24\xa4\x11\x9d\x43\x02\xa9\x49\xaa\xb2\x29\x55\xe9\x90\xa5\x01\x76\x81\xa0\x05\x18\x27\xe0\x12\x5a\x18\x2b\xf9\x94\xc2\xf8\x8b\xd0\x9f\xe3\x53\x91\x1c\x53\xdc\x13\x18\x7b\xd9\x98\x2d\x8c\x3b\x02\xe3\x32\xbd\x94\xeb\xa4\xb1\x3e\xea\xe5\xa0\xaa\xd6\x46\x7b\x98\xdf\xa4\xd7\x8a\xa7\xd5\x5d\xd5\x50\xed\x3d\x43\xc1\x6a\xd1\x5b\xe3\x35\xbf\xd1\xf5\x50\x4d\x9e\x50\xac\x0e\xd2\xf4\x69\x81\x72\x18\x14\xa7\x11\xf1\x6e\x62\xe2\xa7\x75\x98\x30\x91\xd7\xd5\xc2\x4c\x1d\xa7\xa5\xde\x9d\x7e\x5c\xd1\x61\x75\xca\xb9\x84\xeb\x7c\x3a\xe1\x80\xf1\xf6\xd7\xf9\xb7\x51\x8c\xd1\x6c\xa9\xd9\xfd\x2e\xf0\xbd\xb3\xf5\x39\x4f\x65\x01\xe7\xc0\x0f\x8d\xd6\x04\x39\x17\x91\x42\x37\x67\x72\x85\x02\x1e\x6a\xb7\xce\xf5\xed\x43\xb8\x3a\xfa\x32\x64\x70\x5f\x41\x5e\xb6\x99\xba\x19\xe4\x81\x28\xbe\x67\x9c\x34\x61\x63\xb4\xcd\x55\xc6\xa4\xe2\xa1\x80\x8e\x55\x80\xe6\x88\x27\xda\x2d\x73\x3d\x03\x92\x46\x90\x21\xbf\x1b\xd1\x53\x01\x9d\xb5\x81\x61\xd1\xa9\xa4\xa8\xe0\xcd\x54\x1b\x72\x2a\x81\xf7\xcc\x9b\x39\x73\x10\x45\x60\x80\xfe\xac\xb6\xc5\xef\x09\x62\x6e\xfb\xb3\xdc\x0e\xde\xf6\xda\xa9\x59\x7f\xa4\x18\x4c\xfb\xfe\xc8\xa2\xab\x40\x51\xc1\x1f\x3d\xb5\xfe\x48\xe9\x8f\xfe\x84\x98\xfa\x48\x10\x46\x29\xe6\x03\xf4\x49\xb5\x57\x93\x3d\x41\xcd\x6d\x9f\x74\x09\xf1\xa7\xbc\xa8\xd2\xba\x24\xeb\x92\x64\x5c\x2c\xb8\x94\x14\x15\x5c\xd2\x33\xeb\x92\x94\x2e\xe9\x9d\x20\x74\x78\xc8\x60\x12\x52\x1c\x0c\xd0\x27\x19\xed\x54\xcb\x98\xec\x09\x6c\x6e\xfb\x24\x5f\x8c\xf5\x3c\xb7\x07\xeb\x96\x36\xdd\x92\xa6\x48\x52\xcf\x60\xe0\x6e\xc9\xe2\xab\x40\x51\xc1\x2d\x3d\xb7\x6e\x49\xe9\x96\xce\x40\x00\x47\x88\x38\xdf\x47\xc9\x0f\x03\x74\x4a\xcf\x2c\x68\x36\x29\x4a\x9d\x12\x07\x69\x66\x11\x1f\xaa\xd6\x1a\x0d\xd4\x27\x0d\xdc\xa5\x58\x74\x14\x28\xb6\xd1\x01\x63\xc0\x00\xa7\xeb\x0a\xe2\xe4\xfa\x52\x57\x43\x2c\x13\xf7\x1e\x1d\x4a\x9f\x16\xeb\xda\x91\xdf\xda\x1c\x6b\x27\xee\x3d\x37\xc7\x1e\xd5\xd9\xf4\xc9\x1c\xb5\x8f\x84\xe8\x98\xec\x89\x39\xaa\x03\xee\x57\xfd\x31\xc8\x9d\x07\xdc\xbf\xc2\xe5\x28\x8d\x9d\x99\x88\xb2\x86\x1a\x74\xbf\xb0\x61\xc5\x26\x85\x49\xd0\x6d\xc3\xed\x3d\x08\xb7\x2d\x2e\x0a\x14\x26\xb8\x78\x97\xda\x43\xbb\x7d\x80\x46\xe9\xa3\x9f\x2a\x26\xfb\x02\x0d\x45\xe8\xdf\xa3\xa2\xc4\x3e\x45\x30\xaf\x6c\xe8\xbf\x49\x61\x56\xf3\x37\x01\x4b\xe9\x3b\x42\x74\xf2\x76\x69\xb5\xee\x5a\x29\x7b\x91\x95\x97\x40\x83\x97\x82\xe9\xf8\x0d\xb8\xee\xa3\xf6\xd9\x98\x2d\x45\x54\x52\x54\xc8\xf7\x5f\xf6\xc7\x0b\xed\x3c\xdf\xcf\x57\xc5\x2e\x16\xc7\xd7\xde\x6d\xb5\x98\x51\x52\x18\xc7\x6d\xbd\x2e\xde\x2d\x3c\xe6\x5a\x7b\xe9\x2d\xbe\xaf\xb2\x36\x9b\x5a\xaf\xa0\x59\xb3\xda\x0d\xd0\x6a\xfb\xdb\x6e\x3d\xb7\x65\x28\x46\xd7\x90\xd6\xa3\x9a\xc4\x3e\x2d\xfc\x4d\x8b\x64\x87\x9e\x21\x69\xde\xd9\x7a\xd8\x1f\x8b\xbc\x97\xd3\x11\x3f\xe5\x83\x3e\x1e\xb1\x7b\x5d\x05\x8a\x2a\xd0\xe9\x51\xe1\xc5\xce\xa1\xf3\x09\x2c\x9c\x18\x71\x3f\x74\xa6\x90\x0c\xb1\xc0\xfc\xc8\x26\x1b\x05\x8a\x2a\xc0\xe9\x51\x89\xc8\xce\x81\x73\x9e\xdd\x8a\x69\xa2\x7b\x15\x79\x81\x73\x9f\x30\x73\xff\x79\x43\xb7\x31\xb3\x75\xe6\x78\x6d\x0c\x3d\xde\xc8\xb6\xc7\x8e\x3a\x2e\x16\x1e\xea\x5e\x4a\xe1\x91\x07\x18\xa7\x22\xbe\x10\xd1\x86\x45\xc8\x3e\x20\xc4\x06\x5d\x05\x8a\x2a\x41\x57\x8f\x0e\xe7\x77\x1e\x74\xfd\x06\x88\x83\xe9\x55\xb7\x42\xae\x56\x76\xc6\x9a\x56\xb1\xef\x31\x60\x7a\x7d\x2a\xb2\x03\xc0\x84\x62\x82\x07\x88\x18\xfb\x64\x6b\x81\xa2\x0a\x62\x7a\x74\xba\xb1\x73\xc4\xfc\x1e\x25\x0e\x47\x51\x17\xf7\x91\x5b\x81\x4d\xd3\xe7\x53\xf6\x18\x36\xf6\x89\x70\x3d\x6c\x86\x19\x99\xd9\xfa\xfb\x02\x45\x79\xb2\x0f\xc8\xe4\xe6\x3b\xca\xa6\xb2\x0e\x34\xcd\xb7\xc5\x93\x36\xeb\xb9\x21\x68\x07\x5b\xbf\x88\x81\x5b\x70\x59\x70\xc9\xb9\xd8\x04\x49\x49\x51\x0a\x2e\x16\x25\xe7\x48\xfa\x09\x5c\x9d\xb0\x03\x05\xd7\xc0\xb1\x61\xb3\xa0\x02\x85\x09\x36\x6c\x50\x77\xcb\xef\x0c\xf0\x95\x73\xad\x60\xcb\x26\x4c\x05\x0a\xe3\x9a\xe3\x1e\x6d\x2f\xf4\x29\x7f\x6f\xfa\xba\xaa\x3d\xae\x39\x7e\xd1\x1f\x8b\xbc\x97\x0d\xaf\x81\x9e\xac\xd8\x52\xe3\x02\x85\x49\x70\x64\xb3\x72\x1b\x1d\xa9\xb9\x58\x70\x29\x29\xb6\xc0\xf5\x2e\x84\xfe\xc5\x26\xba\x22\xb4\xc8\x9f\x86\x4d\x2a\xe2\x4b\x53\xf3\x8e\x16\x0e\xce\x59\xd6\xd5\x65\x1f\x3e\x92\xdc\x84\x0d\x83\x3e\x44\x97\x30\xf1\x02\x38\x03\xf9\x27\x30\x1a\x89\x76\x17\xea\x5a\xac\xdd\x7b\xed\x5a\xcc\x80\x81\x2b\x0f\x91\x00\xf9\x79\x2c\xde\xbd\xe5\xa6\x15\x5f\xde\xb4\x54\x75\x60\xd1\x6f\x8c\x81\x0f\xb3\x4f\x29\x40\x76\x60\x7b\xb2\x3d\xb5\xd4\xd3\x80\xbf\xf8\x6d\x80\xdf\x12\xec\x6a\x74\xa7\xd6\x9b\xa2\x91\xbc\x81\x94\x78\x35\xf7\x0e\x5f\xc6\x37\xc1\xc0\x16\xb7\xad\xe8\x63\x23\x0f\x4e\x40\x14\x63\xc8\x3e\x90\x19\x9d\x48\x5b\x37\x70\xe1\x4d\x72\x5e\xf3\x38\xe7\x21\xe6\xaf\xa7\x0f\xe7\xfc\xf5\x67\xca\x9d\x09\x05\x01\x0c\xb2\x5b\x07\xf9\xbd\x2a\x5d\x64\xfe\x3b\x02\xec\x22\x8d\xcb\x44\x33\x9e\x1f\x19\xa1\x1c\x25\xb5\x90\x51\x1d\x0d\x95\x11\x20\xb1\xfa\x5b\x72\x16\x65\xdc\xf8\xf3\xdb\x1f\xc7\x07\x88\x70\xc8\x66\x62\x0d\x1a\x3f\xf8\x3f\x00\x00\xff\xff\x00\x64\x11\xe0\x00\x9b\x00\x00")

func data_gui_glade_bytes() ([]byte, error) {
	return bindata_read(
		_data_gui_glade,
		"data/gui.glade",
	)
}

func data_gui_glade() (*asset, error) {
	bytes, err := data_gui_glade_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "data/gui.glade", size: 39680, mode: os.FileMode(420), modTime: time.Unix(1419792286, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"data/gui.glade": data_gui_glade,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"data": &_bintree_t{nil, map[string]*_bintree_t{
		"gui.glade": &_bintree_t{data_gui_glade, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
