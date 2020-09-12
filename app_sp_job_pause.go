package main

import (
	"context"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPJobPause struct {
	*AppSPJob
}

func (app *AppSPJob) AppSPJobPauseComder() cmder.Cmder {
	return &AppSPJobPause{AppSPJob: app}
}

func (app *AppSPJobPause) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "pause",
		Short:        "Pause job",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPJobPause) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	err := app.GetSynchronizationJob(ctx)
	if err != nil {
		return err
	}
	return app.SynchronizationJobRB.Pause(nil).Request().Post(ctx)
}
