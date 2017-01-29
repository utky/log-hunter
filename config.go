package main

type Config struct {
	Config []Entry `json: config`
}

type Entry struct {
	Name     string    `json: name`
	Hosts    []Host    `json: hosts`
	Commands []Command `json: commands`
}

type Host struct {
	Hostname string `json: hostname`
	Port     int    `json: port,omitempty`
	Username string `json: username,omitempty`
	Password string `json: password,omitempty`
	KeyFile  string `json: keyfile,omitempty`
}

type Command struct {
	Action  string `json: action`
	Content string `json: content`
}
