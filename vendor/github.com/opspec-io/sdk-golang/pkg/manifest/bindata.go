// Code generated by go-bindata.
// sources:
// github.com/opspec-io/spec/spec/op.yml.schema.json
// DO NOT EDIT!

package manifest

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

var _githubComOpspecIoSpecSpecOpYmlSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x79\x6f\xdc\xb6\xb6\xff\xdf\x9f\x82\x70\x8d\xd7\x99\x17\x8f\xc6\xd9\xfc\x5a\x17\x85\x61\x64\x79\x2f\x0f\xcd\x82\xa6\xc9\x05\x6a\x4f\x0a\x8e\x74\x66\x86\xb5\x24\xaa\x24\xe5\xa5\xbd\xf9\xee\x17\x24\x35\x5a\x49\x6d\x96\x62\x67\x31\x70\x6f\x63\x91\xe7\x88\xe7\xc7\x1f\x0f\x0f\x0f\x29\xfa\x9f\x1d\x84\x76\xf7\xb8\xbb\x81\x00\xef\x1e\xa1\xdd\x8d\x10\xd1\xd1\x7c\xfe\x27\xa7\xe1\x4c\x3f\x75\x28\x5b\xcf\x3d\x86\x57\x62\x76\xf0\x68\xae\x9f\x7d\xb7\xbb\xaf\xe4\x88\xb7\x95\xe1\x47\xf3\x39\x8d\x78\x04\xae\x43\xe8\xfc\xc0\xb9\xef\x1c\xce\x69\xe4\x5c\x07\xbe\x93\xa8\x91\x2a\xb5\x98\x20\xc2\x07\x29\xa8\x2b\xe8\x87\x1e\x70\x97\x91\x48\x10\x1a\xca\xa2\xa7\xb0\x22\x21\x70\x84\x51\x74\xbe\xd6\x35\x22\x46\x23\x60\x82\x00\xdf\x3d\x42\xb2\xdd\x08\xed\x86\x38\x80\xf4\xb7\xaa\x96\x57\x38\x00\x44\x57\x48\x6c\x20\xd5\xa3\xea\x89\xeb\x48\xb5\x80\x0b\x46\xc2\xf5\xae\x7a\xfc\x51\x97\x96\x74\xd8\x54\x3f\xcd\x7e\x35\xbe\x61\x8f\xc1\x4a\xd6\xfb\x6e\xee\x49\x53\x88\xac\xc8\xe7\x01\x66\xe7\x1e\xbd\x0c\x8b\x6f\x24\x61\x14\x0b\x9e\x7f\x99\x59\x3a\xc2\x0c\x07\x45\x51\x1a\x8b\xde\xb2\x2c\x0e\x9b\xe5\xb8\x5b\x82\xe7\x02\x18\xaf\x87\xe6\xbd\xae\xb1\x85\x85\x46\x4d\xa8\x70\x08\xde\x03\x2b\x19\xa6\xc8\xd4\xee\x2d\xa7\xba\xf2\x62\x52\x61\xe2\x14\xc5\x1c\x3c\xb4\xbc\x6e\xdb\x41\x85\xa6\xec\x24\xcd\xd9\x65\xf0\x57\x4c\x18\x48\xb6\x9f\xe6\x88\xb7\x83\xd0\x42\x95\x63\xcf\x53\xf2\xd8\x7f\x93\x27\xe9\x0a\xfb\x1c\x12\x76\xa7\xaf\xc8\xc8\x8b\x19\xc3\xd7\x6f\x54\xbf\xe4\xcc\x4c\x47\x47\xae\x78\xdf\x82\xc1\x89\xac\x82\x54\xd7\x82\x00\x26\xb1\xc0\x61\x01\xf0\x2d\xd1\xe9\xf2\x4f\x70\x45\xf6\xdc\x30\x98\xb2\x36\x15\x1e\xd9\x2b\xd7\x8c\x96\xb4\xb8\xcd\x30\xd8\xfe\x7c\xdc\x2f\xab\x5e\xe1\xd8\x17\x26\xb5\x5b\xb3\x74\x73\x6b\xb5\x10\xfe\x16\x5c\x06\x46\x35\x25\x38\x5f\x68\xc2\x2a\xa5\x88\x70\xc4\xb5\xe0\xbe\xed\xed\x4b\x4a\x7d\xc0\x0d\x56\xb8\x34\xe4\x82\x61\x12\x8a\x2a\x7a\x56\x80\x54\x13\x9e\xe4\x24\x8b\xaf\xd8\xb1\xbc\xae\x96\x88\x3b\x65\xf1\x54\xb4\x91\xbf\xaa\x52\x65\x10\xa0\x8c\x30\xc9\xef\x8b\xc2\x10\xae\x18\x61\x65\x79\xbe\x52\x6f\xea\x06\xf8\xea\x85\x80\xa0\x8c\x72\xb9\x93\xff\xff\xed\xeb\x57\xe8\xad\x9a\x93\xd0\xe9\x56\x06\x9d\xc3\xf5\x25\x65\x5e\xe6\x42\x04\xa5\x3e\x77\x08\x88\x95\x9a\x00\x37\x22\xf0\x93\x59\xf0\x92\x91\xf5\x46\xcc\x72\x53\xe4\xec\x02\xfb\xc4\xc3\xf2\x0d\xb3\x83\x83\xef\x38\xb8\xea\x9f\x8f\x9d\xfb\x07\xd3\x02\x7b\x52\x9b\x48\x28\x60\x0d\xac\x58\x18\x90\x90\x04\xb1\x74\x06\x07\x3b\x86\xee\x95\xe5\xdd\x0d\x4c\x64\xc6\x32\xf0\xfe\x90\x06\x92\xae\xd6\x91\x81\x4d\xbb\x9f\x9a\x76\xe8\xfc\x58\xb2\x8c\x86\xf0\x7a\x55\xe0\xbe\xfc\x69\x39\x9e\x25\x2c\xf6\xe1\xbc\x5f\xaf\xd2\x04\x4b\xff\xb7\x15\xdd\x47\xf1\xb7\x85\xb1\x5b\x32\xef\xd0\x99\x7e\x25\xd1\x91\xba\xaa\x32\xcc\x3a\xc2\x92\x37\x36\x0e\xc9\x5f\x31\x74\x36\x34\x27\x36\x96\x91\x0f\x2d\x43\xad\x32\x0b\x75\xf3\xef\xc5\xd8\x97\x30\x7b\x40\x92\x16\xda\xc2\x91\xa7\x84\x81\x2b\x28\x1b\x36\x24\xf1\x08\xbb\xfb\x01\x49\x75\xf9\x22\x6b\xa2\x0b\xec\xc7\xf0\x93\x44\x00\x2f\x39\xf5\x63\x01\x28\xc2\x62\x83\x56\x8c\x06\x88\x51\x2a\x24\x3e\xd1\xf9\x1a\x51\x86\x18\xf8\x58\x90\x8b\xa4\x86\x74\xa0\x2c\x62\x20\xc0\xd3\xb5\x65\x64\xe2\x11\x86\x48\x88\x2e\x37\xc4\xdd\x24\xb1\xb5\x8c\x53\x64\x20\x6f\x0d\x52\xf2\x4b\x1c\x9b\x61\xdd\x63\x24\x2f\xed\xea\xde\x71\xd2\x1d\x0a\x62\x24\xc5\x8c\x21\xcc\x8a\xf8\x60\x1f\x10\x59\xa9\x6d\x44\x3c\x27\x3e\x0c\x3a\x18\xe4\x2b\xbf\x8d\x86\xbb\x36\x1a\x64\xaf\x7c\x11\x03\x41\xd1\xcb\x38\x12\xa2\xd8\xf7\x9f\x30\xf0\x8a\x51\xbc\x85\xbd\x25\x94\xa4\x1c\x84\x82\x60\x9f\xeb\xf5\xb8\x17\xcb\x5e\x40\x38\x16\x1b\xf9\xdc\x55\xd3\x1d\xba\x24\x42\xf7\x23\xa7\x31\x73\x21\x19\x2d\x24\xc0\x6b\x90\x8c\x28\xac\xde\x6d\xe3\x23\xe6\xc0\x4a\x49\x21\x53\x8b\xe0\x2a\x62\xc0\x55\xfe\xc0\xa5\xc0\x5c\xb2\xf4\x01\x09\x8a\x34\x3d\x34\x51\xdb\x84\x14\x99\x1e\x73\x34\x11\x61\xce\x65\x2c\x70\x9b\xcd\xa9\xf0\xc3\xdc\xf5\x29\x72\xa6\xe6\x6f\x29\xd1\x3d\xa2\x08\xe3\x60\x09\xac\x69\x15\x58\xad\xd5\x3f\x83\xe1\xfb\x2a\x4e\x6f\x1f\xa3\x4a\x81\x91\xd6\x47\x0f\x1e\x58\x82\x36\xbd\x6c\x2e\x14\x99\xc3\x7c\x4b\x4f\x57\x01\xcb\x7b\x11\x73\x18\x1f\x5e\x77\x04\x46\x0a\x8c\x05\x8c\x2d\x9a\xbd\x05\x60\x20\x8c\x83\x2e\xb8\xc8\xfa\x63\xc1\x62\x4b\x18\xb4\x87\x65\x2b\xa1\x81\x68\xb6\x7e\x45\x59\x80\xcb\x73\x5d\xdb\xd5\x6e\x3a\x80\x4d\xeb\x7d\x13\x90\xbf\x6a\xdf\xc3\x95\x9f\xd7\x4d\x44\x4b\x50\x7e\xde\xa6\xa1\x34\x75\x57\xca\x93\xee\x3b\xad\xae\x8e\xb7\x2a\x4b\x25\x8b\x8e\xeb\xdf\x00\x5f\x25\x69\x8b\x2e\x79\x25\x29\x32\x16\x4b\x2c\x24\x29\x77\x79\x29\x79\xd4\xd9\x08\x2d\x32\x92\x11\x8f\xfa\x18\x11\xfb\x82\x44\x3e\x74\xf3\x63\x99\xd4\x58\x59\xb0\x1e\xa6\x84\xb4\x32\xe6\xea\x6c\x08\xa9\x18\x8b\x4c\x8f\x5b\x25\x4f\x6a\xfc\x6a\xde\xac\xad\xdf\x68\x6d\x98\x12\x18\xcb\x34\x1b\xc7\x3e\xd5\x24\xd3\x29\x34\x37\x84\x4d\xf6\xa5\x67\xbe\xdc\x16\x76\xbf\xd2\xee\x75\xc8\xe5\x67\xc2\xe8\x3b\xbf\x00\xb5\x4f\x82\x43\x2c\xf7\x92\x79\xeb\x56\x77\x88\x6a\x29\x78\xb7\x16\x95\xc5\x4e\x28\x2e\x2b\x35\xf1\x9a\x96\x07\xd5\x5a\x37\xd9\x25\x7a\x63\xa3\x6b\xd3\x94\x9e\x09\x8e\x35\x91\xd8\xa2\xe2\xbe\xfb\x45\x3d\x4d\xcd\x0b\x8e\x65\xaa\xcd\x37\xf7\x32\x35\x47\xbc\xd6\x56\x6e\x65\xc6\x32\xb0\x3c\xaf\xf6\x0f\xe5\xab\x49\x2b\x73\x28\x6f\xf5\xc3\xb5\x38\x44\xa3\x77\xf5\xa1\x05\x89\xd2\x98\x45\x35\x7e\xa6\xd5\xb4\x5c\xb3\xef\x65\x01\x0c\x0b\x01\xac\xe7\x28\xa9\x08\x8f\x05\xdf\xff\xdc\x55\xf8\x5a\xbc\xac\xe5\xee\x60\x3f\x10\x37\x38\xf4\x18\x5c\xf2\x16\x30\x1e\x3a\x8f\x9d\x32\x0d\xdb\xae\x71\x3b\xcc\xe5\x83\x6c\x07\x37\x2e\x48\x3d\x88\x20\xf4\x20\x74\x3b\xe2\x9d\x97\x1b\x8b\xac\x9f\xc5\xb6\x79\xdd\x6e\x40\xf7\xfd\xf1\x2f\x32\x7f\x54\x4e\xa7\xec\x86\xb1\xef\x57\x03\xdd\xc4\x07\x15\x1e\x2f\x9a\x1d\xc7\x57\x91\xa3\xad\x46\xad\xcd\xc0\x7c\x15\x39\xda\x1e\xc0\x7c\x1d\x79\x85\x1e\xc0\x7c\x6e\x79\xa4\x1a\x13\x6f\x90\x2e\xd1\x5a\xed\xe9\x92\x7c\xb9\x2d\x5d\xf2\x5a\xd5\x19\x34\x5d\x92\xd4\xfb\x6c\xd2\x25\x26\x6f\x7e\xf3\x74\x89\xd6\x7a\xbb\xe9\x92\xda\x91\x75\xb7\xd2\x25\xc5\x4e\x28\xa6\x4b\x38\x75\xcf\xa1\x86\xe7\xf9\xf2\x46\xd6\x96\x7a\xeb\xad\x92\xad\xe5\xbf\x8d\xe7\xfa\xb5\x77\x84\xe7\xdd\x09\xaa\x9b\xff\x45\x1c\xe0\x48\x7a\xc2\x4c\x1e\x15\xed\x36\xe5\xda\xaa\xb5\xbe\x6d\xc5\x27\x8f\x2d\x9f\x8c\x54\x00\xfb\x16\xe6\xf5\x05\xe6\x8b\x5c\x4a\xb5\xce\xdf\x99\xb7\xe2\xcb\x6e\x3a\x8e\x80\x71\x50\x67\xe5\x0a\x58\x68\xe9\x51\xd0\x28\x67\xa0\xba\x9e\x0e\xf0\xb0\x80\x99\x20\x01\x34\x9e\x0f\x28\xe6\x2b\xb6\x62\x48\xdb\x36\xac\x4d\xce\xc3\xf2\xc6\xad\xa9\xd3\x3a\x1c\x37\xc8\xac\x2c\x95\x2d\xea\xe6\xab\x1a\xd4\xa4\x2f\x67\x33\x75\x00\x6e\x26\x47\x58\x13\x78\x27\x48\x8b\x24\x67\xe6\x18\xac\x80\x41\xe8\x02\xc2\x1c\xa9\x81\xa9\x3f\x85\x3b\x5d\x13\xb1\x89\x97\x8e\x4b\x83\xb9\x16\x98\x7b\x44\x9a\xbb\x8c\xa5\xa6\x79\x2a\x97\xe1\xdd\x20\x21\x18\xc0\xb6\xe0\xbe\x73\xff\x61\xa6\x62\x58\x80\xcb\x80\x0c\x83\x33\x04\x98\x18\x52\x1b\xb5\x7e\x47\x8a\x8c\xc5\xca\x07\x83\x82\xa6\xad\x1b\x06\xa9\x0d\xe5\xa2\x74\x46\xb0\x05\x58\x5b\xa9\xb1\xf0\x7a\x38\x28\x5e\xa9\x8d\xc3\x40\x46\xa2\x8b\x47\xdd\xe0\x92\x12\x63\x41\xf5\x68\x50\xa8\x94\x6d\x83\xc1\x74\xd8\x19\xa6\xc3\xb1\x60\x7a\x3c\x34\x4c\x87\x03\xc1\x14\x33\xd2\x0d\xa5\x98\x91\xb1\x40\x3a\x1c\x14\x24\x69\xd9\x30\x18\x71\x08\x2e\x5a\x9c\x44\x3c\x41\x1c\x02\x1c\x0a\xe2\xa2\xe4\x93\xf6\xf2\x34\xa9\x15\x49\x8c\x34\x76\x47\xf3\x79\xf6\x68\x3e\xa8\xf5\x49\x9b\xeb\x01\xd8\x31\x95\x94\x0e\x2a\xfe\x02\xe1\x5a\x6c\x3a\x9e\x6b\xd0\x42\x23\xc5\xd1\xb6\xcd\xdf\x86\x7d\xfe\xfb\x66\x0b\x49\xd8\xc3\xc2\xad\xd0\x48\x16\xda\xf6\x67\x9b\x4e\x32\xec\x17\x0d\xd8\x66\xe7\xcc\x27\x1c\x3e\xb7\x3c\x6f\xcd\xe2\xef\xeb\xcb\xeb\xf7\x58\x09\x27\xa7\x0b\x7a\x1c\x48\x18\x09\x9c\x1f\x2c\xd8\x18\x9c\x5d\xb6\x90\xdd\x65\xb0\x86\xab\x41\x52\xfc\xfa\x3d\x35\xa9\xcf\x5c\x79\xe7\xd4\xa7\xfe\xce\xa6\x57\xea\x53\x9b\x7f\x37\x52\x9f\x2d\x52\xfc\xe3\x7c\x00\x97\x7c\xa8\x74\xab\x29\xfe\xda\x41\x76\xc7\xb2\xb4\x85\x4e\x28\x7d\x68\x57\x66\x78\x09\xf1\x37\x7d\xf6\xa7\x6a\xcf\x39\xed\x9e\xce\xfe\x70\xf0\xec\xef\x93\xd9\xef\x07\xb3\x1f\x17\xf7\x7a\x7e\x0b\x52\x73\x93\xc9\x9b\xec\xf2\x21\x4b\x97\xb7\xd4\x96\x7e\x82\x3e\x80\xae\xec\xeb\xdd\x01\x94\xe5\xcf\x63\x0f\xa0\x2e\xbf\x5f\x39\x80\xba\xfc\xb6\xd0\x10\xea\x72\xae\xb6\x4d\x88\xda\xdf\xe9\x97\x0f\x00\x99\x1c\x7f\xb9\x8e\xcd\xc9\x67\x03\xc7\x35\xd5\xae\xb2\xfc\x9f\xe6\x28\xc7\x7e\x4f\x4f\x0e\xd9\x16\x7a\xda\x7d\x5d\xd1\x42\x51\xdd\xf6\x7a\x27\x45\x75\xf1\x9b\xd1\x6f\x71\x77\x6d\x9e\x97\xdd\xb5\xb5\x4b\x4e\xdf\x0a\x2c\x17\x60\x2e\xf6\x7d\xb4\x66\x38\xda\x64\x41\x0b\x84\xce\x25\x39\x27\x11\x78\x44\xdf\x41\x27\x7f\x9b\x3f\xc1\xbe\xff\x87\xaa\x39\x35\xf8\xbf\x5e\x7d\xe8\xd2\x50\x60\x12\x02\x93\xba\x7b\xe3\x1e\xdd\x44\x5a\xfa\x7c\xdf\x07\xff\x26\x3a\x38\x30\x82\xcb\x1a\x8c\x3d\x55\x34\xd8\xd4\x67\xc5\x1a\xbd\xf7\xff\x52\x35\x5d\xc2\x23\x37\x28\x9f\xe2\x36\x31\xe7\x09\x0d\x02\x1c\x7a\x88\xc5\xa1\x5c\xab\x63\x94\xbe\xeb\x27\x44\x2f\x80\x31\xe2\x01\x47\x38\xbc\x46\x1c\x04\xc2\x42\x45\x29\x3a\x2d\xee\xc3\x05\x18\xd2\xbd\xf6\x58\x1f\xd9\xe3\x7d\x53\xd3\x3a\x7f\xcc\x5d\xdb\xad\xa6\x4f\xba\x8b\x9d\x9b\xfc\x56\x0e\x05\x09\x33\xc6\x4e\x35\x87\x87\x4d\xc6\x6c\xaf\x74\x21\xc0\x11\x09\x15\x8a\x59\xaf\x56\x84\x9b\xce\x53\x27\xd5\x3e\x4c\x4e\x75\xb8\xb1\x38\x9a\x1e\xcb\xe0\xe3\xec\x6c\x9e\x8b\x3f\xf6\x8c\x52\xd6\x40\x64\xfb\x63\x12\x31\x99\x34\xb9\x24\xbe\x8f\x96\x80\x96\x34\x0e\x3d\xd5\x33\x38\x48\xef\x99\x40\xd1\xf9\xba\x9a\xda\xa9\xc0\xa7\x4e\x43\x1a\x2b\x7d\x34\xcb\xb6\x6d\x9d\x8d\x3d\x1e\x61\x9a\x3a\xe8\xbf\xe6\x94\x21\xee\xd2\x48\xed\xf0\xa8\xf6\x83\x40\x71\x44\x43\x04\x57\xa4\xda\xa5\xe9\x9b\xba\x12\x2c\xb1\xc7\xf0\x74\x51\x79\x56\xae\x55\x41\xa1\x5d\x94\x6d\x10\xdd\x85\xf0\xe2\x3d\x1e\x84\xcc\xcf\xc2\x0b\xc2\x68\x18\x40\x28\xd0\x05\x66\x04\x2f\xfd\x41\x69\x7d\xfa\xe1\xe7\x5b\x60\x2f\x09\x73\x6c\xb8\x9c\x6b\x36\x87\x38\x30\xec\xbd\x55\x80\xfb\xe4\x34\x6e\x70\x82\x89\xb2\xcf\x94\xa9\x72\x45\x31\x04\x4f\x9f\x93\x61\x79\xf9\xcd\xdd\x5a\x5b\x67\xe3\xa9\xba\xd1\xe7\xeb\xf4\xb7\x2a\x4a\xea\xc3\xe2\x9a\xa8\x0e\xe9\x5c\x48\x39\xd3\x9b\x16\x55\x6e\x89\xd8\x9e\x5c\x10\x34\xbd\x03\xc8\x08\x74\x0f\x90\x0d\x8c\x31\xde\x6e\xd4\xe2\x35\x99\x58\x8f\x0e\x32\xa6\x86\x0a\x60\x95\x9e\x2e\x86\xeb\x63\xc3\x15\x49\xc8\xd0\x0f\xea\xfe\xec\x82\x13\x42\x2e\x0e\xe5\x60\x4e\x0f\x78\xa8\xad\x32\x75\x25\x17\x15\x1b\xbd\xb6\xd6\x35\x79\xe3\xfe\x58\x6d\x0b\x23\xca\xcc\xc9\xbf\xf2\xaa\x5e\xd6\x4b\x7c\x4b\x7a\x33\x58\xd6\x5c\x41\xd5\x83\x0d\xe5\x35\x39\x49\x2b\xa1\xdb\xf9\xd7\x53\xe5\x46\x27\x33\xfd\xdf\xe9\xf1\x44\xb8\xd1\xbf\x63\x2f\x9a\x1e\xb7\xa4\xfb\xff\x51\x2e\x90\x34\x78\xc2\xa7\xb2\xc5\x4b\xa2\x1c\xa5\x99\xf0\x0d\x9b\x8c\xa8\xb8\x79\x50\x69\x5c\x1f\xa6\xf6\xa6\x99\xce\x3b\xf5\x9a\x12\xdb\x62\x7f\x64\x4f\x5c\xa6\x95\x2a\x6b\xc7\x2d\x3b\x92\x83\xbf\xd8\xf3\xa4\xb7\x40\x01\x8e\x22\x50\x53\x14\xde\x16\x99\x8e\x5d\xa1\x26\x2e\x8f\x8d\xaa\xf0\x9e\xb1\xf2\x9a\x7a\x48\x50\x3f\x38\xf6\xa0\xc0\x8e\xa5\xf0\x80\x31\x14\x31\x58\x91\xab\x22\x94\x3a\xe6\x1b\x11\x4a\x0f\x22\x06\x2e\x16\xca\x99\x0a\x16\xc3\xa0\x60\xbf\x8e\xdb\x7c\x34\xf1\xa9\xc1\xa6\xb1\xf8\xe2\xc0\xbe\xa4\xec\xfc\x69\xe5\xba\x57\x13\x14\xff\xa2\xec\x5c\xda\xe9\xe5\xae\x9c\x15\x1b\x34\x29\xe6\x7e\x72\xa7\x39\x54\x08\xd1\x7c\x66\x63\xc7\x86\x45\x71\xcb\xc8\x3a\x77\x27\xd1\x53\xee\xd9\x62\x88\xbd\x26\xf3\x36\x52\xb6\x34\xd8\x29\xbd\xab\xc3\x97\x53\x91\x35\x0b\x98\x14\xf5\x4e\xff\xd1\xa8\x9c\xf7\xab\xfb\x72\xbd\x2e\x27\x18\x9d\x97\x77\x58\x9b\xd4\x35\xa9\x44\x37\x09\x48\x0b\x57\x51\x1a\x5b\x54\x37\x39\x67\x9b\xe3\x31\x23\xb3\x34\x90\xfa\x16\xa8\x1a\xde\x5e\xfd\xd3\x28\x69\x49\x79\x23\x5a\x9a\x8a\x7d\xf2\x37\x70\xf4\xe2\xd5\x9b\x77\xbf\xfd\xf1\xea\xe4\xe5\x33\x1d\x12\xbe\x3f\xf9\xe5\xdd\x33\xb9\xd8\x4c\x8e\xc8\x7f\x9f\x55\x38\xd2\x85\xdf\x3b\xe8\xc5\x6a\x5b\x8f\x23\xb9\xdc\xdc\x47\x44\xa0\x97\xef\xde\xfe\xa6\xee\xa7\xe3\x3c\x0e\xc0\x4b\x6a\xfc\x8c\xf6\x26\x99\x8a\x1a\xa7\x72\xd3\xe0\xa6\x76\x5f\x36\xad\xd6\x73\x9d\x3e\xfc\xda\xfa\x33\x5d\xf0\x56\xff\x86\x4e\x5a\x54\x43\xb2\x8c\x5e\xaf\xdf\xfd\x96\xf2\x2d\x47\x32\x4d\xaf\x5c\xa1\x26\x59\xa1\x76\x0d\xd5\x54\x05\xc9\xb4\x9c\xc0\x37\xaa\x95\x35\xda\xc2\x18\x74\x2b\xc4\xea\x11\x2f\xc8\xa9\xe4\xd3\x44\x0b\x34\xba\x41\x98\x50\xd8\xa1\x34\x05\x0b\x85\x0a\xbd\x43\x86\xad\x16\x5b\xe0\x70\xe3\x83\x78\xae\xed\x33\xac\x96\x20\xa6\x0d\xec\x0f\x65\x6e\xa3\xd6\xb8\x5f\x9e\x15\xf7\x86\x51\xeb\xb8\xb3\x20\x26\xcd\xeb\x0f\x61\x6e\x4a\x31\x41\x98\x2b\xb6\x9d\x3c\x38\x09\xf5\x9f\x79\xda\x4f\x2e\xf3\xdb\x4f\xbe\x52\xdf\x47\x94\x25\x6b\x28\x07\xe9\x63\x81\x3c\x75\xcd\xfa\x16\x7a\xea\xcb\xb5\x0f\xc2\x02\xb1\x38\x14\x24\x80\x7d\xc4\x20\xf2\xb1\x2b\xd7\x23\x7b\x93\x5f\x9f\x3d\x9f\x6e\xff\xe0\x98\x4e\x20\xd3\x15\xfa\xf5\xd9\x73\x47\xfe\x5f\xe6\xe5\x93\x8c\xf2\xde\x44\xed\xd4\x24\xf7\x9b\xa3\x15\x47\x7b\x93\xb9\x5c\xca\x4c\xe5\x12\xce\x41\xaf\x80\x0b\xa9\x37\x8e\x64\xe4\x49\xc3\x64\xfb\x3a\xd5\xc3\xe3\x28\xa2\x4c\x80\xb7\x8f\x88\x03\x8e\x94\x4e\x54\x4e\x2b\xf4\x39\xdd\xf6\x7c\x7a\x25\xdf\x7e\x46\xa9\xd4\x95\x96\x4e\x75\xe8\xbf\x87\x66\x26\xaa\x2a\xb2\x9e\xed\xb8\x78\xe0\x1c\x38\x07\x95\x53\xf6\xa6\xb3\xf4\x3c\x02\x77\xae\xeb\x3b\x1b\x11\xf8\xd5\xb6\x97\x23\xea\x7c\x8a\xeb\xc3\x24\x49\x6e\x9d\x9d\x39\x86\x7f\x4e\x8e\x8f\x26\x67\x67\x2a\x01\x76\x32\xfb\x1d\xcf\xfe\x9e\x2d\xee\x4d\x8e\x8f\xce\xce\x9c\xc2\xa3\xe9\x7f\x4f\xa7\xc7\xea\xf9\xbd\xdc\xf3\xb3\xb3\xd9\xd9\x99\xb3\xb8\x37\x3d\xde\x2b\xfe\x7d\xba\xf4\x40\xa6\x09\x9a\xb4\xd0\x06\xce\xcb\xa4\x82\x8c\x17\x4e\x2f\x0e\x9c\x07\x3f\xa0\x27\x34\x08\x68\x28\x0b\x10\xbf\x0e\x05\xbe\xca\x80\x8a\xc0\x75\x5c\x55\x2c\x15\x2b\xc4\xa4\xc8\x7c\x8a\x48\xe8\xfa\xb1\x27\x09\xf2\xbf\xcf\x5f\x22\x81\x97\x3e\x20\xb8\x12\x10\x16\xc9\x6f\xfc\x73\x87\x3b\xf2\x7f\x1f\x77\x76\xfe\x13\x00\x00\xff\xff\xb6\xf5\x06\x8d\x03\x72\x00\x00")

func githubComOpspecIoSpecSpecOpYmlSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_githubComOpspecIoSpecSpecOpYmlSchemaJson,
		"github.com/opspec-io/spec/spec/op.yml.schema.json",
	)
}

func githubComOpspecIoSpecSpecOpYmlSchemaJson() (*asset, error) {
	bytes, err := githubComOpspecIoSpecSpecOpYmlSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "github.com/opspec-io/spec/spec/op.yml.schema.json", size: 29187, mode: os.FileMode(420), modTime: time.Unix(1520201883, 0)}
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
	"github.com/opspec-io/spec/spec/op.yml.schema.json": githubComOpspecIoSpecSpecOpYmlSchemaJson,
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
	"github.com": &bintree{nil, map[string]*bintree{
		"opspec-io": &bintree{nil, map[string]*bintree{
			"spec": &bintree{nil, map[string]*bintree{
				"spec": &bintree{nil, map[string]*bintree{
					"op.yml.schema.json": &bintree{githubComOpspecIoSpecSpecOpYmlSchemaJson, map[string]*bintree{}},
				}},
			}},
		}},
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
