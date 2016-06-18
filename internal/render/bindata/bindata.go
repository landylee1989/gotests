// Code generated by go-bindata.
// sources:
// templates/call.tmpl
// templates/function.tmpl
// templates/header.tmpl
// templates/inline.tmpl
// templates/inputs.tmpl
// templates/message.tmpl
// templates/results.tmpl
// DO NOT EDIT!

package bindata

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

var _templatesCallTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8e\x51\xca\xc2\x40\x0c\x84\xaf\x12\x4a\x1f\xfe\x1f\x4a\x0e\x20\x78\x00\x5f\x44\x54\xf4\x79\xd9\xa6\x1a\xa8\xab\xa4\x51\x91\x90\xbb\xdb\x2d\xd5\xf5\x29\x30\xf3\x4d\x66\xcc\x5a\xea\x38\x11\x54\x31\xf4\x7d\xe5\x6e\xf6\x64\x3d\x03\x6e\x29\x12\x3f\x48\xb2\xc2\x1d\xa4\xab\x02\xae\x86\x9d\xca\x3d\xaa\xbb\x2a\x9a\x51\x6a\xb3\xfb\x21\x01\xdd\x8b\x8a\xeb\x70\x21\xf7\x3f\x33\x09\xe9\x44\x50\x73\x03\x35\xf5\xb0\x58\x02\x6e\x82\x8c\xa6\x92\x0c\xf3\xf7\x9a\xdd\x1b\xf8\x66\x4b\xdf\x51\x58\xf3\x86\xdf\xbe\x29\x9d\xcb\x26\x10\xf7\xaf\x1b\x8d\xe4\x21\x08\x87\x96\xe3\xb8\x01\x0b\x3b\x9d\xff\xf9\xbe\x03\x00\x00\xff\xff\x9d\x9f\x57\x19\xec\x00\x00\x00")

func templatesCallTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesCallTmpl,
		"templates/call.tmpl",
	)
}

