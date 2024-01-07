package main

import (
	"log"

	"ocicmd/compartments"

	"github.com/spf13/cobra"
)

func run() error {
	rootCmdName := "ocicmd"
	cmd := &cobra.Command{
		Use:       rootCmdName,
		Short:     "Helpful OCI tools",
		Long:      "Additional functionality beyond what is provided by the `oci` CLI provided by Oracle.",
		ValidArgs: []string{"compartments"},
		Args:      cobra.OnlyValidArgs,
	}
	cmd.AddCommand(compartments.ListCmd())
	return cmd.Execute()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}
