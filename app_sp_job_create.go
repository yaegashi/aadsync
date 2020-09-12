package main

import (
	"fmt"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPJobCreate struct {
	*AppSPJob
	Scope string
}

func (app *AppSPJob) AppSPJobCreateComder() cmder.Cmder {
	return &AppSPJobCreate{AppSPJob: app}
}

func (app *AppSPJobCreate) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "create",
		Short:        "Create job (not implemented)",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPJobCreate) RunE(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("Not implemented")
}
