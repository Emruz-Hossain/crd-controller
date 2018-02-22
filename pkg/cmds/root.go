package cmds

import (
	"os"

	"github.com/spf13/cobra"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:               "crd-controller",
		Short:             "run crd-controller",
		Long:              "run crd-controller",
		DisableAutoGenTag: true,
	}

	stopCh := genericapiserver.SetupSignalHandler()
	rootCmd.AddCommand(NewCmdRun(os.Stdout, os.Stderr, stopCh))
	return rootCmd
}
