package kubeconfig

import (
	"errors"

	"github.com/imdario/mergo"
	"github.com/vidimensional/kubeconfigmerge/pkg/file"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

var ErrKubeconfigNotFound = errors.New("kubeconfig not found")

type funcLoadFromFile func(filename string) (*clientcmdapi.Config, error)
type funcWriteToFile func(config clientcmdapi.Config, filename string) error
type funcFileExists func(filename string) bool

type ReadWriter struct {
	loadFromFile funcLoadFromFile
	writeToFile  funcWriteToFile
	fileExists   funcFileExists
}

func NewReadWriter() *ReadWriter {
	return &ReadWriter{
		loadFromFile: clientcmd.LoadFromFile,
		writeToFile:  clientcmd.WriteToFile,
		fileExists:   file.Exists,
	}
}

// Read Loads the config from filename and returns a *Config with the contents.
// If the file is not found on the FileSystem it returns an empty *Config and a ErrKubeconfigNotFound error.
func (kc *ReadWriter) Read(filename string) (*clientcmdapi.Config, error) {
	if !kc.fileExists(filename) {
		return clientcmdapi.NewConfig(), ErrKubeconfigNotFound
	}

	conf, err := kc.loadFromFile(filename)
	if err != nil {
		return nil, err
	}
	conf.SetGroupVersionKind(schema.GroupVersionKind{Version: "v1", Kind: "Config"})

	return conf, err
}

// Write saves the content of config to the Filesystem specified by filename in YAML format.
func (kc *ReadWriter) Write(config clientcmdapi.Config, filename string) error {
	return kc.writeToFile(config, filename)
}

func Merge(dst, src *clientcmdapi.Config) error {
	return mergo.Merge(dst, src)
}
