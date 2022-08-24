package handle

import (
	"encoding/json"
	"zssh/file"

	"github.com/urfave/cli/v2"
)

func InitAction(c *cli.Context) error {
	user := "zhoujx"
	if c.NArg() > 0 {
		user = c.Args().Get(0)
	}

	data := &file.Data{
		User: user,
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = file.Write(string(bytes))
	return err
}
