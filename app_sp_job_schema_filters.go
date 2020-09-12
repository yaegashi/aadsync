package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPJobSchemaFilters struct {
	*AppSPJobSchema
}

func (app *AppSPJobSchema) AppSPJobSchemaFiltersComder() cmder.Cmder {
	return &AppSPJobSchemaFilters{AppSPJobSchema: app}
}

func (app *AppSPJobSchemaFilters) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "filters",
		Aliases:      []string{"filter-operators"},
		Short:        "Get schema filter operators",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPJobSchemaFilters) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	err := app.GetSynchronizationJob(ctx)
	if err != nil {
		return err
	}
	var resObj json.RawMessage
	err = app.SynchronizationJobRB.Schema().Request().JSONRequest(ctx, http.MethodGet, "/filterOperators", nil, &resObj)
	if err != nil {
		return err
	}
	return app.Dump(resObj)
}
