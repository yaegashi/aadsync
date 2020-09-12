package main

import (
	"context"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPJobList struct {
	*AppSPJob
}

func (app *AppSPJob) AppSPJobListComder() cmder.Cmder {
	return &AppSPJobList{AppSPJob: app}
}

func (app *AppSPJobList) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "list",
		Short:        "List jobs",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPJobList) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	err := app.GetSynchronizationJobList(ctx)
	if err != nil {
		return err
	}
	return app.Dump(app.SynchronizationJobList)
}
