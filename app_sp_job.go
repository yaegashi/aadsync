package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
	msgraph "github.com/yaegashi/msgraph.go/beta"
)

type AppSPJob struct {
	*AppSP
	JobID                    string
	SynchronizationJobList   []msgraph.SynchronizationJob
	SynchronizationJobListRB *msgraph.SynchronizationJobsCollectionRequestBuilder
	SynchronizationJob       *msgraph.SynchronizationJob
	SynchronizationJobRB     *msgraph.SynchronizationJobRequestBuilder
}

func (app *AppSP) AppSPJobComder() cmder.Cmder {
	return &AppSPJob{AppSP: app}
}

func (app *AppSPJob) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "job",
		Short:        "Synchronization job commands",
		SilenceUsage: true,
	}
	cmd.PersistentFlags().StringVarP(&app.JobID, "job-id", "", "AD2AADProvisioning", "job ID / tempalte ID")
	return cmd
}

func (app *AppSPJob) GetSynchronizationJobList(ctx context.Context) error {
	if app.ServicePrincipal == nil {
		err := app.GetServicePrincipal(ctx)
		if err != nil {
			return err
		}
	}
	app.SynchronizationJobListRB = app.SynchronizationRB.Jobs()
	jobs, err := app.SynchronizationJobListRB.Request().Get(ctx)
	if err != nil {
		return err
	}
	app.SynchronizationJobList = jobs
	return nil
}

func (app *AppSPJob) GetSynchronizationJob(ctx context.Context) error {
	if app.SynchronizationJobList == nil {
		err := app.GetSynchronizationJobList(ctx)
		if err != nil {
			return err
		}
	}
	for _, job := range app.SynchronizationJobList {
		if *job.ID == app.JobID || *job.TemplateID == app.JobID {
			app.SynchronizationJob = &job
			break
		}
	}
	if app.SynchronizationJob == nil {
		return fmt.Errorf("Synchronization Job for %q is not found", app.JobID)
	}
	app.SynchronizationJobRB = app.SynchronizationJobListRB.ID(*app.SynchronizationJob.ID)
	return nil
}
