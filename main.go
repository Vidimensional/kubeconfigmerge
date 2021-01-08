package main

import (
	"github.com/vidimensional/kubeconfigmerge/pkg/cmd"
	"github.com/vidimensional/kubeconfigmerge/pkg/kubeconfig"
)

func main() {
	kubeRW := kubeconfig.NewReadWriter()
	_ = cmd.Execute(kubeRW)
}
