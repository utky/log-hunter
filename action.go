package main

import (
	"github.com/pkg/sftp"
	"io"
	"log"
	"os"
	"path"
)

type Action interface {
	RunAction(sftp *sftp.Client, logger *log.Logger) error
}

type ScpAction struct {
	RemotePath string
	LocalPath  string
}

func (a ScpAction) RunAction(sftp *sftp.Client, logger *log.Logger) error {

	logger.Printf("Starting copy from %s to %s\n", a.RemotePath, a.LocalPath)
	remote, sftpOpenErr := sftp.Open(a.RemotePath)
	defer remote.Close()

	if sftpOpenErr != nil {
		return sftpOpenErr
	}

	outputDir := path.Dir(a.LocalPath)
	if err := os.MkdirAll(outputDir, 0777); err != nil {
		return err
	}

	local, localOpenErr := os.Create(a.LocalPath)
	defer local.Close()

	if localOpenErr != nil {
		return localOpenErr
	}

	written, copyErr := io.Copy(local, remote)

	if copyErr != nil {
		return copyErr
	}

	logger.Printf("Finished copy (%d bytes)\n", written)

	return nil
}
