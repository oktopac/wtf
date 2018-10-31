package checker

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type checker interface {
	Check(source io.Reader) (bool, error)
	Name() string
}

var checkers = []checker{
	JSON{},
}

// CheckFile runs all the file checkers against a file to see if it is of a type
func CheckFile(file string) error {
	logger := logrus.WithFields(logrus.Fields{
		"file": file,
	})
	for _, c := range checkers {
		// TODO: Don't reopen the file every time
		f, err := os.Open(file)
		if err != nil {
			return err
		}

		isType, err := c.Check(f)
		if err != nil {
			return err
		}

		if isType {
			logger.WithField("type", c.Name).Info("Found possible file type")
		}
	}
	return nil
}
