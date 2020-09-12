package main

import (
	"context"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPList struct {
	*AppSP
}

func (app *AppSP) AppSPListComder() cmder.Cmder {
	return &AppSPList{AppSP: app}
}

func (app *AppSPList) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "list",
		Short:        "List service principals",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPList) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	err := app.GetServicePrincipalList(ctx)
	if err != nil {
		return err
	}
	return app.Dump(app.ServicePrincipalList)
}
