package test

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func unmarshallYamlFile(file string) (interface{}, error) {
	src, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	yamlStr, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, err
	}

	var retYaml interface{}
	if err := yaml.Unmarshal([]byte(yamlStr), &retYaml); err != nil {
		return nil, err
	}

	return retYaml, nil
}

func copyActualConfig(file string) error {
	return copyFile(file, actualConfigFile)
}

func copyNewConfig(file string) error {
	return copyFile(file, newConfigFile)

}

func copyFile(srcFile string, dstFile string) error {
	srcPath := filepath.Join(testdataPath, srcFile)
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	if err := ensureDirExists(localPath); err != nil {
		return err
	}

	dstPath := filepath.Join(localPath, dstFile)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func ensureDirExists(path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return err
	}

	if err := os.Mkdir(localPath, os.ModeDir|0755); err != nil {
		return err
	}

	return nil
}

func cleanUp() error {
	return os.RemoveAll(localPath)
}