func templatesCallTmpl() (*asset, error) {
	bytes, err := templatesCallTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/call.tmpl", size: 236, mode: os.FileMode(436), modTime: time.Unix(1455586613, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\x41\x4f\xdc\x3a\x10\x3e\x27\xbf\x62\x58\x3d\xd0\xe6\x69\x9f\xb9\x83\x38\x3c\x04\xef\x89\x43\x59\x04\xa8\x1c\xaa\xaa\x32\xc9\x84\x5a\x35\x4e\xea\x38\x50\x64\xe5\xbf\x77\xec\x38\x59\x87\x4d\x59\x54\xf5\xc4\x7a\xe2\xf9\xbe\x6f\x66\xbe\x4c\xb0\xb6\xc0\x52\x28\x84\x45\xd9\xaa\xdc\x88\x4a\x2d\xba\x0e\x52\x6b\xff\x81\xbf\x4a\x38\x3a\x01\xd6\x75\x69\xea\x9e\x81\xb5\xec\x16\x1b\x73\xc9\x1f\xb1\xeb\x96\x06\xfe\x36\x74\x12\xea\x81\xdd\x66\x60\xd3\xc4\x9d\x1a\x97\xf1\xe9\x73\x63\x74\x9b\x1b\x17\x4c\x1c\x90\x28\x41\x55\x86\xf0\xd8\x95\x16\xca\x5c\xa8\xba\x35\x0d\xc1\x26\x49\x72\x78\x08\x0e\x13\x0a\x6c\x72\x2d\x6a\xc7\xcf\x5c\x5c\x11\x09\x10\x0c\xc1\x07\x10\x54\x85\x4f\x71\xbf\x9f\x85\xf9\x0a\xec\x1a\x73\x14\x4f\xa8\x7b\xa4\x40\xc4\x2e\x9a\x1b\xcf\xde\x47\xc7\xf0\x7f\x02\x65\x11\x48\x3d\xed\x90\x0d\xa5\x7f\xc2\xc6\xdb\x03\x51\x7f\xd2\x5c\x3d\xe0\xeb\x74\x6b\xfd\xd9\xf5\xc6\x77\xe5\xa5\xc6\x28\x65\x04\xf0\x07\xd9\x0c\xcf\x22\xd2\x81\x6d\x14\x31\x41\x82\x34\xd9\xae\x39\xfe\xed\x0a\x72\x6d\xbb\xe2\x9a\xfa\x64\x50\x6f\xba\xb9\x09\xb1\x99\xc4\x50\xcd\x5c\xae\xb5\x3e\xb2\x55\xd3\x0c\x77\xa5\x7b\x88\x6b\x6c\x5a\x49\x33\xa7\x49\x98\x56\xab\xe6\x5c\xeb\x4a\x8f\x4a\xce\x7f\xd4\x98\x1b\x2c\x40\xf7\xd7\x76\xe8\x09\x60\x83\x98\x3b\xae\xcc\x7b\xb4\xcc\x70\x3f\x53\x2a\x1d\xe1\xbe\xaa\xe4\x34\xa9\x73\x8e\x74\x96\x5b\x9f\xad\x8f\xe0\xdf\xa2\x00\x67\x5a\xc8\x79\x83\x4e\x1e\xdd\x28\xa9\x36\x6b\x83\x61\xc9\x4b\x97\xfc\x1b\x52\x2a\x7c\x59\x81\x31\xce\xdc\xd6\x7a\xac\xa0\xbc\xf7\xbc\xfd\x1d\x5b\x4e\x46\xef\x81\xfd\x54\xa9\x54\x7f\x97\x53\xfe\x41\x20\x0b\x2d\x60\x1f\xb9\x6c\xa9\x11\x16\x76\xb9\x93\xf5\xef\xe8\x11\x69\x66\x91\x57\x57\x30\x63\xf2\x6e\x97\xd9\x02\xc7\xb6\x5f\x86\xc2\xee\xb4\x30\x43\xb9\x13\x1b\x51\x55\x07\xf7\x2f\xd4\x23\x76\xda\x96\x25\x6a\xbb\x93\x8b\x00\xb9\x2a\x60\xe9\xbb\xbf\x56\xf2\x25\x9e\x6e\xb6\x1d\x5f\x2b\xf4\x4d\xc9\x60\x10\x65\xf0\xb1\x96\xdc\xd0\x3e\x0b\xb6\x5b\xd0\xde\xf1\x36\xda\x3c\xc9\xb9\x94\x7d\xf8\xdd\x9e\xa2\x68\x3f\xa0\xd7\xa2\x08\x1a\xc9\x6b\x7e\x80\x73\x0c\xc7\xa3\x63\x96\xee\xde\xde\x09\x28\x21\x33\xf7\x97\x66\x33\x38\xd5\xfa\xd6\x19\xe6\x11\xcb\xe5\x22\x86\x7a\xc4\xa6\xe1\x0f\x18\xca\x40\x77\x03\x4e\x60\xff\x69\x05\x43\xf6\xfe\xd3\x62\x35\x61\x17\x7e\xc1\x6e\x32\x56\x11\x57\x36\xd9\x8a\x5b\x2f\x5e\x92\xe4\x95\xa2\xa5\xde\xe2\xb6\x55\xde\x30\xc7\xcc\x0b\x3c\xeb\x0e\xdf\xc6\xff\x2b\xb3\x31\xfd\xe8\x16\x76\xe3\xb7\xfd\x32\x3b\x8e\xae\xf4\x6d\x8a\xb7\x01\xc4\x8b\x35\x50\x9c\xf2\x46\xe4\xd1\x12\x1e\x86\x45\x1f\x9c\x19\xaf\xb8\xf7\x69\x22\x21\x6e\x9c\xa4\x4f\xe1\xeb\xc1\xbd\x25\x67\x66\xcd\xff\x61\xf6\x3d\x8d\xa5\xa4\x45\xca\xce\x10\xeb\xf3\xef\x2d\x97\xcb\x11\x61\x35\x95\x93\x45\x7a\x86\xa1\xbd\xc7\x54\x83\xdc\x20\xf5\x03\x8d\x51\xd4\x72\x22\x35\xa8\xd9\x18\x6f\x87\xeb\x7e\x29\x71\xc6\x48\x5d\x4a\xff\x61\x04\x82\xf4\x67\x00\x00\x00\xff\xff\xdd\xb3\x82\x11\x90\x08\x00\x00")

func templatesFunctionTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesFunctionTmpl,
		"templates/function.tmpl",
	)
}

func templatesFunctionTmpl() (*asset, error) {
	bytes, err := templatesFunctionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/function.tmpl", size: 2192, mode: os.FileMode(436), modTime: time.Unix(1466271087, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\x48\x4d\x4c\x49\x2d\x52\xaa\xad\xe5\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xa8\xae\xd6\x0b\x80\x30\x81\x82\x5c\x99\xb9\x05\xf9\x45\x25\x0a\x1a\x5c\xd5\xd5\x45\x89\x79\x40\x69\x3d\x4f\xb0\x48\x71\x6d\x2d\x50\xa1\x5f\x62\x2e\x50\x15\x44\x4b\x49\x06\x50\x7d\x75\x75\x6a\x5e\x0a\x90\xd6\x84\xb3\x00\x01\x00\x00\xff\xff\x81\x22\x53\x6f\x6b\x00\x00\x00")

func templatesHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHeaderTmpl,
		"templates/header.tmpl",
	)
}

func templatesHeaderTmpl() (*asset, error) {
	bytes, err := templatesHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/header.tmpl", size: 107, mode: os.FileMode(436), modTime: time.Unix(1455586613, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInlineTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\xcc\xcb\x01\xd2\x4a\xb5\xb5\x0a\xd5\xd5\x25\xa9\xb9\x05\x39\x89\x25\x40\xd1\xe4\xc4\x9c\x1c\x25\x05\x3d\xb0\x68\x6a\x5e\x4a\x6d\x2d\x20\x00\x00\xff\xff\xaa\xeb\x41\xff\x31\x00\x00\x00")

func templatesInlineTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesInlineTmpl,
		"templates/inline.tmpl",
	)
}

func templatesInlineTmpl() (*asset, error) {
	bytes, err := templatesInlineTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/inline.tmpl", size: 49, mode: os.FileMode(420), modTime: time.Unix(1458418263, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInputsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\xcb\x41\x0a\x02\x51\x08\xc6\xf1\xab\xc8\x30\xcb\x78\x07\x08\x3a\x40\xbb\xae\xf0\x60\x3e\x43\x28\x09\xc7\x56\x1f\x73\xf7\x46\x57\xad\x94\x9f\x7f\xc9\x0d\x6a\x0e\x59\xcc\x3f\xdf\xdc\x97\xe3\x20\x57\x95\xeb\x4d\x46\xad\xa6\xb2\xea\x78\x84\x79\xde\x3b\x28\x8c\xe9\x4f\xb4\xcf\x98\x6f\x24\xe2\xe4\xcc\x41\x36\xd4\xe7\x45\x48\xf8\x56\x35\x5e\x3b\xfa\xec\x67\xfb\xef\x35\x7e\x01\x00\x00\xff\xff\x43\x89\x5c\xae\x80\x00\x00\x00")

func templatesInputsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesInputsTmpl,
		"templates/inputs.tmpl",
	)
}

