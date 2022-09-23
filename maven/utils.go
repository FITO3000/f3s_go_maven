package maven

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func fileExists(file string) (bool, error) {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func createFile(file string) error {
	exists, err := fileExists(file)
	if err != nil {
		return err
	} else if exists {
		return fmt.Errorf("file: %s already exists", file)
	} else {
		if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
			return err
		} else {
			if _, err := os.Create(file); err != nil {
				return err
			} else {
				return nil
			}
		}
	}
}
