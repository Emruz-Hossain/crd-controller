package main

import (
	// "k8s.io/code-generator"

	"crd-controller/pkg/cmds"
	"fmt"
)

func main() {
	//controller.StartDeploymentController(1)
	err := cmds.NewRootCmd().Execute()
	if err != nil {
		fmt.Errorf("Error in runnig root command. Reason: ", err.Error())
	}
}
