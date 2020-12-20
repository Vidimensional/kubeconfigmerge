package cmd

import (
	"github.com/spf13/cobra"
	"github.com/viditests/cmdtest/pkg/kubeconfig"
)

var (
	kubeconfigPath string

	rootCmd = &cobra.Command{
		Use:                   cmdName() + " [flags] NEW_CONFIG...",
		Long:                  "Merges kubernetes config files to your local kubeconfig.",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MinimumNArgs(1),
		RunE:                  Run,
	}
)

func Run(_ *cobra.Command, args []string) error {
	kc, err := kubeconfig.NewFromFile(kubeconfigPath)
	if err != nil {
		return err
	}

	for _, configPath := range args {
		config, err := kubeconfig.NewFromFile(configPath)
		if err != nil {
			return err
		}
		kubeconfig.Merge(kc, config)
	}

	err = kubeconfig.Write(*kc, kubeconfigPath)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&kubeconfigPath,
		"kube-config",
		"k",
		defaultKubeConfigPath(),
		"Path to your local kubeconfig",
	)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
