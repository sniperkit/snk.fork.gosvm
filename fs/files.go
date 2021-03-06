package fs // import "github.com/sah4ez/gosvm/fs"

import (
	"bufio"
	"os"
)

var gopath = os.Getenv("GOPATH")

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func PathToGlide(basePath, projectName string) (string, bool) {
	path := gopath + "/src/" + basePath + "/" + projectName + "/glide.yaml"
	return path, Exists(path)
}

func PathToDep(basePath, projectName string) (string, bool) {
	path := gopath + "/src/" + basePath + "/" + projectName + "/Gopkg.toml"
	return path, Exists(path)
}

func PathToGoMod(basePath, projectName string) (string, bool) {
	path := gopath + "/src/" + basePath + "/" + projectName + "/go.mod"
	return path, Exists(path)
}

func ReadFile(path string) (*bufio.Reader, error, func()) {
	var reader *bufio.Reader
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0)
	closer := func() {
		file.Close()
	}
	if err != nil {
		return reader, err, closer
	}

	reader = bufio.NewReader(file)
	return reader, nil, closer
}
