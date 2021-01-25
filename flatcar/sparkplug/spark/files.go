package spark

import (
	"encoding/base64"
	"fmt"
	"github.com/kinvolk/ignition/config/v2_2/types"
	"io/ioutil"
)

func fileFromSource(source string, fileFunc func(string) (types.File, error)) ConfigOpt {
	return ConfigOpt(func(config *types.Config) error {
		file, err := fileFunc(source)
		if err != nil {
			return err
		}
		files := config.Storage.Files
		config.Storage.Files = append(files, file)
		return nil
	})
}

func scriptToSource(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	// TODO: look at using https://godoc.org/github.com/vincent-petithory/dataurl
	return fmt.Sprintf(
		"data:;base64,%s",
		base64.URLEncoding.EncodeToString(data),
	), nil
}
