package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path"
)

func defaultKubeConfigPath() string {
	kubeconfig := os.Getenv("KUBECONFIG")

	if kubeconfig == "" {
		u, err := user.Current()
		if err != nil {
			return ""
		}
		kubeconfig = fmt.Sprintf("%s/.kube/config", u.HomeDir)
	}

	return kubeconfig

}

func cmdName() string {
	return path.Base(os.Args[0])
}
