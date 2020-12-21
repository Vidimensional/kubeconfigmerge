package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vidimensional/kubeconfigmerge/pkg/kubeconfig"
)

var (
	kubeconfigPath       string
	kubeconfigReadWriter *kubeconfig.ReadWriter

	rootCmd = &cobra.Command{
		Use:                   cmdName() + " [flags] NEW_CONFIG...",
		Long:                  "Merges kubernetes config files to your local kubeconfig.",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MinimumNArgs(1),
		RunE:                  Run,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&kubeconfigPath,
		"kube-config",
		"k",
		defaultKubeConfigPath(),
		"Path to your local kubeconfig",
	)
}

func Run(_ *cobra.Command, args []string) error {
	kc, err := kubeconfigReadWriter.Read(kubeconfigPath)
	if err != nil && err != kubeconfig.ErrKubeconfigNotFound {
		return err
	}

	for _, configPath := range args {
		config, err := kubeconfigReadWriter.Read(configPath)
		if err != nil {
			return err
		}
		kubeconfig.Merge(kc, config)
	}

	err = kubeconfigReadWriter.Write(*kc, kubeconfigPath)
	if err != nil {
		return err
	}

	return nil
}

// Execute executes the root command.
func Execute(kubeRW *kubeconfig.ReadWriter) error {
	kubeconfigReadWriter = kubeRW
	return rootCmd.Execute()
}
