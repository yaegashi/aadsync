package main

import (
	"context"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPGet struct {
	*AppSP
}

func (app *AppSP) AppSPGetComder() cmder.Cmder {
	return &AppSPGet{AppSP: app}
}

func (app *AppSPGet) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "get",
		Short:        "Get service principal",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPGet) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	err := app.GetServicePrincipal(ctx)
	if err != nil {
		return err
	}
	return app.Dump(app.ServicePrincipal)
}
