package main

import (
	"fmt"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPJobValidate struct {
	*AppSPJob
	Scope string
}

func (app *AppSPJob) AppSPJobValidateComder() cmder.Cmder {
	return &AppSPJobValidate{AppSPJob: app}
}

func (app *AppSPJobValidate) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "validate",
		Aliases:      []string{"validate-credentials"},
		Short:        "Validate credentials (not implemented)",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPJobValidate) RunE(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("Not implemented")
}
