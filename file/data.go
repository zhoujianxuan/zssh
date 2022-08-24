package file

type Data struct {
	User string    `json:"user"`
	List []SSHInfo `json:"list"`
}

type SSHInfo struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
	Host string `json:"host"`
	Port string `json:"port"`
}
