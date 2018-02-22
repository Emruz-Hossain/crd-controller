package cmds

import (
	"crd-controller/pkg/server"
	"io"

	"github.com/appscode/go/log"
	"github.com/spf13/cobra"
)

func NewCmdRun(out, errOut io.Writer, stopCh <-chan struct{}) *cobra.Command {
	opt := server.NewCrdServerOptions(out, errOut)
	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Run server and controller",
		Long:              "Run server and controller",
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
	return cmd
}
