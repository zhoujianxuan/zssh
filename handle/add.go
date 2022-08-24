package handle

import (
	"encoding/json"
	"fmt"
	"zssh/file"

	"github.com/urfave/cli/v2"
)

func Add(c *cli.Context) error {
	if c.NArg() < 3 {
		fmt.Println("Periodic incomplete")
		return nil
	}

	// 先读
	content, err := file.Read()
	data := &file.Data{User: "zhoujx"}
	if len(content) > 0 {
		err = json.Unmarshal([]byte(content), data)
		if err != nil {
			return err
		}
	}

	data.List = append(data.List, file.SSHInfo{
		Name: c.Args().Get(0),
		Pwd:  c.Args().Get(1),
		Host: c.Args().Get(2),
		Port: "22",
	})
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = file.Write(string(bytes))
	if err != nil {
		return err
	}

	fmt.Println("success!")
	fmt.Println(data.List)
	return nil
}
