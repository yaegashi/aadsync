package main

import (
	"context"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type AppSPTemplateList struct {
	*AppSPTemplate
}

func (app *AppSPTemplate) AppSPTemplateListCmder() cmder.Cmder {
	return &AppSPTemplateList{AppSPTemplate: app}
}

func (app *AppSPTemplateList) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "list",
		Short:        "List synchronization templates",
		RunE:         app.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (app *AppSPTemplateList) RunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	err := app.GetSynchronizationTemplateList(ctx)
	if err != nil {
		return err
	}
	return app.Dump(app.SynchronizationTemplateList)
}
