package client

import (
	"io"
	"os"
	"path"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/pkg/errors"
)

type GetOptions struct {
	SkipAuthRetry bool

	Quiet      bool
	RemoteFile string
	LocalFile  string
}

func (s Service) Get(opts GetOptions) error {
	client, err := s.GetClient(opts.SkipAuthRetry)
	if err != nil {
		return errors.Wrap(err, "Getting client")
	}

	localFilePath := opts.LocalFile

	if localFilePath == "" {
		localFilePath = path.Base(opts.RemoteFile)
	}

	var file io.ReadWriteSeeker

	if localFilePath == "-" {
		fileTemp, err := s.fs.TempFile("ssoca-download-stdout")
		if err != nil {
			return errors.Wrap(err, "Creating temp file for stdout")
		}

		defer os.RemoveAll(fileTemp.Name())

		file = fileTemp
	} else {
		file, err = s.fs.OpenFile(localFilePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
		if err != nil {
			return errors.Wrap(err, "Opening file")
		}
	}

	downloadStatus := pb.New(0).SetRefreshRate(250 * time.Millisecond).SetWidth(80)
	downloadStatus.Output = s.runtime.GetStderr()
	downloadStatus.ShowPercent = false

	err = client.Download(opts.RemoteFile, file, downloadStatus)
	if err != nil {
		return err
	}

	if localFilePath != "-" {
		return nil
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return errors.Wrap(err, "Seeking downloaded temp file")
	}

	_, err = io.Copy(s.runtime.GetStdout(), file)
	if err != nil {
		return errors.Wrap(err, "Writing download to STDOUT")
	}

	return nil
}
