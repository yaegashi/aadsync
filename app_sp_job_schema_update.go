package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPJobSchemaUpdate struct {
	*AppSPJobSchema
}

func (app *AppSPJobSchema) AppSPJobSchemaUpdateComder() cmder.Cmder {
	return &AppSPJobSchemaUpdate{AppSPJobSchema: app}
}

func (app *AppSPJobSchemaUpdate) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "update",
		Short:        "Update schema",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPJobSchemaUpdate) RunE(cmd *cobra.Command, args []string) error {
	b, err := app.ReadInput()
	if err != nil {
		return err
	}
	ctx := context.Background()
	err = app.GetSynchronizationJob(ctx)
	if err != nil {
		return err
	}
	return app.SynchronizationJobRB.Schema().Request().JSONRequest(ctx, http.MethodPut, "", json.RawMessage(b), nil)
}
