package file

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func GetFileName() string {
	current, err := user.Current()
	if err != nil {
		return ""
	}
	return current.HomeDir + "/.zssh.json"
}

//Read 读取数据
func Read() (string, error) {
	file, err := os.OpenFile(GetFileName(), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return "", err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

//Write 写入数据
func Write(content string) error {
	file, err := os.OpenFile(GetFileName(), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	_, err = io.WriteString(file, content)
	return err
}
