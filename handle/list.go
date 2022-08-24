package handle

import (
	"encoding/json"
	"fmt"
	"zssh/file"

	"github.com/urfave/cli/v2"
)

func List(_ *cli.Context) error {
	content, err := file.Read()
	if err != nil {
		return err
	}

	if len(content) == 0 {
		fmt.Println("is nil")
		return nil
	}

	data := &file.Data{}
	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		return err
	}

	for _, info := range data.List {
		fmt.Println(fmt.Sprintf("%s %s", info.Name, info.Host))
	}
	return nil
}
