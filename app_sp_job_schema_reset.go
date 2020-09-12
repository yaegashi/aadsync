package main

import (
	"context"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPJobSchemaReset struct {
	*AppSPJobSchema
}

func (app *AppSPJobSchema) AppSPJobSchemaResetComder() cmder.Cmder {
	return &AppSPJobSchemaReset{AppSPJobSchema: app}
}

func (app *AppSPJobSchemaReset) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "reset",
		Short:        "Reset schema",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPJobSchemaReset) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	err := app.GetSynchronizationJob(ctx)
	if err != nil {
		return err
	}
	return app.SynchronizationJobRB.Schema().Request().Delete(ctx)
}
