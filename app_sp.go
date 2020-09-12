package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
	msgraph "github.com/yaegashi/msgraph.go/beta"
	V "github.com/yaegashi/msgraph.go/val"
)

type AppSP struct {
	*App
	SpID                 string
	ServicePrincipalList []msgraph.ServicePrincipal
	ServicePrincipal     *msgraph.ServicePrincipal
	SynchronizationRB    *msgraph.SynchronizationRequestBuilder
}

func (app *App) AppSPComder() cmder.Cmder {
	return &AppSP{App: app}
}

func (app *AppSP) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "sp",
		Short:        "Service principal commands",
		SilenceUsage: true,
	}
	cmd.PersistentFlags().StringVarP(&app.SpID, "sp-id", "", "", "service principal ID / display name")
	return cmd
}

func (app *AppSP) GetServicePrincipalList(ctx context.Context) error {
	if app.GraphClient == nil {
		err := app.GetGraphClient(ctx)
		if err != nil {
			return err
		}
	}
	spList, err := app.GraphClient.ServicePrincipals().Request().Get(ctx)
	if err != nil {
		return err
	}
	app.ServicePrincipalList = spList
	return nil
}

func (app *AppSP) GetServicePrincipal(ctx context.Context) error {
	if app.SpID == "" {
		return fmt.Errorf("Specify service principal ID (--sp-id)")
	}
	if app.ServicePrincipalList == nil {
		err := app.GetServicePrincipalList(ctx)
		if err != nil {
			return err
		}
	}
	for _, sp := range app.ServicePrincipalList {
		if V.String(sp.AppID) == app.SpID || V.String(sp.DisplayName) == app.SpID {
			app.ServicePrincipal = &sp
			break
		}
	}
	if app.ServicePrincipal == nil {
		return fmt.Errorf("Service principal for app ID %q is not found", app.SpID)
	}
	app.SynchronizationRB = app.GraphClient.ServicePrincipals().ID(*app.ServicePrincipal.ID).Synchronization()
	return nil
}
