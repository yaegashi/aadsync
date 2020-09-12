package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
	msgraph "github.com/yaegashi/msgraph.go/beta"
)

type AppSPTemplate struct {
	*AppSP
	TemplateID                  string
	SynchronizationTemplateList []msgraph.SynchronizationTemplate
	SynchronizationTemplate     *msgraph.SynchronizationTemplate
}

func (app *AppSP) AppSPTemplateCmder() cmder.Cmder {
	return &AppSPTemplate{AppSP: app}
}

func (app *AppSPTemplate) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "template",
		Short:        "Synchronization template commands",
		SilenceUsage: true,
	}
	cmd.PersistentFlags().StringVarP(&app.TemplateID, "template-id", "", "AD2AADProvisioning", "tempalte ID")
	return cmd
}

func (app *AppSPTemplate) GetSynchronizationTemplateList(ctx context.Context) error {
	if app.ServicePrincipal == nil {
		err := app.GetServicePrincipal(ctx)
		if err != nil {
			return err
		}
	}
	templates, err := app.SynchronizationRB.Templates().Request().Get(ctx)
	if err != nil {
		return err
	}
	app.SynchronizationTemplateList = templates
	return nil
}

func (app *AppSPTemplate) GetSynchronizationTemplate(ctx context.Context) error {
	if app.SynchronizationTemplateList == nil {
		err := app.GetSynchronizationTemplateList(ctx)
		if err != nil {
			return err
		}
	}
	for _, template := range app.SynchronizationTemplateList {
		if *template.ID == app.TemplateID {
			app.SynchronizationTemplate = &template
			break
		}
	}
	if app.SynchronizationTemplate == nil {
		return fmt.Errorf("Synchronization template for %q is not found", app.TemplateID)
	}
	return nil
}
