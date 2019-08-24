package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func SaveStructPrettyJson(data interface{}, path string) error {
	arrByte, err :=json.MarshalIndent(data, "", "\t")

	if err != nil {
		return errors.New(fmt.Sprintf("Error convert  %v", err))
	}

	err = ioutil.WriteFile(path, arrByte, 0777)

	if err != nil {
		return errors.New(fmt.Sprintf("Error convert  %v", err))
	}

	return nil;
}

func LoadStructJson(data interface{}, path string) error {
	arrByte, err := ioutil.ReadFile(path)

	if err != nil {
		return errors.New(fmt.Sprintf("Error read %v: %v", path, err))
	}

	err = json.Unmarshal(arrByte, data)

	if err != nil {
		return errors.New(fmt.Sprintf("Error convert struct: %v", err))
	}

	return nil;
}

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
