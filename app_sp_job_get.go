package main

import (
	"context"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPJobGet struct {
	*AppSPJob
}

func (app *AppSPJob) AppSPJobGetComder() cmder.Cmder {
	return &AppSPJobGet{AppSPJob: app}
}

func (app *AppSPJobGet) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "get",
		Short:        "Get job",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPJobGet) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	err := app.GetSynchronizationJob(ctx)
	if err != nil {
		return err
	}
	return app.Dump(app.SynchronizationJob)
}
