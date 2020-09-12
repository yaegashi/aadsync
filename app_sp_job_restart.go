package main

import (
	"context"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
	msgraph "github.com/yaegashi/msgraph.go/beta"
)

type AppSPJobRestart struct {
	*AppSPJob
	Scope string
}

func (app *AppSPJob) AppSPJobRestartComder() cmder.Cmder {
	return &AppSPJobRestart{AppSPJob: app}
}

func (app *AppSPJobRestart) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "restart",
		Short:        "Restart job",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	cmd.Flags().StringVarP(&app.Scope, "scope", "", "Full", "job restart criteria reset scope")
	return cmd
}

func (app *AppSPJobRestart) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	err := app.GetSynchronizationJob(ctx)
	if err != nil {
		return err
	}
	scope := msgraph.SynchronizationJobRestartScope(app.Scope)
	reqObj := &msgraph.SynchronizationJobRestartRequestParameter{Criteria: &msgraph.SynchronizationJobRestartCriteria{ResetScope: &scope}}
	return app.SynchronizationJobRB.Restart(reqObj).Request().Post(ctx)
}
