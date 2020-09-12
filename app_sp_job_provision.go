package main

import (
	"fmt"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPJobProvision struct {
	*AppSPJob
	Scope string
}

func (app *AppSPJob) AppSPJobProvisionComder() cmder.Cmder {
	return &AppSPJobProvision{AppSPJob: app}
}

func (app *AppSPJobProvision) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "provision",
		Aliases:      []string{"provision-on-demand"},
		Short:        "Provision on demand (not implemented)",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPJobProvision) RunE(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("Not implemented")
}
