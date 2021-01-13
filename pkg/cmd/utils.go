package cmd

import (
	"os"
	"os/user"
	"path"
	"path/filepath"
)

func defaultKubeConfigPath() string {
	kubeconfig := os.Getenv("KUBECONFIG")

	if kubeconfig == "" {
		u, err := user.Current()
		if err != nil {
			return ""
		}
		kubeconfig = filepath.Join(u.HomeDir, ".kube/config")
	}

	return kubeconfig

}

func cmdName() string {
	return path.Base(os.Args[0])
}
