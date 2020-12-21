package kubeconfig

import (
	"errors"
	"reflect"
	"testing"

	"k8s.io/client-go/tools/clientcmd"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	filename = "/path/to/kubeconfig"
)

var (
	errGeneric = errors.New("generic error")
)

func TestNewReadWriter(t *testing.T) {
	rw := NewReadWriter()

	assert.Equal(t, funcPointer(clientcmd.WriteToFile), funcPointer(rw.writeToFile))
	assert.Equal(t, funcPointer(clientcmd.LoadFromFile), funcPointer(rw.loadFromFile))
	assert.Equal(t, funcPointer(fileExists), funcPointer(rw.fileExists))
}

func TestReadWhenFileDoesNotExistReturnsEmptyConfigAndErrKubeconfigNotFound(t *testing.T) {
	krw := new(ReadWriter)
	krw.fileExists = fileExistsReturnsFalse

	expectedConfig := &clientcmdapi.Config{
		Preferences: *clientcmdapi.NewPreferences(),
		Clusters:    make(map[string]*clientcmdapi.Cluster),
		AuthInfos:   make(map[string]*clientcmdapi.AuthInfo),
		Contexts:    make(map[string]*clientcmdapi.Context),
		Extensions:  make(map[string]runtime.Object),
	}

	actualConfig, err := krw.Read(filename)

	assert.Equal(t, expectedConfig, actualConfig)
	assert.Equal(t, ErrKubeconfigNotFound, err)
}

func TestReadWhenReaderFailsReturnsNilConfigAndForwardsReaderError(t *testing.T) {
	krw := new(ReadWriter)
	krw.fileExists = fileExistsReturnsTrue
	krw.loadFromFile = loadFromFileReturnsNilAndError

	config, err := krw.Read(filename)

	assert.Equal(t, errGeneric, err)
	assert.Nil(t, config)
}

func TestReadWhenReaderWorksOkReturnsConfigAndNilError(t *testing.T) {
	krw := new(ReadWriter)
	krw.fileExists = fileExistsReturnsTrue
	krw.loadFromFile = loadFromFileReturnsGenericConfigAndNoError

	expectedConfig := &clientcmdapi.Config{
		Kind:       "Config",
		APIVersion: "v1",
	}

	actualConfig, err := krw.Read(filename)

	assert.Equal(t, expectedConfig, actualConfig)
	assert.Nil(t, err)
}

func TestWriteReturnsOk(t *testing.T) {
	krw := new(ReadWriter)
	krw.writeToFile = writeToFileOk

	cfg := clientcmdapi.Config{}

	err := krw.Write(cfg, filename)

	assert.Nil(t, err)
}

func TestWriteReturnsError(t *testing.T) {
	krw := new(ReadWriter)
	krw.writeToFile = writeToFileError

	cfg := clientcmdapi.Config{}

	err := krw.Write(cfg, filename)

	assert.Equal(t, errGeneric, err)
}

func TestMerge(t *testing.T) {
	cfg1 := &clientcmdapi.Config{
		Kind: "Config",
	}

	cfg2 := &clientcmdapi.Config{
		APIVersion: "v1",
	}

	expectedCfg := &clientcmdapi.Config{
		Kind:       "Config",
		APIVersion: "v1",
	}

	err := Merge(cfg1, cfg2)
	assert.Equal(t, expectedCfg, cfg1)
	assert.Nil(t, err)
}

func funcPointer(f interface{}) uintptr {
	return reflect.ValueOf(f).Pointer()
}

func fileExistsReturnsFalse(string) bool {
	return false
}

func fileExistsReturnsTrue(string) bool {
	return true
}

func loadFromFileReturnsNilAndError(string) (*clientcmdapi.Config, error) {
	return nil, errGeneric
}

func loadFromFileReturnsGenericConfigAndNoError(string) (*clientcmdapi.Config, error) {
	return &clientcmdapi.Config{}, nil
}

func writeToFileOk(clientcmdapi.Config, string) error {
	return nil
}

func writeToFileError(clientcmdapi.Config, string) error {
	return errGeneric
}
