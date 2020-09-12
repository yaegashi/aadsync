package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPJobSchemaGet struct {
	*AppSPJobSchema
}

func (app *AppSPJobSchema) AppSPJobSchemaGetComder() cmder.Cmder {
	return &AppSPJobSchemaGet{AppSPJobSchema: app}
}

func (app *AppSPJobSchemaGet) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "get",
		Short:        "Get schema",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPJobSchemaGet) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	err := app.GetSynchronizationJob(ctx)
	if err != nil {
		return err
	}
	var resObj json.RawMessage
	err = app.SynchronizationJobRB.Schema().Request().JSONRequest(ctx, http.MethodGet, "", nil, &resObj)
	if err != nil {
		return err
	}
	return app.Dump(resObj)
}
