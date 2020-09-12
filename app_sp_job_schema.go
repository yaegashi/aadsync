package main

import (
	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPJobSchema struct {
	*AppSPJob
	JobTemplateID string
}

func (app *AppSPJob) AppSPJobSchemaComder() cmder.Cmder {
	return &AppSPJobSchema{AppSPJob: app}
}

func (app *AppSPJobSchema) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "schema",
		Short:        "Synchronization job schema commands",
		SilenceUsage: true,
	}
	return cmd
}
