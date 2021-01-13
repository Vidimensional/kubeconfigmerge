package test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
)

const (
	commandName      = "./bin/kubeconfigmerge"
	testdataPath     = "./test/data/"
	localPath        = testdataPath + "local/"
	actualConfigFile = "actual.yaml"
	newConfigFile    = "new.yaml"
)

func TestCliArgs(t *testing.T) {
	tests := []struct {
		name           string
		actualConfig   string
		mergingConfig  string
		expectedConfig string
	}{
		{
			name:           "Merge blank config into a blank local config",
			actualConfig:   "blank.yaml",
			mergingConfig:  "blank.yaml",
			expectedConfig: "empty.yaml",
		},
		{
			name:           "Merge empty config into a blank local config",
			actualConfig:   "blank.yaml",
			mergingConfig:  "empty.yaml",
			expectedConfig: "empty.yaml",
		},
		{
			name:           "Merge dev config into an blank local config",
			actualConfig:   "blank.yaml",
			mergingConfig:  "kubeconfig-dev.yaml",
			expectedConfig: "kubeconfig-dev.yaml",
		},
		{
			name:           "Merge dev config into an empty local config",
			actualConfig:   "empty.yaml",
			mergingConfig:  "kubeconfig-dev.yaml",
			expectedConfig: "kubeconfig-dev.yaml",
		},
		{
			name:           "Merge non-empty config into an blank local config",
			actualConfig:   "empty.yaml",
			mergingConfig:  "kubeconfig-dev.yaml",
			expectedConfig: "kubeconfig-dev.yaml",
		},
		{
			name:           "Merge non-empty config into non-empty local config",
			actualConfig:   "kubeconfig-dev.yaml",
			mergingConfig:  "kubeconfig-test.yaml",
			expectedConfig: "kubeconfig-dev-test.yaml",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := cleanUp(); err != nil {
				t.Fatal(err)
			}

			if err := copyActualConfig(test.actualConfig); err != nil {
				t.Fatal(err)
			}

			if err := copyNewConfig(test.mergingConfig); err != nil {
				t.Fatal(err)
			}

			actualConfigPath := filepath.Join(localPath, actualConfigFile)
			newConfigPath := filepath.Join(localPath, newConfigFile)
			err := exec.Command(commandName, "-k", actualConfigPath, newConfigPath).Run()
			if err != nil {
				t.Fatal(err)
			}

			expectedConfigPath := filepath.Join(testdataPath + test.expectedConfig)
			expectedConfig, err := unmarshallYamlFile(expectedConfigPath)
			if err != nil {
				t.Fatal(err)
			}

			actualConfig, err := unmarshallYamlFile(actualConfigPath)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(expectedConfig, actualConfig) {
				t.Errorf("Expected: %s\nActual: %s\n", expectedConfig, actualConfig)
			}

		})
	}
}

func TestMain(m *testing.M) {

	if err := os.Chdir(".."); err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}

	if _, err := os.Stat("./bin"); os.IsExist(err) {
		os.RemoveAll("./bin")
	}

	if err := exec.Command("make", "build").Run(); err != nil {
		fmt.Printf("could not make binary for: %v", err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}
