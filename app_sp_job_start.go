package main

import (
	"context"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPJobStart struct {
	*AppSPJob
}

func (app *AppSPJob) AppSPJobStartComder() cmder.Cmder {
	return &AppSPJobStart{AppSPJob: app}
}

func (app *AppSPJobStart) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "start",
		Short:        "Start job",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPJobStart) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	err := app.GetSynchronizationJob(ctx)
	if err != nil {
		return err
	}
	return app.SynchronizationJobRB.Start(nil).Request().Post(ctx)
}
