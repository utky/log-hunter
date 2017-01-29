package main

import (
	"errors"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"path"
)

type Session struct {
	Name    string
	Host    Host
	Config  *ssh.ClientConfig
	Actions []Action
}

func makeAddress(h Host) string {
	port := 22
	if h.Port != 0 {
		port = h.Port
	}
	return fmt.Sprint(h.Hostname, ":", port)
}

func setupClient(s Session, logger *log.Logger) (*ssh.Client, error) {
	address := makeAddress(s.Host)
	return ssh.Dial("tcp", address, s.Config)
}

func runSession(s Session, logger *log.Logger) error {
	ssh, sshErr := setupClient(s, logger)
	if sshErr != nil {
		return sshErr
	}
	sftp, sftpErr := sftp.NewClient(ssh)
	if sftpErr != nil {
		return sftpErr
	}
	defer sftp.Close()

	for _, a := range s.Actions {
		a.RunAction(sftp, logger)
	}
	return nil
}

func buildAction(h Host, c Command, parent string) (Action, error) {
	if c.Action == "scp" {
		remotePath := c.Content
		filename := path.Base(remotePath)
		return ScpAction{
			remotePath,
			path.Join(parent, h.Hostname, filename)}, nil
	}
	return ScpAction{}, errors.New("action must be scp but input: " + c.Action)
}

func actions(h Host, cs []Command, parent string) ([]Action, error) {
	as := make([]Action, 0)

	for _, c := range cs {
		tmpa, buildErr := buildAction(h, c, parent)
		if buildErr != nil {
			return as, buildErr
		}
		as = append(as, tmpa)
	}

	return as, nil
}

func buildClientConfig(h Host) *ssh.ClientConfig {
	config := &ssh.ClientConfig{
		User: h.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(h.Password),
		},
	}
	return config
}

func injectCommands(name string, hs []Host, cs []Command, outPath string) ([]Session, error) {
	ss := make([]Session, 0)
	for _, h := range hs {
		as, err := actions(h, cs, outPath)
		if err != nil {
			return ss, err
		}
		s := Session{name, h, buildClientConfig(h), as}
		ss = append(ss, s)
	}
	return ss, nil
}

func BuildCommand(config Config, outPath string) ([]Session, error) {
	ss := make([]Session, 0)
	for _, e := range config.Config {
		commands := e.Commands
		tss, injErr := injectCommands(e.Name, e.Hosts, commands, outPath)
		if injErr != nil {
			return ss, injErr
		}
		// save result
		ss = append(ss, tss...)
	}
	return ss, nil
}

func RunSessions(sessions []Session, logger *log.Logger) {
	for _, s := range sessions {
		err := runSession(s, logger)
		if err != nil {
			logger.Fatal(err)
		}
	}
}
