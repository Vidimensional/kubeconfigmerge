package test

import (
	"fmt"
	"os"
	"os/exec"
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
		newConfig      string
		expectedConfig string
	}{
		{
			"Merge blank config into a blank local config",
			"blank.yaml",
			"blank.yaml",
			"empty.yaml",
		},
		{
			"Merge empty config into a blank local config",
			"blank.yaml",
			"empty.yaml",
			"empty.yaml",
		},
		{
			"Merge dev config into an blank local config",
			"blank.yaml",
			"kubeconfig-dev.yaml",
			"kubeconfig-dev.yaml",
		},
		{
			"Merge dev config into an empty local config",
			"empty.yaml",
			"kubeconfig-dev.yaml",
			"kubeconfig-dev.yaml",
		},
		{
			"Merge non-empty config into an blank local config",
			"empty.yaml",
			"kubeconfig-dev.yaml",
			"kubeconfig-dev.yaml",
		},
		{
			"Merge non-empty config into non-empty local config",
			"kubeconfig-dev.yaml",
			"kubeconfig-test.yaml",
			"kubeconfig-dev-test.yaml",
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

			if err := copyNewConfig(test.newConfig); err != nil {
				t.Fatal(err)
			}

			//arg := []string{"-k", actualConfigFile, newConfigFile}
			err := exec.Command(commandName, "-k", localPath+actualConfigFile, localPath+newConfigFile).Run()
			if err != nil {
				t.Fatal(err)
			}

			expectedConfig, err := unmarshallYamlFile(testdataPath + test.expectedConfig)
			if err != nil {
				t.Fatal(err)
			}

			actualConfig, err := unmarshallYamlFile(localPath + actualConfigFile)
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
