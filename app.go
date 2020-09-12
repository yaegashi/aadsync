package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/yaegashi/aadsync/store"
	msgraph "github.com/yaegashi/msgraph.go/beta"
	"github.com/yaegashi/msgraph.go/jsonx"
	"github.com/yaegashi/msgraph.go/msauth"
	"golang.org/x/oauth2"
)

const (
	environConfigDir    = "AADSYNC_CONFIG_DIR"
	defaultConfigDir    = "~/.aadsync"
	environTokenFile    = "AADSYNC_TOKEN_FILE"
	defaultTokenFile    = "token.json"
	environClientID     = "AZURE_CLIENT_ID"
	defaultClientID     = "31a04eeb-635f-4b95-82b8-5c3acd45840e"
	environClientSecret = "AZURE_CLIENT_SECRET"
	defaultClientSecret = ""
	environTenantID     = "AZURE_TENANT_ID"
	defaultTenantID     = "common"
)

var defaultScopes = []string{"offline_access", "Directory.ReadWrite.All"}

type App struct {
	Writer       io.WriteCloser
	GraphClient  *msgraph.GraphServiceRequestBuilder
	ConfigStore  *store.Store
	ConfigDir    string
	TokenFile    string
	ClientID     string
	ClientSecret string
	TenantID     string
	Input        string
	Output       string
	IsStdin      bool
	IsStdout     bool
	Quiet        bool
}

func (app *App) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "aadsync",
		Short:             "Azure AD synchronization CLI",
		PersistentPreRunE: app.PersistentPreRunE,
		SilenceUsage:      true,
		Version:           fmt.Sprintf("%s (%-0.7s)", version, commit),
	}
	cmd.PersistentFlags().StringVarP(&app.ConfigDir, "config-dir", "", "", envHelp("config dir", environConfigDir, defaultConfigDir))
	cmd.PersistentFlags().StringVarP(&app.TokenFile, "token-file", "", "", envHelp("token file", environTokenFile, defaultTokenFile))
	cmd.PersistentFlags().StringVarP(&app.ClientID, "client-id", "", "", envHelp("Azure client ID", environClientID, defaultClientID))
	cmd.PersistentFlags().StringVarP(&app.ClientSecret, "client-secret", "", "", envHelp("Azure client secret", environClientSecret, defaultClientSecret))
	cmd.PersistentFlags().StringVarP(&app.TenantID, "tenant-id", "", "", envHelp("Azure tenant ID", environTenantID, defaultTenantID))
	cmd.PersistentFlags().StringVarP(&app.Input, "input", "i", "", "input file")
	cmd.PersistentFlags().StringVarP(&app.Output, "output", "o", "", "output file")
	cmd.PersistentFlags().BoolVarP(&app.Quiet, "quiet", "q", false, "quiet")
	return cmd
}

func envDefault(val, env, def string) string {
	if val == "" {
		val = os.Getenv(env)
	}
	if val == "" {
		val = def
	}
	return val
}

func envHelp(msg, env, def string) string {
	return fmt.Sprintf(`%s (env:%s, default:%q)`, msg, env, def)
}

func (app *App) PersistentPreRunE(cmd *cobra.Command, args []string) error {
	app.ConfigDir = envDefault(app.ConfigDir, environConfigDir, defaultConfigDir)
	app.TokenFile = envDefault(app.TokenFile, environTokenFile, defaultTokenFile)
	app.ClientID = envDefault(app.ClientID, environClientID, defaultClientID)
	app.ClientSecret = envDefault(app.ClientSecret, environClientSecret, defaultClientSecret)
	app.TenantID = envDefault(app.TenantID, environTenantID, defaultTenantID)

	store, err := store.NewStore(app.ConfigDir)
	if err != nil {
		return err
	}
	app.ConfigStore = store

	app.IsStdin = app.Input == "" || app.Input == "-"
	app.IsStdout = app.Output == "" || app.Output == "-"

	return nil
}

func (app *App) Authorize() (oauth2.TokenSource, error) {
	ctx := context.Background()
	m := msauth.NewManager()
	if app.ClientSecret != "" {
		return m.ClientCredentialsGrant(ctx, app.TenantID, app.ClientID, app.ClientSecret, defaultScopes)
	}
	loc, _ := app.ConfigStore.Location(app.TokenFile, true)
	app.Logf("Loading token in %s", loc)
	b, err := app.ConfigStore.ReadFile(app.TokenFile)
	if err == nil {
		err = m.LoadBytes(b)
	}
	if err != nil {
		app.Logf("Warning: %s", err)
	}
	ts, err := m.DeviceAuthorizationGrant(ctx, app.TenantID, app.ClientID, defaultScopes, nil)
	if err != nil {
		return nil, err
	}
	b, err = m.SaveBytes()
	if err == nil {
		app.Logf("Saving token in %s", loc)
		err = app.ConfigStore.WriteFile(app.TokenFile, b, 0600)
	}
	if err != nil {
		app.Logf("Warning: %s", err)
	}
	return ts, nil
}

func (app *App) GetGraphClient(ctx context.Context) error {
	ts, err := app.Authorize()
	if err != nil {
		return err
	}
	httpClient := oauth2.NewClient(ctx, ts)
	graphClient := msgraph.NewClient(httpClient)
	app.GraphClient = graphClient
	return nil
}

func (app *App) ReadInput() ([]byte, error) {
	if app.IsStdin {
		return ioutil.ReadAll(os.Stdin)
	}
	return ioutil.ReadFile(app.Input)
}

func (app *App) WriteOutput(b []byte) error {
	if app.IsStdout {
		_, err := os.Stdout.Write(b)
		return err
	}
	return ioutil.WriteFile(app.Output, b, 0644)
}

func (app *App) Dump(o interface{}) error {
	b, err := jsonx.MarshalIndent(o, "", "  ")
	if err != nil {
		return err
	}
	b = append(b, '\n')
	return app.WriteOutput(b)
}

func (app *App) Log(args ...interface{}) {
	if !app.Quiet {
		log.Print(args...)
	}
}

func (app *App) Logln(args ...interface{}) {
	if !app.Quiet {
		log.Println(args...)
	}
}

func (app *App) Logf(format string, args ...interface{}) {
	if !app.Quiet {
		log.Printf(format, args...)
	}
}
