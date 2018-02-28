package cmds

import (
	"os"

	"github.com/spf13/cobra"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"crd-controller/pkg/server"
	"github.com/appscode/go/log"
)

func NewRootCmd() *cobra.Command {
	opt := server.NewCrdServerOptions(os.Stdout,os.Stderr)
	stopCh := genericapiserver.SetupSignalHandler()
	rootCmd := &cobra.Command{
		Use:               "crd-controller",
		Short:             "run crd-controller",
		Long:              "run crd-controller",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Infof("Starting crd server.......")

			if err := opt.Complete(); err != nil {
				return err
			}
			if err := opt.Validate(args); err != nil {
				return err
			}
			if err := opt.Run(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := rootCmd.Flags()
	opt.RecommendedOptions.AddFlags(flags)
	return rootCmd
}
