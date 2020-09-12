package main

import (
	"fmt"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPJobDelete struct {
	*AppSPJob
	Scope string
}

func (app *AppSPJob) AppSPJobDeleteComder() cmder.Cmder {
	return &AppSPJobDelete{AppSPJob: app}
}

func (app *AppSPJobDelete) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "delete",
		Short:        "Delete job (not implemented)",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPJobDelete) RunE(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("Not implemented")
}
