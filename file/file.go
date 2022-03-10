package file

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(name string) []byte {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatalf("read %s file err is : %s", name, err)
		os.Exit(1)
	}
	return content
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
