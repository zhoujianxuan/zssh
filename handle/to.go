package handle

import (
	"encoding/json"
	"fmt"
	"strconv"
	"zssh/file"

	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
)

func To(c *cli.Context) error {
	if c.NArg() < 1 {
		fmt.Println("Periodic incomplete")
		return nil
	}

	indexStr := c.Args().Get(0)
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return err
	}

	content, err := file.Read()
	if err != nil {
		return err
	}
	if len(content) == 0 {
		fmt.Println("config is nil, please init.")
		return nil
	}

	data := &file.Data{}
	err = json.Unmarshal([]byte(content), data)
	if err != nil {
		return err
	}

	if index >= len(data.List) {
		return nil
	}

	info := data.List[index]
	cmd := fmt.Sprintf("ssh %s@%s", info.Name, info.Host)

	err = clipboard.WriteAll(info.Pwd)
	if err != nil {
		return err
	}
	err = clipboard.WriteAll(cmd)
	if err != nil {
		return err
	}
	return nil
}