func templatesInputsTmpl() (*asset, error) {
	bytes, err := templatesInputsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/inputs.tmpl", size: 128, mode: os.FileMode(436), modTime: time.Unix(1455586613, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMessageTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8e\xd1\xca\x82\x40\x10\x46\x5f\x65\x10\x85\xff\x07\x99\x07\x08\x7a\x80\x6e\x22\x22\xba\x5f\xf2\xd3\x06\x74\xb3\xdd\xd5\x88\x61\xde\x3d\x15\x13\xba\xfa\xe0\x70\x38\x33\xaa\x15\x6a\xf1\xa0\xac\x43\x8c\xae\x41\x66\xa6\x2a\x35\xf9\x47\x22\x3e\x05\xf1\xe9\xe0\xfb\x21\x45\xb3\xe2\xc9\xa4\x0a\x5f\xcd\xc6\x4b\xd2\x9d\xf8\x8c\x1b\x64\x44\x98\x09\x5f\xde\x3d\xf8\xea\xda\x01\x66\xbc\x89\x7c\x74\xdd\x04\xfe\x96\xe8\x6f\x50\x35\x38\xdf\x80\x72\x29\x29\x47\x4b\xbb\xfd\x24\xb8\x30\xf9\x09\x21\xae\x7f\xe4\x62\x56\x7e\xef\x16\xe3\xd6\x5d\xe6\x7f\xdd\x4f\x00\x00\x00\xff\xff\xeb\x6d\x22\x24\xc6\x00\x00\x00")

func templatesMessageTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesMessageTmpl,
		"templates/message.tmpl",
	)
}

func templatesMessageTmpl() (*asset, error) {
	bytes, err := templatesMessageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/message.tmpl", size: 198, mode: os.FileMode(436), modTime: time.Unix(1455586613, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResultsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\xcd\x41\x0a\xc2\x40\x0c\x05\xd0\xab\x7c\x4a\x97\xa5\x07\x10\x5c\x8a\x7b\x6f\x20\x34\x95\x81\x92\x81\x3f\xd3\x55\xc8\xdd\x4d\x6a\x51\x70\x35\x93\xfc\x97\xc4\x6c\x91\xb5\xa8\x60\xa0\xb4\x7d\xeb\x6d\x70\x87\x19\x9f\xfa\x12\x8c\x65\xc2\x28\x1b\x2e\x57\xcc\x8f\x4f\xec\x6e\x56\xd6\x48\xdc\xa7\x70\xa2\x4b\x76\xee\xb5\x63\xce\xcf\x59\x87\x88\x81\xbe\x53\xdb\x8d\xac\x4c\x2c\xe4\x99\xe3\x00\x95\xdf\xa5\xff\x38\x0f\xfe\xec\xf1\xbe\x03\x00\x00\xff\xff\xb0\x4f\xcf\x61\xa8\x00\x00\x00")

func templatesResultsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesResultsTmpl,
		"templates/results.tmpl",
	)
}

func templatesResultsTmpl() (*asset, error) {
	bytes, err := templatesResultsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/results.tmpl", size: 168, mode: os.FileMode(436), modTime: time.Unix(1455586613, 0)}
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
	"templates/call.tmpl": templatesCallTmpl,
	"templates/function.tmpl": templatesFunctionTmpl,
	"templates/header.tmpl": templatesHeaderTmpl,
	"templates/inline.tmpl": templatesInlineTmpl,
	"templates/inputs.tmpl": templatesInputsTmpl,
	"templates/message.tmpl": templatesMessageTmpl,
	"templates/results.tmpl": templatesResultsTmpl,
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
		"call.tmpl": &bintree{templatesCallTmpl, map[string]*bintree{}},
		"function.tmpl": &bintree{templatesFunctionTmpl, map[string]*bintree{}},
		"header.tmpl": &bintree{templatesHeaderTmpl, map[string]*bintree{}},
		"inline.tmpl": &bintree{templatesInlineTmpl, map[string]*bintree{}},
		"inputs.tmpl": &bintree{templatesInputsTmpl, map[string]*bintree{}},
		"message.tmpl": &bintree{templatesMessageTmpl, map[string]*bintree{}},
		"results.tmpl": &bintree{templatesResultsTmpl, map[string]*bintree{}},
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

