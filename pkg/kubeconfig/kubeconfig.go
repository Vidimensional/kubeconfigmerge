package kubeconfig

import (
	"os"

	"github.com/imdario/mergo"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func New() *clientcmdapi.Config {
	return clientcmdapi.NewConfig()
}

func NewFromFile(filename string) (*clientcmdapi.Config, error) {
	if !fileExists(filename) {
		return clientcmdapi.NewConfig(), nil
	}

	conf, err := clientcmd.LoadFromFile(filename)
	conf.SetGroupVersionKind(schema.GroupVersionKind{Version: "v1", Kind: "Config"})
	return conf, err
}

func Write(config clientcmdapi.Config, filename string) error {
	return clientcmd.WriteToFile(config, filename)
}

func Merge(dst, src *clientcmdapi.Config) {
	mergo.Merge(dst, src)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
