package main

import (
	"context"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPTemplateGet struct {
	*AppSPTemplate
}

func (app *AppSPTemplate) AppSPTemplateGetCmder() cmder.Cmder {
	return &AppSPTemplateGet{AppSPTemplate: app}
}

func (app *AppSPTemplateGet) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "get",
		Short:        "Get synchronization template",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPTemplateGet) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	err := app.GetSynchronizationTemplate(ctx)
	if err != nil {
		return err
	}
	return app.Dump(app.SynchronizationTemplate)
}
