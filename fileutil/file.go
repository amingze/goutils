package fileutil

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
)

func PathExist(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func Visit(dir string, suffix string, visitor func(filename string) error) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, suffix) {
			return visitor(path)
		}

		return nil
	})
}

// MD5Hex returns the file md5 hash hex
func MD5Hex(filepath string) string {
	f, err := os.Open(filepath)
	if err != nil {
		return ""
	}
	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return ""
	}

	return hex.EncodeToString(md5hash.Sum(nil)[:])
}

// DetectContentType returns the file content-type
func DetectContentType(filepath string) string {
	mimeType := mime.TypeByExtension(path.Ext(filepath))
	if mimeType != "" {
		return mimeType
	}

	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return ""
	}

	return http.DetectContentType(fileData)
}

func UserHomeAbs(filename string) string {
	if strings.HasPrefix(filename, "~/") {
		u, _ := user.Current()
		return filepath.Join(u.HomeDir, filename[1:])
	}

	return filename
}
func GetSize(f multipart.File) (int, error) {
	content, err := io.ReadAll(f)
	return len(content), err
}

// GetExt get the file ext
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckNotExist check if the file exists
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckAndCreateMkFile  check if exists
func CheckAndCreateMkFile(path string) {
	MkFileAll(path)
}

// CheckAndCreateMkDir  check if exists
func CheckAndCreateMkDir(path string) {
	if CheckNotExist(path) {
		MkDir(path)
	}
}

// CheckExist check if the file exists
func CheckExist(src string) bool {
	return !CheckNotExist(src)
}

// CheckPermission check if the file has permission
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir create a directory if it does not exist
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

// MkDir create a directory
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
func MkFile(src string) error {
	if CheckNotExist(src) {
		err := os.MkdirAll(src, os.ModePerm)
		return err
	}
	return nil
}

func MkFileAll(path string) error {
	MkFilesPathDir(path)
	_, err := os.Create(path)
	return err
}

func MkFilesPathDir(path string) error {
	f := func(c rune) bool {
		if c == '\\' || c == '/' {
			return true
		} else {
			return false
		}
	}
	index := strings.LastIndexFunc(path, f)
	dirs := path[:index]
	err := os.MkdirAll(dirs, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = os.Create(path)
	return err
}

// Open a file according to a specific mode
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// MustOpen maximize trying to open the file
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("fileutil.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("fileutil.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
