package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPJobSchemaFunctions struct {
	*AppSPJobSchema
}

func (app *AppSPJobSchema) AppSPJobSchemaFunctionsComder() cmder.Cmder {
	return &AppSPJobSchemaFunctions{AppSPJobSchema: app}
}

func (app *AppSPJobSchemaFunctions) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "functions",
		Short:        "Get schema functions",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPJobSchemaFunctions) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	err := app.GetSynchronizationJob(ctx)
	if err != nil {
		return err
	}
	var resObj json.RawMessage
	err = app.SynchronizationJobRB.Schema().Request().JSONRequest(ctx, http.MethodGet, "/functions", nil, &resObj)
	if err != nil {
		return err
	}
	return app.Dump(resObj)
}
