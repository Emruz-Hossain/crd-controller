package main

import (
	 //"k8s.io/code-generator"

	"crd-controller/pkg/cmds"

	"flag"
	"github.com/appscode/go/log"
)

func main() {
	//controller.StartDeploymentController(1)
	rootCmd:= cmds.NewRootCmd()
	rootCmd.Flags().AddGoFlagSet(flag.CommandLine)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf(err.Error())
	}
	select {}
}
