package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path"
)

func defaultKubeConfigPath() string {
	u, err := user.Current()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s/.kube/config", u.HomeDir)
}

func cmdName() string {
	return path.Base(os.Args[0])
}
