package lib

import (
	"os"
)

func FileExists(name string) bool {

	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func MustDir(path string, perm os.FileMode) error {
	if FileExists(path) == false {
		if err := os.Mkdir(path, perm); err != nil {
			return err
		}
	}

	return nil
}
