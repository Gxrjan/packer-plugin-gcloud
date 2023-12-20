package main

import (
	"fmt"
	"os"

	"github.com/Gxrjan/packer-plugin-gcloud/customfile"
	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterProvisioner(plugin.DEFAULT_NAME, new(customfile.Provisioner))
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
