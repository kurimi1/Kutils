package kfile

import (
	"io/ioutil"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
